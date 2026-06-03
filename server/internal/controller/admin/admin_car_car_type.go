package admin

import (
	"APT/api/admin/car"
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"
)

func (c *ControllerCar) CarTypeList(ctx context.Context, req *car.CarTypeListReq) (res *car.CarTypeListRes, err error) {
	list, totalCount, err := service.CarCarType().List(ctx, &req.CarCarTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarCarTypeListModel{}
	}

	res = new(car.CarTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) CarTypeView(ctx context.Context, req *car.CarTypeViewReq) (res *car.CarTypeViewRes, err error) {
	data, err := service.CarCarType().View(ctx, &req.CarCarTypeViewInp)
	if err != nil {
		return
	}

	res = new(car.CarTypeViewRes)
	res.CarCarTypeViewModel = data
	return
}
func (c *ControllerCar) CarTypeEdit(ctx context.Context, req *car.CarTypeEditReq) (res *car.CarTypeEditRes, err error) {
	err = service.CarCarType().Edit(ctx, &req.CarCarTypeEditInp)
	return
}
func (c *ControllerCar) CarTypeDelete(ctx context.Context, req *car.CarTypeDeleteReq) (res *car.CarTypeDeleteRes, err error) {
	err = service.CarCarType().Delete(ctx, &req.CarCarTypeDeleteInp)
	return
}
func (c *ControllerCar) CarTypeMaxSort(ctx context.Context, req *car.CarTypeMaxSortReq) (res *car.CarTypeMaxSortRes, err error) {
	data, err := service.CarCarType().MaxSort(ctx, &req.CarCarTypeMaxSortInp)
	if err != nil {
		return
	}

	res = new(car.CarTypeMaxSortRes)
	res.CarCarTypeMaxSortModel = data
	return
}
func (c *ControllerCar) CarTypeStatus(ctx context.Context, req *car.CarTypeStatusReq) (res *car.CarTypeStatusRes, err error) {
	err = service.CarCarType().Status(ctx, &req.CarCarTypeStatusInp)
	return
}
