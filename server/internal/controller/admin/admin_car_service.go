package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) ServiceList(ctx context.Context, req *car.ServiceListReq) (res *car.ServiceListRes, err error) {
	list, totalCount, err := service.CarService().List(ctx, &req.CarServiceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarServiceListModel{}
	}

	res = new(car.ServiceListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) ServiceView(ctx context.Context, req *car.ServiceViewReq) (res *car.ServiceViewRes, err error) {
	data, err := service.CarService().View(ctx, &req.CarServiceViewInp)
	if err != nil {
		return
	}

	res = new(car.ServiceViewRes)
	res.CarServiceViewModel = data
	return
}
func (c *ControllerCar) ServiceEdit(ctx context.Context, req *car.ServiceEditReq) (res *car.ServiceEditRes, err error) {
	err = service.CarService().Edit(ctx, &req.CarServiceEditInp)
	return
}
func (c *ControllerCar) ServiceDelete(ctx context.Context, req *car.ServiceDeleteReq) (res *car.ServiceDeleteRes, err error) {
	err = service.CarService().Delete(ctx, &req.CarServiceDeleteInp)
	return
}
func (c *ControllerCar) ServiceMaxSort(ctx context.Context, req *car.ServiceMaxSortReq) (res *car.ServiceMaxSortRes, err error) {
	data, err := service.CarService().MaxSort(ctx, &req.CarServiceMaxSortInp)
	if err != nil {
		return
	}

	res = new(car.ServiceMaxSortRes)
	res.CarServiceMaxSortModel = data
	return
}
func (c *ControllerCar) ServiceStatus(ctx context.Context, req *car.ServiceStatusReq) (res *car.ServiceStatusRes, err error) {
	err = service.CarService().Status(ctx, &req.CarServiceStatusInp)
	return
}
func (c *ControllerCar) ServiceSort(ctx context.Context, req *car.ServiceSortReq) (res *car.ServiceSortRes, err error) {
	err = service.CarService().ServiceSort(ctx, &req.CarServiceSortInp)
	return
}
