package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) OrderLogList(ctx context.Context, req *car.OrderLogListReq) (res *car.OrderLogListRes, err error) {
	list, totalCount, err := service.CarOrderLog().List(ctx, &req.CarOrderLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarOrderLogListModel{}
	}

	res = new(car.OrderLogListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
