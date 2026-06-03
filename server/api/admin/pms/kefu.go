package pms

import "github.com/gogf/gf/v2/frame/g"

type OssKefuLoginReq struct {
	g.Meta   `path:"/kefu/oss" method:"get" tags:"ADMIN_PMS" summary:"客服_免登跳转固定页面"`
	Redirect string `json:"redirect" v:"required#跳转地址不能为空" dc:"跳转地址"`
}

type OssKefuToken struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result struct {
		Token string `json:"token"`
	} `json:"result"`
}

type OssKefuLoginRes struct {
	IframeUrl string `json:"iframeUrl" dc:"客服联登地址"`
}
