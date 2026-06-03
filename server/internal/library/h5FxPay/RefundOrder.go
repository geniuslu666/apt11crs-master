package h5FxPay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

type RefundOrderRequest struct {
	RefundOrderNo string `json:"refundOrderNo"` // 分销中心退款订单号，标识一次退款请求
	OrderNo       string `json:"orderNo"`       // 分销中心业务订单号，标识退款所属的业务订单
	TradeNo       string `json:"tradeNo"`       // 分销中心业务交易号，标识退款所属的交易号
	RefundAmount  int64  `json:"refundAmount"`  // 本次退款金额，单位元，2位小数，订单可多次进行退款
	TradeTime     string `json:"tradeTime"`     // 用户发起退款的UTC时间戳，10位
}

type RefundOrderResponse struct {
	Code    int64  `json:"code" dc:"响应码	0"`
	Message string `json:"message" dc:"响应信息	成功"`
	Data    struct {
		ThirdRefundOrderNo string `json:"thirdRefundOrderNo" dc:"接入方受理成功后，向平台返回接入方的退款单号"`
	} `json:"data"`
}

func RefundOrder(ctx context.Context, params *RefundOrderRequest) (res *RefundOrderResponse, err error) {
	var (
		body          string
		gHttpClient   = g.Client()
		gHttpResponse *gclient.Response
		PayConfig     *model.PayConfig
	)
	if PayConfig, err = service.BasicsConfig().GetPay(ctx); err != nil {
		return
	}
	gHttpClient.SetHeader("Content-Type", "application/json")
	if gHttpResponse, err = gHttpClient.Post(ctx, PayConfig.H5FxDomain+"/external/apt11/refund", params); err != nil {
		return
	}
	g.Log().Path("logs/SDK/H5_FX").Info(ctx, gHttpResponse.Raw())
	defer gHttpResponse.Close()
	body = gHttpResponse.ReadAllString()
	if err = json.Unmarshal([]byte(body), &res); err != nil {
		return
	}
	if res.Code != 0 {
		err = gerror.New(res.Message)
		return
	}
	return
}
