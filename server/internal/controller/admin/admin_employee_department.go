package admin

import (
	"context"

	"APT/api/admin/employee"
	"APT/internal/service"
)

func (c *ControllerEmployee) DepartmentList(ctx context.Context, req *employee.DepartmentListReq) (res *employee.DepartmentListRes, err error) {
	data, err := service.EmployeeDepartment().List(ctx, &req.EmployeeDepartmentListInp)
	if err != nil || data == nil {
		return
	}

	res = (*employee.DepartmentListRes)(data)
	return
}
func (c *ControllerEmployee) DepartmentView(ctx context.Context, req *employee.DepartmentViewReq) (res *employee.DepartmentViewRes, err error) {
	res = new(employee.DepartmentViewRes)
	res.EmployeeDepartmentViewModel, err = service.EmployeeDepartment().View(ctx, &req.EmployeeDepartmentViewInp)
	return
}
func (c *ControllerEmployee) DepartmentEdit(ctx context.Context, req *employee.DepartmentEditReq) (res *employee.DepartmentEditRes, err error) {
	err = service.EmployeeDepartment().Edit(ctx, &req.EmployeeDepartmentEditInp)
	return
}
func (c *ControllerEmployee) DepartmentDelete(ctx context.Context, req *employee.DepartmentDeleteReq) (res *employee.DepartmentDeleteRes, err error) {
	err = service.EmployeeDepartment().Delete(ctx, &req.EmployeeDepartmentDeleteInp)
	return
}
func (c *ControllerEmployee) DepartmentBatchDelete(ctx context.Context, req *employee.DepartmentBatchDeleteReq) (res *employee.DepartmentBatchDeleteRes, err error) {
	err = service.EmployeeDepartment().BatchDelete(ctx, &req.EmployeeDepartmentBatchDeleteInp)
	return
}
func (c *ControllerEmployee) DepartmentSwitch(ctx context.Context, req *employee.DepartmentSwitchReq) (res *employee.DepartmentSwitchRes, err error) {
	err = service.EmployeeDepartment().Switch(ctx, &req.EmployeeDepartmentSwitchInp)
	return
}
func (c *ControllerEmployee) DepartmentTree(ctx context.Context, req *employee.DepartmentTreeReq) (res *employee.DepartmentTreeRes, err error) {
	data, err := service.EmployeeDepartment().TreeOption(ctx)
	res = (*employee.DepartmentTreeRes)(data)
	return
}
