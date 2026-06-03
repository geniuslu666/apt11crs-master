package h5FxPay

import (
	"APT/internal/model"
	"APT/internal/service"
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateOrderRequest struct {
	UserId          string      `json:"user_id" dc:"在接入方侧的唯一身份标识	10001"`
	OrderNo         string      `json:"order_no" dc:"分销业务中心订单号	342183739628978176"`
	TradeNo         string      `json:"trade_no" dc:"分销业务中心交易号*	14342183983167045632"`
	TotalAmount     int64       `json:"total_amount" dc:"订单金额，单位元，两位小数	10.00"`
	OrderExpireTime *gtime.Time `json:"order_expire_time" dc:"订单过期时间*"`
	OrderResultUrl  string      `json:"order_result_url" dc:"订单结果页地址*	https://xxx/orderResult"`
	OrderDetailUrl  string      `json:"order_detail_url" dc:"订单详情页地址*	https://xxx/orderDetail"`
}

type CreateOrderResponse struct {
	Code    int64  `json:"code" dc:"响应码	0"`
	Message string `json:"message" dc:"响应信息	成功"`
	Data    struct {
		ThirdOrderNo string `json:"third_order_no" dc:"接入方受理成功后，向平台返回接入方的订单号	TH342183740904046592"`
		CashierUrl   string `json:"cashier_url" dc:"接入方的收银台地址	https://xxx/cashier"`
	} `json:"data"`
}

func CreateOrder(ctx context.Context, params *CreateOrderRequest) (res *CreateOrderResponse, err error) {
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
	if gHttpResponse, err = gHttpClient.Post(ctx, PayConfig.H5FxDomain+"/external/apt11/order", params); err != nil {
		return
	}
	g.Log().Path("logs/SDK/H5_FX").Info(ctx, gHttpResponse.Raw())
	defer gHttpResponse.Close()
	body = gHttpResponse.ReadAllString()
	if err = json.Unmarshal([]byte(body), &res); err != nil {
		return
	}
	if res.Code != 0 {
		err = gerror.Newf("H5_FX创建订单失败: %s", res.Message)
		return
	}
	return
}
