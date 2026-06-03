package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) DriverWithdrawList(ctx context.Context, req *car.DriverWithdrawListReq) (res *car.DriverWithdrawListRes, err error) {
	list, totalCount, err := service.CarDriverWithdraw().List(ctx, &req.DriverWithdrawListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.DriverWithdrawListModel{}
	}

	res = new(car.DriverWithdrawListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) DriverWithdrawView(ctx context.Context, req *car.DriverWithdrawViewReq) (res *car.DriverWithdrawViewRes, err error) {
	data, err := service.CarDriverWithdraw().View(ctx, &req.DriverWithdrawViewInp)
	if err != nil {
		return
	}

	res = new(car.DriverWithdrawViewRes)
	res.DriverWithdrawViewModel = data
	return
}
func (c *ControllerCar) DriverWithdrawAgree(ctx context.Context, req *car.DriverWithdrawAgreeReq) (res *car.DriverWithdrawAgreeRes, err error) {
	err = service.CarDriverWithdraw().Agree(ctx, &req.DriverWithdrawAgreeInp)
	return
}
func (c *ControllerCar) DriverWithdrawDisagree(ctx context.Context, req *car.DriverWithdrawDisagreeReq) (res *car.DriverWithdrawDisagreeRes, err error) {
	err = service.CarDriverWithdraw().Disagree(ctx, &req.DriverWithdrawDisagreeInp)
	return
}
func (c *ControllerCar) DriverWithdrawTransfer(ctx context.Context, req *car.DriverWithdrawTransferReq) (res *car.DriverWithdrawTransferRes, err error) {
	err = service.CarDriverWithdraw().Transfer(ctx, &req.DriverWithdrawTransferInp)
	return
}
