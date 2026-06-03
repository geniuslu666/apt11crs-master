package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberLevelListReq struct {
	g.Meta `path:"/pmsMemberLevel/list" method:"get" tags:"ADMIN_PMS" summary:"会员等级_列表"`
	input_app_member.PmsMemberLevelListInp
}

type MemberLevelListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberLevelListModel `json:"list"   dc:"数据列表"`
}

type MemberLevelAllReq struct {
	g.Meta `path:"/pmsMemberLevel/all" method:"get" tags:"ADMIN_PMS" summary:"会员等级_全部列表"`
	input_app_member.PmsMemberLevelAllInp
}

type MemberLevelAllRes struct {
	List []*input_app_member.PmsMemberLevelAllModel `json:"list"   dc:"数据列表"`
}

type MemberLevelViewReq struct {
	g.Meta `path:"/pmsMemberLevel/view" method:"get" tags:"ADMIN_PMS" summary:"会员等级_详情"`
	input_app_member.PmsMemberLevelViewInp
}

type MemberLevelViewRes struct {
	*input_app_member.PmsMemberLevelViewModel
}

type MemberLevelEditReq struct {
	g.Meta `path:"/pmsMemberLevel/edit" method:"post" tags:"ADMIN_PMS" summary:"会员等级_编辑"`
	input_app_member.PmsMemberLevelEditInp
}

type MemberLevelEditRes struct{}

type MemberLevelDeleteReq struct {
	g.Meta `path:"/pmsMemberLevel/delete" method:"post" tags:"ADMIN_PMS" summary:"会员等级_删除"`
	input_app_member.PmsMemberLevelDeleteInp
}

type MemberLevelDeleteRes struct{}
