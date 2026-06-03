package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberCancelListReq struct {
	g.Meta `path:"/pmsMemberCancel/list" method:"get" tags:"ADMIN_PMS" summary:"会员信息_注销列表"`
	input_app_member.PmsMemberCancelListInp
}

type MemberCancelListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberCancelListModel `json:"list"   dc:"数据列表"`
}

type MemberCancelAgreeReq struct {
	g.Meta `path:"/pmsMemberCancel/agree" method:"post" tags:"ADMIN_PMS" summary:"会员注销_同意"`
	input_app_member.PmsMemberCancelAgreeInp
}

type MemberCancelAgreeRes struct {
}

type MemberCancelDisagreeReq struct {
	g.Meta `path:"/pmsMemberCancel/disagree" method:"post" tags:"ADMIN_PMS" summary:"会员注销_拒绝"`
	input_app_member.PmsMemberCancelDisagreeInp
}

type MemberCancelDisagreeRes struct{}
