package hook

import "github.com/gogf/gf/v2/frame/g"

type VisitorMessageCallUrlReq struct {
	g.Meta  `path:"/kefu/VisitorMessageCallUrl" method:"post" tags:"NOTIFY_HOOK" summary:"客服WebHook"`
	Type    string `json:"type"`
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"conntent"`
	OrderSn string `json:"orderSn"`
}

type VisitorMessageCallUrlRes struct{}
