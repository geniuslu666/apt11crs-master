package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) BlacklistList(ctx context.Context, req *basics.BlacklistListReq) (res *basics.BlacklistListRes, err error) {
	list, totalCount, err := service.BasicsBlacklist().List(ctx, &req.BlacklistListInp)
	if err != nil {
		return
	}

	res = new(basics.BlacklistListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) BlacklistView(ctx context.Context, req *basics.BlacklistViewReq) (res *basics.BlacklistViewRes, err error) {
	data, err := service.BasicsBlacklist().View(ctx, &req.BlacklistViewInp)
	if err != nil {
		return
	}

	res = new(basics.BlacklistViewRes)
	res.BlacklistViewModel = data
	return
}
func (c *ControllerBasics) BlacklistEdit(ctx context.Context, req *basics.BlacklistEditReq) (res *basics.BlacklistEditRes, err error) {
	err = service.BasicsBlacklist().Edit(ctx, &req.BlacklistEditInp)
	return
}
func (c *ControllerBasics) BlacklistDelete(ctx context.Context, req *basics.BlacklistDeleteReq) (res *basics.BlacklistDeleteRes, err error) {
	err = service.BasicsBlacklist().Delete(ctx, &req.BlacklistDeleteInp)
	return
}
func (c *ControllerBasics) BlacklistStatus(ctx context.Context, req *basics.BlacklistStatusReq) (res *basics.BlacklistStatusRes, err error) {
	err = service.BasicsBlacklist().Status(ctx, &req.BlacklistStatusInp)
	return
}
