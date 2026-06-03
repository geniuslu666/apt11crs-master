package basics

import "github.com/gogf/gf/v2/frame/g"

type LinkReq struct {
	g.Meta     `path:"/app/:referrer_id" method:"get" tags:"APP_BASICS" summary:"appLink"`
	ReferrerId int `json:"referrer_id"`
}

type LinkRes struct {
}

type WxReturnUrlReq struct {
	g.Meta `path:"/wxReturnUrl" method:"get" tags:"APP_BASICS" summary:"微信支付回调"`
}
type WxReturnUrlRes struct{}
