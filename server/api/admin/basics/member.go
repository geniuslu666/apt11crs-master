package basics

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberUpdateEmailReq struct {
	g.Meta `path:"/member/updateEmail" method:"post" tags:"ADMIN" summary:"用户_换绑邮箱"`
	input_basics.MemberUpdateEmailInp
}

type MemberUpdateEmailRes struct{}

type MemberUpdateMobileReq struct {
	g.Meta `path:"/member/updateMobile" method:"post" tags:"ADMIN" summary:"用户_换绑手机号"`
	input_basics.MemberUpdateMobileInp
}

type MemberUpdateMobileRes struct{}

type MemberUpdateProfileReq struct {
	g.Meta `path:"/member/updateProfile" method:"post" tags:"ADMIN" summary:"用户_更新用户资料"`
	input_basics.MemberUpdateProfileInp
}

type MemberUpdateProfileRes struct{}

type MemberUpdatePwdReq struct {
	g.Meta `path:"/member/updatePwd" method:"post" tags:"ADMIN" summary:"用户_重置密码"`
	input_basics.MemberUpdatePwdInp
}

type MemberUpdatePwdRes struct{}

type MemberResetPwdReq struct {
	g.Meta `path:"/member/resetPwd" method:"post" tags:"ADMIN" summary:"重置密码"`
	input_basics.MemberResetPwdInp
}

type MemberResetPwdRes struct{}

type MemberListReq struct {
	g.Meta `path:"/member/list" method:"get" tags:"ADMIN" summary:"获取用户列表"`
	input_basics.MemberListInp
}

type MemberListRes struct {
	List []*input_basics.MemberListModel `json:"list"   dc:"数据列表"`
	input_form.PageRes
}

type MemberViewReq struct {
	g.Meta `path:"/member/view" method:"get" tags:"ADMIN" summary:"获取指定信息"`
	input_basics.MemberViewInp
}

type MemberViewRes struct {
	*input_basics.MemberViewModel
}

type MemberEditReq struct {
	g.Meta `path:"/member/edit" method:"post" tags:"ADMIN" summary:"修改/新增用户"`
	input_basics.MemberEditInp
}

type MemberEditRes struct{}

type MemberDeleteReq struct {
	g.Meta `path:"/member/delete" method:"post" tags:"ADMIN" summary:"删除用户"`
	input_basics.MemberDeleteInp
}

type MemberDeleteRes struct{}

type MemberStatusReq struct {
	g.Meta `path:"/member/status" method:"post" tags:"ADMIN" summary:"更新用户状态"`
	input_basics.MemberStatusInp
}

type MemberStatusRes struct{}

type MemberSelectReq struct {
	g.Meta `path:"/member/option" method:"get" tags:"ADMIN" summary:"获取可选的后台用户选项"`
	input_basics.MemberSelectInp
}

type MemberSelectRes []*input_basics.MemberSelectModel

type MemberInfoReq struct {
	g.Meta `path:"/member/info" method:"get" tags:"ADMIN" summary:"获取登录用户信息"`
}

type MemberInfoRes struct {
	*input_basics.LoginMemberInfoModel
}
