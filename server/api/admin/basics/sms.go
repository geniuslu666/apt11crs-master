package basics

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type SendTestSmsReq struct {
	g.Meta `path:"/sms/sendTest" tags:"ADMIN" summary:"短信_发送测试短信" method:"post"`
	input_basics.SendCodeInp
}

type SendTestSmsRes struct {
}

type SendBindSmsReq struct {
	g.Meta `path:"/sms/sendBind" tags:"ADMIN" summary:"短信_发送换绑短信" method:"post"`
}

type SendBindSmsRes struct {
}

type SendSmsReq struct {
	g.Meta `path:"/sms/send" tags:"ADMIN" summary:"短信_发送短信" method:"post"`
	input_basics.SendCodeInp
}

type SendSmsRes struct {
}
