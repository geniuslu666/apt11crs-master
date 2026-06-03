package stripePay

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/refund"
)

func (c *Client) Refund(ctx context.Context, RefundOrderSn string, RefundAmount int64, PaymentIntent string) (result *stripe.Refund, err error) {
	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(PaymentIntent),
	}
	params.Amount = stripe.Int64(RefundAmount)
	params.AddMetadata("RefundOrderSn", RefundOrderSn)

	if result, err = refund.New(params); err != nil {
		g.Log().Path("logs/SDK/STRIPE").Info(ctx, err)
		return
	}
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, result.APIResource.LastResponse.StatusCode)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, result.APIResource.LastResponse.Status)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, result.APIResource.LastResponse.Header)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, result.APIResource.LastResponse.RawJSON)
	if g.IsEmpty(result) {
		err = gerror.New("stripe error")
		return
	}
	return
}
