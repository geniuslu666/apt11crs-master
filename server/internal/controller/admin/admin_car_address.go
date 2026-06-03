package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) AddressList(ctx context.Context, req *car.AddressListReq) (res *car.AddressListRes, err error) {
	list, totalCount, err := service.SysCarAddress().List(ctx, &req.CarAddressListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarAddressListModel{}
	}

	res = new(car.AddressListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) AddressExport(ctx context.Context, req *car.AddressExportReq) (res *car.AddressExportRes, err error) {
	err = service.SysCarAddress().Export(ctx, &req.CarAddressListInp)
	return
}
func (c *ControllerCar) AddressView(ctx context.Context, req *car.AddressViewReq) (res *car.AddressViewRes, err error) {
	data, err := service.SysCarAddress().View(ctx, &req.CarAddressViewInp)
	if err != nil {
		return
	}

	res = new(car.AddressViewRes)
	res.CarAddressViewModel = data
	return
}
func (c *ControllerCar) AddressEdit(ctx context.Context, req *car.AddressEditReq) (res *car.AddressEditRes, err error) {
	err = service.SysCarAddress().Edit(ctx, &req.CarAddressEditInp)
	return
}
func (c *ControllerCar) AddressDelete(ctx context.Context, req *car.AddressDeleteReq) (res *car.AddressDeleteRes, err error) {
	err = service.SysCarAddress().Delete(ctx, &req.CarAddressDeleteInp)
	return
}
func (c *ControllerCar) AddressStatus(ctx context.Context, req *car.AddressStatusReq) (res *car.AddressStatusRes, err error) {
	err = service.SysCarAddress().Status(ctx, &req.CarAddressStatusInp)
	return
}
