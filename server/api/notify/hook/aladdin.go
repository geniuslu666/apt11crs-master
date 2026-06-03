package hook

import "github.com/gogf/gf/v2/frame/g"

type AladdinReq struct {
	g.Meta  `path:"/aladdin/checkout" method:"post" tags:"NOTIFY_HOOK" summary:"Aladdin 核销通知"`
	OrderId string `json:"orderId" dc:"订单ID"`
	OrderNo string `json:"orderNo" dc:"订单号"`
	Status  string `json:"status" dc:"核销状态 success-成功 fail-失败"`
}

type AladdinRes struct{}
