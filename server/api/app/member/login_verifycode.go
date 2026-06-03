package member

import "github.com/gogf/gf/v2/frame/g"

type GeeTestValidateReq struct {
	CaptchaOutput string `json:"captcha_output" dc:"极验输出信息"`
	LotNumber     string `json:"lot_number"     dc:"极验验证流水号"`
	PassToken     string `json:"pass_token"     dc:"极验验证通过标识"`
	GenTime       string `json:"gen_time"       dc:"极验验证通过时间戳"`
}

type EmailSendCodeReq struct {
	g.Meta `path:"/member/SendEmailCode" method:"post" tags:"APP_MEMBER" summary:"授权_发送邮箱验证码"`
	EMail  string `json:"email" v:"required#please_fill_in_your_email_address" dc:"邮箱"`
}

type EmailSendCodeRes struct{}

type SmsSendCodeReq struct {
	g.Meta             `path:"/member/SendPhoneCode" method:"post" tags:"APP_MEMBER" summary:"授权_发送手机验证码"`
	Phone              string `json:"phone" v:"required#phone_number_required" dc:"手机"`
	AreaNo             string `json:"area_no" v:"required#phone_area_code_required" dc:"手机区号-例如国内+86 传入+86"`
	GeeTestValidateReq `json:",inline"`
}

type SmsSendCodeRes struct{}
