package mlilifeWxPay

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/guid"
)

// RefundRequest 接口请求参数结构体
type RefundRequest struct {
	Pid           string `json:"pid"`
	TransactionID string `json:"transaction_id"`
	RefundOrderNo string `json:"refund_order_no"`
	RefundFee     string `json:"refund_fee"`
	NonceStr      string `json:"nonce_str"`
	Sign          string `json:"sign"` // 必填 - 加密签名
}

// RefundResponse 接口响应结构体
type RefundResponse struct {
	S    int                `json:"s"`
	Msg  string             `json:"msg"`
	Data RefundResponseData `json:"data"`
}

type RefundResponseData struct {
	OutTradeNo      string `json:"out_trade_no,omitempty"`
	TransactionID   string `json:"transaction_id,omitempty"`
	OutRefundNo     string `json:"out_refund_no,omitempty"`
	RefundID        string `json:"refund_id,omitempty"`
	RefundFee       string `json:"refund_fee,omitempty"`
	CouponRefundFee string `json:"coupon_refund_fee,omitempty"`
	CashFee         string `json:"cash_fee,omitempty"`
	CashRefundFee   string `json:"cash_refund_fee,omitempty"`
}

var (
	refundUrl = "https://pay.mlilife.com/v1/api/refund"
)

func Refund(ctx context.Context, req *RefundRequest) (resp *RefundResponse, err error) {
	var (
		gClientResponse *gclient.Response
	)
	req.Pid = "C03C94ACF4C547748994469DF33EEEDE"
	req.NonceStr = guid.S()
	if req.Sign, err = GetSign(ctx, req, "0C3707F6374745E5A2B5253C6C5A5418"); err != nil {
		return
	}
	if gClientResponse, err = g.Client().Post(ctx, refundUrl, req); err != nil {
		return
	}
	g.Log().Info(ctx, gClientResponse.Raw())
	if gClientResponse.StatusCode != 200 {
		err = gerror.Newf("请求微信退款接口失败，状态码：%d", gClientResponse.StatusCode)
		return
	}
	if err = gjson.New(gClientResponse.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
