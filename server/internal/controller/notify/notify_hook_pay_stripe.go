package notify

import (
	"APT/internal/dao"
	"APT/internal/library/stripePay"
	"APT/internal/model"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/webhook"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/hook"
)

func (c *ControllerHook) StripeWebhook(ctx context.Context, req *hook.StripeWebhookReq) (res *hook.StripeWebhookRes, err error) {
	var (
		r                    = ghttp.RequestFromCtx(ctx)
		PmsTransactionRefund *entity.PmsTransactionRefund
		PmsTransaction       *entity.PmsTransaction
		payload              = r.GetBody()
		event                stripe.Event
		stripeClient         *stripePay.Client
		tx                   gdb.TX
	)
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Info(ctx, req)
	res = new(hook.StripeWebhookRes)
	if stripeClient, err = stripePay.GetClient(ctx, "", ""); err != nil {
		r.Response.WriteStatusExit(http.StatusBadRequest, err.Error())
		return
	}
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Info(ctx, stripeClient.SignKey)
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Info(ctx, r.Request.Header.Get("Stripe-Signature"))
	if event, err = webhook.ConstructEvent(payload, r.Request.Header.Get("Stripe-Signature"), stripeClient.SignKey); err != nil {
		g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
		return
	}
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Debug(ctx, event.Type)
	switch event.Type {
	case "charge.refund.updated", "charge.refunded":
		g.Log().Path("logs/HOOK/STRIPE_HOOK").Info(ctx, "退款通知")
		if req.Data.Object.Status != "succeeded" {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Errorf(ctx, "支付状态%s", req.Data.Object.Status)
			return
		}
		// 查询支付订单
		if err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().CaptureId: req.Data.Object.PaymentIntent,
		}).Scan(&PmsTransaction); err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Errorf(ctx, err.Error())
			return
		}
		if g.IsEmpty(PmsTransaction) {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Errorf(ctx, "支付订单不存在")
			return
		}
		// 更新退款状态
		if _, err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().TransactionSn: PmsTransaction.TransactionSn,
			dao.PmsTransactionRefund.Columns().RefundStatus:  "WAIT",
		}).Data(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().RefundStatus: "DONE",
			dao.PmsTransactionRefund.Columns().RefundTime:   gtime.Now(),
		}).UpdateAndGetAffected(); err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
			return
		}
		// 查询订单号
		if err = dao.PmsTransactionRefund.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransactionRefund.Columns().TransactionSn: PmsTransaction.TransactionSn,
		}).Scan(&PmsTransactionRefund); err != nil {
			return
		}
		tx, err = g.DB().Begin(ctx)
		if err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
			return
		}
		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()
		if err = service.Refund().RefundStatusChange(ctx, tx, PmsTransactionRefund.OrderSn); err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
			r.Response.WriteStatusExit(http.StatusOK, err.Error())
			return
		}
		break
	case "checkout.session.completed":
		g.Log().Path("logs/HOOK/STRIPE_HOOK").Debug(ctx, "支付完成通知")
		if g.IsEmpty(req.Data.Object.ClientReferenceID) {
			r.Response.WriteStatusExit(http.StatusOK, "未识别到订单信息,不是APP创建订单")
			return
		}
		if _, err = dao.PmsTransaction.Ctx(ctx).Where(g.MapStrAny{
			dao.PmsTransaction.Columns().PaymentRequestId: req.Data.Object.ClientReferenceID,
		}).Update(g.MapStrAny{
			dao.PmsTransaction.Columns().CaptureId: req.Data.Object.PaymentIntent,
			dao.PmsTransaction.Columns().PayAmount: req.Data.Object.AmountSubtotal,
			dao.PmsTransaction.Columns().PayCharge: req.Data.Object.TotalDetails.AmountTax,
			dao.PmsTransaction.Columns().PayStatus: "DONE",
		}); err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
			r.Response.WriteStatusExit(http.StatusBadRequest, "PmsTransaction Update Error")
			return
		}
		// TODO 处理支付完成逻辑
		if err = service.PayService().ThirdPayCompleted(ctx, &input_pay.ThirdPayInp{
			TransNo:   req.Data.Object.ClientReferenceID,
			CaptureId: req.Data.Object.PaymentIntent,
		}); err != nil {
			g.Log().Path("logs/HOOK/STRIPE_HOOK").Error(ctx, err)
			err = nil
		}
		g.Log().Path("logs/HOOK/STRIPE_HOOK").Debug(ctx, "处理完成")
		break
	default:
		err = gerror.New(fmt.Sprintf("Error webhook event: %v", event.Type))
		r.Response.WriteStatusExit(http.StatusBadRequest, "Error webhook event")
		return
	}
	r.Response.WriteStatusExit(http.StatusOK)
	return
}
func (c *ControllerHook) StripeSuccess(ctx context.Context, _ *hook.StripeSuccessReq) (res *hook.StripeSuccessRes, err error) {
	var (
		r         = ghttp.RequestFromCtx(ctx)
		PayConfig *model.PayConfig
	)
	res = new(hook.StripeSuccessRes)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Debug(r.Context(), r.GetBody())
	r.Response.RedirectTo(PayConfig.StripeSuccessAPP)
	return
}
func (c *ControllerHook) StripeCancel(ctx context.Context, _ *hook.StripeCancelReq) (res *hook.StripeCancelRes, err error) {
	var (
		r         = ghttp.RequestFromCtx(ctx)
		PayConfig *model.PayConfig
	)
	res = new(hook.StripeCancelRes)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	g.Log().Path("logs/HOOK/STRIPE_HOOK").Debug(r.Context(), r.GetBody())
	r.Response.RedirectTo(PayConfig.StripeCancelAPP)
	return
}
