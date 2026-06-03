package logic_travel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/library/ws"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"APT/utility/encrypt"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sTravelVerifyRecord struct{}

func NewTravelVerifyRecord() *sTravelVerifyRecord {
	return &sTravelVerifyRecord{}
}

func init() {
	service.RegisterTravelVerifyRecord(NewTravelVerifyRecord())
}

// Model 核销记录ORM模型
func (s *sTravelVerifyRecord) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TravelVerifyRecord.Ctx(ctx), option...)
}

// List 获取核销记录列表
func (s *sTravelVerifyRecord) List(ctx context.Context, in *input_travel.TravelVerifyRecordListInp) (list []*input_travel.TravelVerifyRecordListModel, totalCount int, err error) {
	mod := s.Model(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.TravelVerifyRecord.Table(), input_travel.TravelVerifyRecordListModel{})
	mod = mod.Fields(fmt.Sprintf("`%s`.`%s` as `pmsMemberMemberNo`", dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))

	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.TravelVerifyRecord.Table(), dao.TravelVerifyRecord.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.TravelVerifyRecord.Columns().OrderSn, "%"+in.OrderSn+"%")
	}
	if in.VerifyStaffId > 0 {
		mod = mod.Where(dao.TravelVerifyRecord.Columns().VerifyStaffId, in.VerifyStaffId)
	}

	if !g.IsEmpty(in.ProductName) {
		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.ProductName, "title")
		if err == nil {
			productsIds, _ := service.TravelProduct().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.TravelVerifyRecord.Columns().ProductId, productsIds)
		}
	}

	if len(in.BookDate) == 2 {
		mod = mod.WhereBetween(dao.TravelVerifyRecord.Columns().BookDate, in.BookDate[0], in.BookDate[1])
	}
	if len(in.VerifyTime) == 2 {
		mod = mod.WhereBetween(dao.TravelVerifyRecord.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	mod = mod.OrderDesc(dao.TravelVerifyRecord.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取核销记录列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取核销记录列表失败，请稍后重试！")
	}
	// 对于已删除的关联数据，手动加载
	for _, item := range list {
		// 加载已删除的产品
		if item.ProductId > 0 && item.ProductInfo == nil {
			var product *entity.TravelProduct
			if err = dao.TravelProduct.Ctx(ctx).Unscoped().Where(dao.TravelProduct.Columns().Id, item.ProductId).Scan(&product); err == nil && product != nil {
				item.ProductInfo = &struct {
					gmeta.Meta `orm:"table:hg_travel_product"`
					*entity.TravelProduct
				}{TravelProduct: product}
			}
		}
		// 加载已删除的核销人员
		if item.VerifyStaffId > 0 && item.VerifyStaffInfo == nil {
			var staff *entity.TravelVerifyStaff
			if err = dao.TravelVerifyStaff.Ctx(ctx).Unscoped().Where(dao.TravelVerifyStaff.Columns().Id, item.VerifyStaffId).Scan(&staff); err == nil && staff != nil {
				item.VerifyStaffInfo = &struct {
					gmeta.Meta `orm:"table:hg_travel_verify_staff"`
					*entity.TravelVerifyStaff
				}{TravelVerifyStaff: staff}
			}
		}
	}
	return
}

// VerifyList 获取核销端核销记录列表
func (s *sTravelVerifyRecord) VerifyList(ctx context.Context, in *input_travel.VerifyListInp) (list []*input_travel.VerifyListModel, totalCount int, err error) {
	mod := s.Model(ctx).WithAll()

	if in.VerifyStaffId > 0 {
		mod = mod.Where(dao.TravelVerifyRecord.Columns().VerifyStaffId, in.VerifyStaffId)
	}

	mod = mod.OrderDesc(dao.TravelVerifyRecord.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取核销记录列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取核销记录列表失败，请稍后重试！")
	}
	// 对于已删除的关联数据，手动加载
	for _, item := range list {
		// 加载已删除的产品
		if item.ProductId > 0 && item.ProductInfo == nil {
			var product *entity.TravelProduct
			if err = dao.TravelProduct.Ctx(ctx).Unscoped().Where(dao.TravelProduct.Columns().Id, item.ProductId).Scan(&product); err == nil && product != nil {
				item.ProductInfo = &struct {
					gmeta.Meta `orm:"table:hg_travel_product"`
					Id         uint64 `json:"id"              dc:""`
					Title      string `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
					SubTitle   string `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
				}{Id: product.Id, Title: product.Title, SubTitle: product.SubTitle}
			}
		}
	}
	return
}

// VerifyView 获取核销记录信息
func (s *sTravelVerifyRecord) VerifyView(ctx context.Context, in *input_travel.VerifyLogViewInp) (res *input_travel.VerifyLogViewModel, err error) {

	if err = s.Model(ctx).Hook(hook2.PmsFindLanguageValueHook).WithAll().WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取核销记录信息失败，请稍后重试！")
		return
	}

	// 查询核销人员
	if res.VerifyStaffId > 0 {
		var staff *entity.TravelVerifyStaff
		if err = dao.TravelVerifyStaff.Ctx(ctx).Where(dao.TravelVerifyStaff.Columns().Id, res.VerifyStaffId).Scan(&staff); err == nil && staff != nil {
			if staff.Name != "" {
				res.StaffName = staff.Name
			} else {
				res.StaffName = staff.Username
			}
		}
	}

	return
}

func (s *sTravelVerifyRecord) CodeView(ctx context.Context, in *input_travel.CodeViewInp) (res *input_travel.CodeViewModel, err error) {
	Logger := g.Log().Path("logs/TravelStaffVerify")
	Logger.Info(ctx, "--------一日游核销码详情进入----------")
	Logger.Info(ctx, gjson.New(in))

	// 解密动态券码并验证时间戳
	var orderSn string
	//var memberId int
	var timestamp int64

	// 尝试解密动态券码
	// 首先尝试base64解码
	encryptedData, base64Err := base64.StdEncoding.DecodeString(in.Code)
	if base64Err != nil {
		// 如果base64解码失败，可能是旧的券码格式，直接使用原券号查询
		orderSn = in.Code
		Logger.Info(ctx, "base64解码失败")
	} else {
		// base64解码成功，尝试AES解密
		decryptedText, decryptErr := encrypt.AesECBDecrypt(encryptedData, consts.RequestEncryptKey)
		Logger.Info(ctx, "--------核销解码----------")
		Logger.Info(ctx, decryptedText)
		if decryptErr != nil {
			// 如果解密失败，可能是旧的券码格式，直接使用原券号查询
			orderSn = in.Code
			Logger.Info(ctx, "base64解码成功，AES解密失败")
		} else {
			// 解密成功，解析格式：couponNo|memberId|timestamp
			parts := strings.Split(string(decryptedText), "|")
			if len(parts) != 3 {
				err = gerror.New("动态券码格式不正确！")
				return
			}

			orderSn = parts[0]
			//memberId, err = strconv.Atoi(parts[1])
			//if err != nil {
			//	err = gerror.New("动态券码会员ID格式不正确！")
			//	return
			//}

			timestamp, err = strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				err = gerror.New("动态券码时间戳格式不正确！")
				return
			}

			// 验证时间戳是否在30秒内
			currentTime := gtime.Now().Unix()
			if currentTime-timestamp > 60 {
				err = gerror.New("动态券码已过期，请重新生成！")
				return
			}

		}
	}
	Logger.Info(ctx, "--------OrderSn----------")
	Logger.Info(ctx, orderSn)

	// 搜索一日游订单
	if err = dao.TravelOrder.Ctx(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook).
		Where(dao.TravelOrder.Columns().OrderSn, orderSn).
		Scan(&res); err != nil {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	return
}

func (s *sTravelVerifyRecord) CodeVerify(ctx context.Context, in *input_travel.CodeVerifyInp) (res *input_travel.CodeVerifyModel, err error) {
	Logger := g.Log().Path("logs/TravelStaffVerify")
	Logger.Info(ctx, "--------一日游核销进入----------")
	Logger.Info(ctx, gjson.New(in))

	var TravelOrder *entity.TravelOrder

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		MemberInfo := contexts.GetTravelStaffUser(ctx)

		Logger.Info(ctx, "--------OrderSn----------")
		Logger.Info(ctx, in.OrderSn)

		// 搜索一日游订单
		if err = dao.TravelOrder.Ctx(ctx).WithAll().Hook(hook2.PmsFindLanguageValueHook).
			Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).
			Scan(&TravelOrder); err != nil {
			err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
			return
		}

		// 判断会员是否已注销
		memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, TravelOrder.MemberId).Count()
		if memberCount == 0 {
			// 会员已注销，无法操作此订单
			err = gerror.New(gi18n.T(ctx, "member_has_been_cancelled"))
			return
		}

		if TravelOrder.OrderStatus != "WAIT_VERIFY" {
			// 订单状态异常
			err = gerror.New(gi18n.T(ctx, "order_status_incorrect"))
			return
		}

		now := gtime.Now()
		today := gtime.NewFromStr(now.Format("Y-m-d"))
		bookDay := gtime.NewFromStr(TravelOrder.BookDate.Format("Y-m-d"))

		if today.After(bookDay) {
			// 订单已逾期，无法核销
			err = gerror.New(gi18n.T(ctx, "the_order_is_overdue_and_cannot_be_redeemed"))
		}

		// 只有预约当天可以核销，早于或晚于都不行
		if !today.Equal(bookDay) {
			// 订单只能在预约当天核销
			err = gerror.New(gi18n.T(ctx, "order_can_only_be_verified_on_book_date"))
			return
		}

		// 核销权限校验：全部 / 产品全部SKU / 指定SKU
		accessCount, accessErr := dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).
			Where(dao.TravelVerifyStaffScope.Columns().StaffId, MemberInfo.Id).
			Where(dao.TravelVerifyStaffScope.Columns().IsAll, 1).
			Count()
		if accessErr != nil {
			err = gerror.Wrap(accessErr, "校验核销权限失败")
			return
		}
		if accessCount == 0 {
			accessCount, accessErr = dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).
				Where(dao.TravelVerifyStaffScope.Columns().StaffId, MemberInfo.Id).
				Where(dao.TravelVerifyStaffScope.Columns().ProductId, TravelOrder.ProductId).
				Where(dao.TravelVerifyStaffScope.Columns().SkuId, 0).
				Count()
			if accessErr != nil {
				err = gerror.Wrap(accessErr, gi18n.T(ctx, "verify_permission_check_failed"))
				return
			}
		}
		if accessCount == 0 {
			accessCount, accessErr = dao.TravelVerifyStaffScope.Ctx(ctx).TX(tx).
				Where(dao.TravelVerifyStaffScope.Columns().StaffId, MemberInfo.Id).
				Where(dao.TravelVerifyStaffScope.Columns().ProductId, TravelOrder.ProductId).
				Where(dao.TravelVerifyStaffScope.Columns().SkuId, TravelOrder.SkuId).
				Count()
			if accessErr != nil {
				err = gerror.Wrap(accessErr, gi18n.T(ctx, "verify_permission_check_failed"))
				return
			}
		}
		if accessCount == 0 {
			err = gerror.New(gi18n.T(ctx, "verify_staff_no_product_or_sku_permission"))
			return
		}

		// ---开始核销---
		// 修改订单状态
		if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(dao.TravelOrder.Columns().Id, TravelOrder.Id).Data(g.MapStrAny{
			dao.TravelOrder.Columns().OrderStatus:   "DONE",
			dao.TravelOrder.Columns().VerifyTime:    gtime.Now(),
			dao.TravelOrder.Columns().VerifyStaffId: MemberInfo.Id,
		}).Update(); err != nil {
			return
		}

		// 写入订单日志travel_order_log
		if _, err = dao.TravelOrderLog.Ctx(ctx).TX(tx).Insert(&entity.TravelOrderLog{
			OrderId:     int(TravelOrder.Id),
			ActionWay:   "VERIFY",
			OrderStatus: "DONE",
			Remark:      "订单已核销",
			OperateType: "STAFF",
			OperateId:   int(MemberInfo.Id),
		}); err != nil {
			return
		}

		// 写入订单核销记录travel_verify_record
		if _, err = dao.TravelVerifyRecord.Ctx(ctx).TX(tx).Insert(&entity.TravelVerifyRecord{
			OrderId:       TravelOrder.Id,
			OrderSn:       TravelOrder.OrderSn,
			ProductId:     TravelOrder.ProductId,
			MemberId:      TravelOrder.MemberId,
			BookDate:      TravelOrder.BookDate,
			VerifyStaffId: MemberInfo.Id,
			VerifyTime:    gtime.Now(),
		}); err != nil {
			return
		}

		// 转发到返利队列 订单计算佣金/计算经验需要扔队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(TravelOrder.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(TravelOrder.OrderSn).Bytes(),
			Header:       nil,
		})

		// 发送分销订单变更队列
		if TravelOrder.IsFx == "Y" {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
				DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
					OrderNo:      TravelOrder.OrderSn,
					ChangeStatus: "COMPLETE",
				}).MustToJson(),
				Header: nil,
			})
		}

		return
	})

	if err != nil {
		Logger.Info(ctx, "--------核销失败----------")
		Logger.Info(ctx, err)

		// 发送消息
		_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
			MemberId: gvar.New(TravelOrder.MemberId).Int(),
			Code:     -1,
			Event:    "VERIFY",
			Message:  "核销失败",
		})

		return
	}

	Logger.Info(ctx, "--------核销成功----------")
	// 返回结果
	res = new(input_travel.CodeVerifyModel)
	res.VerifyTime = gtime.Now()
	res.OrderSn = TravelOrder.OrderSn

	// 发送消息
	_ = service.BasicsWs().SendMemberWebsocketMessage(ctx, &ws.SendWebsocketMessageInp{
		MemberId: gvar.New(TravelOrder.MemberId).Int(),
		Code:     200,
		Event:    "VERIFY",
		Message:  "核销成功",
	})

	return
}
