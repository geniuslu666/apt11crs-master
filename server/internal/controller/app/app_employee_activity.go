package app

import (
	"context"

	"APT/api/app/employee"
	"APT/internal/model/input/input_employee"
	"APT/internal/service"
)

func (c *ControllerEmployee) EmployeeActivityList(ctx context.Context, req *employee.EmployeeActivityListReq) (res *employee.EmployeeActivityListRes, err error) {
	res = new(employee.EmployeeActivityListRes)
	if res.List, res.Count, err = service.EmployeeActivity().AppList(ctx, &input_employee.EmployeeActivityAppListInp{
		PageReq: req.PageReq,
		Status:  req.Status,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerEmployee) EmployeeActivityView(ctx context.Context, req *employee.EmployeeActivityViewReq) (res *employee.EmployeeActivityViewRes, err error) {
	res = new(employee.EmployeeActivityViewRes)
	if res.EmployeeActivityAppViewModel, err = service.EmployeeActivity().AppView(ctx, &req.EmployeeActivityAppViewInp); err != nil {
		return
	}
	return
}
func (c *ControllerEmployee) EmployeeActivityReceiveCoupon(ctx context.Context, req *employee.EmployeeActivityReceiveCouponReq) (res *employee.EmployeeActivityReceiveCouponRes, err error) {
	err = service.EmployeeActivity().ReceiveActivityCoupon(ctx, &req.EmployeeActivityReceiveCouponInp)
	return
}
