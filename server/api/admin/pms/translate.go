package pms

import "github.com/gogf/gf/v2/frame/g"

type TranslateTextReq struct {
	g.Meta `path:"/translate/text" method:"get" tags:"ADMIN_PMS" summary:"翻译文本"`
	Text   string `json:"text" v:"required#翻译文本不能为空" dc:"翻译文本"`
}

type TranslateTextRes struct {
	Zh   string `json:"zh" dc:"简体中文翻译结果"`
	ZhCN string `json:"zh_CN" dc:"繁体翻译结果"`
	Ko   string `json:"ko" dc:"韩文翻译结果"`
	Ja   string `json:"ja" dc:"日文翻译结果"`
	En   string `json:"en" dc:"英文翻译结果"`
}
