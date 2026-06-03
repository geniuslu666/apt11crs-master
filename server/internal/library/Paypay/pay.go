package Paypay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mythrnr/paypayopa-sdk-go"
)

type PayPayResponse struct {
	CodeId            string `json:"codeId"`
	Url               string `json:"url"`
	Deeplink          string `json:"deeplink"`
	ExpiryDate        int    `json:"expiryDate"`
	MerchantPaymentId string `json:"merchantPaymentId"`
	Amount            struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"amount"`
	OrderDescription    string      `json:"orderDescription"`
	OrderItems          interface{} `json:"orderItems"`
	Metadata            interface{} `json:"metadata"`
	CodeType            string      `json:"codeType"`
	StoreInfo           string      `json:"storeInfo"`
	StoreId             string      `json:"storeId"`
	TerminalId          string      `json:"terminalId"`
	RequestedAt         int         `json:"requestedAt"`
	RedirectUrl         string      `json:"redirectUrl"`
	RedirectType        string      `json:"redirectType"`
	IsAuthorization     bool        `json:"isAuthorization"`
	AuthorizationExpiry interface{} `json:"authorizationExpiry"`
}

func Pay(ctx context.Context, orderId string, amount int) (PayPayResponse *PayPayResponse, err error) {
	var (
		PayConfig *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}

	creds := paypayopa.NewCredentials(
		paypayopa.EnvSandbox,
		PayConfig.PaypayApiKey,
		PayConfig.PaypayApiKeySecret,
		PayConfig.PaypayMerchantID,
	)

	wp := paypayopa.NewWebPayment(creds)
	Params := &paypayopa.CreateQRCodePayload{
		MerchantPaymentID: orderId,
		Amount: &paypayopa.MoneyAmount{
			Amount:   amount,
			Currency: paypayopa.CurrencyJPY,
		},
		CodeType:     paypayopa.CodeTypeOrderQR,
		RedirectURL:  PayConfig.PaypayNotifyUrl,
		RedirectType: paypayopa.RedirectTypeDeepLink,
	}
	g.Log().Info(ctx, Params)
	res, info, err := wp.CreateQRCode(ctx, Params)

	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, res)
	g.Log().Info(ctx, info)
	if !info.Success() {
		g.Log().Error(ctx, info)
		err = gerror.New(info.Message)
		return
	}
	if err = gjson.New(res).Scan(&PayPayResponse); err != nil {
		return
	}
	return
}

func Refund(ctx context.Context, orderId string, PaymentID string, amount int, Reason string) (err error) {
	var (
		PayConfig      *model.PayConfig
		RefundResponse *paypayopa.RefundResponse
		ResultInfo     *paypayopa.ResultInfo
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}

	creds := paypayopa.NewCredentials(
		paypayopa.EnvSandbox,
		PayConfig.PaypayApiKey,
		PayConfig.PaypayApiKeySecret,
		PayConfig.PaypayMerchantID,
	)

	wp := paypayopa.NewWebPayment(creds)
	if RefundResponse, ResultInfo, err = wp.RefundPayment(ctx, &paypayopa.RefundPaymentPayload{
		MerchantRefundID: orderId,
		PaymentID:        PaymentID,
		Amount: &paypayopa.MoneyAmount{
			Amount:   amount,
			Currency: paypayopa.CurrencyJPY,
		},
		RequestedAt: gtime.Now().Unix(),
		Reason:      Reason,
	}); err != nil {
		g.Log().Error(ctx, err)
		return
	}

	if RefundResponse.Status == "SUCCESS" {
		return
	} else {
		err = gerror.New(ResultInfo.Message)
		return
	}
}
