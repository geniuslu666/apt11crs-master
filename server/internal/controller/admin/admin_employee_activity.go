package admin

import (
	"context"

	"APT/api/admin/employee"
	"APT/internal/model/input/input_employee"
	"APT/internal/service"
)

func (c *ControllerEmployee) ActivityList(ctx context.Context, req *employee.ActivityListReq) (res *employee.ActivityListRes, err error) {
	list, totalCount, err := service.EmployeeActivity().List(ctx, &req.EmployeeActivityListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_employee.EmployeeActivityListModel{}
	}

	res = new(employee.ActivityListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerEmployee) ActivityView(ctx context.Context, req *employee.ActivityViewReq) (res *employee.ActivityViewRes, err error) {
	data, err := service.EmployeeActivity().View(ctx, &req.EmployeeActivityViewInp)
	if err != nil {
		return
	}

	res = new(employee.ActivityViewRes)
	res.EmployeeActivityViewModel = data
	return
}
func (c *ControllerEmployee) ActivityEdit(ctx context.Context, req *employee.ActivityEditReq) (res *employee.ActivityEditRes, err error) {
	err = service.EmployeeActivity().Edit(ctx, &req.EmployeeActivityEditInp)
	return
}
func (c *ControllerEmployee) ActivityDelete(ctx context.Context, req *employee.ActivityDeleteReq) (res *employee.ActivityDeleteRes, err error) {
	err = service.EmployeeActivity().Delete(ctx, &req.EmployeeActivityDeleteInp)
	return
}
func (c *ControllerEmployee) ActivityStatus(ctx context.Context, req *employee.ActivityStatusReq) (res *employee.ActivityStatusRes, err error) {
	err = service.EmployeeActivity().Status(ctx, &req.EmployeeActivityStatusInp)
	return
}

func (c *ControllerEmployee) ActivityBatchIssueCoupons(ctx context.Context, req *employee.ActivityBatchIssueCouponsReq) (res *employee.ActivityBatchIssueCouponsRes, err error) {
	data, err := service.EmployeeActivity().BatchIssueCoupons(ctx, &req.EmployeeActivityBatchIssueCouponsInp)
	if err != nil {
		return
	}
	res = new(employee.ActivityBatchIssueCouponsRes)
	res.EmployeeActivityBatchIssueCouponsModel = data
	return
}

func (c *ControllerEmployee) ActivityCouponRecordList(ctx context.Context, req *employee.ActivityCouponRecordListReq) (res *employee.ActivityCouponRecordListRes, err error) {
	list, totalCount, err := service.EmployeeActivity().CouponRecordList(ctx, &req.EmployeeActivityCouponRecordListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_employee.EmployeeActivityCouponRecordListModel{}
	}

	res = new(employee.ActivityCouponRecordListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
