package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberLogListReq struct {
	g.Meta `path:"/pmsMemberLog/list" method:"get" tags:"ADMIN_PMS" summary:"会员登录日志_列表"`
	input_app_member.PmsMemberLogListInp
}

type MemberLogListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberLogListModel `json:"list"   dc:"数据列表"`
}

type MemberLogExportReq struct {
	g.Meta `path:"/pmsMemberLog/export" method:"get" tags:"ADMIN_PMS" summary:"会员登录日志_导出"`
	input_app_member.PmsMemberLogListInp
}

type MemberLogExportRes struct{}

type MemberLogViewReq struct {
	g.Meta `path:"/pmsMemberLog/view" method:"get" tags:"ADMIN_PMS" summary:"会员登录日志_详情"`
	input_app_member.PmsMemberLogViewInp
}

type MemberLogViewRes struct {
	*input_app_member.PmsMemberLogViewModel
}
