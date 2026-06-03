package employee

import (
	"APT/internal/model/input/input_employee"

	"github.com/gogf/gf/v2/frame/g"
)

// DepartmentListReq 获取员工部门列表
type DepartmentListReq struct {
	g.Meta `path:"/employeeDepartment/list" method:"get" summary:"获取员工部门列表" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentListInp
}
type DepartmentListRes input_employee.EmployeeDepartmentListModel

// DepartmentViewReq 获取员工部门详情
type DepartmentViewReq struct {
	g.Meta `path:"/employeeDepartment/view" method:"get" summary:"获取员工部门详情" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentViewInp
}
type DepartmentViewRes struct {
	*input_employee.EmployeeDepartmentViewModel
}

// DepartmentEditReq 编辑员工部门
type DepartmentEditReq struct {
	g.Meta `path:"/employeeDepartment/edit" method:"post" summary:"编辑员工部门" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentEditInp
}
type DepartmentEditRes input_employee.EmployeeDepartmentEditModel

// DepartmentDeleteReq 删除员工部门
type DepartmentDeleteReq struct {
	g.Meta `path:"/employeeDepartment/delete" method:"post" summary:"删除员工部门" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentDeleteInp
}
type DepartmentDeleteRes input_employee.EmployeeDepartmentDeleteModel

// DepartmentBatchDeleteReq 批量删除员工部门
type DepartmentBatchDeleteReq struct {
	g.Meta `path:"/employeeDepartment/batchDelete" method:"post" summary:"批量删除员工部门" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentBatchDeleteInp
}
type DepartmentBatchDeleteRes input_employee.EmployeeDepartmentBatchDeleteModel

// DepartmentSwitchReq 更新员工部门状态
type DepartmentSwitchReq struct {
	g.Meta `path:"/employeeDepartment/switch" method:"post" summary:"更新员工部门状态" tags:"员工部门管理"`
	input_employee.EmployeeDepartmentSwitchInp
}
type DepartmentSwitchRes input_employee.EmployeeDepartmentSwitchModel

// DepartmentTreeReq 获取员工部门树结构
type DepartmentTreeReq struct {
	g.Meta `path:"/employeeDepartment/tree" method:"get" summary:"获取员工部门树结构" tags:"员工部门管理"`
}
type DepartmentTreeRes input_employee.EmployeeDepartmentOptionModel
