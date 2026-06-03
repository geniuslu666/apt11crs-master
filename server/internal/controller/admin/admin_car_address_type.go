package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) AddressTypeList(ctx context.Context, req *car.AddressTypeListReq) (res *car.AddressTypeListRes, err error) {
	list, totalCount, err := service.CarAddressType().List(ctx, &req.CarAddressTypeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarAddressTypeListModel{}
	}

	res = new(car.AddressTypeListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
