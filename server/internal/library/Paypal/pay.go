package Paypal

import (
	"APT/internal/library/cache"
	"APT/internal/model"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/paypal"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/util/guid"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ClientID  string
	Secret    string
	ReturnUrl string
	CancelUrl string
	Client    *paypal.Client
	IsProd    bool
}

func GetClient(ctx context.Context, PaypalReturnUrlApp string, PaypalCleanUrlApp string) (PaypalClient *Client, err error) {
	var (
		PayConfig *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	if !g.IsEmpty(PaypalReturnUrlApp) {
		PayConfig.PaypalReturnUrlApp = PaypalReturnUrlApp
	}
	if !g.IsEmpty(PaypalCleanUrlApp) {
		PayConfig.PaypalCleanUrlApp = PaypalCleanUrlApp
	}
	PaypalClient = &Client{
		ClientID:  PayConfig.PaypalClientID,
		Secret:    PayConfig.PaypalSecret,
		ReturnUrl: PayConfig.PaypalReturnUrlApp,
		CancelUrl: PayConfig.PaypalCleanUrlApp,
		IsProd:    PayConfig.PaypalIsProd,
	}
	// 初始化PayPal客户端
	if PaypalClient.Client, err = paypal.NewClient(PaypalClient.ClientID, PaypalClient.Secret, PaypalClient.IsProd); err != nil {
		return
	}
	return
}

func (c *Client) GetAccessToken(ctx context.Context) (token string, err error) {
	var (
		paypalAccessToken *paypal.AccessToken
		CacheToken        *gvar.Var
	)
	CacheToken = cache.Instance().MustGet(ctx, "paypalToken")
	if !g.IsEmpty(CacheToken) {
		token = CacheToken.String()
		return
	}
	gmlock.Lock("paypalToken")
	defer gmlock.Unlock("paypalToken")
	// 开启调试日志
	c.Client.DebugSwitch = gopay.DebugOn

	// 获取AccessToken
	if paypalAccessToken, err = c.Client.GetAccessToken(); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	// 互斥锁

	token = paypalAccessToken.AccessToken
	if err = cache.Instance().Set(ctx, "paypalToken", token, time.Duration(paypalAccessToken.ExpiresIn-100)*time.Second); err != nil {
		return
	}
	return
}

func (c *Client) CreatOrder(ctx context.Context, language string, amount float64) (resp *paypal.CreateOrderRsp, err error) {
	var (
		bm             = make(gopay.BodyMap)
		paypalLanguage string
	)
	if c.Client.AccessToken, err = c.GetAccessToken(ctx); err != nil {
		return
	}
	switch language {
	case "zh":
		paypalLanguage = "zh-CN"
	case "zh_CN":
		paypalLanguage = "zh-TW"
	case "ja":
		paypalLanguage = "ja-JP"
	case "ko":
		paypalLanguage = "ko-KR"
	case "en":
		paypalLanguage = "en-US"
	}
	var pus []*paypal.PurchaseUnit
	var item = &paypal.PurchaseUnit{
		ReferenceId: guid.S(),
		Amount: &paypal.Amount{
			CurrencyCode: "JPY",
			Value:        gvar.New(amount).String(),
		},
	}
	pus = append(pus, item)
	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus).
		SetBodyMap("payment_source", func(b gopay.BodyMap) {
			b.SetBodyMap("paypal", func(bb gopay.BodyMap) {
				bb.SetBodyMap("experience_context", func(bbb gopay.BodyMap) {
					bbb.Set("brand_name", "gopay").
						Set("locale", paypalLanguage).
						Set("shipping_preference", "NO_SHIPPING").
						Set("user_action", "PAY_NOW").
						Set("return_url", c.ReturnUrl).
						Set("cancel_url", c.CancelUrl)
				})
			})
		})
	g.Log().Debug(ctx, bm)
	// 创建支付订单
	if resp, err = c.Client.CreateOrder(ctx, bm); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, resp.Response)
	if resp.Response.Status != "PAYER_ACTION_REQUIRED" {
		err = gerror.New("paypal支付异常")
	}
	return
}

func (c *Client) CreatCardOrder(ctx context.Context, cardPayParams *input_pay.PaypalCardPayParams) (respBody *paypal.OrderDetail, err error) {
	var (
		bm            = make(gopay.BodyMap)
		respBodyBytes []byte
		resp          *http.Response
		req           *http.Request
		httpclient    = &http.Client{}
		url           string
	)
	if c.Client.AccessToken, err = c.GetAccessToken(ctx); err != nil {
		return
	}
	if c.IsProd == true {
		url = "https://api-m.paypal.com/v2/checkout/orders"
	} else {
		url = "https://api-m.sandbox.paypal.com/v2/checkout/orders"
	}
	var pus []*paypal.PurchaseUnit
	var item = &paypal.PurchaseUnit{
		ReferenceId: guid.S(),
		Amount: &paypal.Amount{
			CurrencyCode: "JPY",
			Value:        "8",
		},
	}
	pus = append(pus, item)

	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus).
		SetBodyMap("payment_source", func(b gopay.BodyMap) {
			b.SetBodyMap("card", func(bb gopay.BodyMap) {
				bb.Set("name", cardPayParams.CardName).
					Set("number", cardPayParams.CardNumber).
					Set("security_code", cardPayParams.CardCvv).
					Set("expiry", cardPayParams.CardExp).SetBodyMap("experience_context", func(bbb gopay.BodyMap) {
					bbb.Set("return_url", c.ReturnUrl).
						Set("cancel_url", c.CancelUrl)
				}).SetBodyMap("billing_address", func(c gopay.BodyMap) {
					c.Set("country_code", cardPayParams.CountryCode).
						Set("postal_code", cardPayParams.PostalCode).
						Set("address_line_1", cardPayParams.Address)
				})
			})
		})

	g.Log().Info(ctx, bm)
	// 创建 HTTP 请求
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(gvar.New(bm).Bytes()))
	if err != nil {
		g.Log().Debug(ctx, "Failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Client.AccessToken)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Paypal-Request-Id", guid.S())

	// 打印请求头和请求体
	g.Log().Debug(ctx, "+---------------------------------------------+\n|                   REQUEST                   |\n+---------------------------------------------+")
	for key, values := range req.Header {
		for _, value := range values {
			g.Log().Debugf(ctx, "%-30s: %s\n", key, value)
		}
	}
	g.Log().Debug(ctx, "Request Body:")
	g.Log().Debug(ctx, gvar.New(bm).String())

	// 使用 HTTP 客户端发送请求
	resp, err = httpclient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Debug(ctx, "Failed to read response body: %v", err)
	}

	// 打印响应头和响应体

	fmt.Println("+---------------------------------------------+\n|                   RESPONSE                  |\n+---------------------------------------------+")
	fmt.Println("\nResponse Status Code:", resp.StatusCode)
	for key, values := range resp.Header {
		for _, value := range values {
			g.Log().Debugf(ctx, "%-30s: %s\n", key, value)
		}
	}
	g.Log().Debug(ctx, "Response Body:")
	g.Log().Debug(ctx, string(respBodyBytes))

	if err = json.Unmarshal(respBodyBytes, &respBody); err != nil {
		return
	}

	if respBody.Status != "COMPLETED" {
		err = gerror.New("PaypalCard支付异常")
	}
	return
}

func (c *Client) CaptureOrder(ctx context.Context, payerId, orderId string) (resp *paypal.OrderCaptureRsp, err error) {
	var (
		bm = make(gopay.BodyMap)
	)
	if c.Client.AccessToken, err = c.GetAccessToken(ctx); err != nil {
		return
	}
	bm.Set("payer_id", payerId)
	g.Log().Debug(ctx, bm)
	// 获取订单详情
	if resp, err = c.Client.OrderCapture(ctx, orderId, bm); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, resp.Response)
	if resp.Response.Status != "COMPLETED" {
		err = gerror.New("批准支付失败")
	}
	return
}

func (c *Client) RefundOrder(ctx context.Context, orderId string, CustomId string, amount float64, isAllRefund bool) (Id string, err error) {
	var (
		resp *paypal.PaymentCaptureRefundRsp
		bm   = make(gopay.BodyMap)
	)
	if c.Client.AccessToken, err = c.GetAccessToken(ctx); err != nil {
		return
	}
	bm = bm.Set("custom_id", CustomId)
	if !isAllRefund {
		bm = bm.SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("value", amount).
				Set("currency_code", "JPY")
		}).Set("custom_id", CustomId)
	}

	g.Log().Debug(ctx, bm)
	if resp, err = c.Client.PaymentCaptureRefund(ctx, orderId, bm); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, resp.Response)
	if resp.Response.Status != "COMPLETED" {
		err = gerror.New("退款失败")
		return
	}
	Id = resp.Response.Id
	return
}
