package basics

import (
	"APT/internal/model"
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type AccountLoginLogoutReq struct {
	g.Meta `path:"/site/logout" method:"post" tags:"ADMIN" summary:"后台基础_注销登录"`
}

type AccountLoginLogoutRes struct{}

type AccountLoginCaptchaReq struct {
	g.Meta `path:"/site/captcha" method:"get" tags:"ADMIN" summary:"后台基础_获取登录验证码"`
}

type AccountLoginCaptchaRes struct {
	Cid    string `json:"cid" dc:"验证码ID"`
	Base64 string `json:"base64" dc:"验证码"`
}

type AccountLoginReq struct {
	g.Meta `path:"/site/accountLogin" method:"post" tags:"ADMIN" summary:"账号登录"`
	input_basics.AccountLoginInp
}

type AccountLoginRes struct {
	*input_basics.LoginModel
}

type AccountSiteConfigReq struct {
	g.Meta `path:"/site/config" method:"get" tags:"ADMIN" summary:"获取配置"`
}

type AccountSiteConfigRes struct {
	Version string `json:"version"        dc:"系统版本"`
	WsAddr  string `json:"wsAddr"         dc:"客户端websocket地址"`
	Domain  string `json:"domain"         dc:"对外域名"`
	Mode    string `json:"mode"           dc:"运行模式"`
}

type AccountSiteLoginConfigReq struct {
	g.Meta `path:"/site/loginConfig" method:"get" tags:"ADMIN" summary:"获取登录配置"`
}

type AccountSiteLoginConfigRes struct {
	*model.LoginConfig
}

type AccountSitePingReq struct {
	g.Meta `path:"/site/ping" method:"get" tags:"ADMIN" summary:"ping"`
}

type AccountSitePingRes struct{}
