package basics

import "github.com/gogf/gf/v2/frame/g"

type ChatReq struct {
	g.Meta   `path:"/chat/:kefu_id/:ent_id/:lang/:member_id" method:"get" tags:"APP_BASICS" summary:"客服_客服客户端地址"`
	KefuId   string `json:"KefuId"`
	EntId    int    `json:"ent_id"`
	Lang     string `json:"lang"`
	MemberId int    `json:"member_id"`
}

type ChatRes struct{}
