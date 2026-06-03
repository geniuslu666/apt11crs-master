package admin

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) StaffList(ctx context.Context, req *pms.StaffListReq) (res *pms.StaffListRes, err error) {
	list, totalCount, err := service.BasicsStaff().List(ctx, &req.PmsStaffListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsStaffListModel{}
	}

	res = new(pms.StaffListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) StaffView(ctx context.Context, req *pms.StaffViewReq) (res *pms.StaffViewRes, err error) {
	data, err := service.BasicsStaff().View(ctx, &req.PmsStaffViewInp)
	if err != nil {
		return
	}

	res = new(pms.StaffViewRes)
	res.PmsStaffViewModel = data
	return
}
func (c *ControllerPms) StaffEdit(ctx context.Context, req *pms.StaffEditReq) (res *pms.StaffEditRes, err error) {
	err = service.BasicsStaff().Edit(ctx, &req.PmsStaffEditInp)
	return
}
func (c *ControllerPms) StaffDelete(ctx context.Context, req *pms.StaffDeleteReq) (res *pms.StaffDeleteRes, err error) {
	err = service.BasicsStaff().Delete(ctx, &req.PmsStaffDeleteInp)
	return
}
func (c *ControllerPms) StaffStatus(ctx context.Context, req *pms.StaffStatusReq) (res *pms.StaffStatusRes, err error) {
	err = service.BasicsStaff().Status(ctx, &req.PmsStaffStatusInp)
	return
}
func (c *ControllerPms) StaffBind(ctx context.Context, req *pms.StaffBindReq) (res *pms.StaffBindRes, err error) {
	err = service.BasicsStaff().Bind(ctx, &req.PmsStaffBindInp)
	return
}
func (c *ControllerPms) StaffUnbind(ctx context.Context, req *pms.StaffUnbindReq) (res *pms.StaffUnbindRes, err error) {
	err = service.BasicsStaff().Unbind(ctx, &req.PmsStaffUnbindInp)
	return
}
