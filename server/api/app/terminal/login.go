package terminal

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TerminalLoginReq struct {
	g.Meta     `path:"/login" method:"post" tags:"APP_TERMINAL" summary:"[终端]登录"`
	TerminalSn string `json:"terminalSn"       v:"required#device_number_unknown" dc:"设备号"`
	Account    string `json:"account"       v:"required#account_unknown" dc:"账号"`
	Password   string `json:"password"   v:"required#password_unknown" dc:"密码"`
}

type TerminalLoginRes struct {
	Token   string `json:"token" dc:"用户token令牌身份"`
	Expires int64  `json:"expires" dc:"用户token令牌有效期"`
}

type TerminalLogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"APP_TERMINAL" summary:"[终端]登出"`
}

type TerminalLogoutRes struct{}
