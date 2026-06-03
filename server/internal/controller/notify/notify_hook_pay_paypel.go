package notify

import (
	"APT/internal/dao"
	"APT/internal/library/Paypal"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"context"
	"github.com/go-pay/gopay/paypal"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/hook"
)

func (c *ControllerHook) PaypelPayReturn(ctx context.Context, req *hook.PaypelPayReturnReq) (res *hook.PaypelPayReturnRes, err error) {

	var (
		r                = ghttp.RequestFromCtx(ctx)
		client           *Paypal.Client
		CaptureOrderResp *paypal.OrderCaptureRsp
		PayConfig        *model.PayConfig
	)
	g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), r.GetBody())
	if client, err = Paypal.GetClient(ctx, "", ""); err != nil {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
		return
	}
	if CaptureOrderResp, err = client.CaptureOrder(ctx, req.PayerID, req.Token); err != nil {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
		return
	}
	if _, err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
		dao.PmsTransaction.Columns().PaymentRequestId: req.Token,
	}).Update(g.MapStrAny{
		dao.PmsTransaction.Columns().CaptureId: CaptureOrderResp.Response.PurchaseUnits[0].Payments.Captures[0].Id,
		dao.PmsTransaction.Columns().PayAmount: CaptureOrderResp.Response.PurchaseUnits[0].Payments.Captures[0].SellerReceivableBreakdown.NetAmount.Value,
		dao.PmsTransaction.Columns().PayCharge: CaptureOrderResp.Response.PurchaseUnits[0].Payments.Captures[0].SellerReceivableBreakdown.PaypalFee.Value,
	}); err != nil {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
		return
	}
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	r.Response.RedirectTo(PayConfig.PaypalReturnApp)
	return
}
func (c *ControllerHook) PaypelPayCancel(ctx context.Context, _ *hook.PaypelPayCancelReq) (res *hook.PaypelPayCancelRes, err error) {
	var (
		r         = ghttp.RequestFromCtx(ctx)
		PayConfig *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), r.GetBody())
	r.Response.RedirectTo(PayConfig.PaypalCleanApp)
	return
}
func (c *ControllerHook) PaypelPayHook(ctx context.Context, req *hook.PaypelPayHookReq) (res *hook.PaypelPayHookRes, err error) {
	var (
		r                    = ghttp.RequestFromCtx(ctx)
		AffectedId           int64
		PmsTransactionRefund *entity.PmsTransactionRefund
		tx                   gdb.TX
	)
	g.DB().SetLogger(g.Log().Path("logs/HOOK/PAYPAL_HOOK"))
	g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), r.GetBody())
	if req.EventType == "PAYMENT.CAPTURE.COMPLETED" {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), "支付等待，待划账")
	}
	if req.EventType == "PAYMENT.CAPTURE.COMPLETED" {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), "支付完成")

		if _, err = dao.PmsTransaction.Ctx(ctx).
			Where(dao.PmsTransaction.Columns().TransactionSn, req.Resource.ID).Data(g.Map{
			dao.PmsTransaction.Columns().PayAmount: req.Resource.Amount,
			dao.PmsTransaction.Columns().PayStatus: "DONE",
		}).Update(); err != nil {
			return
		}

		if err = service.PayService().ThirdPayCompleted(ctx, &input_pay.ThirdPayInp{
			//TransNo:   req.Resource.SupplementaryData.RelatedIds.OrderID,
			TransNo: req.Resource.ID,
		}); err != nil {
			g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
			return
		}
	}
	if req.EventType == "PAYMENT.CAPTURE.REFUNDED" {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Debug(r.Context(), "退款完成")
		// 更新退款状态
		if AffectedId, err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().TransNo:      req.Resource.ID,
			dao.PmsTransactionRefund.Columns().RefundStatus: "WAIT",
		}).Data(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
			dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
		}).UpdateAndGetAffected(); err != nil {
			g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
			err = gerror.New("处理退款失败")
			return
		}
		if AffectedId != 1 {
			err = gerror.New("处理退款更新失败")
			g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
			err = gerror.New("处理退款失败")
			return
		}

		// 查询订单号
		if err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().TransNo: req.Resource.ID,
		}).Scan(&PmsTransactionRefund); err != nil {
			return
		}
		tx, err = g.DB().Begin(ctx)
		if err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
		}
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			} else {
				_ = tx.Commit()
			}
		}()
		if err = service.Refund().RefundStatusChange(ctx, tx, PmsTransactionRefund.OrderSn); err != nil {
			return
		}

		ghttp.RequestFromCtx(ctx).Response.WriteExit("SUCCESS")
	}
	return
}
