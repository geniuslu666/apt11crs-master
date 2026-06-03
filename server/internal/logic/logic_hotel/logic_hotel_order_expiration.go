package logic_hotel

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/h5FxPay"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderExpiration 订单过期
func (s *sHotelService) OrderExpiration(ctx context.Context, OrderSn string) (err error) {
	var (
		tx                   gdb.TX
		PmsAppStay           entity.PmsAppStay
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
	if err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn:     OrderSn,
		dao.PmsAppStay.Columns().OrderStatus: "WAIT_PAY",
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		err = gerror.New("该订单无需处理")
		return
	}
	// 更新当前住宿订单状态-取消
	if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn:     OrderSn,
		dao.PmsAppStay.Columns().OrderStatus: "WAIT_PAY",
	}).Data(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderStatus: "CANCEL",
		dao.PmsAppStay.Columns().CancelTime:  gtime.Now(),
	}).Update(); err != nil {
		return
	}

	// 更新房间预约状态-取消
	if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsAppReservation.Columns().OrderSn:     OrderSn,
		dao.PmsAppReservation.Columns().OrderStatus: "WAIT_PAY",
	}).Data(g.MapStrAny{
		dao.PmsAppReservation.Columns().Status: "cancelled",
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

	// 酒店订单取消
	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   "CANCEL",
		Remark:      "订单取消",
		OperateType: "SYSTEM",
	}); err != nil {
		return
	}

	// 发送到消息队列
	systemMessageTitle := map[string]string{
		"zh":    "订单已取消",
		"en":    "The order has been cancelled.",
		"ja":    "注文はキャンセルされました。",
		"ko":    "주문이 취소되었습니다.",
		"zh_CN": "訂單已取消",
	}
	systemMessageContent := map[string]string{
		"zh":    OrderSn + "订单已取消",
		"en":    "Order " + OrderSn + " has been cancelled.",
		"ja":    "注文" + OrderSn + "はキャンセルされました。",
		"ko":    "주문 " + OrderSn + "은 취소되었습니다.",
		"zh_CN": OrderSn + "訂單已取消",
	}
	// 查询用户的手机号区号 来判断用户语言
	phoneArea, _ := dao.PmsMember.Ctx(ctx).Fields(dao.PmsMember.Columns().PhoneArea).WherePri(PmsAppStay.MemberId).Value()
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
		"orderSn": OrderSn,
	}
	pushDataJson, _ := json.Marshal(pushData)
	appPushData := g.MapStrStr{
		"type":  "2",
		"param": string(pushDataJson),
	}
	service.BasicsSystemMessage().SendMessage(ctx, &input_basics.SendMessageInp{
		SystemMessageTitle:   systemMessageTitle,
		SystemMessageContent: systemMessageContent,
		Scene:                "hotel",
		Type:                 "order",
		MemberId:             int(PmsAppStay.MemberId),
		Language:             memberLanguage,
		AppPushData:          appPushData,
		AppLink:              "/order/order-detail",
		WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", OrderSn),
		EnablePush:           true,
		EnableSms:            false,
		PushTitle:            systemMessageTitle[memberLanguage],
		PushContent:          systemMessageContent[memberLanguage],
		OperatorId:           0,
		OperatorRole:         "SYSTEM",
		OrderSn:              OrderSn,
	})

	if PmsAppStay.IsFx == "Y" {
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameFxChangeOrderPush,
			DataByte: gjson.New(&h5FxPay.ChangeOrderPushRequest{
				OrderNo:      PmsAppStay.OrderSn,
				ChangeStatus: "CANCEL",
			}).MustToJson(),
			Header: nil,
		}); err != nil {
			g.Log().Error(ctx, "发送分销订单消息变更失败", err)
			err = nil
		}
	}

	return
}
