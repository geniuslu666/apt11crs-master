package logic_travel

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_refund"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

type sTravelOrderRefundService struct{}

func NewTravelOrderRefundService() *sTravelOrderRefundService {
	return &sTravelOrderRefundService{}
}

func init() {
	service.RegisterTravelOrderRefundService(NewTravelOrderRefundService())
}

// ApplyRefundDetail 预退款订单详情
func (s *sTravelOrderRefundService) ApplyRefundDetail(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (out *input_travel.TravelOrderApplyRefundDetailModel, err error) {
	var (
		TravelOrderInfo *input_travel.TravelOrderInfo
		NowTime         = gtime.Now().Format("Y-m-d")
		RefundAmount    float64
		PayAmount       float64
	)
	out = new(input_travel.TravelOrderApplyRefundDetailModel)

	//查询订单信息
	if err = dao.TravelOrder.Ctx(ctx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).WithAll().Scan(&TravelOrderInfo); err != nil && errors.Is(err, sql.ErrNoRows) {
		g.Log().Error(ctx, err)
		return
	}
	if g.IsEmpty(TravelOrderInfo) {
		// 订单不存在
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}
	if TravelOrderInfo.BookDate.Format("Y-m-d") < NowTime {
		// 当前不可取消预约
		err = gerror.New(gi18n.T(ctx, "the_appointment_cannot_be_cancelled"))
		return
	}
	if TravelOrderInfo.OrderStatus != "WAIT_VERIFY" {
		// 当前订单状态不可取消
		err = gerror.New(gi18n.T(ctx, "the_current_order_status_cannot_be_cancelled"))
		return
	}
	out.OrderCreateAt = TravelOrderInfo.CreatedAt.Format("Y-m-d H:i:s")
	out.OrderSn = TravelOrderInfo.OrderSn
	out.OrderAmount = TravelOrderInfo.OrderAmount

	var Transaction []*entity.PmsTransaction
	if err = dao.PmsTransaction.Ctx(ctx).Where(dao.PmsTransaction.Columns().PayStatus, "DONE").Where(dao.PmsTransaction.Columns().OrderSn, TravelOrderInfo.OrderSn).OrderDesc(`
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
	for _, v := range TravelOrderInfo.TransactionRefund {
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

	// 获取取消费用 - 从SKU获取取消政策
	var productSkuInfo *entity.TravelProductSku
	if err = dao.TravelProductSku.Ctx(ctx).Unscoped().
		WherePri(TravelOrderInfo.SkuId).
		Scan(&productSkuInfo); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "server_exception"))
		return
	}

	// 检查SKU是否存在
	if productSkuInfo == nil {
		err = gerror.New(gi18n.T(ctx, "get_car_type_info_failed"))
		return
	}

	allowCancel := productSkuInfo.AllowCancel
	if allowCancel == 1 {
		// 允许取消
		// 距离服务开始前几小时可免费取消
		freeCancelHours := gconv.Int(productSkuInfo.FreeCancelHours)
		// 取消费用比例
		cancelFeePercent := gconv.Int(productSkuInfo.CancelFeePercent)

		// 计算取消手续费
		if TravelOrderInfo.BookDate != nil {
			// 使用已查询的SKU集合时间
			meetingTimeStr := productSkuInfo.MeetingTime

			// 服务开始时间 = 预定日期 + 集合时间
			serviceStartStr := TravelOrderInfo.BookDate.Format("Y-m-d") + " 00:00:00"
			if meetingTimeStr != "" {
				serviceStartStr = TravelOrderInfo.BookDate.Format("Y-m-d") + " " + meetingTimeStr + ":00"
			}
			serviceStart, parseErr := gtime.StrToTime(serviceStartStr)
			if parseErr != nil {
				err = gerror.New(gi18n.T(ctx, "server_exception"))
				return
			}

			// 免费取消截止时间 = 服务开始时间 - freeCancelHours 小时
			cancelDeadline := serviceStart.Add(-time.Duration(freeCancelHours) * time.Hour)
			now := gtime.Now()

			// 已过服务开始时间，无法取消
			if now.After(serviceStart) {
				err = gerror.New(gi18n.T(ctx, "the_appointment_time_has_passed_and_cannot_be_cancelled"))
				return
			}

			if now.Before(cancelDeadline) {
				// 在免费取消截止时间之前，免费取消
				out.CancelFee = 0
			} else {
				// 已过免费取消截止时间，收取手续费
				out.CancelFee = decimal.NewFromInt(gvar.New(cancelFeePercent).Int64()).Mul(decimal.NewFromFloat(TravelOrderInfo.OrderAmount)).Div(decimal.NewFromInt(100)).Round(0).InexactFloat64()
			}
		}
	}
	CancelFee := out.CancelFee

	for _, v := range Transaction {
		if v.PayStatus != "DONE" {
			continue
		}
		item := &input_travel.PreRefundTransaction{
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
				out.RefundAmount += item.ActualRefund
			} else if CancelFee > item.Refundable {
				out.RefundAmount += 0
				item.ActualRefund = 0
				CancelFee = CancelFee - item.Refundable
			} else {
				out.RefundAmount += item.Refundable - CancelFee
				item.ActualRefund = item.Refundable - CancelFee
				CancelFee = 0
			}
		}
	}

	// 避免后台先部分退款 手续费超过剩余支付金额
	if CancelFee > 0 {
		out.CancelFee = out.CancelFee - CancelFee
	}
	return
}

// RefundOrder 退款
func (s *sTravelOrderRefundService) RefundOrder(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (err error) {
	var (
		PreRefundOut *input_travel.TravelOrderApplyRefundDetailModel
		TravelOrder  entity.TravelOrder
		TX           gdb.TX
	)

	// 开启事务
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

	// 获取最新的退款信息
	if PreRefundOut, err = s.ApplyRefundDetail(ctx, in); err != nil {
		err = gerror.New(gi18n.T(ctx, "server_exception"))
		return
	}

	// 查询旅行订单信息
	if err = dao.TravelOrder.Ctx(ctx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Scan(&TravelOrder); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(TravelOrder) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	// 根据总退款金额设置订单状态
	orderStatus := "CANCEL"
	totalRefundAmount := PreRefundOut.RefundAmount + PreRefundOut.RefundBalance
	if totalRefundAmount > 0 {
		orderStatus = "REFUND" // 退款取消
	}

	// 更新订单状态
	totalRefund := TravelOrder.RefundAmount + totalRefundAmount
	var refundStatus string
	if totalRefund <= 0 {
		refundStatus = "WAIT" // 无退款
	} else if totalRefund < TravelOrder.OrderAmount {
		refundStatus = "PART" // 部分退款
	} else {
		refundStatus = "DONE" // 全额退款
	}

	// 准备更新数据
	updateData := g.MapStrAny{
		dao.TravelOrder.Columns().OrderStatus:     orderStatus,
		dao.TravelOrder.Columns().CancelFee:       PreRefundOut.CancelFee,
		dao.TravelOrder.Columns().RefundAmount:    totalRefund,
		dao.TravelOrder.Columns().RefundBalAmount: TravelOrder.RefundBalAmount + PreRefundOut.RefundBalance,
		dao.TravelOrder.Columns().RefundStatus:    refundStatus,
		dao.TravelOrder.Columns().PayStatus:       orderStatus,
		dao.TravelOrder.Columns().CancelTime:      gtime.Now(),
	}

	// 只有实际有退款时才设置退款时间
	if refundStatus != "WAIT" {
		updateData[dao.TravelOrder.Columns().RefundTime] = gtime.Now()
	}

	if _, err = dao.TravelOrder.Ctx(ctx).TX(TX).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Data(updateData).Update(); err != nil {
		return
	}

	// 退款成功减少产品销量字段
	if _, err = dao.TravelProduct.Ctx(ctx).TX(TX).
		Where(dao.TravelProduct.Columns().Id, TravelOrder.ProductId).Update(g.MapStrAny{
		dao.TravelProduct.Columns().SalesNum: gdb.Raw(fmt.Sprintf("sales_num-%d", TravelOrder.BookingNum)),
	}); err != nil {
		return
	}

	// 退款成功减少产品Sku销量字段
	if _, err = dao.TravelProductSku.Ctx(ctx).TX(TX).
		Where(dao.TravelProductSku.Columns().Id, TravelOrder.SkuId).Update(g.MapStrAny{
		dao.TravelProductSku.Columns().SalesNum: gdb.Raw(fmt.Sprintf("sales_num-%d", TravelOrder.BookingNum)),
	}); err != nil {
		return
	}

	// 写入订单日志
	if _, err = dao.TravelOrderLog.Ctx(ctx).TX(TX).OmitEmptyData().Insert(&entity.TravelOrderLog{
		OrderId:     int(TravelOrder.Id),
		ActionWay:   "REFUND",
		OrderStatus: orderStatus,
		Remark:      "订单已退款",
		OperateType: "USER",
		OperateId:   int(TravelOrder.MemberId),
	}); err != nil {
		return
	}

	// 如果有退款金额，调用退款服务
	if totalRefundAmount > 0 {
		if err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      in.OrderSn,
			RefundAmount: totalRefundAmount,
		}, TX); err != nil {
			return
		}
	} else {
		// 插入退款记录到 order_refund_log 表
		if _, err = dao.OrderRefundLog.Ctx(ctx).TX(TX).Insert(&entity.OrderRefundLog{
			OrderSn:      in.OrderSn,
			Scene:        "TRAVEL",
			RefundType:   "AMOUNT",
			RefundAmount: 0,
			RefundTime:   gtime.Now(),
			RefundStatus: "DONE",
			OperateType:  "USER",
			OperateId:    int(TravelOrder.MemberId),
			Remark:       "",
		}); err != nil {
			return
		}
	}

	return
}

// RefundOrderDetail 退款订单详情
func (s *sTravelOrderRefundService) RefundOrderDetail(ctx context.Context, in *input_travel.TravelOrderApplyRefundDetailInp) (out *input_travel.RefundDetailModel, err error) {
	var (
		Order      *entity.TravelOrder
		refundLogs []*entity.OrderRefundLog
	)

	out = new(input_travel.RefundDetailModel)

	// 查询订单信息
	if err = dao.TravelOrder.Ctx(ctx).Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).Scan(&Order); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(Order) {
		err = gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		return
	}

	// 查询退款记录
	if err = dao.OrderRefundLog.Ctx(ctx).
		Where(dao.OrderRefundLog.Columns().OrderSn, in.OrderSn).
		Where(dao.OrderRefundLog.Columns().RefundStatus, "DONE").
		OrderAsc(dao.OrderRefundLog.Columns().Id).
		Scan(&refundLogs); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	// 处理退款记录
	for _, refund := range refundLogs {
		out.RefundRecordList = append(out.RefundRecordList, &input_travel.RefundRecordDetailItem{
			RefundType:   refund.RefundType,
			RefundAmount: refund.RefundAmount,
			ApplyTime:    refund.CreatedAt.Format("Y-m-d H:i:s"),
			OperateType:  refund.OperateType,
			RefundTime:   refund.RefundTime.Format("Y-m-d H:i:s"),
			RefundStatus: refund.RefundStatus,
			Remark:       refund.Remark,
		})
	}

	// 取消政策
	var cancelPolicyConfigs []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err = dao.SysConfig.Ctx(ctx).
		Fields(dao.SysConfig.Columns().Key, dao.SysConfig.Columns().Value).
		Where(dao.SysConfig.Columns().Group, "travelCancelPolicy").
		Scan(&cancelPolicyConfigs); err != nil {
		err = gerror.Wrap(err, gi18n.T(ctx, "server_exception"))
		return
	}
	cancelConfigMap := make(map[string]string)
	for _, cfg := range cancelPolicyConfigs {
		cancelConfigMap[cfg.Key] = cfg.Value
	}
	g.Log().Debug(ctx, "cancelConfigMap", cancelConfigMap)

	allowCancel := cancelConfigMap["allowCancel"]
	if allowCancel == "2" {
		// 不可取消
		out.CancelPolicy = gi18n.T(ctx, "cancel_policy_no_cancel")
	} else if allowCancel == "1" {
		freeCancelHours := gconv.Int(cancelConfigMap["freeCancelDays"])
		cancelFeePercent := cancelConfigMap["cancelFeePercent"]
		out.CancelPolicy = gi18n.Tf(ctx, "cancel_policy_free_before", gvar.New(freeCancelHours).String(), cancelFeePercent)
	}

	// 设置基本信息
	out.OrderSn = in.OrderSn
	out.OrderCreateAt = Order.CreatedAt.Format("Y-m-d H:i:s")
	out.OrderAmount = Order.OrderAmount
	out.BalanceAmount = Order.BalAmount
	out.CouponAmount = Order.CouponAmount
	out.ActualAmount = Order.OrderAmount - Order.BalAmount - Order.CouponAmount
	out.CancelFee = Order.CancelFee
	out.RefundAmount = Order.RefundAmount
	out.RefundActualAmount, _ = decimal.NewFromFloat(Order.RefundAmount).Sub(decimal.NewFromFloat(Order.RefundBalAmount)).Float64()
	out.RefundBalance = Order.RefundBalAmount
	out.RefundTime = Order.RefundTime.Format("Y-m-d H:i:s")

	return
}
