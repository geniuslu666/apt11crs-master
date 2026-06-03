package logic_th

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm"
	"APT/internal/library/hgorm/handler"
	"APT/internal/library/hgorm/hook"
	"APT/internal/library/ws"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_th"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/encrypt"
	"APT/utility/excel"
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sThMemberCoupon struct{}

func NewThMemberCoupon() *sThMemberCoupon {
	return &sThMemberCoupon{}
}

func init() {
	service.RegisterThMemberCoupon(NewThMemberCoupon())
}

func (s *sThMemberCoupon) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ThMemberCoupon.Ctx(ctx), option...)
}

func (s *sThMemberCoupon) List(ctx context.Context, in *input_th.ThMemberCouponListInp) (list []*input_th.ThMemberCouponListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.FieldsPrefix(dao.ThMemberCoupon.Table(), input_th.ThMemberCouponListModel{})

	mod = mod.Fields(hgorm.JoinFields(ctx, input_th.ThMemberCouponListModel{}, &dao.PmsMember, "pmsMember"))
	mod = mod.LeftJoinOnFields(dao.PmsMember.Table(), dao.ThMemberCoupon.Columns().MemberId, "=", dao.PmsMember.Columns().Id)

	if !g.IsEmpty(in.CouponNo) {
		mod = mod.WhereLike(dao.ThMemberCoupon.Columns().CouponNo, "%"+in.CouponNo+"%")
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.ThMemberCoupon.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.CouponName) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.CouponName, "coupon_name")
		if err == nil {
			couponIds, _ := service.ThCoupon().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.ThMemberCoupon.Columns().CouponId, couponIds)
		}
	}

	if in.CouponId > 0 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().CouponId, in.CouponId)
	}

	if in.State > 0 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().State, in.State)
	}

	if len(in.VerifyTime) == 2 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().State, 2)
		mod = mod.WhereBetween(dao.ThMemberCoupon.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	if !g.IsEmpty(in.StoreName) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.StoreName, "store_name")
		if err == nil {
			storeIds, _ := service.ThMchStore().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.ThMemberCoupon.Columns().VerifyStoreId, storeIds)
		}
	}

	if !g.IsEmpty(in.VerifyMchId) {
		mod = mod.Where(dao.ThMemberCoupon.Columns().VerifyMchId, in.VerifyMchId)
	}

	if !g.IsEmpty(in.VerifyStoreId) {
		mod = mod.Where(dao.ThMemberCoupon.Columns().VerifyStoreId, in.VerifyStoreId)
	}

	if in.EmployeeId > 0 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().EmployeeId, in.EmployeeId)
	}

	if in.ActivityId > 0 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().ActivityId, in.ActivityId)
	}

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	if in.IsVerifyTimeDesc {
		mod = mod.OrderDesc(dao.ThMemberCoupon.Columns().VerifyTime)
	} else {
		mod = mod.OrderDesc(dao.ThMemberCoupon.Columns().Id)
	}

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return
	}

	for _, v := range list {
		if !g.IsEmpty(v.ThCouponMch) {
			for _, item := range v.ThCouponMch {
				if gvar.New(item.MchId).Int() == v.VerifyMchId {
					v.ThCouponMchName = item.Name
				}
			}
		}
	}

	return
}

func (s *sThMemberCoupon) RefreshCode(ctx context.Context, in *input_th.ThMemberCouponRefreshCodeInp) (code string, verifyStatus int, err error) {

	var thMemberCouponInfo *struct {
		Id       int
		CouponNo string
		MemberId int
		State    int
	}
	if err = s.Model(ctx).WherePri(in.Id).Scan(&thMemberCouponInfo); err != nil {
		err = gerror.Wrap(err, "获取会员礼品券信息失败，请稍后重试！")
		return
	}
	if thMemberCouponInfo.MemberId != in.MemberId {
		err = gerror.New("会员信息不匹配")
		return
	}

	// 状态 1待生效 2未使用 3已核销 4已过期  5已失效  6已回收
	verifyStatus = thMemberCouponInfo.State

	// 生成券码，采用aes加密，将couponNo|会员ID|时间戳进行加密
	timestamp := gtime.Now().Unix()
	plainText := fmt.Sprintf("%s|%d|%d", thMemberCouponInfo.CouponNo, thMemberCouponInfo.MemberId, timestamp)

	// 使用AES加密生成动态券码
	code = encrypt.MustAesECBEncryptToString(plainText, string(consts.RequestEncryptKey))

	return
}

func (s *sThMemberCoupon) Verify(ctx context.Context, in *input_th.ThMemberCouponVerifyInp) (err error) {

	Logger := g.Log().Path("logs/Verify")
	Logger.Info(ctx, "--------礼品券核销进入----------")
	Logger.Info(ctx, gjson.New(in))

	if g.IsEmpty(in.Sn) {
		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
	}

	if g.IsEmpty(in.Val) {
		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
	}

	var terminalModel *input_basics.TerminalViewModel
	var thMemberCouponInfo *input_th.ThMemberCouponViewModel
	var thMemberCouponModel *input_th.ThMemberCouponViewModel
	var thCouponMch *entity.ThCouponMch
	var thStoreTerminal *entity.ThStoreTerminal

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 终端信息
		if err = dao.SysTerminal.Ctx(ctx).TX(tx).
			Where(dao.SysTerminal.Columns().Sn, in.Sn).
			WithAll().
			Scan(&terminalModel); err != nil {
			err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
			return
		}

		// 验签
		//sign := util2.Sign(in.Sn + gvar.New(in.Timestamp).String() + in.Val + terminalModel.BrandInfo.ClientSecret)

		//if sign != in.Sign {
		//	err = gerror.New("签名不正确")
		//	return
		//}

		if terminalModel.StoreId <= 0 {
			err = gerror.New("不可核销礼品券")
			return
		}

		// 解密动态券码并验证时间戳
		var couponNo string
		var memberId int
		var timestamp int64
		var isDynamicCode bool = false

		// 尝试解密动态券码
		// 首先尝试base64解码
		encryptedData, base64Err := base64.StdEncoding.DecodeString(in.Val)
		if base64Err != nil {
			// 如果base64解码失败，可能是旧的券码格式，直接使用原券号查询
			couponNo = in.Val
			Logger.Info(ctx, "base64解码失败")
		} else {
			// base64解码成功，尝试AES解密
			decryptedText, decryptErr := encrypt.AesECBDecrypt(encryptedData, consts.RequestEncryptKey)
			Logger.Info(ctx, "--------核销解码----------")
			Logger.Info(ctx, decryptedText)
			if decryptErr != nil {
				// 如果解密失败，可能是旧的券码格式，直接使用原券号查询
				couponNo = in.Val
				Logger.Info(ctx, "base64解码成功，AES解密失败")
			} else {
				// 解密成功，解析格式：couponNo|memberId|timestamp
				parts := strings.Split(string(decryptedText), "|")
				if len(parts) != 3 {
					err = gerror.New("动态券码格式不正确！")
					return
				}

				couponNo = parts[0]
				memberId, err = strconv.Atoi(parts[1])
				if err != nil {
					err = gerror.New("动态券码会员ID格式不正确！")
					return
				}

				timestamp, err = strconv.ParseInt(parts[2], 10, 64)
				if err != nil {
					err = gerror.New("动态券码时间戳格式不正确！")
					return
				}

				// 验证时间戳是否在30秒内
				currentTime := gtime.Now().Unix()
				if currentTime-timestamp > 30 {
					err = gerror.New("动态券码已过期，请重新生成！")
					return
				}

				isDynamicCode = true
			}
		}

		Logger.Info(ctx, "--------CouponNo----------")
		Logger.Info(ctx, couponNo)

		if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).WithAll().
			Where(dao.ThMemberCoupon.Columns().CouponNo, couponNo).
			Scan(&thMemberCouponInfo); err != nil {
			err = gerror.New("礼品券信息不存在！")
			return
		}

		// 如果使用的是动态券码，还需要验证会员ID是否匹配
		if isDynamicCode && thMemberCouponInfo.MemberId != memberId {
			err = gerror.New("动态券码会员信息不匹配！")
			return
		}

		if thMemberCouponInfo.ThCoupon.UseStatus != 1 {
			err = gerror.New("礼品券已停止使用！")
			return
		}

		if thMemberCouponInfo.State != 2 {
			err = gerror.New("礼品券状态不正确！")
			return
		}

		// 如果是员工活动领取的券，限制同个活动同个券一天只能核销一张
		if thMemberCouponInfo.Source == 3 && thMemberCouponInfo.ActivityId > 0 {
			var ActivityCoupon *entity.EmployeeActivityCoupon
			if err = dao.EmployeeActivityCoupon.Ctx(ctx).
				Where(dao.EmployeeActivityCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.EmployeeActivityCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Scan(&ActivityCoupon); err != nil {
				return
			}

			var todayVerifiedCount int
			todayStart := gtime.Now().Format("Y-m-d") + " 00:00:00"
			todayEnd := gtime.Now().Format("Y-m-d") + " 23:59:59"

			if todayVerifiedCount, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
				Where(dao.ThMemberCoupon.Columns().Source, 3).
				Where(dao.ThMemberCoupon.Columns().MemberId, thMemberCouponInfo.MemberId).
				Where(dao.ThMemberCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.ThMemberCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Where(dao.ThMemberCoupon.Columns().State, 3).
				WhereBetween(dao.ThMemberCoupon.Columns().VerifyTime, todayStart, todayEnd).
				Count(); err != nil {
				err = gerror.Wrap(err, "查询今日核销记录失败！")
				return
			}

			if todayVerifiedCount >= ActivityCoupon.PerDayVerify {
				err = gerror.New("今日不可再进行核销！")
				return
			}
		}

		if err = dao.ThCouponMch.Ctx(ctx).TX(tx).WithAll().
			Where(dao.ThCouponMch.Columns().CouponId, thMemberCouponInfo.CouponId).
			Where(dao.ThCouponMch.Columns().MchId, terminalModel.StoreInfo.MchId).
			Scan(&thCouponMch); err != nil {
			err = gerror.New("礼品券商户关联信息不存在！")
			return
		}

		if g.IsEmpty(thCouponMch) {
			err = gerror.New("您的商户不能核销该礼品券")
			return
		}

		// 获取门店终端信息
		if err = dao.ThStoreTerminal.Ctx(ctx).TX(tx).
			Where(dao.ThStoreTerminal.Columns().StoreId, terminalModel.StoreId).
			Where(dao.ThStoreTerminal.Columns().TerminalId, terminalModel.Id).
			Scan(&thStoreTerminal); err != nil {
			err = gerror.New("门店终端信息不存在！")
			return
		}

		// 核销礼品券
		if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			Where(dao.ThMemberCoupon.Columns().CouponNo, couponNo).Data(input_th.ThMemberCouponVerifyFields{
			State:         3,
			VerifyTime:    gtime.Now(),
			VerifyMchId:   terminalModel.StoreInfo.MchId,
			VerifyStoreId: terminalModel.StoreInfo.Id,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 获取修改后的礼品券信息
		if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			WithAll().
			Where(dao.ThMemberCoupon.Columns().CouponNo, couponNo).
			Scan(&thMemberCouponModel); err != nil {
			err = gerror.New("礼品券信息不存在！")
			return
		}

		if g.IsEmpty(thMemberCouponModel) {
			err = gerror.Wrap(err, "礼品券信息不存在！")
			return
		}

		// 更新门店核销数
		if _, err = dao.ThMchStore.Ctx(ctx).Data(g.Map{
			dao.ThMchStore.Columns().VerifyNum: gdb.Raw(fmt.Sprintf("verify_num+%d", 1)),
		}).WherePri(terminalModel.StoreInfo.Id).Update(); err != nil {
			err = gerror.Wrap(err, "修改门店核销数，请稍后重试！")
			return
		}

		// 更新礼品券使用数
		if _, err = dao.ThCoupon.Ctx(ctx).Data(g.Map{
			dao.ThCoupon.Columns().UsedCount: gdb.Raw(fmt.Sprintf("used_count+%d", 1)),
		}).WherePri(thMemberCouponInfo.CouponId).Update(); err != nil {
			err = gerror.Wrap(err, "修改券使用数失败，请稍后重试！")
			return
		}

		// 写入终端核销日志
		if _, err = dao.SysTerminalVerify.Ctx(ctx).OmitEmptyData().Insert(&entity.SysTerminalVerify{
			TerminalId:     terminalModel.Id,
			VerifyType:     "TH_COUPON",
			MchId:          gvar.New(terminalModel.StoreInfo.MchId).Int(),
			StoreId:        terminalModel.StoreInfo.Id,
			MemberCouponId: thMemberCouponModel.Id,
			CouponMchName:  thCouponMch.Name,
			VerifyMemberId: gvar.New(thMemberCouponModel.MemberId).Int(),
			VerifyTime:     gtime.Now(),
		}); err != nil {
			return err
		}

		return
	})

	if err != nil {
		// 核销失败
		Logger.Info(ctx, "--------核销失败----------")
		Logger.Info(ctx, err)

		if g.IsEmpty(thMemberCouponInfo) {
			ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
				"data": "OK",
			})
			return
		}

		// 发送消息
		_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
			MemberId: thMemberCouponInfo.MemberId,
			Code:     -1,
			Event:    "VERIFY",
			Message:  "核销失败",
		})

		ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
			"data": "OK",
		})
		return
	}

	Logger.Info(ctx, "--------核销成功----------")

	printTimes := 1
	if !g.IsEmpty(thStoreTerminal) {
		printTimes = int(thStoreTerminal.PrintTimes)
	}

	// 打印
	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"
	printContent = printContent + "<CB>引換券<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.Now().Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><HB>" + thCouponMch.Name + "*1<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>--------------------------------"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>引換券:" + thMemberCouponModel.ThCoupon.NameLanguage.Content + "<BR>"
	printContent = printContent + "顧客情報：" + thMemberCouponModel.Member.MemberNo + "<BR>"
	printContent = printContent + "使用時間：" + gtime.Now().Format("Y-m-d H:i:s") + "<BR>"
	printContent = printContent + "場所：" + thMemberCouponModel.VerifyStore.NameLanguage.Content + "<BR>"
	printContent = printContent + "<N>〒" + thMemberCouponModel.VerifyStore.DetailAddress + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR><BR>"

	Logger.Info(ctx, "--------打印内容----------")
	Logger.Info(ctx, printContent)

	if printTimes > 0 {
		Logger.Info(ctx, "--------开始打印----------")
		err = service.BasicsTerminal().Printer(ctx, &input_basics.PrinterInp{
			Sn:           terminalModel.Sn,
			PrintContent: printContent,
			PrintTimes:   printTimes,
		})
		if err != nil {
			return
		}
	}

	// 发送消息
	_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
		MemberId: thMemberCouponInfo.MemberId,
		Code:     200,
		Event:    "VERIFY",
		Message:  "核销成功",
	})

	ghttp.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"data": "OK",
	})
	return
}

func (s *sThMemberCoupon) AppList(ctx context.Context, in *input_th.ThMemberCouponAppListInp) (list []*input_th.ThMemberCouponAppListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll().Hook(hook.PmsFindLanguageValueHook)

	mod = mod.Fields(input_th.ThMemberCouponAppListModel{})

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.ThMemberCoupon.Columns().MemberId, in.MemberId)
	}

	if in.State > 0 {
		if in.State == 3 {
			mod = mod.WhereIn(dao.ThMemberCoupon.Columns().State, []int{4, 5})
		} else {
			mod = mod.Where(dao.ThMemberCoupon.Columns().State, in.State)
		}
	} else {
		mod = mod.WhereNot(dao.ThMemberCoupon.Columns().State, 6)
	}

	if in.Source > 0 {
		mod = mod.Where(dao.ThMemberCoupon.Columns().Source, in.Source)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.ThMemberCoupon.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (s *sThMemberCoupon) AppView(ctx context.Context, in *input_th.ThMemberCouponAppViewInp) (res *input_th.ThMemberCouponAppViewModel, err error) {
	if err = s.Model(ctx).Hook(hook.PmsFindLanguageValueHook).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取礼品券信息失败，请稍后重试！")
		return
	}

	// 生成券码，采用aes加密，将couponNo|会员ID|时间戳进行加密
	timestamp := gtime.Now().Unix()
	plainText := fmt.Sprintf("%s|%d|%d", res.CouponNo, res.MemberId, timestamp)

	// 使用AES加密生成动态券码
	res.Code = encrypt.MustAesECBEncryptToString(plainText, string(consts.RequestEncryptKey))

	if res.ThCoupon.NeedReservation == 1 && !g.IsEmpty(res.ThCoupon.ReservationRestaurantIds) {
		if err = dao.FoodRestaurant.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).Fields("id,name").
			WhereIn(dao.FoodRestaurant.Columns().Id, strings.Split(res.ThCoupon.ReservationRestaurantIds, ",")).Scan(&res.ThCoupon.RestaurantList); err != nil {
			err = gerror.Wrap(err, "获取餐厅列表信息失败，请稍后重试！")
			return
		}
	}

	return
}

func (s *sThMemberCoupon) CouponEffect(ctx context.Context, CouponOn string) (err error) {
	var (
		tx             gdb.TX
		ThMemberCoupon entity.ThMemberCoupon
	)
	if tx, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.ThMemberCoupon.Columns().CouponNo: CouponOn,
		dao.ThMemberCoupon.Columns().State:    1,
	}).Scan(&ThMemberCoupon); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(ThMemberCoupon) {
		err = gerror.New("该礼品券无需处理")
		return
	}

	// 更新当前礼品券为待使用
	if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.ThMemberCoupon.Columns().CouponNo: CouponOn,
		dao.ThMemberCoupon.Columns().State:    1,
	}).Data(g.MapStrAny{
		dao.ThMemberCoupon.Columns().State: 2,
	}).Update(); err != nil {
		return
	}

	return
}

func (s *sThMemberCoupon) InvalidMemberCoupon(ctx context.Context, in *input_th.ThInvalidMemberCouponInp) (err error) {
	var (
		memberCouponInfo *entity.ThMemberCoupon
	)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if in.SourceOrderId > 0 {
			if err = s.Model(ctx).Where(dao.ThMemberCoupon.Columns().SourceOrderId, in.SourceOrderId).Scan(&memberCouponInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}
			// 如果礼品券不存在，则返回
			if g.IsEmpty(memberCouponInfo) {
				g.Log().Info(ctx, "礼品券不存在")
				return
			}
		}
		if in.Id > 0 {
			if err = s.Model(ctx).WherePri(in.Id).Scan(&memberCouponInfo); err != nil {
				err = gerror.Wrap(err, consts.ErrorORM)
				return
			}
		}

		if g.IsEmpty(memberCouponInfo) {
			err = gerror.New("礼品券不存在")
			return
		}

		if memberCouponInfo.State == 3 || memberCouponInfo.State == 4 || memberCouponInfo.State == 5 {
			err = gerror.New("该礼品券状态不正确")
			return
		}

		// 更新当前礼品券为已失效
		if in.SourceOrderId > 0 {
			if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.ThMemberCoupon.Columns().SourceOrderId: in.SourceOrderId,
			}).Data(g.MapStrAny{
				dao.ThMemberCoupon.Columns().State:       5,
				dao.ThMemberCoupon.Columns().InvalidTime: gtime.Now(),
			}).Update(); err != nil {
				return
			}
		}
		if in.Id > 0 {
			if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.ThMemberCoupon.Columns().Id: in.Id,
			}).Data(g.MapStrAny{
				dao.ThMemberCoupon.Columns().State:       5,
				dao.ThMemberCoupon.Columns().InvalidTime: gtime.Now(),
			}).Update(); err != nil {
				return
			}
		}

		return
	}); err != nil {
		return
	}

	return
}

func (s *sThMemberCoupon) Export(ctx context.Context, in *input_th.ThMemberCouponListInp) (err error) {
	var (
		list       []*input_th.ThMemberCouponListModel
		exportList []*input_th.ThMemberCouponExportModel
		tags       []string
	)
	in.Pagination = false

	if list, _, err = s.List(ctx, in); err != nil {
		return
	}

	for _, v := range list {
		var state string
		switch v.State {
		case 1:
			// 待生效
			state = "利用不可"
			break
		case 2:
			// 未使用
			state = "未使用"
			break
		case 3:
			// 已核销
			state = "使用済み"
			break
		case 4:
			// 已过期
			state = "期限切れ"
			break
		case 5:
			// 已失效
			state = "無効"
			break
		case 6:
			// 已回收
			state = "回収済み"
			break
		}

		CouponName := ""
		if !g.IsEmpty(v.ThCoupon) {
			CouponName = v.ThCoupon.CouponName
		}

		IdentityName := ""
		if !g.IsEmpty(v.ThCoupon) {
			IdentityName = v.ThCoupon.IdentityName
		}

		MemberNo := ""
		FullName := ""
		if !g.IsEmpty(v.Member) {
			MemberNo = v.Member.MemberNo
			FullName = v.Member.FullName
		}
		MchName := ""
		if !g.IsEmpty(v.VerifyMch) {
			MchName = v.VerifyMch.Name
		}
		StoreName := ""
		if !g.IsEmpty(v.VerifyStore) {
			StoreName = v.VerifyStore.StoreName
		}

		EmployeeActivityName := ""
		if !g.IsEmpty(v.EmployeeActivity) {
			EmployeeActivityName = v.EmployeeActivity.Name
		}

		IndexActivityTitle := ""
		if !g.IsEmpty(v.IndexActivity) {
			IndexActivityTitle = v.IndexActivity.Title
		}

		EmployeeName := ""
		if !g.IsEmpty(v.EmployeeInfo) {
			EmployeeName = v.EmployeeInfo.Name
		}

		exportList = append(exportList, &input_th.ThMemberCouponExportModel{
			Id:                   v.Id,
			CouponNo:             v.CouponNo,
			CouponName:           CouponName,
			IdentityName:         IdentityName,
			MemberNo:             MemberNo,
			MemberName:           FullName,
			State:                state,
			VerifyTime:           v.VerifyTime,
			VerifyMchName:        MchName,
			VerifyStoreName:      StoreName,
			ThCouponMchName:      v.ThCouponMchName,
			StartTime:            v.StartTime,
			EndTime:              v.EndTime,
			CreatedAt:            v.CreateAt,
			EmployeeActivityName: EmployeeActivityName,
			EmployeeName:         EmployeeName,
			IndexActivityTitle:   IndexActivityTitle,
		})
	}

	if tags, err = convert.GetEntityDescTags(input_th.ThMemberCouponExportModel{}); err != nil {
		return
	}

	var (
		fileName  = "导出会员礼品券-" + gctx.CtxId(ctx)
		sheetName = "会员礼品券"
		exports   []input_th.ThMemberCouponExportModel
	)

	if err = gconv.Scan(exportList, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

func (s *sThMemberCoupon) View(ctx context.Context, in *input_th.ThMemberCouponViewInp) (res *input_th.ThMemberCouponAdminViewModel, err error) {
	if err = s.Model(ctx).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取会员礼品券信息失败，请稍后重试！")
		return
	}

	return
}

func (s *sThMemberCoupon) Recycle(ctx context.Context, in *input_th.ThMemberCouponRecycleInp) (err error) {
	// 查询礼品券
	var coupon *entity.ThMemberCoupon
	if err = s.Model(ctx).WherePri(in.Id).Scan(&coupon); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.ThMemberCoupon.Columns().State:        6,
		dao.ThMemberCoupon.Columns().OperatorId:   int(contexts.GetUserId(ctx)),
		dao.ThMemberCoupon.Columns().RecoveryTime: gtime.Now(),
	}).Update(); err != nil {
		err = gerror.Wrap(err, "礼品券回收失败，请稍后重试！")
		return
	}

	// 回收后已发放数量要减1
	if _, err = dao.ThCoupon.Ctx(ctx).Where(dao.ThCoupon.Columns().Id, coupon.CouponId).Decrement(dao.ThCoupon.Columns().Count, 1); err != nil {
		err = gerror.Wrap(err, "操作失败，请稍后重试！")
		return
	}
	return
}

// ManualVerify 会员手动核销
func (s *sThMemberCoupon) ManualVerify(ctx context.Context, in *input_th.ThMemberCouponManualVerifyInp) (err error) {

	Logger := g.Log().Path("logs/Verify")
	Logger.Info(ctx, "--------手动礼品券核销进入----------"+gvar.New(contexts.GetUserId(ctx)).String())
	Logger.Info(ctx, gjson.New(in))

	if g.IsEmpty(in.StoreId) {
		err = gerror.New("核销门店不存在")
		return
	}

	if g.IsEmpty(in.Id) {
		err = gerror.New("会员礼品券ID不存在")
		return
	}
	// 不传核销时间时，默认当前时间
	if g.IsEmpty(in.VerifyTime) {
		in.VerifyTime = gtime.Now()
	}

	var terminalModel *input_basics.TerminalViewModel
	var thMemberCouponInfo *input_th.ThMemberCouponViewModel
	var thMemberCouponModel *input_th.ThMemberCouponViewModel
	var thCouponMch *entity.ThCouponMch
	var thStoreTerminal *entity.ThStoreTerminal

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 终端信息
		if err = dao.SysTerminal.Ctx(ctx).TX(tx).
			Where(dao.SysTerminal.Columns().StoreId, in.StoreId).
			WithAll().
			Scan(&terminalModel); err != nil {
			err = gerror.Wrap(err, "获取终端信息失败，请稍后重试！")
			return
		}

		Logger.Info(ctx, "--------ThMemberCouponId----------")
		Logger.Info(ctx, in.Id)

		if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).WithAll().
			Where(dao.ThMemberCoupon.Columns().Id, in.Id).
			Scan(&thMemberCouponInfo); err != nil {
			err = gerror.New("礼品券信息不存在！")
			return
		}

		if thMemberCouponInfo.ThCoupon.UseStatus != 1 {
			err = gerror.New("礼品券已停止使用！")
			return
		}

		if thMemberCouponInfo.State != 2 {
			err = gerror.New("礼品券状态不正确！")
			return
		}

		// 如果是员工活动领取的券，限制同个活动同个券一天只能核销一张
		if thMemberCouponInfo.Source == 3 && thMemberCouponInfo.ActivityId > 0 {
			var ActivityCoupon *entity.EmployeeActivityCoupon
			if err = dao.EmployeeActivityCoupon.Ctx(ctx).
				Where(dao.EmployeeActivityCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.EmployeeActivityCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Scan(&ActivityCoupon); err != nil {
				return
			}

			var todayVerifiedCount int
			todayStart := gtime.New(in.VerifyTime).Format("Y-m-d") + " 00:00:00"
			todayEnd := gtime.New(in.VerifyTime).Format("Y-m-d") + " 23:59:59"

			if todayVerifiedCount, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
				Where(dao.ThMemberCoupon.Columns().Source, 3).
				Where(dao.ThMemberCoupon.Columns().MemberId, thMemberCouponInfo.MemberId).
				Where(dao.ThMemberCoupon.Columns().ActivityId, thMemberCouponInfo.ActivityId).
				Where(dao.ThMemberCoupon.Columns().CouponId, thMemberCouponInfo.CouponId).
				Where(dao.ThMemberCoupon.Columns().State, 3).
				WhereBetween(dao.ThMemberCoupon.Columns().VerifyTime, todayStart, todayEnd).
				Count(); err != nil {
				err = gerror.Wrap(err, "查询当日核销记录失败！")
				return
			}

			if todayVerifiedCount >= ActivityCoupon.PerDayVerify {
				err = gerror.New("该核销时间不可再进行核销！")
				return
			}
		}

		if err = dao.ThCouponMch.Ctx(ctx).TX(tx).WithAll().
			Where(dao.ThCouponMch.Columns().CouponId, thMemberCouponInfo.CouponId).
			Where(dao.ThCouponMch.Columns().MchId, terminalModel.StoreInfo.MchId).
			Scan(&thCouponMch); err != nil {
			err = gerror.New("礼品券商户关联信息不存在！")
			return
		}

		if g.IsEmpty(thCouponMch) {
			err = gerror.New("您的商户不能核销该礼品券")
			return
		}

		// 获取门店终端信息
		if err = dao.ThStoreTerminal.Ctx(ctx).TX(tx).
			Where(dao.ThStoreTerminal.Columns().StoreId, terminalModel.StoreId).
			Where(dao.ThStoreTerminal.Columns().TerminalId, terminalModel.Id).
			Scan(&thStoreTerminal); err != nil {
			err = gerror.New("门店终端信息不存在！")
			return
		}

		// 核销礼品券
		if _, err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			Where(dao.ThMemberCoupon.Columns().Id, in.Id).Data(input_th.ThMemberCouponVerifyFields{
			State:         3,
			VerifyTime:    in.VerifyTime,
			VerifyMchId:   terminalModel.StoreInfo.MchId,
			VerifyStoreId: terminalModel.StoreInfo.Id,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 获取修改后的礼品券信息
		if err = dao.ThMemberCoupon.Ctx(ctx).TX(tx).
			WithAll().
			Where(dao.ThMemberCoupon.Columns().Id, in.Id).
			Scan(&thMemberCouponModel); err != nil {
			err = gerror.New("礼品券信息不存在！")
			return
		}

		if g.IsEmpty(thMemberCouponModel) {
			err = gerror.Wrap(err, "礼品券信息不存在！")
			return
		}

		// 更新门店核销数
		if _, err = dao.ThMchStore.Ctx(ctx).Data(g.Map{
			dao.ThMchStore.Columns().VerifyNum: gdb.Raw(fmt.Sprintf("verify_num+%d", 1)),
		}).WherePri(terminalModel.StoreInfo.Id).Update(); err != nil {
			err = gerror.Wrap(err, "修改门店核销数，请稍后重试！")
			return
		}

		// 更新礼品券使用数
		if _, err = dao.ThCoupon.Ctx(ctx).Data(g.Map{
			dao.ThCoupon.Columns().UsedCount: gdb.Raw(fmt.Sprintf("used_count+%d", 1)),
		}).WherePri(thMemberCouponInfo.CouponId).Update(); err != nil {
			err = gerror.Wrap(err, "修改券使用数失败，请稍后重试！")
			return
		}

		// 写入终端核销日志
		if _, err = dao.SysTerminalVerify.Ctx(ctx).OmitEmptyData().Insert(&entity.SysTerminalVerify{
			TerminalId:     terminalModel.Id,
			VerifyType:     "TH_COUPON",
			MchId:          gvar.New(terminalModel.StoreInfo.MchId).Int(),
			StoreId:        terminalModel.StoreInfo.Id,
			MemberCouponId: thMemberCouponModel.Id,
			CouponMchName:  thCouponMch.Name,
			VerifyMemberId: gvar.New(thMemberCouponModel.MemberId).Int(),
			VerifyTime:     in.VerifyTime,
		}); err != nil {
			return err
		}

		return
	})

	if err != nil {
		// 核销失败
		Logger.Info(ctx, "--------核销失败----------")
		Logger.Info(ctx, err)
		return
	}

	Logger.Info(ctx, "--------核销成功----------")

	printTimes := 1
	if !g.IsEmpty(thStoreTerminal) {
		printTimes = int(thStoreTerminal.PrintTimes)
	}

	// 打印
	printContent := "<IMG30></IMG>"
	printContent = printContent + "<BR><BR>"
	printContent = printContent + "<CB>引換券<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><HB>---" + gtime.New(in.VerifyTime).Format("Y-m-d") + "---<BR>"
	printContent = printContent + "<L><N>********************************<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><HB>" + thCouponMch.Name + "*1<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>--------------------------------"
	printContent = printContent + "<BR>"
	printContent = printContent + "<L><N>引換券:" + thMemberCouponModel.ThCoupon.NameLanguage.Content + "<BR>"
	printContent = printContent + "顧客情報：" + thMemberCouponModel.Member.MemberNo + "<BR>"
	printContent = printContent + "使用時間：" + gtime.Now().Format("Y-m-d H:i:s") + "<BR>"
	printContent = printContent + "場所：" + thMemberCouponModel.VerifyStore.NameLanguage.Content + "<BR>"
	printContent = printContent + "<N>〒" + thMemberCouponModel.VerifyStore.DetailAddress + "<BR>"
	printContent = printContent + "<BR>"
	printContent = printContent + "<C><B>**終了**"
	printContent = printContent + "<BR>"

	Logger.Info(ctx, "--------打印内容----------")
	Logger.Info(ctx, printContent)

	// 开始创建打印任务
	if printTimes > 0 {
		err = service.BasicsTerminal().Printer(ctx, &input_basics.PrinterInp{
			Sn:           terminalModel.Sn,
			PrintContent: printContent,
			PrintTimes:   printTimes,
		})
		if err != nil {
			return err
		}
	}
	return
}
