package admin

import (
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/admin/basics"
)

func (c *ControllerBasics) LogClear(_ context.Context, _ *basics.LogClearReq) (res *basics.LogClearRes, err error) {
	err = gerror.New("暂时考虑到安全问题，请到数据库清空")
	return
}
func (c *ControllerBasics) LogExport(ctx context.Context, req *basics.LogExportReq) (res *basics.LogExportRes, err error) {
	err = service.BasicsLog().Export(ctx, &req.LogListInp)
	return
}
func (c *ControllerBasics) LogList(ctx context.Context, req *basics.LogListReq) (res *basics.LogListRes, err error) {
	list, totalCount, err := service.BasicsLog().List(ctx, &req.LogListInp)
	if err != nil {
		return
	}

	res = new(basics.LogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerBasics) LogDelete(ctx context.Context, req *basics.LogDeleteReq) (res *basics.LogDeleteRes, err error) {
	err = service.BasicsLog().Delete(ctx, &req.LogDeleteInp)
	return
}
func (c *ControllerBasics) LogView(ctx context.Context, req *basics.LogViewReq) (res *basics.LogViewRes, err error) {
	res = new(basics.LogViewRes)
	res.LogViewModel, err = service.BasicsLog().View(ctx, &req.LogViewInp)
	return
}
