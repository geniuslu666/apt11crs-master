package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberGroupListReq 查询会员分组列表
type MemberGroupListReq struct {
	g.Meta `path:"/pmsMemberGroup/list" method:"get" tags:"ADMIN_PMS" summary:"会员分组_列表"`
	input_app_member.PmsMemberGroupListInp
}

type MemberGroupListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberGroupListModel `json:"list"   dc:"数据列表"`
}

type MemberGroupAllReq struct {
	g.Meta `path:"/pmsMemberGroup/all" method:"get" tags:"ADMIN_PMS" summary:"会员分组_全部列表"`
	input_app_member.PmsMemberGroupAllInp
}

type MemberGroupAllRes struct {
	List []*input_app_member.PmsMemberGroupAllModel `json:"list"   dc:"数据列表"`
}

// MemberGroupExportReq 导出会员分组列表
type MemberGroupExportReq struct {
	g.Meta `path:"/pmsMemberGroup/export" method:"get" tags:"ADMIN_PMS" summary:"会员分组_导出"`
	input_app_member.PmsMemberGroupListInp
}

type MemberGroupExportRes struct{}

// MemberGroupViewReq 获取会员分组指定信息
type MemberGroupViewReq struct {
	g.Meta `path:"/pmsMemberGroup/view" method:"get" tags:"ADMIN_PMS" summary:"会员分组_详情"`
	input_app_member.PmsMemberGroupViewInp
}

type MemberGroupViewRes struct {
	*input_app_member.PmsMemberGroupViewModel
}

// MemberGroupEditReq 修改/新增会员分组
type MemberGroupEditReq struct {
	g.Meta `path:"/pmsMemberGroup/edit" method:"post" tags:"ADMIN_PMS" summary:"会员分组_修改/新增"`
	input_app_member.PmsMemberGroupEditInp
}

type MemberGroupEditRes struct{}

// MemberGroupDeleteReq 删除会员分组
type MemberGroupDeleteReq struct {
	g.Meta `path:"/pmsMemberGroup/delete" method:"post" tags:"ADMIN_PMS" summary:"会员分组_删除"`
	input_app_member.PmsMemberGroupDeleteInp
}

type MemberGroupDeleteRes struct{}
