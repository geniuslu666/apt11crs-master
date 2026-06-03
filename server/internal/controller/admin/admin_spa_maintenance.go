package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) MaintenanceList(ctx context.Context, req *spa.MaintenanceListReq) (res *spa.MaintenanceListRes, err error) {
	list, totalCount, err := service.SpaMaintenance().List(ctx, &req.SpaMaintenanceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaMaintenanceListModel{}
	}

	res = new(spa.MaintenanceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) MaintenanceView(ctx context.Context, req *spa.MaintenanceViewReq) (res *spa.MaintenanceViewRes, err error) {
	data, err := service.SpaMaintenance().View(ctx, &req.SpaMaintenanceViewInp)
	if err != nil {
		return
	}

	res = new(spa.MaintenanceViewRes)
	res.SpaMaintenanceViewModel = data
	return
}
func (c *ControllerSpa) MaintenanceEdit(ctx context.Context, req *spa.MaintenanceEditReq) (res *spa.MaintenanceEditRes, err error) {
	err = service.SpaMaintenance().Edit(ctx, &req.SpaMaintenanceEditInp)
	return
}
func (c *ControllerSpa) MaintenanceDelete(ctx context.Context, req *spa.MaintenanceDeleteReq) (res *spa.MaintenanceDeleteRes, err error) {
	err = service.SpaMaintenance().Delete(ctx, &req.SpaMaintenanceDeleteInp)
	return
}
func (c *ControllerSpa) MaintenanceLanguageList(ctx context.Context, req *spa.MaintenanceLanguageListReq) (res *spa.MaintenanceLanguageListRes, err error) {
	list, err := service.SpaMaintenance().LanguageList(ctx, &req.SpaMaintenanceLanguageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaMaintenanceLanguageListModel{}
	}

	res = new(spa.MaintenanceLanguageListRes)
	res.List = list
	return
}
