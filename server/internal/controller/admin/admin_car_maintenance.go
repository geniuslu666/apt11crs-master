package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) MaintenanceList(ctx context.Context, req *car.MaintenanceListReq) (res *car.MaintenanceListRes, err error) {
	list, totalCount, err := service.CarMaintenance().List(ctx, &req.CarMaintenanceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarMaintenanceListModel{}
	}

	res = new(car.MaintenanceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) MaintenanceView(ctx context.Context, req *car.MaintenanceViewReq) (res *car.MaintenanceViewRes, err error) {
	data, err := service.CarMaintenance().View(ctx, &req.CarMaintenanceViewInp)
	if err != nil {
		return
	}

	res = new(car.MaintenanceViewRes)
	res.CarMaintenanceViewModel = data
	return
}
func (c *ControllerCar) MaintenanceEdit(ctx context.Context, req *car.MaintenanceEditReq) (res *car.MaintenanceEditRes, err error) {
	err = service.CarMaintenance().Edit(ctx, &req.CarMaintenanceEditInp)
	return
}
func (c *ControllerCar) MaintenanceDelete(ctx context.Context, req *car.MaintenanceDeleteReq) (res *car.MaintenanceDeleteRes, err error) {
	err = service.CarMaintenance().Delete(ctx, &req.CarMaintenanceDeleteInp)
	return
}
func (c *ControllerCar) MaintenanceLanguageList(ctx context.Context, req *car.MaintenanceLanguageListReq) (res *car.MaintenanceLanguageListRes, err error) {
	list, err := service.CarMaintenance().LanguageList(ctx, &req.CarMaintenanceLanguageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarMaintenanceLanguageListModel{}
	}

	res = new(car.MaintenanceLanguageListRes)
	res.List = list
	return
}
