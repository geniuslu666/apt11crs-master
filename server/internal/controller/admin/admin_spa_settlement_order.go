package admin

import (
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
	"context"

	"APT/api/admin/spa"
)

func (c *ControllerSpa) SettlementOrderList(ctx context.Context, req *spa.SettlementOrderListReq) (res *spa.SettlementOrderListRes, err error) {
	list, totalCount, err := service.SpaSettlementOrder().List(ctx, &req.SpaSettlementOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaSettlementOrderListModel{}
	}

	res = new(spa.SettlementOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) SettlementOrderStat(ctx context.Context, req *spa.SettlementOrderStatReq) (res *spa.SettlementOrderStatRes, err error) {
	data, err := service.SpaSettlementOrder().SettlementOrderStat(ctx, &req.SpaSettlementOrderStatInp)

	res = new(spa.SettlementOrderStatRes)
	res.SpaSettlementOrderStatModel = data
	return
}
func (c *ControllerSpa) SettlementOrderView(ctx context.Context, req *spa.SettlementOrderViewReq) (res *spa.SettlementOrderViewRes, err error) {
	data, err := service.SpaSettlementOrder().View(ctx, &req.SpaSettlementOrderViewInp)
	if err != nil {
		return
	}

	res = new(spa.SettlementOrderViewRes)
	res.SpaSettlementOrderViewModel = data
	return
}
func (c *ControllerSpa) SettlementOrderVerify(ctx context.Context, req *spa.SettlementOrderVerifyReq) (res *spa.SettlementOrderVerifyRes, err error) {
	err = service.SpaSettlementOrder().Verify(ctx, &req.SpaSettlementOrderVerifyInp)
	return
}
