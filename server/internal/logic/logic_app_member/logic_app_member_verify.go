package logic_app_member

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"APT/utility/encrypt"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	memberCodeLogger = g.Log().Path("logs/MemberCode")
)

// verifyResult 事务内核销成功后，需要在事务外执行的 IO 数据
type verifyResult struct {
	memberCouponId int64
	couponId       uint64
	couponMchName  string
	member         *entity.PmsMember
	employee       *entity.Employee
}

// VerifyMemberCode 会员码验证（发券+核销）
func (s *sAppMember) VerifyMemberCode(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error) {
	memberCodeLogger.Info(ctx, "--------会员码验证进入----------")

	valArr := strings.Split(in.Val, "|")

	// 解析会员码：格式为 {AES加密(memberNo|timestamp|MEMBER)}
	encryptedData, base64Err := base64.StdEncoding.DecodeString(valArr[0])
	if base64Err != nil {
		err = gerror.Wrap(base64Err, "会员码解码失败")
		memberCodeLogger.Errorf(ctx, "会员码解码失败: %v", base64Err)
		return
	}
	decryptedText, decryptErr := encrypt.AesECBDecrypt(encryptedData, consts.RequestEncryptKey)
	if decryptErr != nil {
		err = gerror.Wrap(decryptErr, "会员码解密失败")
		memberCodeLogger.Errorf(ctx, "会员码解密失败: %v", decryptErr)
		return
	}
	parts := strings.Split(string(decryptedText), "|")

	memberNo := parts[0]
	timestampStr := parts[1]

	ts, parseErr := strconv.ParseInt(timestampStr, 10, 64)
	if parseErr != nil {
		err = gerror.New("会员码时间戳格式不正确！")
		memberCodeLogger.Errorf(ctx, "会员码时间戳解析失败: %v", parseErr)
		return
	}
	if gtime.Now().Unix()-ts > 60 {
		err = gerror.New("二维码已过期，请刷新后重试！")
		memberCodeLogger.Warningf(ctx, "会员码已过期, sn=%s ts=%d", in.Sn, ts)
		return
	}

	// 获取终端信息
	var terminalModel *input_basics.TerminalViewModel
	if err = dao.SysTerminal.Ctx(ctx).
		Where(dao.SysTerminal.Columns().Sn, in.Sn).
		WithAll().
		Scan(&terminalModel); err != nil {
		err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
		memberCodeLogger.Errorf(ctx, "获取终端信息失败, sn=%s: %v", in.Sn, err)
		return
	}
	if terminalModel == nil || terminalModel.StoreId <= 0 {
		err = gerror.New("终端未绑定门店，无法使用会员码")
		memberCodeLogger.Errorf(ctx, "终端未绑定门店, sn=%s", in.Sn)
		return
	}

	// 提前查 ThStoreTerminal（整个请求只需一次）
	var thStoreTerminal *entity.ThStoreTerminal
	_ = dao.ThStoreTerminal.Ctx(ctx).
		Where(dao.ThStoreTerminal.Columns().StoreId, terminalModel.StoreId).
		Where(dao.ThStoreTerminal.Columns().TerminalId, terminalModel.Id).
		Scan(&thStoreTerminal)

	// 事务内：发券 + 核销 DB 操作，收集需要事务外执行的 IO 数据
	var results []*verifyResult
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var txErr error
		results, txErr = s.handleMemberCodeIssue(ctx, tx, memberNo, terminalModel)
		return txErr
	})
	if err != nil {
		memberCodeLogger.Errorf(ctx, "会员码发券/核销事务失败, memberNo=%s sn=%s: %v", memberNo, in.Sn, err)
		return
	}

	// 事务外：打印 + WebSocket（IO 操作不占事务锁）
	for _, r := range results {
		s.printMemberCouponReceipt(ctx, terminalModel, thStoreTerminal, r.couponMchName, r.memberCouponId)
		memberCodeLogger.Infof(ctx, "员工 %d 券 %d (id=%d) 核销成功", r.employee.Id, r.couponId, r.memberCouponId)
	}
	return
}

// handleMemberCodeIssue 处理会员码发券逻辑（事务内）
func (s *sAppMember) handleMemberCodeIssue(ctx context.Context, tx gdb.TX, memberNo string, terminalModel *input_basics.TerminalViewModel) (results []*verifyResult, err error) {

	// 根据 memberNo 查询会员
	var member *entity.PmsMember
	if err = dao.PmsMember.Ctx(ctx).TX(tx).Where(dao.PmsMember.Columns().MemberNo, memberNo).Scan(&member); err != nil {
		err = gerror.Wrap(err, "查询会员信息失败！")
		return
	}
	if g.IsEmpty(member) {
		err = gerror.New("会员不存在！")
		return
	}

	// 验证是否是员工
	var employee *entity.Employee
	if err = dao.Employee.Ctx(ctx).TX(tx).
		Where(dao.Employee.Columns().MemberId, member.Id).
		Where(dao.Employee.Columns().Status, 1).
		Scan(&employee); err != nil {
		err = gerror.Wrap(err, "查询员工信息失败！")
		return
	}
	if g.IsEmpty(employee) {
		err = gerror.New("非员工无法使用此功能！")
		return
	}

	storeId := terminalModel.StoreInfo.Id
	mchId := terminalModel.StoreInfo.MchId
	terminalId := int(terminalModel.Id)

	// 查询该商户关联的所有礼品券（含 name）
	type couponMchRow struct {
		CouponId int    `json:"couponId" orm:"coupon_id"`
		Name     string `json:"name"     orm:"name"`
	}
	var couponMchList []*couponMchRow
	if err = dao.ThCouponMch.Ctx(ctx).TX(tx).
		Fields("coupon_id, name").
		Where(dao.ThCouponMch.Columns().MchId, mchId).
		Scan(&couponMchList); err != nil {
		err = gerror.Wrap(err, "查询商户礼品券失败！")
		return
	}
	if len(couponMchList) == 0 {
		err = gerror.New("没有券可以发放")
		return
	}
	couponIds := make([]int, 0, len(couponMchList))
	couponMchNameMap := make(map[int]string, len(couponMchList))
	for _, c := range couponMchList {
		couponIds = append(couponIds, c.CouponId)
		couponMchNameMap[c.CouponId] = c.Name
	}

	// 查询员工活动券（有效且不需要预约）
	type activityCouponRow struct {
		ActivityId        uint64 `json:"activityId"        orm:"activity_id"`
		CouponId          uint64 `json:"couponId"          orm:"coupon_id"`
		AvailableQuantity int    `json:"availableQuantity" orm:"available_quantity"`
		PerDayVerify      int    `json:"perDayVerify"      orm:"per_day_verify"`
	}
	var activityCouponList []*activityCouponRow

	mod := dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).
		Fields("hg_employee_activity_coupon.activity_id, hg_employee_activity_coupon.coupon_id, hg_employee_activity_coupon.available_quantity, hg_employee_activity_coupon.per_day_verify").
		InnerJoin("hg_employee_activity", "hg_employee_activity.id = hg_employee_activity_coupon.activity_id").
		InnerJoin("hg_th_coupon", "hg_th_coupon.id = hg_employee_activity_coupon.coupon_id").
		Where("hg_employee_activity_coupon.coupon_id IN(?)", couponIds).
		Where("hg_employee_activity.status", 2).
		Where("hg_th_coupon.need_reservation", 0).
		Where("hg_th_coupon.use_status", 1).
		Where(
			"(hg_employee_activity.restriction_type = 1"+
				" OR (hg_employee_activity.restriction_type = 2 AND EXISTS (SELECT 1 FROM hg_employee_activity_department WHERE activity_id = hg_employee_activity.id AND department_id = ?))"+
				" OR (hg_employee_activity.restriction_type = 3 AND EXISTS (SELECT 1 FROM hg_employee_activity_employee WHERE activity_id = hg_employee_activity.id AND employee_id = ?)))",
			employee.DepartmentId, employee.Id,
		)

	if err = mod.Scan(&activityCouponList); err != nil {
		err = gerror.Wrap(err, "查询员工活动礼品券失败！")
		return
	}
	if len(activityCouponList) == 0 {
		err = gerror.New("没有券可以发放")
		return
	}

	// 每次扫码只领取并核销一张券，找到第一个可处理的活动券即停止
	for _, ac := range activityCouponList {
		couponMchName := couponMchNameMap[int(ac.CouponId)]
		issueResult, issueErr := s.autoIssueMemberCoupon(ctx, tx, employee, member, ac.ActivityId, ac.CouponId, ac.AvailableQuantity, ac.PerDayVerify, storeId, mchId, terminalId, couponMchName)
		if issueErr != nil {
			memberCodeLogger.Warningf(ctx, "用户[%d]发放/核销券 %d 失败: %v", member.Id, ac.CouponId, issueErr)
			continue
		}
		if issueResult != nil {
			results = []*verifyResult{issueResult}
			break
		}
	}

	return
}

// autoIssueMemberCoupon 每次扫码只领取一张券并核销（全程在外层 tx 内）
// 返回本次核销成功的 verifyResult，nil 表示该活动券无可处理项（非错误）
func (s *sAppMember) autoIssueMemberCoupon(
	ctx context.Context,
	tx gdb.TX,
	employee *entity.Employee,
	member *entity.PmsMember,
	activityId uint64,
	couponId uint64,
	availableQuantity int,
	perDayVerify int,
	storeId, mchId, terminalId int,
	couponMchName string,
) (result *verifyResult, err error) {

	// 今日核销额度检查
	if perDayVerify > 0 {
		todayStart := gtime.Now().Format("Y-m-d") + " 00:00:00"
		todayEnd := gtime.Now().Format("Y-m-d") + " 23:59:59"
		var todayVerifiedCount int
		if todayVerifiedCount, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			Where(dao.ThMemberCoupon.Columns().Source, 3).
			Where(dao.ThMemberCoupon.Columns().MemberId, member.Id).
			Where(dao.ThMemberCoupon.Columns().ActivityId, activityId).
			Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
			Where(dao.ThMemberCoupon.Columns().State, 3).
			WhereBetween(dao.ThMemberCoupon.Columns().VerifyTime, todayStart, todayEnd).
			Count(); err != nil {
			err = gerror.Wrap(err, "查询今日核销记录失败！")
			return
		}
		if todayVerifiedCount >= perDayVerify {
			memberCodeLogger.Infof(ctx, "员工 %d 券 %d 今日核销额度已满，跳过", employee.Id, couponId)
			return
		}
	}

	// FOR UPDATE：防止并发重复领取，同时查一张未核销的券
	var pendingCoupon *entity.ThMemberCoupon
	if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
		Where(dao.ThMemberCoupon.Columns().Source, 3).
		Where(dao.ThMemberCoupon.Columns().EmployeeId, employee.Id).
		Where(dao.ThMemberCoupon.Columns().ActivityId, activityId).
		Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
		Where(dao.ThMemberCoupon.Columns().State, 2).
		LockUpdate().
		Scan(&pendingCoupon); err != nil {
		err = gerror.Wrap(err, "查询员工未核销券失败！")
		return
	}

	// 优先核销已有的未使用券
	if pendingCoupon != nil {
		if ok := s.doVerifyCoupon(ctx, tx, int64(pendingCoupon.Id), mchId, storeId, terminalId, couponMchName); ok {
			s.batchUpdateCounters(ctx, tx, couponId, storeId, 1)
			result = &verifyResult{
				memberCouponId: int64(pendingCoupon.Id),
				couponId:       couponId,
				couponMchName:  couponMchName,
				member:         member,
				employee:       employee,
			}
		}
		return
	}

	// 无未核销的券：用 Count 检查已发放总数是否达上限
	var totalIssued int
	if totalIssued, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
		Where(dao.ThMemberCoupon.Columns().Source, 3).
		Where(dao.ThMemberCoupon.Columns().EmployeeId, employee.Id).
		Where(dao.ThMemberCoupon.Columns().ActivityId, activityId).
		Where(dao.ThMemberCoupon.Columns().CouponId, couponId).
		Count(); err != nil {
		err = gerror.Wrap(err, "查询已发放券数量失败！")
		return
	}
	if availableQuantity > 0 && totalIssued >= availableQuantity {
		memberCodeLogger.Infof(ctx, "员工 %d 券 %d 已达发放上限 %d，跳过", employee.Id, couponId, availableQuantity)
		return
	}

	// 发放一张新券并核销
	couponStartTime := gtime.Now().Format("Y-m-d H:i:s")
	memberCouponId, issueErr := s.issueCouponInTx(ctx, tx, int(couponId), int(member.Id), couponStartTime, activityId, employee.Id)
	if issueErr != nil {
		err = issueErr
		return
	}
	memberCodeLogger.Infof(ctx, "员工 %d 券 %d 发放成功", employee.Id, couponId)

	if ok := s.doVerifyCoupon(ctx, tx, memberCouponId, mchId, storeId, terminalId, couponMchName); ok {
		s.batchUpdateCounters(ctx, tx, couponId, storeId, 1)
		result = &verifyResult{
			memberCouponId: memberCouponId,
			couponId:       couponId,
			couponMchName:  couponMchName,
			member:         member,
			employee:       employee,
		}
	}
	return
}

// issueCouponInTx 在外层事务内发放一张礼品券，返回 memberCouponId
func (s *sAppMember) issueCouponInTx(
	ctx context.Context,
	tx gdb.TX,
	couponId, memberId int,
	startTime string,
	activityId uint64,
	employeeId uint64,
) (memberCouponId int64, err error) {
	var couponInfo *entity.ThCoupon
	if err = dao.ThCoupon.Ctx(ctx).TX(tx).WherePri(couponId).Scan(&couponInfo); err != nil || g.IsEmpty(couponInfo) {
		err = gerror.New("礼品券不存在")
		return
	}
	if couponInfo.Status != 1 {
		err = gerror.New("该礼品券已停止发行")
		return
	}

	var activityInfo *entity.EmployeeActivity
	if activityId > 0 {
		if err = dao.EmployeeActivity.Ctx(ctx).TX(tx).WherePri(activityId).Scan(&activityInfo); err != nil || g.IsEmpty(activityInfo) {
			err = gerror.New("活动不存在")
			return
		}
	}

	// 判断员工活动是否限制当天领取（1-7 分别对应周一到周日）
	if activityInfo != nil && activityInfo.LimitWeek != "" {
		// Go Weekday: 0=周日,1=周一,...,6=周六 → 转换为 LimitWeek 约定: 周日=7, 其余不变
		goWeekday := int(gtime.Now().Weekday())
		limitWeekDay := goWeekday
		if goWeekday == 0 {
			limitWeekDay = 7
		}
		limits := gstr.Split(activityInfo.LimitWeek, ",")
		if gstr.InArray(limits, gvar.New(limitWeekDay).String()) {
			err = gerror.New("今日不可领取")
			return
		}
	}

	todayStart := startTime
	endTime := gtime.New(todayStart).Add(time.Duration(24*couponInfo.FixedTerm) * time.Hour)
	endTime = gtime.New(endTime.StartOfDay()).Add(time.Duration(23)*time.Hour + time.Duration(59)*time.Minute + time.Duration(59)*time.Second)

	if activityInfo != nil && activityInfo.CouponValidity == 1 {
		todayStart = gtime.New(activityInfo.StartTime).String()
		endTime = activityInfo.EndTime
	}

	timeDiff := int(gtime.New(todayStart).Sub(gtime.Now()).Seconds())
	couponStatus := 2
	if timeDiff > 0 {
		couponStatus = 1
	}

	couponNo := uuid.CreateOrderCode(couponInfo.CouponNoPrefix)

	if memberCouponId, err = g.Model(dao.ThMemberCoupon.Table()).Ctx(ctx).TX(tx).Safe().
		Fields(input_th.ThMemberCouponInsertFields{}).
		Data(g.MapStrAny{
			dao.ThMemberCoupon.Columns().CouponNo:   couponNo,
			dao.ThMemberCoupon.Columns().CouponId:   couponId,
			dao.ThMemberCoupon.Columns().MemberId:   memberId,
			dao.ThMemberCoupon.Columns().State:      couponStatus,
			dao.ThMemberCoupon.Columns().StartTime:  gtime.New(todayStart).Format("Y-m-d H:i:s"),
			dao.ThMemberCoupon.Columns().EndTime:    endTime,
			dao.ThMemberCoupon.Columns().Source:     3,
			dao.ThMemberCoupon.Columns().CountDown:  timeDiff,
			dao.ThMemberCoupon.Columns().ActivityId: activityId,
			dao.ThMemberCoupon.Columns().EmployeeId: employeeId,
		}).OmitEmptyData().InsertAndGetId(); err != nil {
		err = gerror.Wrap(err, "发放礼品券失败！")
		return
	}

	if _, err = dao.ThCoupon.Ctx(ctx).TX(tx).WherePri(couponId).Increment(dao.ThCoupon.Columns().Count, 1); err != nil {
		err = gerror.Wrap(err, "更新券发放数失败！")
		return
	}
	if activityId > 0 {
		if _, err = dao.EmployeeActivityCoupon.Ctx(ctx).TX(tx).
			Where(dao.EmployeeActivityCoupon.Columns().ActivityId, activityId).
			Where(dao.EmployeeActivityCoupon.Columns().CouponId, couponId).
			Increment(dao.EmployeeActivityCoupon.Columns().TotalReceived, 1); err != nil {
			err = gerror.Wrap(err, "更新活动领取数失败！")
			return
		}
	}

	// 待生效券投递延迟队列
	if timeDiff > 0 {
		if mqErr := rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeDelayedName,
			QueueName:    consts.RabbitMQQueueNameThCouponEffect,
			DataByte:     gvar.New(couponNo).Bytes(),
			Header: amqp.Table{
				"x-delay": gvar.New(timeDiff * 1000).String(),
			},
		}); mqErr != nil {
			memberCodeLogger.Warningf(ctx, "发送自动生效礼品券MQ失败: %v", mqErr)
		}
	}
	return
}

// doVerifyCoupon 在事务内核销一张券（仅 DB 操作，不做 IO）
func (s *sAppMember) doVerifyCoupon(
	ctx context.Context,
	tx gdb.TX,
	memberCouponId int64,
	mchId, storeId, terminalId int,
	couponMchName string,
) bool {
	var targetCoupon *entity.ThMemberCoupon
	if err := dao.ThMemberCoupon.Ctx(ctx).TX(tx).WherePri(memberCouponId).Scan(&targetCoupon); err != nil || g.IsEmpty(targetCoupon) {
		memberCodeLogger.Warningf(ctx, "查询待核销券 %d 失败", memberCouponId)
		return false
	}
	if targetCoupon.State != 2 {
		memberCodeLogger.Infof(ctx, "券 %d 状态非未使用(%d)，跳过核销", memberCouponId, targetCoupon.State)
		return false
	}

	if _, verifyErr := dao.ThMemberCoupon.Ctx(ctx).TX(tx).
		WherePri(memberCouponId).
		Data(input_th.ThMemberCouponVerifyFields{
			State:         3,
			VerifyTime:    gtime.Now(),
			VerifyMchId:   mchId,
			VerifyStoreId: storeId,
		}).Update(); verifyErr != nil {
		memberCodeLogger.Warningf(ctx, "核销券 %d 失败: %v", memberCouponId, verifyErr)
		return false
	}

	if _, logErr := dao.SysTerminalVerify.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SysTerminalVerify{
		TerminalId:     terminalId,
		VerifyType:     "TH_COUPON",
		MchId:          mchId,
		StoreId:        storeId,
		MemberCouponId: int(memberCouponId),
		CouponMchName:  couponMchName,
		VerifyTime:     gtime.Now(),
	}); logErr != nil {
		memberCodeLogger.Warningf(ctx, "写入核销日志失败: %v", logErr)
	}
	return true
}

// batchUpdateCounters 批量更新券使用数和门店核销数（一次 UPDATE）
func (s *sAppMember) batchUpdateCounters(ctx context.Context, tx gdb.TX, couponId uint64, storeId, count int) {
	if _, err := dao.ThCoupon.Ctx(ctx).TX(tx).Data(g.Map{
		dao.ThCoupon.Columns().UsedCount: gdb.Raw(fmt.Sprintf("used_count+%d", count)),
	}).WherePri(couponId).Update(); err != nil {
		memberCodeLogger.Warningf(ctx, "批量更新券使用数失败: %v", err)
	}
	if _, err := dao.ThMchStore.Ctx(ctx).TX(tx).Data(g.Map{
		dao.ThMchStore.Columns().VerifyNum: gdb.Raw(fmt.Sprintf("verify_num+%d", count)),
	}).WherePri(storeId).Update(); err != nil {
		memberCodeLogger.Warningf(ctx, "批量更新门店核销数失败: %v", err)
	}
}

// printMemberCouponReceipt 打印会员码核销小票（事务外调用）
func (s *sAppMember) printMemberCouponReceipt(
	ctx context.Context,
	terminalModel *input_basics.TerminalViewModel,
	thStoreTerminal *entity.ThStoreTerminal,
	couponMchName string,
	memberCouponId int64,
) {
	printTimes := 1
	if !g.IsEmpty(thStoreTerminal) {
		printTimes = int(thStoreTerminal.PrintTimes)
	}
	if printTimes <= 0 {
		return
	}

	var couponModel *input_th.ThMemberCouponViewModel
	_ = dao.ThMemberCoupon.Ctx(ctx).WithAll().WherePri(memberCouponId).Scan(&couponModel)

	couponName, storeName, storeAddress, memberNo := "", "", "", ""
	if !g.IsEmpty(couponModel) {
		if couponModel.ThCoupon != nil {
			couponName = couponModel.ThCoupon.NameLanguage.Content
		}
		if couponModel.VerifyStore != nil {
			storeName = couponModel.VerifyStore.NameLanguage.Content
			storeAddress = couponModel.VerifyStore.DetailAddress
		}
		if couponModel.Member != nil {
			memberNo = couponModel.Member.MemberNo
		}
	}

	printContent := "<IMG30></IMG>"
	printContent += "<BR><BR><CB>引換券<BR><BR>"
	printContent += "<C><HB>---" + gtime.Now().Format("Y-m-d") + "---<BR>"
	printContent += "<L><N>********************************<BR><BR>"
	printContent += "<L><HB>" + couponMchName + "*1<BR><BR>"
	printContent += "<L><N>--------------------------------<BR>"
	printContent += "<L><N>引換券:" + couponName + "<BR>"
	printContent += "顧客情報：" + memberNo + "<BR>"
	printContent += "使用時間：" + gtime.Now().Format("Y-m-d H:i:s") + "<BR>"
	printContent += "場所：" + storeName + "<BR>"
	printContent += "<N>〒" + storeAddress + "<BR><BR>"
	printContent += "<C><B>**終了**<BR><BR>"

	if err := service.BasicsTerminal().Printer(ctx, &input_basics.PrinterInp{
		Sn:           terminalModel.Sn,
		PrintContent: printContent,
		PrintTimes:   printTimes,
	}); err != nil {
		memberCodeLogger.Warningf(ctx, "打印失败: %v", err)
	}
}
