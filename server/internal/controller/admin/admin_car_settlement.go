package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

func (c *ControllerCar) SettlementList(ctx context.Context, req *car.SettlementListReq) (res *car.SettlementListRes, err error) {
	list, totalCount, err := service.CarSettlement().List(ctx, &req.CarSettlementListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarSettlementListModel{}
	}

	res = new(car.SettlementListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerCar) SettlementView(ctx context.Context, req *car.SettlementViewReq) (res *car.SettlementViewRes, err error) {
	data, err := service.CarSettlement().View(ctx, &req.CarSettlementViewInp)
	if err != nil {
		return
	}

	res = new(car.SettlementViewRes)
	res.CarSettlementViewModel = data
	return
}
func (c *ControllerCar) SettlementEdit(ctx context.Context, req *car.SettlementEditReq) (res *car.SettlementEditRes, err error) {
	err = service.CarSettlement().Edit(ctx, &req.CarSettlementEditInp)
	return
}
func (c *ControllerCar) SettlementDelete(ctx context.Context, req *car.SettlementDeleteReq) (res *car.SettlementDeleteRes, err error) {
	err = service.CarSettlement().Delete(ctx, &req.CarSettlementDeleteInp)
	return
}
func (c *ControllerCar) SettlementStatus(ctx context.Context, req *car.SettlementStatusReq) (res *car.SettlementStatusRes, err error) {
	err = service.CarSettlement().Status(ctx, &req.CarSettlementStatusInp)
	return
}
