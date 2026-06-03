package travel

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CancelOrderNoPayReq struct {
	g.Meta  `path:"/travel/CancelOrder" method:"post" tags:"APP_HOTEL" summary:"[一日游]_取消未支付订单"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type CancelOrderNoPayRes struct{}
