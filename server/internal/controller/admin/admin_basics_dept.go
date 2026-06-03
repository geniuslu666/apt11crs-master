package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) DeptList(ctx context.Context, req *basics.DeptListReq) (res *basics.DeptListRes, err error) {
	data, err := service.BasicsAdminDept().List(ctx, &req.DeptListInp)
	if err != nil || data == nil {
		return
	}

	res = (*basics.DeptListRes)(&data)
	return
}
func (c *ControllerBasics) DeptView(ctx context.Context, req *basics.DeptViewReq) (res *basics.DeptViewRes, err error) {
	res = new(basics.DeptViewRes)
	res.DeptViewModel, err = service.BasicsAdminDept().View(ctx, &req.DeptViewInp)
	return
}
func (c *ControllerBasics) DeptEdit(ctx context.Context, req *basics.DeptEditReq) (res *basics.DeptEditRes, err error) {
	err = service.BasicsAdminDept().Edit(ctx, &req.DeptEditInp)
	return
}
func (c *ControllerBasics) DeptDelete(ctx context.Context, req *basics.DeptDeleteReq) (res *basics.DeptDeleteRes, err error) {
	err = service.BasicsAdminDept().Delete(ctx, &req.DeptDeleteInp)
	return
}
func (c *ControllerBasics) DeptMaxSort(ctx context.Context, req *basics.DeptMaxSortReq) (res *basics.DeptMaxSortRes, err error) {
	res = new(basics.DeptMaxSortRes)
	res.DeptMaxSortModel, err = service.BasicsAdminDept().MaxSort(ctx, &req.DeptMaxSortInp)
	return
}
func (c *ControllerBasics) DeptOption(ctx context.Context, req *basics.DeptOptionReq) (res *basics.DeptOptionRes, err error) {
	list, totalCount, err := service.BasicsAdminDept().Option(ctx, &req.DeptOptionInp)
	if err != nil {
		return
	}

	res = new(basics.DeptOptionRes)
	res.DeptOptionModel = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) DeptTreeOption(ctx context.Context, _ *basics.DeptTreeOptionReq) (res *basics.DeptTreeOptionRes, err error) {
	data, err := service.BasicsAdminDept().TreeOption(ctx)
	res = (*basics.DeptTreeOptionRes)(&data)
	return
}
