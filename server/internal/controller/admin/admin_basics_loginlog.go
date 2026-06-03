package admin

import (
	"APT/api/admin/basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerBasics) LoginLogList(ctx context.Context, req *basics.LoginLogListReq) (res *basics.LoginLogListRes, err error) {
	list, totalCount, err := service.BasicsLoginLog().List(ctx, &req.LoginLogListInp)
	if err != nil {
		return
	}

	res = new(basics.LoginLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) LoginLogExport(ctx context.Context, req *basics.LoginLogExportReq) (res *basics.LoginLogExportRes, err error) {
	err = service.BasicsLoginLog().Export(ctx, &req.LoginLogListInp)
	return
}
func (c *ControllerBasics) LoginLogDelete(ctx context.Context, req *basics.LoginLogDeleteReq) (res *basics.LoginLogDeleteRes, err error) {
	err = service.BasicsLoginLog().Delete(ctx, &req.LoginLogDeleteInp)
	return
}
