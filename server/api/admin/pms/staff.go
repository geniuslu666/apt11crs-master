package pms

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type StaffListReq struct {
	g.Meta `path:"/pmsStaff/list" method:"get" tags:"ADMIN_PMS" summary:"员工管理_列表"`
	input_basics.PmsStaffListInp
}

type StaffListRes struct {
	input_form.PageRes
	List []*input_basics.PmsStaffListModel `json:"list"   dc:"数据列表"`
}

type StaffViewReq struct {
	g.Meta `path:"/pmsStaff/view" method:"get" tags:"ADMIN_PMS" summary:"员工管理_详情"`
	input_basics.PmsStaffViewInp
}

type StaffViewRes struct {
	*input_basics.PmsStaffViewModel
}

type StaffEditReq struct {
	g.Meta `path:"/pmsStaff/edit" method:"post" tags:"ADMIN_PMS" summary:"员工管理_修改/新增"`
	input_basics.PmsStaffEditInp
}

type StaffEditRes struct{}

type StaffDeleteReq struct {
	g.Meta `path:"/pmsStaff/delete" method:"post" tags:"ADMIN_PMS" summary:"员工管理_删除"`
	input_basics.PmsStaffDeleteInp
}

type StaffDeleteRes struct{}

type StaffStatusReq struct {
	g.Meta `path:"/pmsStaff/status" method:"post" tags:"ADMIN_PMS" summary:"员工管理_更新状态"`
	input_basics.PmsStaffStatusInp
}

type StaffStatusRes struct{}

type StaffBindReq struct {
	g.Meta `path:"/pmsStaff/bind" method:"post" tags:"ADMIN_PMS" summary:"员工管理_员工绑定会员"`
	input_basics.PmsStaffBindInp
}

type StaffBindRes struct{}

type StaffUnbindReq struct {
	g.Meta `path:"/pmsStaff/unbind" method:"post" tags:"ADMIN_PMS" summary:"员工管理_员工解绑会员"`
	input_basics.PmsStaffUnbindInp
}

type StaffUnbindRes struct{}
