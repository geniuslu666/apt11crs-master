package app

import (
	"APT/internal/consts"
	"APT/internal/dao"
	"APT/internal/library/airhousePublicApi"
	"APT/internal/library/cache"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/rabbitmq"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) PreOrderRefundDetail(ctx context.Context, req *hotel.PreOrderRefundDetailReq) (res *hotel.PreOrderRefundDetailRes, err error) {
	res = new(hotel.PreOrderRefundDetailRes)
	if res.PreRefundOut, err = service.HotelService().PreRefundOrderDetail(ctx, &req.PreRefundIn); err != nil {
		//err = gerror.New("服务器异常请稍后再试")
		return
	}

	// 查询退款记录列表
	res.RefundRecordList = getRefundRecordDetailList(ctx, req.OrderSn)

	if err = cache.Instance().Set(ctx, res.PreCancelOrderSn, res, gtime.M*10); err != nil {
		//err = gerror.New("服务器异常请稍后再试")
		return
	}
	return
}
func (c *ControllerHotel) RefundOrder(ctx context.Context, req *hotel.RefundOrderReq) (res *hotel.RefundOrderRes, err error) {
	var (
		PreRefundOut       *input_hotel.PreRefundOut
		NowPreRefundOut    *input_hotel.PreRefundOut
		ProCancelOrderInfo *gvar.Var
		CancelOrderSn      = uuid.CreateOrderCode("C")
		RoomReservation    []*entity.PmsAppReservation
		TX                 gdb.TX
		PmsAppStay         entity.PmsAppStay
	)
	if TX, err = g.DB().Begin(ctx); err != nil {
		return
	}
	defer func() {
		if err != nil {
			g.Log().Error(ctx, err)
			_ = TX.Rollback()
		} else {
			_ = TX.Commit()
		}
	}()
	if ProCancelOrderInfo, err = cache.Instance().Get(ctx, req.PreCancelOrderSn); err != nil {
		// 服务器异常请稍后再试
		err = gerror.New(gi18n.T(ctx, "server_exception"))
		return
	}
	if err = ProCancelOrderInfo.Struct(&NowPreRefundOut); err != nil {
		// 服务器异常请稍后再试
		err = gerror.New(gi18n.T(ctx, "server_exception"))
		return
	}
	if PreRefundOut, err = service.HotelService().PreRefundOrderDetail(ctx, &input_hotel.PreRefundIn{
		OrderSn: req.OrderSn,
	}); err != nil {
		// 服务器异常请稍后再试
		err = gerror.New(gi18n.T(ctx, "server_exception"))
		return
	}
	if g.IsEmpty(NowPreRefundOut) {
		// 取消退款订单信息已过期，请重新申请取消退款
		err = gerror.New(gi18n.T(ctx, "cancel_refund_order_info_has_expired"))
		return
	}
	if PreRefundOut.RefundFee != NowPreRefundOut.RefundFee || PreRefundOut.RefundBalance != NowPreRefundOut.RefundBalance {
		// 退款信息发生变化
		err = gerror.New(gi18n.T(ctx, "refund_info_has_changed"))
		return
	}

	// 酒店订单退款
	if err = dao.PmsAppStay.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsAppStay.Columns().OrderSn: req.OrderSn,
	}).Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsAppStay) {
		// 该订单无需处理
		err = gerror.New(gi18n.T(ctx, "order_does_not_need_handle"))
		return
	}

	// 创建取消订单信息
	if _, err = dao.PmsAppCancelOrder.Ctx(ctx).TX(TX).Insert(g.MapStrAny{
		dao.PmsAppCancelOrder.Columns().CancelOrderSn: CancelOrderSn,
		dao.PmsAppCancelOrder.Columns().OrderSn:       PreRefundOut.OrderSn,
		dao.PmsAppCancelOrder.Columns().OutOrderSn:    PreRefundOut.OutOrderSn,
		dao.PmsAppCancelOrder.Columns().CancelAmount:  PreRefundOut.CancelFee,
		dao.PmsAppCancelOrder.Columns().CancelRate:    gvar.New(PreRefundOut.CancelRate).String(),
		dao.PmsAppCancelOrder.Columns().CancelAt:      gtime.Now().Format("Y-m-d H:i:s"),
		dao.PmsAppCancelOrder.Columns().CreatedAt:     gtime.Now().Format("Y-m-d H:i:s"),
		dao.PmsAppCancelOrder.Columns().UpdatedAt:     gtime.Now().Format("Y-m-d H:i:s"),
	}); err != nil {
		return
	}
	// 写入取消金额到订单中去
	if _, err = dao.PmsAppStay.Ctx(ctx).TX(TX).Where(dao.PmsAppStay.Columns().OrderSn, req.OrderSn).Data(g.MapStrAny{
		dao.PmsAppStay.Columns().CleanFee:     PreRefundOut.CancelFee,
		dao.PmsAppStay.Columns().OrderStatus:  "CANCEL",
		dao.PmsAppStay.Columns().RefundStatus: "PART",
		dao.PmsAppStay.Columns().RefundAmount: PmsAppStay.RefundAmount + PreRefundOut.RefundFee + PreRefundOut.RefundBalance + PreRefundOut.RefundCouponAmount,
		dao.PmsAppStay.Columns().RefundTime:   gtime.Now(),
	}).Update(); err != nil {
		return
	}
	if (PreRefundOut.RefundFee + PreRefundOut.RefundBalance + PreRefundOut.RefundCouponAmount) > 0 {
		if err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      PreRefundOut.OrderSn,
			RefundAmount: PreRefundOut.RefundFee + PreRefundOut.RefundBalance + PreRefundOut.RefundCouponAmount,
		}, TX); err != nil {
			return
		}
	}

	// 查询入住订单信息 进行取消操作
	if err = dao.PmsAppReservation.Ctx(ctx).Where(dao.PmsAppReservation.Columns().OrderSn, req.OrderSn).Scan(&RoomReservation); err != nil {
		return
	}
	for _, RoomReservationItem := range RoomReservation {
		if _, err = airhousePublicApi.UpdateRoomReservationPost(ctx, RoomReservationItem.Uuid, g.MapStrAny{
			"status": "cancelled",
			//"checkin_date":     RoomReservationItem.CheckinDate,
			//"checkout_date":    RoomReservationItem.CheckoutDate,
			//"cancellation_fee": PreRefundOut.CancelFee,
			//"booking_fee":      0,
		}); err != nil {
			return
		}
		// 更新入住订单信息
		if _, err = dao.PmsAppReservation.Ctx(ctx).TX(TX).Where(dao.PmsAppReservation.Columns().Uuid, RoomReservationItem.Uuid).Data(g.MapStrAny{
			dao.PmsAppReservation.Columns().OrderStatus:     "CANCEL",
			dao.PmsAppReservation.Columns().Status:          "cancelled",
			dao.PmsAppReservation.Columns().CancellationFee: PreRefundOut.CancelFee,
		}).Update(); err != nil {
			return
		}
	}

	if _, err = dao.PmsAppStayLog.Ctx(ctx).OmitEmptyData().Insert(&entity.PmsAppStayLog{
		OrderId:     PmsAppStay.Id,
		ActionWay:   "REFUND",
		Remark:      "订单退款",
		OperateType: "USER",
		OperateId:   PmsAppStay.MemberId,
	}); err != nil {
		return
	}

	// 发送到消息队列
	systemMessageTitle := map[string]string{
		"zh":    "订单退款成功",
		"en":    "Order refund successful",
		"ja":    "注文の払い戻しが完了しました",
		"ko":    "주문 환불 성공",
		"zh_CN": "訂單退款成功",
	}
	systemMessageContent := map[string]string{
		"zh":    PmsAppStay.OrderSn + "订单已成功退款",
		"en":    "Order " + PmsAppStay.OrderSn + " has been successfully refunded.",
		"ja":    "注文 " + PmsAppStay.OrderSn + " の払い戻しが正常に完了しました。",
		"ko":    "주문 " + PmsAppStay.OrderSn + "이 성공적으로 환불되었습니다.",
		"zh_CN": PmsAppStay.OrderSn + "訂單已成功退款",
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
		OperatorId:           int(PmsAppStay.MemberId),
		OperatorRole:         "MEMBER",
		OrderSn:              PmsAppStay.OrderSn,
	})

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

	return
}

func (c *ControllerHotel) OrderRefundDetail(ctx context.Context, req *hotel.OrderRefundDetailReq) (res *hotel.OrderRefundDetailRes, err error) {
	res = new(hotel.OrderRefundDetailRes)
	if res.RefundDetailModel, err = service.HotelService().RefundOrderDetail(ctx, req.RefundDetailInp); err != nil {
		return
	}

	// 查询退款记录列表
	if res.RefundDetailModel != nil {
		res.RefundDetailModel.RefundRecordList = getRefundRecordDetailList(ctx, req.OrderSn)
	}

	return
}

// getRefundRecordDetailList 获取退款记录详情列表（包含更多字段）
// 逻辑：
// 1. 查询 transaction_refund 表中的退款记录
// 2. 查询 app_cancel_order 表中的取消记录
// 3. 如果 cancel_order 中有记录但 refund 表中没有 operate_type 为 USER 的记录，说明有一笔退款为0的记录
// 4. 合并结果并按申请时间升序排序
func getRefundRecordDetailList(ctx context.Context, orderSn string) []*input_hotel.RefundRecordDetailItem {
	var (
		refundRecords []*entity.PmsTransactionRefund
		cancelOrder   *entity.PmsAppCancelOrder
		result        []*input_hotel.RefundRecordDetailItem
		hasUserRefund bool
	)

	// 1. 查询 transaction_refund 表中的退款记录（已完成的退款）
	_ = dao.PmsTransactionRefund.Ctx(ctx).
		Where(dao.PmsTransactionRefund.Columns().OrderSn, orderSn).
		Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").
		OrderAsc(dao.PmsTransactionRefund.Columns().CreatedAt).
		Scan(&refundRecords)

	// 转换退款记录
	for _, v := range refundRecords {
		refundType := "AMOUNT"
		if v.RefundType == "BAL" {
			refundType = "BAL"
		}
		refundTime := ""
		if v.RefundTime != nil {
			refundTime = v.RefundTime.String()
		}
		result = append(result, &input_hotel.RefundRecordDetailItem{
			RefundType:   refundType,
			RefundAmount: v.RefundAmount,
			ApplyTime:    v.CreatedAt.String(),
			OperateType:  v.OperateType,
			RefundTime:   refundTime,
			RefundStatus: v.RefundStatus,
			Remark:       v.Remark,
		})
		// 检查是否有用户发起的退款
		if v.OperateType == "USER" {
			hasUserRefund = true
		}
	}

	// 2. 查询 app_cancel_order 表中的取消记录
	_ = dao.PmsAppCancelOrder.Ctx(ctx).
		Where(dao.PmsAppCancelOrder.Columns().OrderSn, orderSn).
		Scan(&cancelOrder)

	// 3. 如果有取消记录但没有用户发起的退款记录，说明是退款金额为0的情况
	if !g.IsEmpty(cancelOrder) && !hasUserRefund {
		cancelAt := ""
		if cancelOrder.CancelAt != nil {
			cancelAt = cancelOrder.CancelAt.String()
		}
		result = append(result, &input_hotel.RefundRecordDetailItem{
			RefundType:   "AMOUNT",
			RefundAmount: 0,
			ApplyTime:    cancelOrder.CreatedAt.String(),
			OperateType:  "USER",
			RefundTime:   cancelAt,
			RefundStatus: "DONE",
			Remark:       "",
		})
	}

	// 4. 按申请时间升序排序
	if len(result) > 1 {
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].ApplyTime > result[j].ApplyTime {
					result[i], result[j] = result[j], result[i]
				}
			}
		}
	}

	return result
}
