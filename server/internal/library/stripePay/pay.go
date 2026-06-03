package stripePay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type Client struct {
	Key        string
	SignKey    string
	SuccessURL string
	CancelURL  string
}

func GetClient(ctx context.Context, PaypalReturnUrlApp string, PaypalCleanUrlApp string) (StripeClient *Client, err error) {
	var (
		PayConfig *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	if !g.IsEmpty(PaypalReturnUrlApp) {
		PayConfig.StripeSuccessURL = PaypalReturnUrlApp
	}
	if !g.IsEmpty(PaypalCleanUrlApp) {
		PayConfig.StripeCancelURL = PaypalCleanUrlApp
	}

	StripeClient = &Client{
		Key:        PayConfig.StripeKey,
		SignKey:    PayConfig.StripeSignKey,
		SuccessURL: PayConfig.StripeSuccessURL,
		CancelURL:  PayConfig.StripeCancelURL,
	}
	stripe.Key = StripeClient.Key
	return
}

func (c *Client) Pay(ctx context.Context, Amount int64, Currency string) (s *stripe.CheckoutSession, err error) {
	var (
		params  *stripe.CheckoutSessionParams
		orderSn = guid.S()
	)
	params = &stripe.CheckoutSessionParams{
		ClientReferenceID: stripe.String(orderSn),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String(Currency),
					UnitAmount: stripe.Int64(Amount),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Apartment 11"),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:   stripe.String(c.SuccessURL),
		CancelURL:    stripe.String(c.CancelURL),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}
	g.Log().Path("logs/SDK/NX_CLOUD").Info(ctx, params)
	if s, err = session.New(params); err != nil {
		g.Log().Path("logs/SDK/STRIPE").Info(ctx, err)
		return
	}
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, s.APIResource.LastResponse.StatusCode)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, s.APIResource.LastResponse.Status)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, s.APIResource.LastResponse.Header)
	g.Log().Path("logs/SDK/STRIPE").Info(ctx, s.APIResource.LastResponse.RawJSON)
	if g.IsEmpty(s) {
		err = gerror.New("stripe error")
		return
	}
	return
}
