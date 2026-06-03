// Package pmsAdminApi

package pms

import (
	"APT/internal/model/input/input_app_member"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberIntentionListReq 查询会员意向表列表
type MemberIntentionListReq struct {
	g.Meta `path:"/pmsMemberIntention/list" method:"get" tags:"ADMIN_PMS" summary:"会员意向表_列表"`
	input_app_member.PmsMemberIntentionListInp
}

type MemberIntentionListRes struct {
	input_form.PageRes
	List []*input_app_member.PmsMemberIntentionListModel `json:"list"   dc:"数据列表"`
}

// MemberIntentionExportReq 导出会员意向表列表
type MemberIntentionExportReq struct {
	g.Meta `path:"/pmsMemberIntention/export" method:"get" tags:"ADMIN_PMS" summary:"会员意向表_导出列表"`
	input_app_member.PmsMemberIntentionListInp
}

type MemberIntentionExportRes struct{}

// MemberIntentionViewReq 获取会员意向表指定信息
type MemberIntentionViewReq struct {
	g.Meta `path:"/pmsMemberIntention/view" method:"get" tags:"ADMIN_PMS" summary:"会员意向表_详情"`
	input_app_member.PmsMemberIntentionViewInp
}

type MemberIntentionViewRes struct {
	*input_app_member.PmsMemberIntentionViewModel
}
