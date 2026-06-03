package admin

import (
	"APT/api/admin/pms"
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/contexts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerPms) RefundOrder(ctx context.Context, req *pms.RefundOrderReq) (res *pms.RefundOrderRes, err error) {
	var (
		RoomReservation []*entity.PmsAppReservation
		//PmsTransaction  []*entity.PmsTransaction
	)

	// 酒店订单退款
	var (
		PmsAppStay entity.PmsAppStay
	)
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: req.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		err = gerror.New("该订单无需处理")
		return
	}

	// 判断会员是否已注销
	memberCount, _ := dao.PmsMember.Ctx(ctx).Where(dao.PmsMember.Columns().Id, PmsAppStay.MemberId).Count()
	if memberCount == 0 {
		err = gerror.New("会员已注销，无法操作此订单")
		return
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if req.IsCancelOrder == "Y" {
			//if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(dao.PmsTransaction.Columns().OrderSn, req.OrderSn).Scan(&PmsTransaction); err != nil {
			//	return
			//}
			//for _, PmsTransactionItem := range PmsTransaction {
			//	if PmsTransactionItem.RefundStatus != "DONE" && PmsTransactionItem.PayType != "COUPON" {
			//		err = gerror.New("已退款，是部分退款无法进行取消订单")
			//		return
			//	}
			//}
			// 查询入住订单信息 进行取消操作
			if err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(g.MapStrAny{
				dao.PmsAppReservation.Columns().OrderSn: req.OrderSn,
				//dao.PmsAppReservation.Columns().Id:      req.Id,
			}).Scan(&RoomReservation); err != nil {
				return
			}
			if g.IsEmpty(RoomReservation) {
				err = gerror.New("不存在入住单")
				return
			}

			for _, RoomReservationItem := range RoomReservation {
				//if RoomReservationItem.Status != "confirmed" {
				//	err = gerror.New("已退款，订单不是确认状态，无法取消")
				//	return
				//}
				if _, err = airhousePublicApi.UpdateRoomReservationPost(ctx, RoomReservationItem.Uuid, g.MapStrAny{
					"status":           "cancelled",
					"checkin_date":     RoomReservationItem.CheckinDate,
					"checkout_date":    RoomReservationItem.CheckoutDate,
					"cancellation_fee": RoomReservationItem.BookingFee,
					"booking_fee":      0,
				}); err != nil {
					return
				}
				// 更新入住订单信息
				if _, err = dao.PmsAppReservation.Ctx(ctx).TX(tx).Where(dao.PmsAppReservation.Columns().Uuid, RoomReservationItem.Uuid).Data(g.MapStrAny{
					dao.PmsAppReservation.Columns().OrderStatus:     "CANCEL",
					dao.PmsAppReservation.Columns().Status:          "cancelled",
					dao.PmsAppReservation.Columns().CancellationFee: 0,
					dao.PmsAppReservation.Columns().CancelRemake:    "退款取消:" + req.Reason,
				}).Update(); err != nil {
					return
				}
			}

			// 修改主订单为已取消
			if _, err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, req.OrderSn).Data(g.MapStrAny{
				dao.PmsAppReservation.Columns().OrderStatus: "CANCEL",
			}).Update(); err != nil {
				return
			}

		}
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      req.OrderSn,
			RefundAmount: req.RefundAmount,
			Remark:       req.Reason,
			OperateType:  "ADMIN",
			OperateId:    int(contexts.GetUserId(ctx)),
		}, tx)

		// 失效
		invalidSendMsg, _ := json.Marshal(g.Map{
			"type": "INVALID",
			"id":   gvar.New(PmsAppStay.Id).Int(),
		})
		if err = rabbitmq.SendMqMessage(ctx, &rabbitmq.SendMqMessageParams{
			ExchangeName: consts.RabbitMQExchangeName,
			QueueName:    consts.RabbitMQQueueNameOrderAward,
			DataByte:     invalidSendMsg,
			Header:       nil,
		}); err != nil {
			g.Log().Error(ctx, "发送下单奖励失效MQ失败", err)
		}
		//err = service.HotelService().HotelOrderAwardInvalid(ctx, tx, gvar.New(PmsAppStay.Id).Int())

		return
	}); err != nil {
		return
	}

	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   "REFUND",
		Remark:      "订单退款",
		OperateType: "ADMIN",
		OperateId:   int(contexts.GetUserId(ctx)),
	}); err != nil {
		return
	}

	// 发送通知
	if req.IsCancelOrder == "Y" {
		// 发送到消息队列
		systemMessageTitle := map[string]string{
			"zh":    "订单已取消",
			"en":    "Order has been canceled",
			"ja":    "注文はキャンセルされました",
			"ko":    "주문이 취소되었습니다",
			"zh_CN": "訂單已取消",
		}
		var systemMessageContent map[string]string
		if req.RefundAmount > 0 {
			systemMessageContent = map[string]string{
				"zh":    "订单已取消，并成功退款" + gvar.New(req.RefundAmount).String() + "JPY",
				"en":    "Order has been canceled, and a refund of " + gvar.New(req.RefundAmount).String() + "JPY has been issued",
				"ja":    "注文はキャンセルされ、" + gvar.New(req.RefundAmount).String() + "JPY の返金が完了しました",
				"ko":    "주문이 취소되었으며 " + gvar.New(req.RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
				"zh_CN": "訂單已取消，並成功退款 " + gvar.New(req.RefundAmount).String() + "JPY",
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
			"orderSn": PmsAppStay.OrderSn,
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
			WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", PmsAppStay.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              PmsAppStay.OrderSn,
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

		if req.RefundAmount > 0 {
			systemMessageContent = map[string]string{
				"zh":    "订单成功退款" + gvar.New(req.RefundAmount).String() + "JPY",
				"en":    "Refund of " + gvar.New(req.RefundAmount).String() + "JPY completed",
				"ja":    gvar.New(req.RefundAmount).String() + "JPY の返金が完了しました",
				"ko":    gvar.New(req.RefundAmount).String() + "JPY가 성공적으로 환불되었습니다",
				"zh_CN": "訂單成功退款" + gvar.New(req.RefundAmount).String() + "JPY",
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
			"orderSn": PmsAppStay.OrderSn,
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
			WxLink:               fmt.Sprintf("/pages/orders/details?orderSn=%s", PmsAppStay.OrderSn),
			EnablePush:           true,
			EnableSms:            false,
			PushTitle:            systemMessageTitle[memberLanguage],
			PushContent:          systemMessageContent[memberLanguage],
			OperatorId:           int(contexts.GetUserId(ctx)),
			OperatorRole:         "ADMIN",
			OrderSn:              PmsAppStay.OrderSn,
		})
	}

	return
}
