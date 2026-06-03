package employee

import (
	"APT/internal/model/input/input_employee"
	"APT/internal/model/input/input_form"

	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeListReq 查询员工管理列表
type EmployeeListReq struct {
	g.Meta `path:"/employee/list" method:"get" tags:"ADMIN_EMPLOYEE" summary:"获取员工管理列表"`
	input_employee.EmployeeListInp
}

type EmployeeListRes struct {
	input_form.PageRes
	List []*input_employee.EmployeeListModel `json:"list"   dc:"数据列表"`
}

// EmployeeViewReq 获取员工详情
type EmployeeViewReq struct {
	g.Meta `path:"/employee/view" method:"get" tags:"ADMIN_EMPLOYEE" summary:"获取员工详情"`
	input_employee.EmployeeViewInp
}

type EmployeeViewRes struct {
	*input_employee.EmployeeViewModel
}

// EmployeeEditReq 修改/新增员工
type EmployeeEditReq struct {
	g.Meta `path:"/employee/edit" method:"post" tags:"ADMIN_EMPLOYEE" summary:"修改/新增员工"`
	input_employee.EmployeeEditInp
}

type EmployeeEditRes struct{}

// EmployeeDeleteReq 删除员工
type EmployeeDeleteReq struct {
	g.Meta `path:"/employee/delete" method:"post" tags:"ADMIN_EMPLOYEE" summary:"删除员工"`
	input_employee.EmployeeDeleteInp
}

type EmployeeDeleteRes struct{}

// EmployeeStatusReq 更新员工状态
type EmployeeStatusReq struct {
	g.Meta `path:"/employee/status" method:"post" tags:"ADMIN_EMPLOYEE" summary:"更新员工状态"`
	input_employee.EmployeeStatusInp
}

type EmployeeStatusRes struct{}

type EmployeeBindReq struct {
	g.Meta `path:"/employee/bind" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工绑定用户"`
	input_employee.EmployeeBindInp
}

type EmployeeBindRes struct{}

type EmployeeUnbindReq struct {
	g.Meta `path:"/employee/unbind" method:"post" tags:"ADMIN_EMPLOYEE" summary:"员工解绑用户"`
	input_employee.EmployeeUnbindInp
}

type EmployeeUnbindRes struct{}
