package logic_spa

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_refund"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"APT/utility/convert"
	"APT/utility/excel"
	"APT/utility/format"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/shopspring/decimal"
)

type sSpaOrder struct{}

func NewSpaOrder() *sSpaOrder {
	return &sSpaOrder{}
}

func init() {
	service.RegisterSpaOrder(NewSpaOrder())
}

func (s *sSpaOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.SpaOrder.Ctx(ctx), option...)
}

func (s *sSpaOrder) List(ctx context.Context, in *input_spa.SpaOrderListInp) (list []*input_spa.SpaOrderListModel, totalCount int, err error) {
	mod := dao.SpaOrder.Ctx(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.SpaOrder.Table(), input_spa.SpaOrderListModel{})
	mod = mod.Fields(fmt.Sprintf("`%s`.`%s` as `pmsMemberMemberNo`", dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))

	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.SpaOrder.Table(), dao.SpaOrder.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.SpaOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}

	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.SpaOrder.Columns().MemberId, in.MemberId)
	}

	if !g.IsEmpty(in.TechnicianName) {
		var (
			technicianIds []gdb.Value
			orderIds      []gdb.Value
		)
		technicianIds, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Ctx(ctx).Builder().
			WhereLike(dao.SpaTechnician.Columns().Name, "%"+in.TechnicianName+"%").
			WhereOrLike(dao.SpaTechnician.Columns().Nickname, "%"+in.TechnicianName+"%")).Array(dao.SpaTechnician.Columns().Id)

		if len(technicianIds) > 0 {
			orderIds, err = dao.SpaOrderTechnician.Ctx(ctx).WhereIn(dao.SpaOrderTechnician.Columns().TechnicianId, technicianIds).Array(dao.SpaOrderTechnician.Columns().OrderId)
			if len(orderIds) > 0 {
				mod = mod.WherePrefixIn(dao.SpaOrder.Table(), dao.SpaOrder.Columns().Id, orderIds)
			} else {
				mod = mod.WherePrefix(dao.SpaOrder.Table(), dao.SpaOrder.Columns().Id, 0)
			}
		} else {
			mod = mod.WherePrefix(dao.SpaOrder.Table(), dao.SpaOrder.Columns().Id, 0)
		}
	}
	if !g.IsEmpty(in.MemberSearch) {
		mod = mod.Where(mod.Builder().
			WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Id, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.SpaOrder.Table(), dao.SpaOrder.Columns().BookingName, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.SpaOrder.Table(), dao.SpaOrder.Columns().BookingMobile, "%"+in.MemberSearch+"%"))
	}

	if !g.IsEmpty(in.OrderStatus) && in.OrderStatus != "ALL" && in.OrderStatus != "ABNORMAL" {
		mod = mod.Where(dao.SpaOrder.Columns().OrderStatus, in.OrderStatus)
	}

	if in.OrderStatus == "ABNORMAL" {
		mod = mod.WhereGT(dao.SpaOrder.Columns().AbnormalStatus, 1)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.SpaOrder.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取预订单列表失败，请稍后重试！")
		return
	}

	// 对于已删除的关联数据，手动加载
	for _, item := range list {
		// 加载已删除的服务
		if item.ServiceId > 0 && item.ServiceDetail == nil {
			var svc *entity.SpaService
			if err = dao.SpaService.Ctx(ctx).Unscoped().Where(dao.SpaService.Columns().Id, item.ServiceId).Scan(&svc); err == nil && svc != nil {
				item.ServiceDetail = &struct {
					gmeta.Meta `orm:"table:hg_spa_service"`
					*entity.SpaService
				}{SpaService: svc}
			}
		}
		// 加载已删除的商品
		if item.GoodsId > 0 && item.GoodsDetail == nil {
			var goods *entity.SpaServiceGoods
			if err = dao.SpaServiceGoods.Ctx(ctx).Unscoped().Where(dao.SpaServiceGoods.Columns().Id, item.GoodsId).Scan(&goods); err == nil && goods != nil {
				item.GoodsDetail = &struct {
					gmeta.Meta `orm:"table:hg_spa_service_goods"`
					*entity.SpaServiceGoods
				}{SpaServiceGoods: goods}
			}
		}
	}

	return
}

func (s *sSpaOrder) View(ctx context.Context, in *input_spa.SpaOrderViewInp) (res *input_spa.SpaOrderViewModel, err error) {
	if err = dao.SpaOrder.Ctx(ctx).Unscoped().WithAll().Where(dao.SpaOrder.Columns().OrderSn, in.OrderSn).Hook(hook2.PmsFindLanguageValueHook).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取预订单信息，请稍后重试！")
		return
	}

	if res == nil {
		return
	}

	// 检查会员是否已删除
	if res.MemberId > 0 {
		memberDeletedAt, _ := dao.PmsMember.Ctx(ctx).Unscoped().Fields(dao.PmsMember.Columns().DeletedAt).
			Where(dao.PmsMember.Columns().Id, res.MemberId).Value()
		if !memberDeletedAt.IsNil() && !memberDeletedAt.IsEmpty() {
			res.MemberDeleted = true
			// 手动加载已删除的会员信息
			if res.MemberDetail == nil {
				var member struct {
					Id       int    `json:"id"`
					FullName string `json:"fullName"`
					MemberNo string `json:"memberNo"`
				}
				if err = dao.PmsMember.Ctx(ctx).Unscoped().
					Fields("id, full_name, member_no").
					Where(dao.PmsMember.Columns().Id, res.MemberId).Scan(&member); err == nil {
					res.MemberDetail = &struct {
						gmeta.Meta `orm:"table:hg_pms_member"`
						Id         int    `json:"id"    orm:"id"      dc:"id"`
						FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
						MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
					}{Id: member.Id, FullName: member.FullName, MemberNo: member.MemberNo}
				}
			}
		}
	}

	// 手动加载已删除的服务
	if res.ServiceId > 0 && res.ServiceDetail == nil {
		var svc *entity.SpaService
		if err = dao.SpaService.Ctx(ctx).Unscoped().Where(dao.SpaService.Columns().Id, res.ServiceId).Scan(&svc); err == nil && svc != nil {
			res.ServiceDetail = &struct {
				gmeta.Meta `orm:"table:hg_spa_service"`
				*entity.SpaService
			}{SpaService: svc}
		}
	}

	// 手动加载已删除的商品
	if res.GoodsId > 0 && res.GoodsDetail == nil {
		var goods *entity.SpaServiceGoods
		if err = dao.SpaServiceGoods.Ctx(ctx).Unscoped().Where(dao.SpaServiceGoods.Columns().Id, res.GoodsId).Scan(&goods); err == nil && goods != nil {
			res.GoodsDetail = &struct {
				gmeta.Meta `orm:"table:hg_spa_service_goods"`
				*entity.SpaServiceGoods
			}{SpaServiceGoods: goods}
		}
	}

	for _, v := range res.LogList {
		if v.OperateType == "SYSTEM" {
			v.OperateName = "系统"
		}
		if v.OperateType == "ADMIN" {
			var AdminMemberInfo *entity.AdminMember
			if err = dao.AdminMember.Ctx(ctx).Unscoped().Where(dao.AdminMember.Columns().Id, v.OperateId).Scan(&AdminMemberInfo); err != nil {
				return
			}
			if AdminMemberInfo != nil {
				v.OperateName = AdminMemberInfo.Username
			}
		}
		if v.OperateType == "USER" {
			var PmsMemberInfo *entity.PmsMember
			if err = dao.PmsMember.Ctx(ctx).Unscoped().Where(dao.PmsMember.Columns().Id, v.OperateId).Scan(&PmsMemberInfo); err != nil {
				return
			}
			if PmsMemberInfo != nil {
				v.OperateName = PmsMemberInfo.FullName
			}
		}

		if v.OperateType == "TECHNICIAN" {
			var TechnicianInfo *entity.SpaTechnician
			if err = dao.SpaTechnician.Ctx(ctx).Unscoped().Where(dao.SpaTechnician.Columns().Id, v.OperateId).Scan(&TechnicianInfo); err != nil {
				return
			}
			if TechnicianInfo != nil {
				v.OperateName = TechnicianInfo.Name
			}
		}

		if v.OperateType == "ISP" {
			var IspInfo *entity.SpaIsp
			if err = dao.SpaIsp.Ctx(ctx).Where(dao.SpaIsp.Columns().Id, v.OperateId).Scan(&IspInfo); err != nil {
				return
			}
			if IspInfo != nil {
				v.OperateName = IspInfo.Name
			}
		}
	}

	return
}

func (s *sSpaOrder) ConfirmAgree(ctx context.Context, in *input_spa.SpaOrderConfirmAgreeInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderConfirmAgreeFields{
			OrderStatus: "WAIT_SERVE",
			ConfirmTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().OrderStatus: "WAIT_SERVE",
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   "CONFIRMED",
			Remark:      "系统已自动接单",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已确认",
			"en":    "Order confirmed",
			"ja":    "注文確定",
			"ko":    "주문 확인됨",
			"zh_CN": "訂單已確認",
		}
		systemMessageContent := map[string]string{
			"zh":    models.OrderSn + "订单后台已接单",
			"en":    "Order " + models.OrderSn + " has been received in the system.",
			"ja":    "システムで注文" + models.OrderSn + "が受信されました。",
			"ko":    "시스템에 주문" + models.OrderSn + "이 접수되었습니다.",
			"zh_CN": models.OrderSn + "訂單後台已接單",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

// ConfirmDisagree 确认失败，全额退款
func (s *sSpaOrder) ConfirmDisagree(ctx context.Context, in *input_spa.SpaOrderConfirmDisagreeInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_CONFIRM" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.RefundStatus != "WAIT" {
		err = gerror.New("订单已退款")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 改变订单确认状态
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(input_spa.SpaOrderConfirmDisagreeFields{
			BookingStatus:       "CANCEL",
			ConfirmRefuseReason: in.ConfirmRefuseReason,
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "CANCEL",
			ActionWay:   "DISCONFIRMED",
			Remark:      fmt.Sprintf("系统拒绝接单，原因：%s", in.ConfirmRefuseReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 确认失败、全额退款
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).Scan(&Transaction); err != nil {
			return
		}
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

		for _, v := range Transaction {
			// 可退款金额
			Refundable := v.Amount - v.RefundAmount
			if v.PayType == "BAL" {
				RefundBalance = RefundBalance + Refundable
			} else if v.PayType == "COUPON" {

			} else {
				RefundFee = RefundFee + Refundable
			}
		}
		// 退款金额
		RefundAmount := RefundBalance + RefundFee

		// 修改订单状态
		RefundStatus := "DONE"
		if models.OrderAmount > RefundAmount {
			RefundStatus = "PART"
		}
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(g.MapStrAny{
			dao.SpaOrder.Columns().RefundAmount:       RefundAmount,
			dao.SpaOrder.Columns().RefundBalAmount:    RefundBalance,
			dao.SpaOrder.Columns().RefundCouponAmount: 0,
			dao.SpaOrder.Columns().RefundStatus:       RefundStatus,
			dao.SpaOrder.Columns().RefundTime:         gtime.Now(),
			dao.SpaOrder.Columns().OrderStatus:        "CANCEL",
			dao.SpaOrder.Columns().PayStatus:          "REFUND",
			dao.SpaOrder.Columns().CancelTime:         gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).TX(tx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().RefundAmount:       RefundAmount,
			dao.SpaOrderGoods.Columns().RefundBalAmount:    RefundBalance,
			dao.SpaOrderGoods.Columns().RefundCouponAmount: 0,
			dao.SpaOrderGoods.Columns().RefundStatus:       RefundStatus,
			dao.SpaOrderGoods.Columns().RefundTime:         gtime.Now(),
			dao.SpaOrderGoods.Columns().OrderStatus:        "CANCEL",
			dao.SpaOrderGoods.Columns().PayStatus:          "REFUND",
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 更新服务预定量和预定金额
		if _, err = dao.SpaService.Ctx(ctx).TX(tx).Where(dao.SpaService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
			dao.SpaService.Columns().PayOrderNum:    gdb.Raw("pay_order_num-1"),
			dao.SpaService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
			return
		}

		// 更新项目预定量和预定金额
		if _, err = dao.SpaServiceGoods.Ctx(ctx).TX(tx).Where(dao.SpaServiceGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
			dao.SpaServiceGoods.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", models.GoodsNum)),
			dao.SpaServiceGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
		}); err != nil {
			err = gerror.Wrap(err, "更新项目信息失败，请稍后重试！")
			return
		}

		//// 更新技师预定量和预定金额
		//if !g.IsEmpty(models.TechnicianIds) {
		//	for _, Technician := range strings.Split(models.TechnicianIds, ",") {
		//		if !g.IsEmpty(Technician) {
		//			TechnicianId, _ := strconv.ParseInt(Technician, 10, 64)
		//			// 更新技师预定量和预定金额
		//			if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, TechnicianId).Update(g.MapStrAny{
		//				dao.SpaTechnician.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
		//				dao.SpaTechnician.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", models.OrderAmount/float64(models.GoodsNum))),
		//			}); err != nil {
		//				err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
		//				return
		//			}
		//		}
		//	}
		//}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "CANCEL",
			ActionWay:   "REFUND",
			Remark:      "订单已退款",
			OperateType: "SYSTEM",
		}); err != nil {
			return err
		}

		// 全额退款
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      models.OrderSn,
			RefundAmount: RefundAmount,
			OperateType:  "ADMIN",
			OperateId:    int(contexts.GetUserId(ctx)),
		}, tx)
		if err != nil {
			err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已被拒",
			"en":    "The order has been rejected.",
			"ja":    "注文は拒否されました。",
			"ko":    "주문이 거부되었습니다.",
			"zh_CN": "訂單已被拒",
		}
		systemMessageContent := map[string]string{
			"zh":    "订单" + models.OrderSn + "已被拒，已全额退款",
			"en":    "Order " + models.OrderSn + " has been rejected and fully refunded.",
			"ja":    "注文 " + models.OrderSn + " は拒否され、全額返金されました。",
			"ko":    "주문 " + models.OrderSn + "은 거부되었으며 전액 환불되었습니다.",
			"zh_CN": "訂單" + models.OrderSn + "已被拒，已全額退款",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

func (s *sSpaOrder) TechnicianList(ctx context.Context, in *input_spa.SpaOrderTechnicianInp) (list []*input_spa.SpaOrderTechnicianModel, totalCount int, err error) {
	var (
		OrderInfo       *entity.SpaOrder
		NoTechnicianIds []string
		CheckOrderList  []*entity.SpaOrder
		Config          *input_basics.GetConfigModel
		BeforeTime      int
		OrderBeforeTime int
	)
	if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().Id, in.Id).Scan(&OrderInfo); err != nil {
		return
	}

	techMod := dao.SpaTechnician.Ctx(ctx)
	techMod = techMod.Fields(input_spa.SpaOrderTechnicianModel{})

	if !g.IsEmpty(OrderInfo.IspId) {
		techMod = techMod.Where(dao.SpaTechnician.Columns().IspId, OrderInfo.IspId)
	}

	//if OrderInfo.ConfirmType == 1 {
	// 手动确认
	Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
		Group: "spaordersetting",
	})
	if err != nil {
		return
	}
	OrderMode := Config.List["orderMode"]
	ServiceTimeDuration := Config.List["serviceTimeDuration"]
	OutWorkTimeDuration := Config.List["outWorkTimeDuration"]
	if gvar.New(OrderMode).Int() == 2 {
		if OrderInfo.ServiceType == 1 {
			BeforeTime = gvar.New(ServiceTimeDuration).Int()
		} else {
			BeforeTime = gvar.New(ServiceTimeDuration).Int() + gvar.New(OutWorkTimeDuration).Int()
		}
		StartTime := gtime.New(OrderInfo.BookStartTime).Add(time.Duration(-BeforeTime) * gtime.M).Timestamp()
		// 找出订单预约当天的已被派单的且处于待服务或服务中状态的订单
		if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().BookDate, OrderInfo.BookDate).WhereNot(dao.SpaOrder.Columns().Id, in.Id).WhereIn(dao.SpaOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Where(dao.SpaOrder.Columns().DispatchStatus, "DONE").Scan(&CheckOrderList); err != nil {
			return
		}

		// 遍历订单
		for _, v := range CheckOrderList {
			if v.ServiceType == 1 {
				OrderBeforeTime = gvar.New(ServiceTimeDuration).Int()
			} else {
				OrderBeforeTime = gvar.New(ServiceTimeDuration).Int() + gvar.New(OutWorkTimeDuration).Int()
			}
			OrderStartTime := gtime.New(v.BookStartTime).Add(time.Duration(-OrderBeforeTime) * gtime.M).Timestamp()
			if gtime.New(v.BookEndTime).Timestamp() <= StartTime || OrderStartTime >= gtime.New(OrderInfo.BookEndTime).Timestamp() {

			} else {
				technicianId := strings.Split(v.TechnicianIds, ",")
				for _, vv := range technicianId {
					technicianIdString := gvar.New(vv).String()
					NoTechnicianIds = append(NoTechnicianIds, technicianIdString)
				}
			}
		}
		NoTechnicianIdsStr := strings.Join(NoTechnicianIds, ",")
		if g.IsEmpty(NoTechnicianIdsStr) {
			NoTechnicianIdsStr = "0"
		}
		techMod = techMod.WhereNot(dao.SpaTechnician.Columns().WorkStatus, "REST").Where(dao.SpaTechnician.Columns().Status, 1)
		techMod = techMod.Fields(fmt.Sprintf(`
	CASE
        WHEN work_status='REST' THEN 1
        WHEN id IN(%s) THEN 2
        ELSE 3
    END AS work_status_enum
`, NoTechnicianIdsStr))
		techMod = techMod.Where(dao.SpaTechnician.Columns().Status, 1)
		techMod = techMod.OrderDesc(fmt.Sprintf(`
	CASE
        WHEN work_status='REST' THEN 1
        WHEN id IN(%s) THEN 2
        ELSE 3
    END
`, NoTechnicianIdsStr))
	} else {
		techMod = techMod.Where(dao.SpaTechnician.Columns().Status, 1)
		techMod = techMod.Fields(fmt.Sprintf(`
	CASE
        WHEN id > %d THEN 3
        ELSE 3
    END AS work_status_enum
`, 0))
	}

	//} else {
	//	techMod = techMod.Where(dao.SpaTechnician.Columns().Status, 1)
	//}
	// 分页
	if in.Pagination {
		techMod = techMod.Page(in.Page, in.PerPage)
	}
	if in.Pagination {
		if err = techMod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取技师列表失败，请稍后重试！")
			return
		}
	} else {
		if err = techMod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取技师列表失败，请稍后重试！")
			return
		}
	}
	return
}

func (s *sSpaOrder) Dispatch(ctx context.Context, in *input_spa.SpaOrderDispatchInp) (err error) {

	var (
		models                *entity.SpaOrder
		OrderGoods            *entity.SpaOrderGoods
		TechnicianInfo        *entity.SpaTechnician
		TechnicianName        string
		CheckOrderList        []*entity.SpaOrder
		Config                *input_basics.GetConfigModel
		IspSettlementRate     float64
		IspSettlementObject   string
		BeforeTime            int
		OrderBeforeTime       int
		CantOrderNum          int
		OrderTechnicianInfo   []*entity.SpaOrderTechnician
		TechnicianSettlement  *input_spa.SpaOrderDispatchTechnicianModel
		IspSettlement         *input_spa.SpaOrderDispatchIspModel
		ActualSettlementRate  float64
		ActualSettlementType  int
		ActualSettlementCycle int
		ActualSettlementCost  string
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" || models.DispatchStatus != "WAIT" {
		err = gerror.New("订单状态不正确")
		return
	}

	if err = dao.SpaOrderGoods.Ctx(ctx).Where("order_id", in.Id).Scan(&OrderGoods); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models.IspId > 0 {
		if err = dao.SpaIsp.Ctx(ctx).WithAll().Where("id", models.IspId).Scan(&IspSettlement); err != nil {
			err = gerror.Wrap(err, consts.ErrorORM)
			return
		}

		if IspSettlement.SettlementType == 1 {
			// 跟随系统
			if IspSettlement.SettlementInfo.Type == 1 {
				IspSettlementRate = 0
			} else {
				IspSettlementRate = IspSettlement.SettlementInfo.Rate
			}
		} else {
			// 自己设置
			IspSettlementRate = IspSettlement.SettlementRate
		}

		IspSettlementObject = IspSettlement.SettlementObject
	} else {
		IspSettlementObject = "TECHNICIAN"
	}

	// 重新判断下技师
	if !g.IsEmpty(in.TechnicianIds) {
		Config, err = service.BasicsConfig().GetConfigByGroup(ctx, &input_basics.GetConfigInp{
			Group: "spaordersetting",
		})
		if err != nil {
			return
		}
		ServiceTimeDuration := Config.List["serviceTimeDuration"]
		OutWorkTimeDuration := Config.List["outWorkTimeDuration"]
		OrderMode := Config.List["orderMode"]
		if models.ServiceType == 1 {
			BeforeTime = gvar.New(ServiceTimeDuration).Int()
		} else {
			BeforeTime = gvar.New(ServiceTimeDuration).Int() + gvar.New(OutWorkTimeDuration).Int()
		}
		StartTime := gtime.New(models.BookStartTime).Add(time.Duration(-BeforeTime) * gtime.M).Timestamp()
		for _, Technician := range strings.Split(in.TechnicianIds, ",") {
			CantOrderNum = 0
			if !g.IsEmpty(Technician) {
				TechnicianId, _ := strconv.ParseInt(Technician, 10, 64)
				if err = dao.SpaTechnician.Ctx(ctx).WherePri(TechnicianId).Scan(&TechnicianInfo); err != nil {
					return
				}
				if gvar.New(OrderMode).Int() == 2 {
					OrderIds, _ := dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().TechnicianId, TechnicianId).Array(dao.SpaOrderTechnician.Columns().OrderId)
					if err = dao.SpaOrder.Ctx(ctx).Where(dao.SpaOrder.Columns().BookDate, models.BookDate).WhereIn(dao.SpaOrder.Columns().Id, OrderIds).WhereNot(dao.SpaOrder.Columns().Id, in.Id).WhereIn(dao.SpaOrder.Columns().OrderStatus, g.Slice{"WAIT_SERVE", "SERVING"}).Scan(&CheckOrderList); err != nil {
						return
					}
					for _, v := range CheckOrderList {
						if v.ServiceType == 1 {
							OrderBeforeTime = gvar.New(ServiceTimeDuration).Int()
						} else {
							OrderBeforeTime = gvar.New(ServiceTimeDuration).Int() + gvar.New(OutWorkTimeDuration).Int()
						}
						OrderStartTime := gtime.New(v.BookStartTime).Add(time.Duration(-OrderBeforeTime) * gtime.M).Timestamp()
						if gtime.New(v.BookEndTime).Timestamp() <= StartTime || OrderStartTime >= gtime.New(models.BookEndTime).Timestamp() {
						} else {
							CantOrderNum++
						}
					}
				} else {
					if TechnicianInfo.Status == 2 {
						CantOrderNum++
					}
				}
				if CantOrderNum > 0 {
					err = gerror.New(fmt.Sprintf("技师%s不可派单，请重新选择", TechnicianInfo.Nickname))
					return
				}
				TechnicianName = TechnicianName + TechnicianInfo.Nickname + ","
			}
		}
	}
	// TechnicianName 去掉最后一个逗号
	TechnicianName = strings.TrimSuffix(TechnicianName, ",")

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderDispatchInp{
			DispatchStatus:     "DONE",
			DispatchTime:       gtime.Now(),
			DispatchOperatorId: in.DispatchOperatorId,
			DispatchDesc:       in.DispatchDesc,
			TechnicianIds:      in.TechnicianIds,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).TX(tx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().DispatchStatus:     "DONE",
			dao.SpaOrderGoods.Columns().DispatchTime:       gtime.Now(),
			dao.SpaOrderGoods.Columns().DispatchOperatorId: in.DispatchOperatorId,
			dao.SpaOrderGoods.Columns().DispatchDesc:       in.DispatchDesc,
			dao.SpaOrderGoods.Columns().TechnicianIds:      in.TechnicianIds,
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   "DISPATCH",
			Remark:      "店长派单",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 每次调度，需重新计算技师订单金额
		preTechnicianOrderAmount := format.Round2Float64(models.OrderAmount / float64(models.GoodsNum))
		preTechnicianCouponAmount := format.Round2Float64(models.CouponAmount / float64(models.GoodsNum))
		preTechnicianBalAmount := format.Round2Float64(models.BalAmount / float64(models.GoodsNum))

		// 存订单技师表
		// 记录结算比例和金额
		if !g.IsEmpty(in.TechnicianIds) {
			for _, Technician := range strings.Split(in.TechnicianIds, ",") {
				if !g.IsEmpty(Technician) {
					TechnicianId, _ := strconv.ParseInt(Technician, 10, 64)
					_ = dao.SpaTechnician.Ctx(ctx).WithAll().WherePri(TechnicianId).Scan(&TechnicianSettlement)
					var SettlementRate float64
					if TechnicianSettlement.SettlementType == 1 {
						// 跟随系统
						if TechnicianSettlement.SettlementInfo.Type == 1 {
							SettlementRate = 0
						} else {
							SettlementRate = TechnicianSettlement.SettlementInfo.Rate
						}
					} else {
						// 自己设置
						SettlementRate = TechnicianSettlement.SettlementRate
					}

					if IspSettlementObject == "TECHNICIAN" {
						ActualSettlementRate = SettlementRate
						ActualSettlementType = TechnicianSettlement.SettlementInfo.Type
						ActualSettlementCycle = TechnicianSettlement.SettlementInfo.Cycle
						ActualSettlementCost = TechnicianSettlement.SettlementInfo.Cost
					} else {
						ActualSettlementRate = IspSettlementRate
						ActualSettlementType = IspSettlement.SettlementInfo.Type
						ActualSettlementCycle = IspSettlement.SettlementInfo.Cycle
						ActualSettlementCost = IspSettlement.SettlementInfo.Cost
					}

					OrderTechnicianInfo = append(OrderTechnicianInfo, &entity.SpaOrderTechnician{
						IspId:            models.IspId,
						OrderId:          int64(in.Id),
						OrderGoodsId:     OrderGoods.Id,
						TechnicianId:     TechnicianId,
						OrderAmount:      preTechnicianOrderAmount,
						CouponAmount:     preTechnicianCouponAmount,
						BalAmount:        preTechnicianBalAmount,
						SettlementObject: IspSettlementObject,
						SettlementRate:   ActualSettlementRate,
						SettlementStatus: "WAIT",
						SettlementType:   ActualSettlementType,
						SettlementCycle:  ActualSettlementCycle,
						SettlementCost:   ActualSettlementCost,
					})

					// 更新技师预定量和预定金额
					if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, TechnicianId).Update(g.MapStrAny{
						dao.SpaTechnician.Columns().TotalOrderNum:    gdb.Raw(fmt.Sprintf("total_order_num+%d", 1)),
						dao.SpaTechnician.Columns().TotalOrderAmount: gdb.Raw(fmt.Sprintf("total_order_amount+%f", models.OrderAmount/float64(models.GoodsNum))),
						dao.SpaTechnician.Columns().PayOrderNum:      gdb.Raw(fmt.Sprintf("pay_order_num+%d", 1)),
						dao.SpaTechnician.Columns().PayOrderAmount:   gdb.Raw(fmt.Sprintf("pay_order_amount+%f", models.OrderAmount/float64(models.GoodsNum))),
					}); err != nil {
						err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
						return
					}
				}
			}
		}
		if !g.IsEmpty(OrderTechnicianInfo) {
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Delete(); err != nil {
				err = gerror.Wrap(err, "清理旧数据失败，请稍后重试！")
				return
			}
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).OmitEmptyData().Insert(OrderTechnicianInfo); err != nil {
				return err
			}
		} else {
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Delete(); err != nil {
				err = gerror.Wrap(err, "清理旧数据失败，请稍后重试！")
				return
			}
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "技师已接单",
			"en":    "Technician has accepted the order.",
			"ja":    "技術者が注文を承諾しました。",
			"ko":    "기술자가 주문을 수락했습니다.",
			"zh_CN": "技師已接單",
		}
		systemMessageContent := map[string]string{
			"zh":    "您的订单" + models.OrderSn + "已派单给技师" + TechnicianName,
			"en":    "Your order " + models.OrderSn + " has been dispatched to technician " + TechnicianName,
			"ja":    "ご注文番号" + models.OrderSn + "は技術者" + TechnicianName + "に発送されました",
			"ko":    "귀하의 주문 " + models.OrderSn + "이 기술자 " + TechnicianName + "에게 발송되었습니다.",
			"zh_CN": "您的訂單" + models.OrderSn + "已派給技師" + TechnicianName,
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

// TechnicianOrderList 结算订单列表
func (s *sSpaOrder) TechnicianOrderList(ctx context.Context, in *input_spa.SettleSpaOrderTechnicianListInp) (list []*input_spa.SettleSpaOrderTechnicianListModel, totalCount int, err error) {
	mod := dao.SpaOrderTechnician.Ctx(ctx).Safe().WithAll()

	mod = mod.Fields(input_spa.SettleSpaOrderTechnicianListModel{})

	if !g.IsEmpty(in.SettlementOrderId) {
		mod = mod.Where(dao.SpaOrderTechnician.Columns().SettlementOrderId, in.SettlementOrderId)
	}

	if !g.IsEmpty(in.OrderSn) {
		var OrderIds []int
		columns, _ := dao.SpaOrder.Ctx(ctx).
			Fields(dao.SpaOrder.Columns().Id).
			WhereLike(dao.SpaOrder.Columns().OrderSn, "%"+in.OrderSn+"%").
			Array()
		OrderIds = g.NewVar(columns).Ints()
		mod = mod.WhereIn(dao.SpaOrderTechnician.Columns().OrderId, OrderIds)
	}

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.SpaOrderTechnician.Columns().Id)
	//mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取结算订单列表失败，请稍后重试！")
		return
	}
	return
}

func (s *sSpaOrder) GoOut(ctx context.Context, in *input_spa.SpaOrderGoOutInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.ServiceType != 2 {
		err = gerror.New("服务方式不正确")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.DispatchStatus != "DONE" {
		err = gerror.New("调度状态不正确")
		return
	}

	if !g.IsEmpty(models.TechnicianGoTime) {
		err = gerror.New("技师已出发")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderGoOutFields{
			TechnicianGoTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().TechnicianGoTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   "OUT",
			Remark:      "技师出发",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 技师状态改为工作中
		// 获取关联订单技师ids
		orderTechnicianIds, _ := service.SpaOrderTechnician().GetTechnicianIds(ctx, in.Id)
		if _, err = dao.SpaTechnician.Ctx(ctx).
			WhereIn(dao.SpaTechnician.Columns().Id, orderTechnicianIds).
			Data(g.MapStrAny{
				dao.SpaTechnician.Columns().WorkStatus: "WORKING",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改技师工作状态失败，请稍后重试！")
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "技师已出发",
			"en":    "The technicians have set off.",
			"ja":    "技術者たちは出発しました。",
			"ko":    "기술자들이 출발했습니다.",
			"zh_CN": "技師已出發",
		}
		systemMessageContent := map[string]string{
			"zh":    "您的订单" + models.OrderSn + "技师已出发，正在赶往您所住的民宿",
			"en":    "Your order's technician " + models.OrderSn + " has departed and is on her way to your apartment.",
			"ja":    "ご注文の技術者 " + models.OrderSn + " は出発し、ゲストハウスへ向かっています。",
			"ko":    "귀하의 주문에 대한 기술자 " + models.OrderSn + "이 출발하여 귀하의 게스트하우스로 향하고 있습니다.",
			"zh_CN": "您的訂單" + models.OrderSn + "技師已出發，正在趕往您所住的民宿",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		// 发送确认短信(队列)
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderRemind,
			DataByte: gvar.New(gjson.New(g.Map{
				"orderSn": models.OrderSn,
				"event":   "spa_order_out",
			})).Bytes(),
			Header: nil,
		})

		return
	})

}

func (s *sSpaOrder) Arrive(ctx context.Context, in *input_spa.SpaOrderGoOutInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.ServiceType != 1 {
		err = gerror.New("服务方式不正确")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.DispatchStatus != "DONE" {
		err = gerror.New("调度状态不正确")
		return
	}

	if !g.IsEmpty(models.MemberArriveTime) {
		err = gerror.New("客人已到店")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改客人到店时间
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderArriveFields{
			MemberArriveTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "WAIT_SERVE",
			ActionWay:   "ARRIVE",
			Remark:      "客人到店",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 技师状态改为工作中
		// 获取关联订单技师ids
		orderTechnicianIds, _ := service.SpaOrderTechnician().GetTechnicianIds(ctx, in.Id)
		if _, err = dao.SpaTechnician.Ctx(ctx).
			WhereIn(dao.SpaTechnician.Columns().Id, orderTechnicianIds).
			Data(g.MapStrAny{
				dao.SpaTechnician.Columns().WorkStatus: "WORKING",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改技师工作状态失败，请稍后重试！")
			return
		}

		return
	})

}

func (s *sSpaOrder) StartService(ctx context.Context, in *input_spa.SpaOrderStartServiceInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "WAIT_SERVE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.DispatchStatus != "DONE" {
		err = gerror.New("调度状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderStartServiceFields{
			OrderStatus:     "SERVING",
			ActualStartTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().OrderStatus:     "SERVING",
			dao.SpaOrderGoods.Columns().ActualStartTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "SERVING",
			ActionWay:   "SERVING",
			Remark:      "按摩服务开始",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 技师状态改为工作中
		// 获取关联订单技师ids
		orderTechnicianIds, _ := service.SpaOrderTechnician().GetTechnicianIds(ctx, in.Id)
		if _, err = dao.SpaTechnician.Ctx(ctx).
			WhereIn(dao.SpaTechnician.Columns().Id, orderTechnicianIds).
			Data(g.MapStrAny{
				dao.SpaTechnician.Columns().WorkStatus: "WORKING",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改技师工作状态失败，请稍后重试！")
			return
		}

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "开始服务",
			"en":    "Start Service",
			"ja":    "サービスを開始",
			"ko":    "서비스 시작",
			"zh_CN": "開始服務",
		}
		systemMessageContent := map[string]string{
			"zh":    "开始服务",
			"en":    "Start Service",
			"ja":    "サービスを開始",
			"ko":    "서비스 시작",
			"zh_CN": "開始服務",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		return
	})

}

func (s *sSpaOrder) EndService(ctx context.Context, in *input_spa.SpaOrderEndServiceInp) (err error) {

	var (
		models              *entity.SpaOrder
		OrderTechnicianInfo *entity.SpaOrderTechnician
		TechnicianInfo      *entity.SpaTechnician
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.OrderStatus != "SERVING" {
		err = gerror.New("订单状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderEndServiceFields{
			OrderStatus:   "DONE",
			ActualEndTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().OrderStatus:   "DONE",
			dao.SpaOrderGoods.Columns().ActualEndTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderTechnician.Columns().ActualEndTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "DONE",
			ActionWay:   "DONE",
			Remark:      "按摩服务结束",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 技师状态改为可约
		// 获取关联订单技师ids
		orderTechnicianIds, _ := service.SpaOrderTechnician().GetTechnicianIds(ctx, in.Id)
		if _, err = dao.SpaTechnician.Ctx(ctx).
			WhereIn(dao.SpaTechnician.Columns().Id, orderTechnicianIds).
			Data(g.MapStrAny{
				dao.SpaTechnician.Columns().WorkStatus: "CANORDER",
			}).Update(); err != nil {
			err = gerror.Wrap(err, "修改技师工作状态失败，请稍后重试！")
			return
		}

		// 记录分佣
		for _, orderTechnicianId := range orderTechnicianIds {
			_ = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Where(dao.SpaOrderTechnician.Columns().TechnicianId, orderTechnicianId).Scan(&OrderTechnicianInfo)
			_ = dao.SpaTechnician.Ctx(ctx).WherePri(orderTechnicianId).Scan(&TechnicianInfo)
			cost := OrderTechnicianInfo.SettlementCost
			settlementAmount := OrderTechnicianInfo.OrderAmount
			if !g.IsEmpty(cost) {
				costArr := strings.Split(cost, ",")
				constIsCoupon := false
				constIsBal := false
				for _, costItem := range costArr {
					if gvar.New(costItem).Int() == 1 {
						constIsCoupon = true
						//settlementAmount = settlementAmount - OrderTechnicianInfo.CouponAmount
					}
					if gvar.New(costItem).Int() == 2 {
						constIsBal = true
						//settlementAmount = settlementAmount - OrderTechnicianInfo.BalAmount
					}
				}
				if !constIsCoupon {
					settlementAmount = settlementAmount - models.CouponAmount
				}
				if !constIsBal {
					settlementAmount = settlementAmount - models.BalAmount
				}
			}
			settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(OrderTechnicianInfo.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()
			if _, err = dao.SpaOrderTechnician.Ctx(ctx).TX(tx).Data(g.Map{
				dao.SpaOrderTechnician.Columns().SettlementAmount: settlementAmount,
			}).Where(dao.SpaOrderTechnician.Columns().Id, OrderTechnicianInfo.Id).Update(); err != nil {
				return
			}
		}

		// 转发到返利队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})

		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "服务结束",
			"en":    "Service ended",
			"ja":    "サービス終了",
			"ko":    "서비스 종료",
			"zh_CN": "服務結束",
		}
		systemMessageContent := map[string]string{
			"zh":    "服务结束",
			"en":    "Service ended",
			"ja":    "サービス終了",
			"ko":    "서비스 종료",
			"zh_CN": "服務結束",
		}
		// 查询用户的手机号区号 来判断用户语言
		phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
		var memberLanguage string
		if phoneArea.String() == "+86" {
			memberLanguage = "zh"
		} else if phoneArea.String() == "+81" {
			memberLanguage = "ja"
		} else if phoneArea.String() == "+82" {
			memberLanguage = "ko"
		} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
			memberLanguage = "zh_CN"
		} else {
			memberLanguage = "en"
		}
		pushData := g.MapStrAny{
			"type":    0,
			"orderNo": models.OrderSn,
		}
		pushDataJson, _ := json.Marshal(pushData)
		appPushData := g.MapStrStr{
			"type":  "2",
			"param": string(pushDataJson),
		}
		service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
			SystemMessageTitle:   systemMessageTitle,
			SystemMessageContent: systemMessageContent,
			Scene:                "spa",
			Type:                 "order",
			MemberId:             int(models.MemberId),
			Language:             memberLanguage,
			AppPushData:          appPushData,
			AppLink:              "/spaOrderDetailPage",
			WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              models.OrderSn,
		})

		if models.IsFx == "Y" {
			_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
				ExchangeName: consts.RabbitMQExchangeName,
				QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
				DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
					OrderNo:      models.OrderSn,
					ChangeStatus: "COMPLETE",
				}).MustToJson(),
				Header: nil,
			})
		}

		return
	})

}

func (s *sSpaOrder) Abnormal(ctx context.Context, in *input_spa.SpaOrderAbnormalInp) (err error) {
	var (
		models              *entity.SpaOrder
		OrderTechnicianInfo *entity.SpaOrderTechnician
		TechnicianInfo      *entity.SpaTechnician
	)
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.AbnormalStatus != 2 {
		err = gerror.New("订单异常处理状态不正确")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 修改订单异常状态
		if _, err = s.Model(ctx).TX(tx).
			WherePri(in.Id).Data(input_spa.SpaOrderAbnormalInp{
			AbnormalStatus:     3,
			AbnormalTime:       gtime.Now(),
			AbnormalOperatorId: in.AbnormalOperatorId,
			AbnormalReason:     in.AbnormalReason,
		}).OmitEmptyData().Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 如果当前订单状态为已完成DONE,则不进行后续操作
		if models.OrderStatus == "DONE" {
			return
		}

		// 修改订单状态为已完成DONE
		if _, err = s.Model(ctx).
			WherePri(in.Id).Data(input_spa.SpaOrderEndServiceFields{
			OrderStatus:   "DONE",
			ActualEndTime: gtime.Now(),
		}).Update(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderGoods.Ctx(ctx).Where(dao.SpaOrderGoods.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderGoods.Columns().OrderStatus:   "DONE",
			dao.SpaOrderGoods.Columns().ActualEndTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		if _, err = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Update(g.MapStrAny{
			dao.SpaOrderTechnician.Columns().ActualEndTime: gtime.Now(),
		}); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     in.Id,
			OrderStatus: "DONE",
			ActionWay:   "DONE",
			Remark:      "按摩服务结束",
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		// 技师状态改为可约
		// 获取关联订单技师ids
		orderTechnicianIds, _ := service.SpaOrderTechnician().GetTechnicianIds(ctx, in.Id)
		if !g.IsEmpty(orderTechnicianIds) {
			if _, err = dao.SpaTechnician.Ctx(ctx).
				WhereIn(dao.SpaTechnician.Columns().Id, orderTechnicianIds).
				Data(g.MapStrAny{
					dao.SpaTechnician.Columns().WorkStatus: "CANORDER",
				}).Update(); err != nil {
				err = gerror.Wrap(err, "修改技师工作状态失败，请稍后重试！")
				return
			}
		}

		// 记录分佣
		for _, orderTechnicianId := range orderTechnicianIds {
			_ = dao.SpaOrderTechnician.Ctx(ctx).Where(dao.SpaOrderTechnician.Columns().OrderId, in.Id).Where(dao.SpaOrderTechnician.Columns().TechnicianId, orderTechnicianId).Scan(&OrderTechnicianInfo)
			_ = dao.SpaTechnician.Ctx(ctx).WherePri(orderTechnicianId).Scan(&TechnicianInfo)
			cost := OrderTechnicianInfo.SettlementCost
			settlementAmount := OrderTechnicianInfo.OrderAmount
			if !g.IsEmpty(cost) {
				costArr := strings.Split(cost, ",")
				constIsCoupon := false
				constIsBal := false
				for _, costItem := range costArr {
					if gvar.New(costItem).Int() == 1 {
						constIsCoupon = true
						//settlementAmount = settlementAmount - OrderTechnicianInfo.CouponAmount
					}
					if gvar.New(costItem).Int() == 2 {
						constIsBal = true
						//settlementAmount = settlementAmount - OrderTechnicianInfo.BalAmount
					}
				}
				if !constIsCoupon {
					settlementAmount = settlementAmount - models.CouponAmount
				}
				if !constIsBal {
					settlementAmount = settlementAmount - models.BalAmount
				}
			}
			settlementAmount = decimal.NewFromFloat(settlementAmount).Mul(decimal.NewFromFloat(OrderTechnicianInfo.SettlementRate)).Div(decimal.NewFromFloat(100)).Round(0).InexactFloat64()

			if _, err = dao.SpaOrderTechnician.Ctx(ctx).TX(tx).Data(g.Map{
				dao.SpaOrderTechnician.Columns().SettlementAmount: settlementAmount,
			}).Where(dao.SpaOrderTechnician.Columns().Id, OrderTechnicianInfo.Id).Update(); err != nil {
				return
			}
		}

		// 转发到返利队列
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameRebate,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})
		_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameExp,
			DataByte:     gvar.New(models.OrderSn).Bytes(),
			Header:       nil,
		})

		return
	})

}

// CancelPay 取消订单
func (s *sSpaOrder) CancelPay(ctx context.Context, in *input_spa.SpaOrderCancelPayInp) (err error) {

	var models *entity.SpaOrder
	if err = s.Model(ctx).Where("id", in.Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return
	}

	if models == nil {
		err = gerror.New("订单信息不存在或已被删除")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, models.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if models.PayStatus != "HAVE_PAID" && g.IsEmpty(models.PayTime) {
		err = gerror.New("订单支付状态不正确")
		return
	}

	if models.OrderStatus == "DONE" {
		err = gerror.New("订单状态不正确")
		return
	}

	if models.AdminCancelNum > 0 {
		err = gerror.New("后台不可多次退款")
		return
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 确认失败、全额退款
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			CancelFee         float64 // 退款手续费
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		CancelFee = models.OrderAmount - models.CouponAmount - in.RefundMoney
		CancelFeeCalc := CancelFee

		if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().PayStatus, "DONE").Where(dao.PmsTransaction.Columns().OrderSn, models.OrderSn).OrderDesc(`
			CASE pay_type
				WHEN 'StripeCard' THEN 1
				WHEN 'PaypalCard' THEN 2
				WHEN 'Paypal' THEN 3
				WHEN 'WeChatPay' THEN 4
				WHEN 'Alipay+' THEN 5
				WHEN 'WeChatMiniPay' THEN 6
				WHEN 'BAL' THEN 7
				ELSE 8
			END
		`).Scan(&Transaction); err != nil {
			return
		}
		if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().OrderSn, models.OrderSn).Scan(&TransactionRefund); err != nil {
			return
		}

		for _, v := range Transaction {
			// 可退款金额
			Refundable := v.Amount - v.RefundAmount
			if v.PayType == "BAL" {
				if g.IsEmpty(CancelFeeCalc) {
					RefundBalance += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundBalance += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundBalance += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			} else if v.PayType == "COUPON" {

			} else {
				if g.IsEmpty(CancelFeeCalc) {
					RefundFee += Refundable
				} else if CancelFeeCalc > Refundable {
					RefundFee += 0
					CancelFeeCalc = CancelFeeCalc - Refundable
				} else {
					RefundFee += Refundable - CancelFeeCalc
					CancelFeeCalc = 0
				}
			}
		}
		// 退款金额
		RefundAmount := RefundBalance + RefundFee

		OrderTotalRefundAmount := RefundAmount + models.RefundAmount
		OrderTotalRefundBalance := RefundBalance + models.RefundBalAmount

		AdminOrderTotalRefundAmount := RefundAmount + models.AdminRefundAmount
		AdminOrderTotalRefundBalance := RefundBalance + models.AdminRefundBalAmount

		// 修改订单状态
		RefundStatus := "DONE"
		if models.OrderAmount > OrderTotalRefundAmount {
			RefundStatus = "PART"
		}
		if in.RefundType == 1 {
			// 仅退款
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				//dao.SpaOrder.Columns().RefundFee:          CancelFee,
				dao.SpaOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.SpaOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.SpaOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.SpaOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.SpaOrder.Columns().RefundCouponAmount:   0,
				dao.SpaOrder.Columns().RefundStatus:         RefundStatus,
				dao.SpaOrder.Columns().RefundTime:           gtime.Now(),
				dao.SpaOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 更新服务预定量和预定金额
			if _, err = dao.SpaService.Ctx(ctx).TX(tx).Where(dao.SpaService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
				dao.SpaService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
				return
			}

			// 更新项目预定量和预定金额
			if _, err = dao.SpaServiceGoods.Ctx(ctx).TX(tx).Where(dao.SpaServiceGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
				dao.SpaServiceGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新项目信息失败，请稍后重试！")
				return
			}

			// 更新技师预定量和预定金额
			if !g.IsEmpty(models.TechnicianIds) {
				for _, Technician := range strings.Split(models.TechnicianIds, ",") {
					if !g.IsEmpty(Technician) {
						TechnicianId, _ := strconv.ParseInt(Technician, 10, 64)
						// 更新技师预定量和预定金额
						if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, TechnicianId).Update(g.MapStrAny{
							dao.SpaTechnician.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount/float64(models.GoodsNum))),
						}); err != nil {
							err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
							return
						}
					}
				}
			}
		} else {
			// 退款并取消
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				//dao.SpaOrder.Columns().RefundFee:          CancelFee,
				dao.SpaOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.SpaOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.SpaOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.SpaOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.SpaOrder.Columns().RefundCouponAmount:   0,
				dao.SpaOrder.Columns().RefundStatus:         RefundStatus,
				dao.SpaOrder.Columns().RefundTime:           gtime.Now(),
				dao.SpaOrder.Columns().OrderStatus:          "CANCEL",
				dao.SpaOrder.Columns().PayStatus:            "REFUND",
				dao.SpaOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
				dao.SpaOrder.Columns().CancelTime:           gtime.Now(),
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

			// 更新服务预定量和预定金额
			if _, err = dao.SpaService.Ctx(ctx).TX(tx).Where(dao.SpaService.Columns().Id, models.ServiceId).Update(g.MapStrAny{
				dao.SpaService.Columns().PayOrderNum:    gdb.Raw("pay_order_num-1"),
				dao.SpaService.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新服务信息失败，请稍后重试！")
				return
			}

			// 更新项目预定量和预定金额
			if _, err = dao.SpaServiceGoods.Ctx(ctx).TX(tx).Where(dao.SpaServiceGoods.Columns().Id, models.GoodsId).Update(g.MapStrAny{
				dao.SpaServiceGoods.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", models.GoodsNum)),
				dao.SpaServiceGoods.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount)),
			}); err != nil {
				err = gerror.Wrap(err, "更新项目信息失败，请稍后重试！")
				return
			}

			// 更新技师预定量和预定金额
			if !g.IsEmpty(models.TechnicianIds) {
				for _, Technician := range strings.Split(models.TechnicianIds, ",") {
					if !g.IsEmpty(Technician) {
						TechnicianId, _ := strconv.ParseInt(Technician, 10, 64)
						// 更新技师预定量和预定金额
						if _, err = dao.SpaTechnician.Ctx(ctx).Where(dao.SpaTechnician.Columns().Id, TechnicianId).Update(g.MapStrAny{
							dao.SpaTechnician.Columns().PayOrderNum:    gdb.Raw(fmt.Sprintf("pay_order_num-%d", 1)),
							dao.SpaTechnician.Columns().PayOrderAmount: gdb.Raw(fmt.Sprintf("pay_order_amount-%f", RefundAmount/float64(models.GoodsNum))),
						}); err != nil {
							err = gerror.Wrap(err, "更新技师信息失败，请稍后重试！")
							return
						}
					}
				}
			}
		}

		if _, err = dao.SpaOrder.Ctx(ctx).TX(tx).WherePri(in.Id).Update(g.MapStrAny{
			dao.SpaOrder.Columns().AdminCancelNum: gdb.Raw("admin_cancel_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新失败，请稍后重试！")
			return
		}

		if RefundAmount > 0 {
			// 退款
			err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
				OrderSn:      models.OrderSn,
				RefundAmount: RefundAmount,
				OperateType:  "ADMIN",
				Remark:       in.AdminCancelReason,
				OperateId:    int(contexts.GetUserId(ctx)),
			}, tx)
			if err != nil {
				err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
			OrderId:     int(in.Id),
			OrderStatus: models.OrderStatus,
			ActionWay:   "ADMIN_REFUND",
			Remark:      fmt.Sprintf("后台退款，原因：%s", in.AdminCancelReason),
			OperateType: "ADMIN",
			OperateId:   int(contexts.GetUserId(ctx)),
		}); err != nil {
			return err
		}

		if in.RefundType == 2 {
			// 订单日志
			if _, err = dao.SpaOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.SpaOrderLog{
				OrderId:     int(in.Id),
				OrderStatus: "CANCEL",
				ActionWay:   "CANCEL",
				Remark:      "后台订单取消",
				OperateType: "ADMIN",
				OperateId:   int(contexts.GetUserId(ctx)),
			}); err != nil {
				return err
			}

			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单已取消",
				"en":    "Order has been canceled",
				"ja":    "注文はキャンセルされました",
				"ko":    "주문이 취소되었습니다",
				"zh_CN": "訂單已取消",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Order has been canceled, and a refund of " + gvar.New(RefundFee).String() + "JPY has been issued",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消，并成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Order has been canceled, and " + gvar.New(RefundBalance).String() + " points have been refunded",
			// 		"ja":    "注文はキャンセルされ、" + gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    "주문이 취소되었으며 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單已取消，並成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单已取消",
			// 		"en":    "Order has been canceled",
			// 		"ja":    "注文はキャンセルされました",
			// 		"ko":    "주문이 취소되었습니다",
			// 		"zh_CN": "訂單已取消",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消，并成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Order has been canceled, and a refund of " + gvar.New(RefundAmount).String() + "JPY has been issued",
					"ja":    "注文はキャンセルされ、" + gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    "주문이 취소되었으며 " + gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單已取消，並成功退款 " + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单已取消",
					"en":    "Order has been canceled",
					"ja":    "注文はキャンセルされました",
					"ko":    "주문이 취소되었습니다",
					"zh_CN": "訂單已取消",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
			var memberLanguage string
			if phoneArea.String() == "+86" {
				memberLanguage = "zh"
			} else if phoneArea.String() == "+81" {
				memberLanguage = "ja"
			} else if phoneArea.String() == "+82" {
				memberLanguage = "ko"
			} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
				memberLanguage = "zh_CN"
			} else {
				memberLanguage = "en"
			}
			pushData := g.MapStrAny{
				"type":    0,
				"orderNo": models.OrderSn,
			}
			pushDataJson, _ := json.Marshal(pushData)
			appPushData := g.MapStrStr{
				"type":  "2",
				"param": string(pushDataJson),
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "spa",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/spaOrderDetailPage",
				WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		} else {
			// 发送到消息队列
			systemMessageTitle := map[string]string{
				"zh":    "订单成功退款",
				"en":    "Refund Successful",
				"ja":    "返金が正常に完了しました",
				"ko":    "환불이 성공적으로 완료되었습니다",
				"zh_CN": "訂單成功退款",
			}

			var systemMessageContent map[string]string

			// if RefundFee > 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY和" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY and " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY と " + gvar.New(RefundBalance).String() + " ポイントの返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY와 " + gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款 " + gvar.New(RefundFee).String() + "JPY 及 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else if RefundFee > 0 && RefundBalance <= 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 		"en":    "Refund of " + gvar.New(RefundFee).String() + "JPY completed",
			// 		"ja":    gvar.New(RefundFee).String() + "JPY の返金が完了しました",
			// 		"ko":    gvar.New(RefundFee).String() + "JPY가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退款" + gvar.New(RefundFee).String() + "JPY",
			// 	}
			// } else if RefundFee <= 0 && RefundBalance > 0 {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退还" + gvar.New(RefundBalance).String() + "积分",
			// 		"en":    "Refund of " + gvar.New(RefundBalance).String() + " points completed",
			// 		"ja":    gvar.New(RefundBalance).String() + "ポイントの返還が完了しました",
			// 		"ko":    gvar.New(RefundBalance).String() + "포인트가 성공적으로 환불되었습니다",
			// 		"zh_CN": "訂單成功退還 " + gvar.New(RefundBalance).String() + " 積分",
			// 	}
			// } else {
			// 	systemMessageContent = map[string]string{
			// 		"zh":    "订单成功退款",
			// 		"en":    "Refund Successful",
			// 		"ja":    "返金が正常に完了しました",
			// 		"ko":    "환불이 성공적으로 완료되었습니다",
			// 		"zh_CN": "訂單成功退款",
			// 	}
			// }

			if RefundAmount > 0 {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款" + gvar.New(RefundAmount).String() + "JPY",
					"en":    "Refund of " + gvar.New(RefundAmount).String() + "JPY completed",
					"ja":    gvar.New(RefundAmount).String() + "JPY の返金が完了しました",
					"ko":    gvar.New(RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
					"zh_CN": "訂單成功退款" + gvar.New(RefundAmount).String() + "JPY",
				}
			} else {
				systemMessageContent = map[string]string{
					"zh":    "订单成功退款",
					"en":    "Refund Successful",
					"ja":    "返金が正常に完了しました",
					"ko":    "환불이 성공적으로 완료되었습니다",
					"zh_CN": "訂單成功退款",
				}
			}

			// 查询用户的手机号区号 来判断用户语言
			phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(models.MemberId).Value()
			var memberLanguage string
			if phoneArea.String() == "+86" {
				memberLanguage = "zh"
			} else if phoneArea.String() == "+81" {
				memberLanguage = "ja"
			} else if phoneArea.String() == "+82" {
				memberLanguage = "ko"
			} else if phoneArea.String() == "+886" || phoneArea.String() == "+852" || phoneArea.String() == "+853" {
				memberLanguage = "zh_CN"
			} else {
				memberLanguage = "en"
			}
			pushData := g.MapStrAny{
				"type":    0,
				"orderNo": models.OrderSn,
			}
			pushDataJson, _ := json.Marshal(pushData)
			appPushData := g.MapStrStr{
				"type":  "2",
				"param": string(pushDataJson),
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "spa",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/spaOrderDetailPage",
				WxLink:               fmt.Sprintf("/subpackages/spa/pages/order?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})
		}

		return
	})

}

func (s *sSpaOrder) ExportOrder(ctx context.Context, in *input_spa.SpaOrderExportInp) (err error) {
	var (
		lastInsertId int64
	)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		resultJson, _ := json.Marshal(in)

		if lastInsertId, err = dao.OrderExport.Ctx(ctx).
			Data(entity.OrderExport{
				Scene:     3,
				Condition: string(resultJson),
			}).OmitEmptyData().InsertAndGetId(); err != nil {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
		}

		if lastInsertId < 1 {
			err = gerror.Wrap(err, "操作失败，请稍后重试！")
			return
		}

		return
	})

	if err != nil {
		return
	}

	// 导出(队列)
	_ = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
		ExchangeName: consts.RabbitMQExchangeName,
		QueueName:    consts.RabbitMQQueueNameOrderExport,
		DataByte:     gvar.New(lastInsertId).Bytes(),
		Header:       nil,
	})

	return
}

func (s *sSpaOrder) StartExport(ctx context.Context, in *input_spa.SpaOrderExportInp) (path string, err error) {
	var (
		ChangeList  []*input_spa.SpaOrderExportModel
		OrderListIn *input_spa.SpaOrderListInp
	)

	OrderListIn = &input_spa.SpaOrderListInp{
		PageReq: input_form.PageReq{
			Pagination: false,
		},
		MemberId:       in.MemberId,
		OrderSn:        in.OrderSn,
		TechnicianName: in.TechnicianName,
		MemberSearch:   in.MemberSearch,
		OrderStatus:    in.OrderStatus,
	}
	list, _, err := s.List(ctx, OrderListIn)
	if err != nil {
		return
	}
	for _, item := range list {
		itemData := &input_spa.SpaOrderExportModel{
			OrderSn:           item.OrderSn,
			MemberNo:          item.PmsMemberMemberNo,
			BookingName:       item.BookingName,
			BookStartTime:     item.BookDate + " " + item.BookTime,
			OrderTime:         item.CreatedAt.Format("Y-m-d H:i:s"),
			TotalAmount:       item.OrderAmount,
			RefundTotalAmount: item.RefundAmount,
			RefundTime:        item.RefundTime.Format("Y-m-d H:i:s"),
		}
		if item.ServiceType == 1 {
			itemData.ServiceType = "到店"
		} else {
			itemData.ServiceType = "上门"
		}
		for _, TransactionDetail := range item.TransactionDetail {
			if TransactionDetail.PayStatus == "DONE" {
				switch TransactionDetail.PayType {
				case "BAL":
					itemData.PointsPayment += TransactionDetail.PayAmount
					break
				case "COUPON":
					itemData.CouponPayment += TransactionDetail.PayAmount
					break
				case "WeChatPay":
					itemData.PaycloudWechatPay += TransactionDetail.PayAmount
					break
				case "Alipay+":
					itemData.PaycloudAlipayPay += TransactionDetail.PayAmount
					break
				case "Paypal":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "PaypalCard":
					itemData.PaypelCreditPay += TransactionDetail.PayAmount
					break
				case "StripeCard":
					itemData.StripeCreditPay += TransactionDetail.PayAmount
					break
				case "WeChatMiniPay":
					itemData.MlilifeWeChatMiniPay += TransactionDetail.PayAmount
					break
				}
			}
		}

		for _, TransactionRefundDetail := range item.TransactionRefundDetail {
			switch TransactionRefundDetail.RefundType {
			case "BAL":
				itemData.PointsRefund += TransactionRefundDetail.RefundAmount
				break
			case "COUPON":
				itemData.CouponRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatPay":
				itemData.PaycloudWechatRefund += TransactionRefundDetail.RefundAmount
				break
			case "Alipay+":
				itemData.PaycloudAlipayRefund += TransactionRefundDetail.RefundAmount
				break
			case "Paypal":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "PaypalCard":
				itemData.PaypelCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "StripeCard":
				itemData.StripeCreditRefund += TransactionRefundDetail.RefundAmount
				break
			case "WeChatMiniPay":
				itemData.MlilifeWeChatMiniRefund += TransactionRefundDetail.RefundAmount
				break
			}
		}
		ChangeList = append(ChangeList, itemData)
	}
	tags, err := convert.GetEntityDescTags(input_spa.SpaOrderExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出按摩订单-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("按摩订单")
		exports   []input_spa.SpaOrderExportModel
	)

	if err = gconv.Scan(ChangeList, &exports); err != nil {
		return
	}

	path, err = excel.ExportByStructsFile(ctx, tags, exports, fileName, sheetName)

	return
}

func (s *sSpaOrder) ExportList(ctx context.Context, in *input_spa.SpaOrderExportListInp) (list []*input_spa.SpaOrderExportListModel, totalCount int, err error) {
	mod := dao.OrderExport.Ctx(ctx)

	mod = mod.Fields(input_spa.SpaOrderExportListModel{})

	mod = mod.Where(dao.OrderExport.Columns().Scene, 3)

	mod = mod.Page(in.Page, in.PerPage)

	mod = mod.OrderDesc(dao.OrderExport.Columns().Id)

	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取列表失败，请稍后重试！")
		return
	}

	for _, v := range list {
		v.Path = g.Cfg().MustGet(ctx, "localUploadDomain").String() + "/" + v.Path
	}

	return
}
