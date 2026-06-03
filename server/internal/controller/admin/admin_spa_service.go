package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) ServiceList(ctx context.Context, req *spa.ServiceListReq) (res *spa.ServiceListRes, err error) {
	list, totalCount, err := service.SpaService().List(ctx, &req.SpaServiceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaServiceListModel{}
	}

	res = new(spa.ServiceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) ServiceView(ctx context.Context, req *spa.ServiceViewReq) (res *spa.ServiceViewRes, err error) {
	data, err := service.SpaService().View(ctx, &req.SpaServiceViewInp)
	if err != nil {
		return
	}

	res = new(spa.ServiceViewRes)
	res.SpaServiceViewModel = data
	return
}
func (c *ControllerSpa) ServiceEdit(ctx context.Context, req *spa.ServiceEditReq) (res *spa.ServiceEditRes, err error) {
	err = service.SpaService().Edit(ctx, &req.SpaServiceEditInp)
	return
}
func (c *ControllerSpa) ServiceDelete(ctx context.Context, req *spa.ServiceDeleteReq) (res *spa.ServiceDeleteRes, err error) {
	err = service.SpaService().Delete(ctx, &req.SpaServiceDeleteInp)
	return
}
func (c *ControllerSpa) ServiceMaxSort(ctx context.Context, req *spa.ServiceMaxSortReq) (res *spa.ServiceMaxSortRes, err error) {
	data, err := service.SpaService().MaxSort(ctx, &req.SpaServiceMaxSortInp)
	if err != nil {
		return
	}

	res = new(spa.ServiceMaxSortRes)
	res.SpaServiceMaxSortModel = data
	return
}
func (c *ControllerSpa) ServiceStatus(ctx context.Context, req *spa.ServiceStatusReq) (res *spa.ServiceStatusRes, err error) {
	err = service.SpaService().Status(ctx, &req.SpaServiceStatusInp)
	return
}
func (c *ControllerSpa) ServiceAllList(ctx context.Context, req *spa.ServiceAllListReq) (res *spa.ServiceAllListRes, err error) {
	list, err := service.SpaService().All(ctx, &req.SpaServiceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaServiceAllListModel{}
	}

	res = new(spa.ServiceAllListRes)
	res.List = list
	return
}
func (c *ControllerSpa) ServiceRecycle(ctx context.Context, req *spa.ServiceRecycleReq) (res *spa.ServiceRecycleRes, err error) {
	err = service.SpaService().Recycle(ctx, &req.SpaServiceDeleteInp)
	return
}
