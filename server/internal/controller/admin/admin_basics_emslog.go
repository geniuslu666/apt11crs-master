package admin

import (
	"APT/internal/service"
	"context"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) EmsLogList(ctx context.Context, req *basics.EmsLogListReq) (res *basics.EmsLogListRes, err error) {
	list, totalCount, err := service.BasicsEmsLog().List(ctx, &req.EmsLogListInp)
	if err != nil {
		return
	}

	res = new(basics.EmsLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) EmsLogView(ctx context.Context, req *basics.EmsLogViewReq) (res *basics.EmsLogViewRes, err error) {
	data, err := service.BasicsEmsLog().View(ctx, &req.EmsLogViewInp)
	if err != nil {
		return
	}

	res = new(basics.EmsLogViewRes)
	res.EmsLogViewModel = data
	return
}
func (c *ControllerBasics) EmsLogEdit(ctx context.Context, req *basics.EmsLogEditReq) (res *basics.EmsLogEditRes, err error) {
	err = service.BasicsEmsLog().Edit(ctx, &req.EmsLogEditInp)
	return
}
func (c *ControllerBasics) EmsLogDelete(ctx context.Context, req *basics.EmsLogDeleteReq) (res *basics.EmsLogDeleteRes, err error) {
	err = service.BasicsEmsLog().Delete(ctx, &req.EmsLogDeleteInp)
	return
}
func (c *ControllerBasics) EmsLogStatus(ctx context.Context, req *basics.EmsLogStatusReq) (res *basics.EmsLogStatusRes, err error) {
	err = service.BasicsEmsLog().Status(ctx, &req.EmsLogStatusInp)
	return
}
