package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) DriverList(ctx context.Context, req *car.DriverListReq) (res *car.DriverListRes, err error) {
	list, totalCount, err := service.CarDriver().List(ctx, &req.CarDriverListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarDriverListModel{}
	}

	res = new(car.DriverListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) DriverAllList(ctx context.Context, req *car.DriverAllListReq) (res *car.DriverAllListRes, err error) {
	list, err := service.CarDriver().All(ctx, &req.CarDriverListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarDriverAllListModel{}
	}

	res = new(car.DriverAllListRes)
	res.List = list
	return
}
func (c *ControllerCar) DriverView(ctx context.Context, req *car.DriverViewReq) (res *car.DriverViewRes, err error) {
	data, err := service.CarDriver().View(ctx, &req.CarDriverViewInp)
	if err != nil {
		return
	}

	res = new(car.DriverViewRes)
	res.CarDriverViewModel = data
	return
}
func (c *ControllerCar) DriverEdit(ctx context.Context, req *car.DriverEditReq) (res *car.DriverEditRes, err error) {
	err = service.CarDriver().Edit(ctx, &req.CarDriverEditInp)
	return
}
func (c *ControllerCar) DriverDelete(ctx context.Context, req *car.DriverDeleteReq) (res *car.DriverDeleteRes, err error) {
	err = service.CarDriver().Delete(ctx, &req.CarDriverDeleteInp)
	return
}
func (c *ControllerCar) DriverStatus(ctx context.Context, req *car.DriverStatusReq) (res *car.DriverStatusRes, err error) {
	err = service.CarDriver().Status(ctx, &req.CarDriverStatusInp)
	return
}
func (c *ControllerCar) DriverWorkStatus(ctx context.Context, req *car.DriverWorkStatusReq) (res *car.DriverWorkStatusRes, err error) {
	err = service.CarDriver().WorkStatus(ctx, &req.CarDriverWorkStatusInp)
	return
}
func (c *ControllerCar) DriverBind(ctx context.Context, req *car.DriverBindReq) (res *car.DriverBindRes, err error) {
	err = service.CarDriver().Bind(ctx, &req.CarDriverBindInp)
	return
}
func (c *ControllerCar) DriverUnbind(ctx context.Context, req *car.DriverUnbindReq) (res *car.DriverUnbindRes, err error) {
	err = service.CarDriver().Unbind(ctx, &req.CarDriverUnbindInp)
	return
}
func (c *ControllerCar) DriverBindCar(ctx context.Context, req *car.DriverBindCarReq) (res *car.DriverBindCarRes, err error) {
	err = service.CarDriver().BindCar(ctx, &req.CarDriverBindCarInp)
	return
}
