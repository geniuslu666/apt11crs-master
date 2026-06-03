package logic_hotel

import (
	"APT/internal/dao"
	"APT/internal/library/hgorm/hook"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// PreRefundOrderDetail 预退款订单详情
func (s *sHotelService) PreRefundOrderDetail(ctx context.Context, in *input_hotel.PreRefundIn) (out *input_hotel.PreRefundOut, err error) {
	var (
		PmsAppStay        *input_hotel.PmsAppStayInfo
		PmsAppReservation []*entity.PmsAppReservation
		NowTime           = gtime.Now().Format("Y-m-d")
		RefundAmount      float64
		PayAmount         float64
	)
	out = new(input_hotel.PreRefundOut)
	if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&out.CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(ctx, err)
		return
	}
	//查询订单信息
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).WithAll().Scan(&PmsAppStay); err != nil && errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(ctx, err)
		return
	}
	if g.IsEmpty(PmsAppStay) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	if PmsAppStay.CheckInDate < NowTime {
		// 当前不可取消预约
		err = gerror.New(gi18n.T(ctx, "the_appointment_cannot_be_cancelled"))
		return
	}
	if PmsAppStay.OrderStatus != "HAVE_PAID" {
		// 当前订单状态不可取消
		err = gerror.New(gi18n.T(ctx, "the_current_order_status_cannot_be_cancelled"))
		return
	}
	out.OrderCreateAt = PmsAppStay.CreatedAt.Format("Y-m-d H:i:s")
	// 查询入住单
	if err = dao.PmsAppReservation.Ctx(ctx).
		Where(dao.PmsAppReservation.Columns().OrderSn, in.OrderSn).
		Scan(&PmsAppReservation); err != nil && errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(ctx, err)
		return
	}
	if g.IsEmpty(PmsAppReservation) {
		// 入住单不存在，无法进行退款
		err = gerror.New(gi18n.T(ctx, "checkin_order_does_not_exist_cannot_refund"))
		return
	}

	for _, v := range PmsAppReservation {
		if v.CheckinStatus != "before_checkin" {
			// 当前订单无法取消
			err = gerror.New(gi18n.T(ctx, "the_current_order_cannot_be_cancelled"))
			return
		}
	}

	out.PreCancelOrderSn = uuid.CreateOrderCode("PC")
	out.OrderSn = PmsAppStay.OrderSn
	out.OutOrderSn = PmsAppStay.OutOrderSn
	out.OrderFee = PmsAppStay.OrderAmount

	var Transaction []*entity.PmsTransaction
	if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().PayStatus, "DONE").Where(dao.PmsTransaction.Columns().OrderSn, PmsAppStay.OrderSn).OrderDesc(`
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

	// 获取已退款金额
	for _, v := range PmsAppStay.TransactionRefund {
		if v.RefundStatus == "DONE" {
			RefundAmount = RefundAmount + v.RefundAmount
		}
	}

	// 获取支付金额
	for _, v := range Transaction {
		if v.PayStatus == "DONE" {
			PayAmount = PayAmount + v.PayAmount
		}
	}
	PayAmount = PayAmount - RefundAmount

	for k, v := range out.CancelRate {
		switch v.Mode {
		case "after":
			startDate := gtime.New(PmsAppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			if startDate <= NowTime {
				out.CancelRate[k].Selected = true
			}
			out.CancelRate[k].Date = startDate
			break
		case "middle":
			startDay := gtime.New(PmsAppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
			endDate := gtime.New(PmsAppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if startDay <= NowTime && endDate >= NowTime {
				out.CancelRate[k].Selected = true
			}
			out.CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
			break
		case "before":
			endDate := gtime.New(PmsAppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
			if endDate >= NowTime {
				out.CancelRate[k].Selected = true
			}
			out.CancelRate[k].Date = endDate
			break
		}
		if out.CancelRate[k].Selected {
			out.CancelFee = decimal.NewFromInt(gvar.New(v.Rate).Int64()).Mul(decimal.NewFromFloat(PmsAppStay.OrderAmount)).Div(decimal.NewFromInt(100)).Round(0).InexactFloat64()
		}
	}
	CancelFee := out.CancelFee

	for _, v := range Transaction {
		if v.PayStatus != "DONE" {
			continue
		}
		item := &input_hotel.PreRefundTransaction{
			TransactionSn: v.TransactionSn,           // 流水号
			PayType:       v.PayType,                 // 支付方式
			Amount:        v.Amount,                  // 支付金额
			RefundAmount:  v.RefundAmount,            // 已退款金额
			Refundable:    v.Amount - v.RefundAmount, // 可退款金额
			ActualRefund:  0,                         // 需退款金额
			PriceCurrency: v.PriceCurrency,           // 币种
			PayStatus:     v.PayStatus,               // 支付状态
		}
		if v.PayType == "BAL" {
			out.BalanceAmount += v.Amount
			if g.IsEmpty(CancelFee) {
				item.ActualRefund = item.Refundable
				out.RefundBalance += item.ActualRefund
			} else if CancelFee > item.Refundable {
				out.RefundBalance += 0
				item.ActualRefund = 0
				CancelFee = CancelFee - item.Refundable
			} else {
				out.RefundBalance += item.Refundable - CancelFee
				item.ActualRefund = item.Refundable - CancelFee
				CancelFee = 0
			}
		} else if v.PayType == "COUPON" {
			out.CouponAmount = v.PayAmount
			out.RefundCouponAmount = 0
			//out.RefundFee += v.PayAmount
			if g.IsEmpty(CancelFee) {

			} else if CancelFee > item.Refundable {
				CancelFee = CancelFee - item.Refundable
			} else {
				CancelFee = 0
			}

		} else {
			out.ActualAmount += v.Amount
			if g.IsEmpty(CancelFee) {
				item.ActualRefund = item.Refundable
				out.RefundFee += item.ActualRefund
			} else if CancelFee > item.Refundable {
				out.RefundFee += 0
				item.ActualRefund = 0
				CancelFee = CancelFee - item.Refundable
			} else {
				out.RefundFee += item.Refundable - CancelFee
				item.ActualRefund = item.Refundable - CancelFee
				CancelFee = 0
			}
		}
		if item.PayType != "COUPON" {
			out.Transaction = append(out.Transaction, item)
		}
	}

	// 避免后台先部分退款 手续费超过剩余支付金额
	if CancelFee > 0 {
		out.CancelFee = out.CancelFee - CancelFee
	}
	return
}

// RefundOrderDetail 退款订单详情
func (s *sHotelService) RefundOrderDetail(ctx context.Context, in *input_hotel.RefundDetailInp) (out *input_hotel.RefundDetailModel, err error) {
	var (
		Transaction       []*entity.PmsTransaction
		TransactionRefund []*entity.PmsTransactionRefund
		AppStay           *entity.PmsAppStay
		CancelOrderInfo   *struct {
			*entity.PmsAppCancelOrder
			Transaction []*struct {
				g.Meta `orm:"table:hg_pms_transaction"`
			} `json:"transaction" orm:"with:order_sn=order_sn, where:pay_status='DONE', where:refund_status!='DONE'" dc:"支付流水"`
			TransactionRefund []*struct {
				g.Meta `orm:"table:hg_pms_transaction_refund"`
				*entity.PmsTransactionRefund
			} `json:"transaction_refund" orm:"with:order_sn=order_sn, where:refund_status='DONE'" dc:"支付流水"`
		}
	)
	out = new(input_hotel.RefundDetailModel)
	if err = dao.PmsAppStay.Ctx(ctx).Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).WithAll().Scan(&AppStay); err != nil {
		return
	}
	if err = dao.PmsAppCancelOrder.Ctx(ctx).Where(dao.PmsAppCancelOrder.Columns().OrderSn, in.OrderSn).WithAll().Scan(&CancelOrderInfo); err != nil {
		return
	}
	//if g.IsEmpty(CancelOrderInfo) {
	//	// 系统取消
	//	err = gerror.New(gi18n.T(ctx, "system_cancel"))
	//	return
	//}
	if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().OrderSn, in.OrderSn).Scan(&Transaction); err != nil {
		return
	}

	for _, v := range Transaction {
		if v.PayStatus != "DONE" {
			continue
		}
		if v.PayType == "BAL" {
			out.BalanceAmount += v.Amount
		} else if v.PayType == "COUPON" {
			out.CouponAmount = v.PayAmount
		} else {
			out.ActualAmount += v.Amount
		}
	}

	if err = dao.PmsTransactionRefund.Ctx(ctx).Where(dao.PmsTransactionRefund.Columns().OrderSn, in.OrderSn).Where(dao.PmsTransactionRefund.Columns().RefundStatus, "DONE").Scan(&TransactionRefund); err != nil {
		return
	}

	if !g.IsEmpty(CancelOrderInfo) {
		if err = gvar.New(CancelOrderInfo.CancelRate).Struct(&out.CancelRate); err != nil {
			return
		}

		out.CancelFee = CancelOrderInfo.CancelAmount
		out.CancelOrderSn = CancelOrderInfo.CancelOrderSn
		out.OutOrderSn = CancelOrderInfo.OutOrderSn
		out.RefundAt = gtime.New(CancelOrderInfo.CancelAt).Format("Y-m-d H:i:s")
	} else {
		if err = dao.PmsCancelRate.Ctx(ctx).Hook(hook.PmsFindLanguageValueHook).OrderAsc(dao.PmsCancelRate.Columns().Sort).Scan(&out.CancelRate); err != nil && errors.Is(err, sql.ErrNoRows) {
			g.Log().Error(ctx, err)
			return
		}

		out.CancelFee = decimal.NewFromFloat(AppStay.OrderAmount).Sub(decimal.NewFromFloat(AppStay.RefundAmount)).InexactFloat64()
		out.RefundAt = gtime.New(AppStay.RefundTime).Format("Y-m-d H:i:s")

		for k, v := range out.CancelRate {
			switch v.Mode {
			case "after":
				startDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
				if startDate <= out.RefundAt {
					out.CancelRate[k].Selected = true
				}
				out.CancelRate[k].Date = startDate
				break
			case "middle":
				startDay := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.StartDays*24) * time.Hour).Format("Y-m-d")
				endDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
				if startDay <= out.RefundAt && endDate >= out.RefundAt {
					out.CancelRate[k].Selected = true
				}
				out.CancelRate[k].Date = fmt.Sprintf("%s ~ %s", startDay, endDate)
				break
			case "before":
				endDate := gtime.New(AppStay.CheckInDate).Add(time.Duration(-v.EndDays*24) * time.Hour).Format("Y-m-d")
				if endDate >= out.RefundAt {
					out.CancelRate[k].Selected = true
				}
				out.CancelRate[k].Date = endDate
				break
			}
		}
	}

	out.OrderSn = in.OrderSn

	out.Transaction = Transaction
	out.TransactionRefund = TransactionRefund
	out.OrderCreateAt = gtime.New(AppStay.CreatedAt).Format("Y-m-d H:i:s")
	out.OrderFee = AppStay.OrderAmount
	for _, v := range TransactionRefund {
		if v.RefundType == "BAL" {
			out.RefundBalance += v.RefundAmount
		} else {
			out.RefundFee += v.RefundAmount
		}
	}

	return
}
