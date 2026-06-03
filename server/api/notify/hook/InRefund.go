package hook

import "github.com/gogf/gf/v2/frame/g"

type InRefundReq struct {
	g.Meta  `path:"/in/refund" method:"post" tags:"NOTIFY_HOOK" summary:"内部_退款"`
	OrderSn string  `json:"orderSn" v:"required#order_number_unknown" dc:"订单号"`
	Amount  float64 `json:"amount" v:"required#refund_amount_unknown" dc:"退款金额"`
}

type InRefundRes struct{}
