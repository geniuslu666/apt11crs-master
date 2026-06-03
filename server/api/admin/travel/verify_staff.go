package travel

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

// VerifyStaffListReq 获取核销人员列表
type VerifyStaffListReq struct {
	g.Meta `path:"/travel/verifyStaff/list" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游核销人员列表"`
	input_travel.TravelVerifyStaffListInp
}

type VerifyStaffListRes struct {
	input_form.PageRes
	List []*input_travel.TravelVerifyStaffListModel `json:"list" dc:"数据列表"`
}

// VerifyStaffViewReq 获取核销人员详情
type VerifyStaffViewReq struct {
	g.Meta `path:"/travel/verifyStaff/view" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游核销人员详情"`
	input_travel.TravelVerifyStaffViewInp
}

type VerifyStaffViewRes struct {
	*input_travel.TravelVerifyStaffViewModel
}

// VerifyStaffEditReq 新增/编辑核销人员
type VerifyStaffEditReq struct {
	g.Meta `path:"/travel/verifyStaff/edit" method:"post" tags:"ADMIN_TRAVEL" summary:"新增/编辑一日游核销人员"`
	input_travel.TravelVerifyStaffEditInp
}

type VerifyStaffEditRes struct {
	Id int64 `json:"id" dc:"核销人员ID"`
}

// VerifyStaffDeleteReq 删除核销人员
type VerifyStaffDeleteReq struct {
	g.Meta `path:"/travel/verifyStaff/delete" method:"post" tags:"ADMIN_TRAVEL" summary:"删除一日游核销人员"`
	input_travel.TravelVerifyStaffDeleteInp
}

type VerifyStaffDeleteRes struct{}

// VerifyStaffStatusReq 启用/禁用核销人员
type VerifyStaffStatusReq struct {
	g.Meta `path:"/travel/verifyStaff/status" method:"post" tags:"ADMIN_TRAVEL" summary:"启用/禁用一日游核销人员"`
	input_travel.TravelVerifyStaffStatusInp
}

type VerifyStaffStatusRes struct{}

// VerifyStaffScopeOptionsReq 获取核销权限范围选项
type VerifyStaffScopeOptionsReq struct {
	g.Meta `path:"/travel/verifyStaff/scopeOptions" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游核销权限范围选项"`
}

type VerifyStaffScopeOptionsRes struct {
	List []*input_travel.TravelVerifyStaffScopeOptionModel `json:"list" dc:"数据列表"`
}
