package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) ServeLogList(ctx context.Context, req *basics.ServeLogListReq) (res *basics.ServeLogListRes, err error) {
	list, totalCount, err := service.BasicsServeLog().List(ctx, &req.ServeLogListInp)
	if err != nil {
		return
	}

	res = new(basics.ServeLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) ServeLogExport(ctx context.Context, req *basics.ServeLogExportReq) (res *basics.ServeLogExportRes, err error) {
	err = service.BasicsServeLog().Export(ctx, &req.ServeLogListInp)
	return
}
func (c *ControllerBasics) ServeLogView(ctx context.Context, req *basics.ServeLogViewReq) (res *basics.ServeLogViewRes, err error) {
	data, err := service.BasicsServeLog().View(ctx, &req.ServeLogViewInp)
	if err != nil {
		return
	}

	res = new(basics.ServeLogViewRes)
	res.ServeLogViewModel = data
	return
}
func (c *ControllerBasics) ServeLogDelete(ctx context.Context, req *basics.ServeLogDeleteReq) (res *basics.ServeLogDeleteRes, err error) {
	err = service.BasicsServeLog().Delete(ctx, &req.ServeLogDeleteInp)
	return
}
