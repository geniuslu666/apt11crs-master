package kefu

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WelcomeReq struct {
	g.Meta `path:"/welcome/info" method:"post" tags:"NOTIFY_KEFU" summary:"欢迎语和热门问题"`
	Type   string `json:"type" v:"in:WELCOME,OPTION#welcome_content_type_error"`
}

type WelcomeRes struct {
	Welcome string `json:"welcome" dc:"欢迎语"`
	Options []string
}
