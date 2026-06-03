package mlilifeWxPay

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/guid"
)

type WxappQueryParams struct {
	// 支付系统分配的唯一标识符（必须）
	Pid string `json:"pid"`

	// 微信支付订单号（必填，与商户订单号二选一）
	//TransactionID string `json:"transaction_id,omitempty"`

	// 随机字符串，防止重复请求（必填）
	NonceStr string `json:"nonce_str"`

	// 商户系统自定义的订单号（非必填，与微信支付订单号二选一）
	OutTradeNo string `json:"out_trade_no"`

	Sign string `json:"sign"` // 必填，加密串
}

type WxappQueryResponse struct {
	S    int         `json:"s"`    // 必填，状态码。200:成功;400:失败
	Msg  string      `json:"msg"`  // 必填，状态描述
	Data QueryDetail `json:"data"` // 必填，支付明细数据
}

type QueryDetail struct {
	TradeType      string `json:"trade_type"`       // 必填，交易类型（JSAPI/NATIVE/APP/MICROPAY）
	BankType       string `json:"bank_type"`        // 必填，银行类型
	TotalFee       int    `json:"total_fee"`        // 必填，标价金额（单位为最小单位）
	FeeType        string `json:"fee_type"`         // 必填，标价币种
	TransactionID  string `json:"transaction_id"`   // 必填，微信支付订单号
	TimeEnd        string `json:"time_end"`         // 必填，支付完成时间
	TradeState     string `json:"trade_state"`      // 必填，交易状态（SUCCESS/REFUND/NOTPAY等）
	SubAppID       string `json:"sub_appid"`        // 必填，子商户公众账号ID
	SubOpenID      string `json:"sub_openid"`       // 必填，用户子标识
	SubIsSubscribe string `json:"sub_is_subscribe"` // 必填，是否关注子公众账号
	CashFee        string `json:"cash_fee"`         // 必填，用户支付金额
	CashFeeType    string `json:"cash_fee_type"`    // 必填，现金支付货币类型
	TradeStateDesc string `json:"trade_state_desc"` // 必填，支付结果描述
	Rate           string `json:"rate"`             // 必填，支付费率
	Pid            string `json:"pid"`              // 必填，支付系统分配
	OutTradeNo     string `json:"out_trade_no"`     // 必填，商户订单号
	NonceStr       string `json:"nonce_str"`        // 必填，随机字符串
	Sign           string `json:"sign"`             // 必填，加密串
	Attach         string `json:"attach"`           // 必填，附加数据
}

var (
	wxappQueryUrl = "https://pay.mlilife.com/v1/api/query"
)

func WxappQuery(ctx context.Context, req *WxappQueryParams) (resp *WxappQueryResponse, err error) {
	var (
		gClientResponse *gclient.Response
		ParamsByte      []byte
	)
	req.Pid = "C03C94ACF4C547748994469DF33EEEDE"
	req.NonceStr = guid.S()
	if req.Sign, err = GetSign(ctx, req, "0C3707F6374745E5A2B5253C6C5A5418"); err != nil {
		return
	}
	if ParamsByte, err = json.Marshal(req); err != nil {
		return
	}
	if gClientResponse, err = g.Client().Post(ctx, wxappQueryUrl, ParamsByte); err != nil {
		return
	}
	g.Log().Info(ctx, gClientResponse.Raw())
	if gClientResponse.StatusCode != 200 {
		err = gerror.Newf("微信支付操作失败，状态码：%d", gClientResponse.StatusCode)
		return
	}
	if err = gjson.New(gClientResponse.ReadAllString()).Scan(&resp); err != nil {
		return
	}
	return
}
