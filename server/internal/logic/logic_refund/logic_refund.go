package logic_refund

import (
	"APT/internal/dao"
	"APT/internal/library/Paypal"
	"APT/internal/library/h5FxPay"
	"APT/internal/library/mlilifeWxPay"
	"APT/internal/library/paycloud"
	"APT/internal/library/stripePay"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"APT/utility/uuid"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"github.com/stripe/stripe-go/v79"
)

type sRefund struct{}

func NewRefund() *sRefund {
	return &sRefund{}
}

func init() {
	service.RegisterRefund(NewRefund())
}

// RefundOrder 金额 订单退款
func (s *sRefund) RefundOrder(ctx context.Context, in *input_refund.RefundAmountInp, tx gdb.TX) (err error) {
	var (
		PmsAppStay            *entity.PmsAppStay
		FoodOrder             *entity.FoodOrder
		SpaOrder              *entity.SpaOrder
		CarOrder              *entity.CarOrder
		CabinetOrder          *entity.CabinetOrder
		TravelOrder           *entity.TravelOrder
		PmsTransactions       []*entity.PmsTransaction
		PmsTransactionRefunds []*entity.PmsTransactionRefund
		RefundHandleList      []*input_refund.RefundAmountList
		AllRefundAmount       float64
		AllPayAmount          float64
		CouponPayAmount       float64
		IsAllRefund           bool
		MemberId              int
		Scene                 string
	)
	// 退款金额验证 必须为整数
	if gvar.New(in.RefundAmount).Int64()*100 != gvar.New(in.RefundAmount*100).Int64() {
		// 退款金额错误
		err = gerror.New(gi18n.T(ctx, "refund_amount_incorrect"))
		return
	}
	// TODO 订单退款验证
	if in.OrderSn[:1] == "H" {
		if err = dao.PmsAppStay.Ctx(ctx).
			Where(dao.PmsAppStay.Columns().OrderSn, in.OrderSn).
			Scan(&PmsAppStay); err != nil {
			return
		}
		if g.IsEmpty(PmsAppStay) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = PmsAppStay.MemberId
		Scene = "HOTEL"
	} else if in.OrderSn[:1] == "F" {
		if err = dao.FoodOrder.Ctx(ctx).
			Where(dao.FoodOrder.Columns().OrderSn, in.OrderSn).
			Scan(&FoodOrder); err != nil {
			return
		}
		if g.IsEmpty(FoodOrder) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = gvar.New(FoodOrder.MemberId).Int()
		Scene = "FOOD"
	} else if in.OrderSn[:1] == "S" {
		if err = dao.SpaOrder.Ctx(ctx).
			Where(dao.SpaOrder.Columns().OrderSn, in.OrderSn).
			Scan(&SpaOrder); err != nil {
			return
		}
		if g.IsEmpty(SpaOrder) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = gvar.New(SpaOrder.MemberId).Int()
		Scene = "SPA"
	} else if in.OrderSn[:1] == "C" {
		if err = dao.CarOrder.Ctx(ctx).
			Where(dao.CarOrder.Columns().OrderSn, in.OrderSn).
			Scan(&CarOrder); err != nil {
			return
		}
		if g.IsEmpty(CarOrder) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = gvar.New(CarOrder.MemberId).Int()
		Scene = "CAR"
	} else if in.OrderSn[:1] == "B" {
		if err = dao.CabinetOrder.Ctx(ctx).
			Where(dao.CabinetOrder.Columns().OrderSn, in.OrderSn).
			Scan(&CabinetOrder); err != nil {
			return
		}
		if g.IsEmpty(CabinetOrder) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = gvar.New(CabinetOrder.MemberId).Int()
		Scene = "CABINET"
	} else if in.OrderSn[:1] == "T" {
		if err = dao.TravelOrder.Ctx(ctx).
			Where(dao.TravelOrder.Columns().OrderSn, in.OrderSn).
			Scan(&TravelOrder); err != nil {
			return
		}
		if g.IsEmpty(TravelOrder) {
			return gerror.New(gi18n.T(ctx, "order_does_not_exist"))
		}
		MemberId = gvar.New(TravelOrder.MemberId).Int()
		Scene = "TRAVEL"
	}

	//if err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
	//	dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
	//	dao.PmsTransaction.Columns().PayStatus: "DONE",
	//}).OrderDesc(dao.PmsTransaction.Columns().Id).Scan(&PmsTransactions); err != nil && !errors.Is(err, sql.ErrNoRows) {
	//	return
	//}
	if err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   in.OrderSn,
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).OrderAsc(`
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
`).Scan(&PmsTransactions); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsTransactions) {
		return gerror.New(gi18n.T(ctx, "order_does_not_exist_or_has_not_paid"))
	}

	if err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().OrderSn:      in.OrderSn,
		dao.PmsTransactionRefund.Columns().RefundStatus: g.Slice{"DONE", "WAIT"},
	}).OrderDesc(dao.PmsTransaction.Columns().Id).Scan(&PmsTransactionRefunds); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}

	// 退款数据组合
	for _, Transactions := range PmsTransactions {
		if Transactions.PayType == "COUPON" {
			CouponPayAmount += Transactions.PayAmount
		}
		AllPayAmount += Transactions.PayAmount
		RefundAmountItem := &input_refund.RefundAmountList{}
		if err = gvar.New(Transactions).Scan(RefundAmountItem); err != nil {
			return
		}
		RefundAmountItem.PaymentRequestId = Transactions.PaymentRequestId
		RefundAmountItem.RefundAmount = 0
		RefundAmountItem.BalanceAmount = Transactions.PayAmount
		RefundAmountItem.MemberId = MemberId
		RefundHandleList = append(RefundHandleList, RefundAmountItem)
	}

	for _, TransactionRefund := range PmsTransactionRefunds {
		for index, RefundAmountItem := range RefundHandleList {
			if TransactionRefund.TransactionSn == RefundAmountItem.TransactionSn {
				RefundHandleList[index].RefundAmount += TransactionRefund.RefundAmount
				RefundHandleList[index].BalanceAmount -= TransactionRefund.RefundAmount
				AllRefundAmount += TransactionRefund.RefundAmount
			}
		}
	}

	// 使用 decimal 处理金额计算，避免浮点数精度问题
	allPayDec := decimal.NewFromFloat(AllPayAmount)
	couponPayDec := decimal.NewFromFloat(CouponPayAmount)
	allRefundDec := decimal.NewFromFloat(AllRefundAmount)
	refundAmountDec := decimal.NewFromFloat(in.RefundAmount)

	if allPayDec.Equal(refundAmountDec) {
		IsAllRefund = true
	}
	if in.RefundAmount == 0 {
		in.RefundAmount = allPayDec.Sub(couponPayDec).Sub(allRefundDec).InexactFloat64()
		refundAmountDec = decimal.NewFromFloat(in.RefundAmount)
	}
	if allPayDec.Sub(couponPayDec).Sub(allRefundDec).LessThan(refundAmountDec) {
		return gerror.New(gi18n.T(ctx, "refund_amount_exceeds_limit_and_cannot_be_refunded"))
	}
	if in.RefundAmount == 0 {
		g.Log().Warning(ctx, "无需退款")
		return
	}

	WantRefundAmount := in.RefundAmount
	for index, RefundAmountItem := range RefundHandleList {
		if RefundAmountItem.BalanceAmount <= 0 || RefundAmountItem.PayType == "COUPON" {
			continue
		}
		if RefundAmountItem.BalanceAmount < WantRefundAmount {
			RefundHandleList[index].WantRefundAmount = RefundAmountItem.BalanceAmount
			WantRefundAmount = WantRefundAmount - RefundAmountItem.BalanceAmount
			continue
		}
		if RefundAmountItem.BalanceAmount >= WantRefundAmount {
			RefundHandleList[index].WantRefundAmount = WantRefundAmount
			WantRefundAmount = 0
			break
		}
	}
	g.Log().Info(ctx, RefundHandleList)
	for _, RefundAmountItem := range RefundHandleList {
		if g.IsEmpty(RefundAmountItem.WantRefundAmount) {
			continue
		}
		itemInp := input_refund.RefundInp{
			MemberId:         RefundAmountItem.MemberId,
			OrderSn:          RefundAmountItem.OrderSn,
			TransactionSn:    RefundAmountItem.TransactionSn,
			PayChannel:       RefundAmountItem.PayChannel,
			RefundType:       RefundAmountItem.PayType,
			RefundAmount:     RefundAmountItem.WantRefundAmount,
			Scene:            RefundAmountItem.Scene,
			PriceCurrency:    RefundAmountItem.PriceCurrency,
			IsAllRefund:      IsAllRefund,
			CaptureId:        RefundAmountItem.CaptureId,
			PaymentRequestId: RefundAmountItem.PaymentRequestId,
			OperateId:        in.OperateId,
			OperateType:      in.OperateType,
			Remark:           in.Remark,
		}
		RefundType := "AMOUNT"

		switch RefundAmountItem.PayChannel {
		case "SYSTEM":
			err = s.BalanceRefund(ctx, tx, itemInp)
			RefundType = "BAL"
			break
		case "PAYCLOUD":
			err = s.PayCloudRefund(ctx, tx, itemInp)
			break
		case "PAYPAL":
			err = s.PaypalRefund(ctx, tx, itemInp)
			break
		case "STRIPE":
			err = s.StripeRefund(ctx, tx, itemInp)
			break
		case "MLILIFE":
			err = s.MlilifeRefund(ctx, tx, itemInp)
			break
		case "H5_FX":
			err = s.H5FXRefund(ctx, tx, itemInp)
			break
		case "COUPON":
			g.Log().Warning(ctx, "优惠券不支持退款")
			break
		default:
			// 退款方式错误
			return gerror.New(gi18n.T(ctx, "refund_method_incorrect"))
		}

		// 退款写入order_refund_log（非H开头的订单号）
		if len(in.OrderSn) > 0 && in.OrderSn[:1] != "H" {
			if _, err = dao.OrderRefundLog.Ctx(ctx).TX(tx).OmitEmptyData().Insert(entity.OrderRefundLog{
				OrderSn:      in.OrderSn,
				Scene:        Scene,
				RefundType:   RefundType,
				RefundAmount: RefundAmountItem.WantRefundAmount,
				RefundTime:   gtime.Now(),
				RefundStatus: "DONE",
				OperateType:  in.OperateType,
				OperateId:    in.OperateId,
				Remark:       in.Remark,
			}); err != nil {
				g.Log().Error(ctx, "写入退款日志失败", "orderSn", in.OrderSn, "error", err)
				// 不阻断退款流程，只记录错误
			}
		}

		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}
	// 回写订单退款状态和支付订单状态
	if err = s.RefundStatusChange(ctx, tx, in.OrderSn); err != nil {
		return
	}

	// 接送机订单退款发短信
	if in.OrderSn[:1] == "C" {
		if CarOrder.DispatchStatus == "DONE" {
			// 发送预警短信
			_ = service.BasicsSmsLog().SendPlaceOrderMsg(ctx, &input_basics.SendMsgInp{
				Event:   "carRefundOrder",
				OrderSn: in.OrderSn,
			})
		}
		// 打印
		_ = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
			OrderSn:  in.OrderSn,
			IsRefund: 1,
		})
	}
	return
}

// BalanceRefund 余额退款
func (s *sRefund) BalanceRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {
	var (
		RefundSn = uuid.CreateOrderCode("R")
		InsertId int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
	}); err != nil {
		// 退款失败
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	if err = service.AppMember().MemberBalanceChange(ctx, &input_app_member.MemberBalanceInp{
		MemberId:      in.MemberId,
		Scene:         in.Scene,
		Type:          "REFUND",
		ChangeBalance: in.RefundAmount,
		OrderSn:       in.OrderSn,
		Reason:        "退款",
	}, tx); err != nil {
		return
	}
	// 退款成功
	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
		dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
	}); err != nil {
		return
	}
	return
}

// PayCloudRefund 云支付退款
func (s *sRefund) PayCloudRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {

	var (
		RefundSn                  = uuid.CreateOrderCode("R")
		paycloudClient            *paycloud.Client
		orderRefundSubmitResponse *paycloud.OrderRefundSubmitResponse
		InsertId                  int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
	}); err != nil {
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	// 发送退款申请
	if paycloudClient, err = paycloud.NewClient(ctx, ""); err != nil {
		return
	}

	if orderRefundSubmitResponse, err = paycloudClient.OrderRefundPayCloudSubmit(ctx, &paycloud.OrderRefundSubmitRequest{
		OrigMerchantOrderNo: in.TransactionSn,
		MerchantOrderNo:     RefundSn,
		PriceCurrency:       in.PriceCurrency,
		OrderAmount:         gvar.New(in.RefundAmount).String(),
		Description:         "退款",
	}); err != nil {
		return
	}
	if orderRefundSubmitResponse.Code != "0" {
		err = gerror.New(orderRefundSubmitResponse.Msg)
		return
	}
	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().TransNo: orderRefundSubmitResponse.Data.TransNo,
	}); err != nil {
		return
	}
	return
}

// PaypalRefund paypal退款
func (s *sRefund) PaypalRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {

	var (
		RefundSn       = uuid.CreateOrderCode("R")
		PaypalClient   *Paypal.Client
		PaypalRefundId string
		InsertId       int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
	}); err != nil {
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	// 发送退款申请
	if PaypalClient, err = Paypal.GetClient(ctx, "", ""); err != nil {
		return
	}
	if PaypalRefundId, err = PaypalClient.RefundOrder(ctx, in.CaptureId, RefundSn, in.RefundAmount, in.IsAllRefund); err != nil {
		return
	}
	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().TransNo: PaypalRefundId,
	}); err != nil {
		return
	}
	return
}
func (s *sRefund) MlilifeRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {

	var (
		RefundSn           = uuid.CreateOrderCode("R")
		mlilifeWxPayRefund *mlilifeWxPay.RefundResponse
		InsertId           int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
	}); err != nil {
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	if mlilifeWxPayRefund, err = mlilifeWxPay.Refund(ctx, &mlilifeWxPay.RefundRequest{
		TransactionID: in.PaymentRequestId,
		RefundOrderNo: RefundSn,
		RefundFee:     gvar.New(in.RefundAmount).String(),
	}); err != nil {
		return
	}
	if mlilifeWxPayRefund.S != 200 {
		err = gerror.New(mlilifeWxPayRefund.Msg)
		return
	}
	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
		dao.PmsTransactionRefund.Columns().TransNo:      mlilifeWxPayRefund.Data.RefundID,
		dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
	}); err != nil {
		return
	}
	return
}

// StripeRefund stripe退款
func (s *sRefund) StripeRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {

	var (
		RefundSn     = uuid.CreateOrderCode("R")
		StripeClient *stripePay.Client
		stripeRefund *stripe.Refund
		InsertId     int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
	}); err != nil {
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	// 发送退款申请
	if StripeClient, err = stripePay.GetClient(ctx, "", ""); err != nil {
		return
	}
	stripe.Key = StripeClient.Key
	if stripeRefund, err = StripeClient.Refund(ctx, RefundSn, gvar.New(in.RefundAmount).Int64(), in.CaptureId); err != nil {
		return
	}
	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().TransNo:    stripeRefund.ID,
		dao.PmsTransactionRefund.Columns().RefundTime: gtime.Now(),
	}); err != nil {
		return
	}
	return
}

func (s *sRefund) RefundStatusChange(ctx context.Context, tx gdb.TX, OrderSn string) (err error) {
	var (
		PmsTransactions       []*entity.PmsTransaction
		PmsTransactionRefunds []*entity.PmsTransactionRefund
		AllRefundAmount       float64
	)
	if err = dao.PmsTransaction.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().OrderSn:   OrderSn,
		dao.PmsTransaction.Columns().PayStatus: "DONE",
	}).OrderDesc(dao.PmsTransaction.Columns().Id).Scan(&PmsTransactions); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	if g.IsEmpty(PmsTransactions) {
		return gerror.New(gi18n.T(ctx, "order_does_not_exist_or_has_not_paid"))
	}

	if err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().OrderSn:      OrderSn,
		dao.PmsTransactionRefund.Columns().RefundStatus: g.Slice{"DONE", "WAIT"},
	}).OrderDesc(dao.PmsTransaction.Columns().Id).Scan(&PmsTransactionRefunds); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return
	}
	for index, Transactions := range PmsTransactions {
		PmsTransactions[index].RefundAmount = 0
		for _, Refund := range PmsTransactionRefunds {
			if Transactions.TransactionSn == Refund.TransactionSn {
				PmsTransactions[index].RefundAmount += Refund.RefundAmount
				AllRefundAmount += Refund.RefundAmount
			}
		}
		//if PmsTransactions[index].RefundAmount == 0 {
		//	continue
		//}
		// 更新支付流水订单状态
		if _, err = dao.PmsTransaction.Ctx(ctx).TX(tx).
			Where(dao.PmsTransaction.Columns().Id, Transactions.Id).
			Update(g.MapStrAny{
				dao.PmsTransaction.Columns().RefundAmount: PmsTransactions[index].RefundAmount,
				dao.PmsTransaction.Columns().RefundStatus: gdb.Raw(fmt.Sprintf(`
			CASE 
				WHEN %v >= pay_amount THEN 'DONE'
				WHEN %v < pay_amount AND %v > 0 THEN 'PART'
				ELSE  'WAIT'
			END`, PmsTransactions[index].RefundAmount, PmsTransactions[index].RefundAmount, PmsTransactions[index].RefundAmount)),
			}); err != nil {
			return
		}
	}
	if OrderSn[:1] == "H" {
		// 修改住宿主订单退款状态
		if _, err = dao.PmsAppStay.Ctx(ctx).TX(tx).Where(dao.PmsAppStay.Columns().OrderSn, OrderSn).Data(g.Map{
			dao.PmsAppStay.Columns().RefundAmount: AllRefundAmount,
			dao.PmsAppStay.Columns().RefundStatus: gdb.Raw(fmt.Sprintf(`
			CASE 
				WHEN %v >= order_amount THEN 'DONE'
				WHEN %v < order_amount AND %v > 0 THEN 'PART'
				ELSE  'WAIT'
			END`, AllRefundAmount, AllRefundAmount, AllRefundAmount)),
			dao.PmsAppStay.Columns().RefundTime: gtime.Now(),
		}).Update(); err != nil {
			return
		}
	}

	return
}

// H5FXRefund 分销退款
func (s *sRefund) H5FXRefund(ctx context.Context, tx gdb.TX, in input_refund.RefundInp) (err error) {

	var (
		RefundSn                = uuid.CreateOrderCode("R")
		H5FxRefundOrderResponse *h5FxPay.RefundOrderResponse
		InsertId                int64
	)

	if InsertId, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).OmitEmptyData().InsertAndGetId(entity.PmsTransactionRefund{
		CancelOrderSn: in.CancelOrderSn,
		OrderSn:       in.OrderSn,
		TransactionSn: in.TransactionSn,
		RefundChannel: in.PayChannel,
		RefundType:    in.RefundType,
		RefundSn:      RefundSn,
		RefundAmount:  in.RefundAmount,
		OperateId:     in.OperateId,
		OperateType:   in.OperateType,
		Remark:        in.Remark,
		IsFx:          "Y",
	}); err != nil {
		err = gerror.New(gi18n.T(ctx, "refund_failed"))
		return
	}

	// 发送退款申请

	if H5FxRefundOrderResponse, err = h5FxPay.RefundOrder(ctx, &h5FxPay.RefundOrderRequest{
		RefundOrderNo: RefundSn,
		OrderNo:       in.OrderSn,
		TradeNo:       in.TransactionSn,
		RefundAmount:  int64(in.RefundAmount),
		TradeTime:     gtime.Now().Format("Y-m-d H:i:s"),
	}); err != nil {
		return
	}

	if _, err = dao.PmsTransactionRefund.Ctx(ctx).TX(tx).Where(dao.PmsTransactionRefund.Columns().Id, InsertId).Update(g.MapStrAny{
		dao.PmsTransactionRefund.Columns().TransNo: H5FxRefundOrderResponse.Data.ThirdRefundOrderNo,
	}); err != nil {
		return
	}
	return
}
