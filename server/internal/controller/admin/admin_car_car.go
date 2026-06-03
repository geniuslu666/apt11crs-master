package admin

import (
	"APT/api/admin/car"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
)

func (c *ControllerCar) CarList(ctx context.Context, req *car.CarListReq) (res *car.CarListRes, err error) {
	list, totalCount, err := service.CarCar().List(ctx, &req.CarCarListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarCarListModel{}
	}

	res = new(car.CarListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) CarView(ctx context.Context, req *car.CarViewReq) (res *car.CarViewRes, err error) {
	data, err := service.CarCar().View(ctx, &req.CarCarViewInp)
	if err != nil {
		return
	}

	res = new(car.CarViewRes)
	res.CarCarViewModel = data
	return
}
func (c *ControllerCar) CarEdit(ctx context.Context, req *car.CarEditReq) (res *car.CarEditRes, err error) {
	err = service.CarCar().Edit(ctx, &req.CarCarEditInp)
	return
}
func (c *ControllerCar) CarDelete(ctx context.Context, req *car.CarDeleteReq) (res *car.CarDeleteRes, err error) {
	err = service.CarCar().Delete(ctx, &req.CarCarDeleteInp)
	return
}
func (c *ControllerCar) CarMaxSort(ctx context.Context, req *car.CarMaxSortReq) (res *car.CarMaxSortRes, err error) {
	data, err := service.CarCar().MaxSort(ctx, &req.CarCarMaxSortInp)
	if err != nil {
		return
	}

	res = new(car.CarMaxSortRes)
	res.CarCarMaxSortModel = data
	return
}
func (c *ControllerCar) CarStatus(ctx context.Context, req *car.CarStatusReq) (res *car.CarStatusRes, err error) {
	err = service.CarCar().Status(ctx, &req.CarCarStatusInp)
	return
}
func (c *ControllerCar) WorkStatus(ctx context.Context, req *car.WorkStatusReq) (res *car.WorkStatusRes, err error) {
	err = service.CarCar().WorkStatus(ctx, &req.CarCarWorkStatusInp)
	return
}
