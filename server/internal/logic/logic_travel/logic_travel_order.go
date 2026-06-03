package logic_travel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/contexts"
	"APT/internal/library/hgorm/handler"
	hook2 "APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_refund"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"APT/utility/encrypt"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type sTravelOrder struct{}

func NewTravelOrder() *sTravelOrder {
	return &sTravelOrder{}
}

func init() {
	service.RegisterTravelOrder(NewTravelOrder())
}

// Model 订单ORM模型
func (s *sTravelOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TravelOrder.Ctx(ctx), option...)
}

// List 获取订单列表
func (s *sTravelOrder) List(ctx context.Context, in *input_travel.TravelOrderListInp) (list []*input_travel.TravelOrderListModel, totalCount int, err error) {
	mod := s.Model(ctx).Unscoped().WithAll()

	mod = mod.FieldsPrefix(dao.TravelOrder.Table(), input_travel.TravelOrderListModel{})
	mod = mod.Fields(fmt.Sprintf("`%s`.`%s` as `pmsMemberMemberNo`", dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo))
	mod = mod.Fields(fmt.Sprintf("IF(`%s`.`%s` IS NOT NULL, 1, 0) as `member_deleted`", dao.PmsMember.Table(), dao.PmsMember.Columns().DeletedAt))

	mod = mod.LeftJoin(dao.PmsMember.Table(), fmt.Sprintf("`%s`.`%s` = `%s`.`%s`", dao.TravelOrder.Table(), dao.TravelOrder.Columns().MemberId, dao.PmsMember.Table(), dao.PmsMember.Columns().Id))

	if !g.IsEmpty(in.OrderSn) {
		mod = mod.WhereLike(dao.TravelOrder.Columns().OrderSn, "%"+in.OrderSn+"%")
	}
	if !g.IsEmpty(in.OrderStatus) {
		mod = mod.Where(dao.TravelOrder.Columns().OrderStatus, in.OrderStatus)
	}
	if !g.IsEmpty(in.MemberId) {
		mod = mod.Where(dao.TravelOrder.Columns().MemberId, in.MemberId)
	}
	if !g.IsEmpty(in.MemberSearch) {
		mod = mod.Where(mod.Builder().
			WherePrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().Id, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.PmsMember.Table(), dao.PmsMember.Columns().MemberNo, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.TravelOrder.Table(), dao.TravelOrder.Columns().BookingName, "%"+in.MemberSearch+"%").
			WhereOrPrefixLike(dao.TravelOrder.Table(), dao.TravelOrder.Columns().BookingMobile, "%"+in.MemberSearch+"%"))
	}

	if !g.IsEmpty(in.ProductName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.ProductName, "title")
		if err == nil {
			productsIds, _ := service.TravelProduct().GetIds(ctx, uuIds)
			mod = mod.WhereIn(dao.TravelOrder.Columns().ProductId, productsIds)
		}
	}

	if !g.IsEmpty(in.SkuName) {

		uuIds, err := service.BasicsLanguage().GetUuids(ctx, in.SkuName, "name")
		if err == nil {
			skuIds, _ := service.TravelProduct().GetSkuIds(ctx, uuIds)
			mod = mod.WhereIn(dao.TravelOrder.Columns().SkuId, skuIds)
		}
	}

	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TravelOrder.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}
	if len(in.BookDate) == 2 {
		mod = mod.WhereBetween(dao.TravelOrder.Columns().BookDate, gtime.New(in.BookDate[0]).Format("Y-m-d"), gtime.New(in.BookDate[1]).Format("Y-m-d"))
	}
	if len(in.VerifyTime) == 2 {
		mod = mod.WhereBetween(dao.TravelOrder.Columns().VerifyTime, in.VerifyTime[0], in.VerifyTime[1])
	}

	mod = mod.OrderDesc(dao.TravelOrder.Columns().Id)
	mod = mod.Hook(hook2.PmsFindLanguageValueHook)

	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
		if err = mod.ScanAndCount(&list, &totalCount, false); err != nil && !errors.Is(err, sql.ErrNoRows) {
			err = gerror.Wrap(err, "获取订单列表失败，请稍后重试！")
		}
		return
	}
	if err = mod.Scan(&list); err != nil && !errors.Is(err, sql.ErrNoRows) {
		err = gerror.Wrap(err, "获取订单列表失败，请稍后重试！")
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

// View 获取订单详情
func (s *sTravelOrder) View(ctx context.Context, in *input_travel.TravelOrderViewInp) (res *input_travel.TravelOrderViewModel, err error) {
	if err = s.Model(ctx).Unscoped().Hook(hook2.PmsFindLanguageValueHook).WithAll().Where(dao.TravelOrder.Columns().Id, in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取订单详情失败，请稍后重试！")
		return
	}
	if res == nil {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
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

	// 加载已删除的产品
	if res.ProductId > 0 && res.ProductInfo == nil {
		var product *entity.TravelProduct
		if err = dao.TravelProduct.Ctx(ctx).Unscoped().Where(dao.TravelProduct.Columns().Id, res.ProductId).Scan(&product); err == nil && product != nil {
			res.ProductInfo = &struct {
				gmeta.Meta `orm:"table:hg_travel_product"`
				*entity.TravelProduct
			}{TravelProduct: product}
		}
	}

	// 加载已删除的核销人员
	if res.VerifyStaffId > 0 && res.VerifyStaffInfo == nil {
		var staff *entity.TravelVerifyStaff
		if err = dao.TravelVerifyStaff.Ctx(ctx).Unscoped().Where(dao.TravelVerifyStaff.Columns().Id, res.VerifyStaffId).Scan(&staff); err == nil && staff != nil {
			res.VerifyStaffInfo = &struct {
				gmeta.Meta `orm:"table:hg_travel_verify_staff"`
				*entity.TravelVerifyStaff
			}{TravelVerifyStaff: staff}
		}
	}

	// 创建新的切片来存储过滤后的日志
	var filteredLogList []*struct {
		gmeta.Meta  `orm:"table:hg_travel_order_log"`
		OrderId     int         `json:"orderId"     description:"订单ID"`
		ActionWay   string      `json:"actionWay"   description:"操作名"`
		Remark      string      `json:"remark"      description:"备注"`
		OperateType string      `json:"operateType" description:"操作员类型"`
		OperateId   int         `json:"operateId"   description:"操作员ID"`
		OperateName string      `json:"operateName"      dc:"操作人姓名"`
		CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
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

		if v.OperateType == "STAFF" {
			var StaffInfo *entity.TravelVerifyStaff
			if err = dao.TravelVerifyStaff.Ctx(ctx).Unscoped().Where(dao.TravelVerifyStaff.Columns().Id, v.OperateId).Scan(&StaffInfo); err != nil {
				return
			}
			if StaffInfo != nil {
				v.OperateName = StaffInfo.Name
			}
		}

		// 将处理后的记录添加到过滤后的列表中
		filteredLogList = append(filteredLogList, v)
	}

	// 更新结果中的日志列表
	res.LogList = filteredLogList

	return
}

// Refund 订单退款
func (s *sTravelOrder) Refund(ctx context.Context, in *input_travel.TravelOrderRefundInp) (err error) {

	var models *entity.TravelOrder
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

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// 计算退款金额
		var (
			Transaction       []*entity.PmsTransaction
			TransactionRefund []*entity.PmsTransactionRefund
			CancelFee         float64 // 退款手续费
			RefundBalance     float64 // 可退款积分
			RefundFee         float64 // 第三方支付
		)
		CancelFee = models.OrderAmount - models.CouponAmount - models.RefundAmount - in.RefundMoney // 订单金额100， 优惠券10， 余额支付30，三方支付 60， 申请退款5， fee 85
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
				dao.TravelOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.TravelOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.TravelOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.TravelOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.TravelOrder.Columns().RefundCouponAmount:   0,
				dao.TravelOrder.Columns().RefundStatus:         RefundStatus,
				dao.TravelOrder.Columns().RefundTime:           gtime.Now(),
				dao.TravelOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}

		} else {
			// 退款并取消
			if _, err = s.Model(ctx).TX(tx).
				WherePri(in.Id).Data(g.MapStrAny{
				dao.TravelOrder.Columns().RefundAmount:         OrderTotalRefundAmount,
				dao.TravelOrder.Columns().RefundBalAmount:      OrderTotalRefundBalance,
				dao.TravelOrder.Columns().AdminRefundAmount:    AdminOrderTotalRefundAmount,
				dao.TravelOrder.Columns().AdminRefundBalAmount: AdminOrderTotalRefundBalance,
				dao.TravelOrder.Columns().RefundCouponAmount:   0,
				dao.TravelOrder.Columns().RefundStatus:         RefundStatus,
				dao.TravelOrder.Columns().RefundTime:           gtime.Now(),
				dao.TravelOrder.Columns().AdminCancelReason:    in.AdminCancelReason,
				dao.TravelOrder.Columns().OrderStatus:          "CANCEL",
				dao.TravelOrder.Columns().PayStatus:            "REFUND",
				dao.TravelOrder.Columns().CancelTime:           gtime.Now(),
			}).Update(); err != nil {
				err = gerror.Wrap(err, "操作失败，请稍后重试！")
				return
			}
		}
		if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).WherePri(in.Id).Update(g.MapStrAny{
			dao.TravelOrder.Columns().AdminCancelNum: gdb.Raw("admin_cancel_num+1"),
		}); err != nil {
			err = gerror.Wrap(err, "更新失败，请稍后重试！")
			return
		}

		if RefundAmount > 0 {
			// 退款
			err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
				OrderSn:      models.OrderSn,
				RefundAmount: RefundAmount,
				Remark:       in.AdminCancelReason,
				OperateType:  "ADMIN",
				OperateId:    int(contexts.GetUserId(ctx)),
			}, tx)
			if err != nil {
				err = gerror.Wrap(err, "退款操作失败，请稍后重试！")
				return
			}
		}

		// 订单日志
		if _, err = dao.TravelOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.TravelOrderLog{
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
			if _, err = dao.TravelOrderLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(&entity.TravelOrderLog{
				OrderId:     int(in.Id),
				OrderStatus: "CANCEL",
				ActionWay:   "CANCEL",
				Remark:      "后台订单取消",
				OperateType: "ADMIN",
				OperateId:   int(contexts.GetUserId(ctx)),
			}); err != nil {
				return err
			}

			// 发送到消息队列 todo
			/*systemMessageTitle := map[string]string{
				"zh":    "订单已取消",
				"en":    "Order has been canceled",
				"ja":    "注文はキャンセルされました",
				"ko":    "주문이 취소되었습니다",
				"zh_CN": "訂單已取消",
			}

			var systemMessageContent map[string]string

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
			appPushData := g.MapStrStr{
				"type":   "1",
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "car",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/transferDetailScene",
				WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})*/
		} else {
			// 发送到消息队列
			/*systemMessageTitle := map[string]string{
				"zh":    "订单成功退款",
				"en":    "Refund Successful",
				"ja":    "返金が正常に完了しました",
				"ko":    "환불이 성공적으로 완료되었습니다",
				"zh_CN": "訂單成功退款",
			}

			var systemMessageContent map[string]string
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
			appPushData := g.MapStrStr{
				"type":   "1",
				"string": models.OrderSn,
			}
			service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
				SystemMessageTitle:   systemMessageTitle,
				SystemMessageContent: systemMessageContent,
				Scene:                "car",
				Type:                 "order",
				MemberId:             int(models.MemberId),
				Language:             memberLanguage,
				AppPushData:          appPushData,
				AppLink:              "/transferDetailScene",
				WxLink:               fmt.Sprintf("/pages/car/order-detail?orderSn=%s", models.OrderSn),
				EnablePush:           true,
				EnableSms:            false,
				PushTitle:            systemMessageTitle[memberLanguage],
				PushContent:          systemMessageContent[memberLanguage],
				OperatorId:           int(contexts.GetUserId(ctx)),
				OperatorRole:         "ADMIN",
				OrderSn:              models.OrderSn,
			})*/
		}

		return
	})

}

// OrderExpiration 订单过期
func (s *sTravelOrder) OrderExpiration(ctx context.Context, OrderSn string) (err error) {
	var (
		tx                   gdb.TX
		TravelOrder          entity.TravelOrder
		PmsCouponTransaction *entity.PmsTransaction
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
	if err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.TravelOrder.Columns().OrderSn:     OrderSn,
		dao.TravelOrder.Columns().OrderStatus: "WAIT_PAY",
	}).Scan(&TravelOrder); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(TravelOrder) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	// 更新当前订单状态-取消
	if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.TravelOrder.Columns().OrderSn:     OrderSn,
		dao.TravelOrder.Columns().OrderStatus: "WAIT_PAY",
	}).Data(g.MapStrAny{
		dao.TravelOrder.Columns().PayStatus:   "CANCEL",
		dao.TravelOrder.Columns().OrderStatus: "CANCEL",
		dao.TravelOrder.Columns().CancelTime:  gtime.Now(),
	}).Update(); err != nil {
		return
	}

	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).
		Where(dao.PmsTransaction.Columns().OrderSn, OrderSn).
		Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
		Where(dao.PmsTransaction.Columns().PayType, "COUPON").
		Scan(&PmsCouponTransaction); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if !g.IsEmpty(PmsCouponTransaction) {
		// 释放优惠券
		if _, err = dao.PmsCoupon.Ctx(ctx).TX(tx).
			Where(dao.PmsCoupon.Columns().Id, PmsCouponTransaction.CouponId).
			Update(g.MapStrAny{
				dao.PmsCoupon.Columns().State:   1,
				dao.PmsCoupon.Columns().UseTime: nil,
			}); err != nil {
			return
		}
		// 修改支付订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).OmitEmptyData().
			Where(dao.PmsTransaction.Columns().TransactionSn, PmsCouponTransaction.TransactionSn).
			Where(dao.PmsTransaction.Columns().PayStatus, "DONE").
			Update(entity.PmsTransaction{
				PayAmount: 0,
				PayStatus: "CANCEL",
				PayTime:   gtime.Now(),
			}); err != nil {
			return
		}
	}

	// 更新付费信息状态-取消
	if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   OrderSn,
		dao.PmsTransaction.Columns().PayStatus: "WAIT",
	}).Data(g.MapStrAny{
		dao.PmsTransaction.Columns().PayStatus: "CANCEL",
	}).Update(); err != nil {
		return
	}

	// 订单取消
	if _, err = dao.TravelOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.CabinetOrderLog{
		OrderId:     int(TravelOrder.Id),
		OrderStatus: "CANCEL",
		ActionWay:   "CANCEL",
		Remark:      "订单未支付取消",
		OperateType: "SYSTEM",
	}); err != nil {
		return
	}

	/*if TravelOrder.IsFx == "Y" {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
			DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
				OrderNo:      TravelOrder.OrderSn,
				ChangeStatus: "CANCEL",
			}).MustToJson(),
			Header: nil,
		}); err != nil {
			g.Log().Error(ctx, "发送分销订单消息变更失败", err)
			err = nil
		}
	}*/

	return
}

// OrderOverdue 订单过期
func (s *sTravelOrder) OrderOverdue(ctx context.Context, OrderSn string) (err error) {
	var (
		tx          gdb.TX
		TravelOrder entity.TravelOrder
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

	// 查询订单
	if err = dao.TravelOrder.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.TravelOrder.Columns().OrderSn: OrderSn,
	}).Scan(&TravelOrder); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}

	// 订单不存在
	if g.IsEmpty(TravelOrder) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	// 订单状态不为待核销，直接返回
	if TravelOrder.OrderStatus != "WAIT_VERIFY" {
		return
	}

	// 更新当前订单状态-逾期
	if _, err = dao.TravelOrder.Ctx(ctx).TX(tx).WherePri(TravelOrder.Id).Data(g.MapStrAny{
		dao.TravelOrder.Columns().OrderStatus: "OVERDUE",
	}).Update(); err != nil {
		return
	}

	// 创建订单日志
	if _, err = dao.TravelOrderLog.Ctx(ctx).OmitEmptyData().Insert(&entity.TravelOrderLog{
		OrderId:     int(TravelOrder.Id),
		OrderStatus: "OVERDUE",
		ActionWay:   "OVERDUE",
		Remark:      "订单已逾期",
		OperateType: "SYSTEM",
	}); err != nil {
		return
	}

	// TODO 发消息队列

	return
}

func (s *sTravelOrder) RefreshCode(ctx context.Context, in *input_travel.TravelOrderRefreshCodeInp) (res *input_travel.TravelOrderRefreshCodeModel, err error) {

	var TravelOrder *struct {
		Id          int
		OrderSn     string
		MemberId    int
		OrderStatus string
	}
	if err = s.Model(ctx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Scan(&TravelOrder); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "get_order_info_failed"))
		return
	}
	if TravelOrder.MemberId != in.MemberId {
		// 会员信息不匹配
		err = gerror.New(gi18n.T(ctx, "member_information_mismatch"))
		return
	}

	res = new(input_travel.TravelOrderRefreshCodeModel)

	res.OrderStatus = TravelOrder.OrderStatus

	// 生成券码，采用aes加密，将couponNo|会员ID|时间戳进行加密
	timestamp := gtime.Now().Unix()
	plainText := fmt.Sprintf("%s|%d|%d", TravelOrder.OrderSn, TravelOrder.MemberId, timestamp)

	// 使用AES加密生成动态券码
	res.Code = encrypt.MustAesECBEncryptToString(plainText, string(consts.RequestEncryptKey))

	return
}
