package hotel

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CancelOrderNoPayReq struct {
	g.Meta  `path:"/home/CancelOrder" method:"post" tags:"APP_HOTEL" summary:"酒店订单_取消未支付订单"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type CancelOrderNoPayRes struct{}
