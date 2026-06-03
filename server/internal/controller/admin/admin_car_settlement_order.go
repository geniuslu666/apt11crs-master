package admin

import (
	"APT/internal/model/input/input_car"
	"APT/internal/service"
	"context"

	"APT/api/admin/car"
)

// SettlementOrderList 结算单列表
func (c *ControllerCar) SettlementOrderList(ctx context.Context, req *car.SettlementOrderListReq) (res *car.SettlementOrderListRes, err error) {
	list, totalCount, err := service.CarSettlementOrder().List(ctx, &req.CarSettlementOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_car.CarSettlementOrderListModel{}
	}

	res = new(car.SettlementOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// SettlementOrderStat 结算单概况
func (c *ControllerCar) SettlementOrderStat(ctx context.Context, req *car.SettlementOrderStatReq) (res *car.SettlementOrderStatRes, err error) {
	data, err := service.CarSettlementOrder().SettlementOrderStat(ctx, &req.CarSettlementOrderStatInp)

	res = new(car.SettlementOrderStatRes)
	res.CarSettlementOrderStatModel = data
	return
}
func (c *ControllerCar) SettlementOrderView(ctx context.Context, req *car.SettlementOrderViewReq) (res *car.SettlementOrderViewRes, err error) {
	data, err := service.CarSettlementOrder().View(ctx, &req.CarSettlementOrderViewInp)
	if err != nil {
		return
	}

	res = new(car.SettlementOrderViewRes)
	res.CarSettlementOrderViewModel = data
	return
}
func (c *ControllerCar) SettlementOrderVerify(ctx context.Context, req *car.SettlementOrderVerifyReq) (res *car.SettlementOrderVerifyRes, err error) {
	err = service.CarSettlementOrder().Verify(ctx, &req.CarSettlementOrderVerifyInp)
	return
}
