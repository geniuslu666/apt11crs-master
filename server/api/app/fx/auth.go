package fx

import (
	"APT/internal/model/input/input_app_member"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/member/login" method:"post" tags:"APP_MEMBER_FX" summary:"授权_登录"`
	*input_app_member.FxMemberInfo
}

type LoginRes struct {
	LoginUrl string `json:"login_url" dc:"用户登录链接"`
}
