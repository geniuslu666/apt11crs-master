package hook

import "github.com/gogf/gf/v2/frame/g"

type OrderPrintReq struct {
	g.Meta  `path:"/order/print" method:"post" tags:"NOTIFY_HOOK" summary:"内部_打印订单"`
	OrderSn string `json:"orderSn" v:"required#order_number_unknown" dc:"订单号"`
	Scene   string `json:"scene" v:"required#scene_unknown" dc:"订单场景值"`
}

type OrderPrintRes struct{}
