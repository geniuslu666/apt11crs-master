package admin

import (
	"APT/api/admin/pms"
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
)

func (c *ControllerPms) AppConfigList(ctx context.Context, req *pms.AppConfigListReq) (res *pms.AppConfigListRes, err error) {
	list, totalCount, err := service.BasicsPmsAppconfig().List(ctx, &req.PmsAppconfigListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_basics.PmsAppconfigListModel{}
	}

	res = new(pms.AppConfigListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) AppConfigView(ctx context.Context, req *pms.AppConfigViewReq) (res *pms.AppConfigViewRes, err error) {
	data, err := service.BasicsPmsAppconfig().View(ctx, &req.PmsAppconfigViewInp)
	if err != nil {
		return
	}

	res = new(pms.AppConfigViewRes)
	res.PmsAppconfigViewModel = data
	return
}
func (c *ControllerPms) AppConfigEdit(ctx context.Context, req *pms.AppConfigEditReq) (res *pms.AppConfigEditRes, err error) {
	err = service.BasicsPmsAppconfig().Edit(ctx, &req.PmsAppconfigEditInp)
	return
}
func (c *ControllerPms) AppConfigDelete(ctx context.Context, req *pms.AppConfigDeleteReq) (res *pms.AppConfigDeleteRes, err error) {
	err = service.BasicsPmsAppconfig().Delete(ctx, &req.PmsAppconfigDeleteInp)
	return
}
