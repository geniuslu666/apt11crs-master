package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) IspList(ctx context.Context, req *spa.IspListReq) (res *spa.IspListRes, err error) {
	list, totalCount, err := service.SpaIsp().List(ctx, &req.SpaIspListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaIspListModel{}
	}

	res = new(spa.IspListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) IspAllList(ctx context.Context, req *spa.IspAllListReq) (res *spa.IspAllListRes, err error) {
	list, err := service.SpaIsp().All(ctx, &req.SpaIspListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaIspAllListModel{}
	}

	res = new(spa.IspAllListRes)
	res.List = list
	return
}
func (c *ControllerSpa) IspView(ctx context.Context, req *spa.IspViewReq) (res *spa.IspViewRes, err error) {
	data, err := service.SpaIsp().View(ctx, &req.SpaIspViewInp)
	if err != nil {
		return
	}

	res = new(spa.IspViewRes)
	res.SpaIspViewModel = data
	return
}
func (c *ControllerSpa) IspEdit(ctx context.Context, req *spa.IspEditReq) (res *spa.IspEditRes, err error) {
	err = service.SpaIsp().Edit(ctx, &req.SpaIspEditInp)
	return
}
func (c *ControllerSpa) IspDelete(ctx context.Context, req *spa.IspDeleteReq) (res *spa.IspDeleteRes, err error) {
	err = service.SpaIsp().Delete(ctx, &req.SpaIspDeleteInp)
	return
}
func (c *ControllerSpa) IspStatus(ctx context.Context, req *spa.IspStatusReq) (res *spa.IspStatusRes, err error) {
	err = service.SpaIsp().Status(ctx, &req.SpaIspStatusInp)
	return
}
func (c *ControllerSpa) IspWorkStatus(ctx context.Context, req *spa.IspWorkStatusReq) (res *spa.IspWorkStatusRes, err error) {
	err = service.SpaIsp().WorkStatus(ctx, &req.SpaIspWorkStatusInp)
	return
}
func (c *ControllerSpa) IspBind(ctx context.Context, req *spa.IspBindReq) (res *spa.IspBindRes, err error) {
	err = service.SpaIsp().Bind(ctx, &req.SpaIspBindInp)
	return
}
func (c *ControllerSpa) IspUnbind(ctx context.Context, req *spa.IspUnbindReq) (res *spa.IspUnbindRes, err error) {
	err = service.SpaIsp().Unbind(ctx, &req.SpaIspUnbindInp)
	return
}
