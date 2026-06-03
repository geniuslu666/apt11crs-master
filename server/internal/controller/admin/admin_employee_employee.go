package admin

import (
	"APT/internal/model/input/input_employee"
	"APT/internal/service"
	"context"

	"APT/api/admin/employee"
)

func (c *ControllerEmployee) EmployeeList(ctx context.Context, req *employee.EmployeeListReq) (res *employee.EmployeeListRes, err error) {
	list, totalCount, err := service.Employee().List(ctx, &req.EmployeeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_employee.EmployeeListModel{}
	}

	res = new(employee.EmployeeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerEmployee) EmployeeView(ctx context.Context, req *employee.EmployeeViewReq) (res *employee.EmployeeViewRes, err error) {
	data, err := service.Employee().View(ctx, &req.EmployeeViewInp)
	if err != nil {
		return
	}

	res = new(employee.EmployeeViewRes)
	res.EmployeeViewModel = data
	return
}
func (c *ControllerEmployee) EmployeeEdit(ctx context.Context, req *employee.EmployeeEditReq) (res *employee.EmployeeEditRes, err error) {
	err = service.Employee().Edit(ctx, &req.EmployeeEditInp)
	return
}
func (c *ControllerEmployee) EmployeeDelete(ctx context.Context, req *employee.EmployeeDeleteReq) (res *employee.EmployeeDeleteRes, err error) {
	err = service.Employee().Delete(ctx, &req.EmployeeDeleteInp)
	return
}
func (c *ControllerEmployee) EmployeeStatus(ctx context.Context, req *employee.EmployeeStatusReq) (res *employee.EmployeeStatusRes, err error) {
	err = service.Employee().Status(ctx, &req.EmployeeStatusInp)
	return
}
func (c *ControllerEmployee) EmployeeBind(ctx context.Context, req *employee.EmployeeBindReq) (res *employee.EmployeeBindRes, err error) {
	err = service.Employee().Bind(ctx, &req.EmployeeBindInp)
	return
}
func (c *ControllerEmployee) EmployeeUnbind(ctx context.Context, req *employee.EmployeeUnbindReq) (res *employee.EmployeeUnbindRes, err error) {
	err = service.Employee().Unbind(ctx, &req.EmployeeUnbindInp)
	return
}
