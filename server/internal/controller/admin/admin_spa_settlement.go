package admin

import (
	"context"

	"APT/api/admin/spa"
	"APT/internal/model/input/input_spa"
	"APT/internal/service"
)

func (c *ControllerSpa) SettlementList(ctx context.Context, req *spa.SettlementListReq) (res *spa.SettlementListRes, err error) {
	list, totalCount, err := service.SpaSettlement().List(ctx, &req.SpaSettlementListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_spa.SpaSettlementListModel{}
	}

	res = new(spa.SettlementListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerSpa) SettlementView(ctx context.Context, req *spa.SettlementViewReq) (res *spa.SettlementViewRes, err error) {
	data, err := service.SpaSettlement().View(ctx, &req.SpaSettlementViewInp)
	if err != nil {
		return
	}

	res = new(spa.SettlementViewRes)
	res.SpaSettlementViewModel = data
	return
}
func (c *ControllerSpa) SettlementEdit(ctx context.Context, req *spa.SettlementEditReq) (res *spa.SettlementEditRes, err error) {
	err = service.SpaSettlement().Edit(ctx, &req.SpaSettlementEditInp)
	return
}
func (c *ControllerSpa) SettlementDelete(ctx context.Context, req *spa.SettlementDeleteReq) (res *spa.SettlementDeleteRes, err error) {
	err = service.SpaSettlement().Delete(ctx, &req.SpaSettlementDeleteInp)
	return
}
func (c *ControllerSpa) SettlementStatus(ctx context.Context, req *spa.SettlementStatusReq) (res *spa.SettlementStatusRes, err error) {
	err = service.SpaSettlement().Status(ctx, &req.SpaSettlementStatusInp)
	return
}
