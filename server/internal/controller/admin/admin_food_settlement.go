package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) SettlementList(ctx context.Context, req *food.SettlementListReq) (res *food.SettlementListRes, err error) {
	list, totalCount, err := service.FoodSettlement().List(ctx, &req.FoodSettlementListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodSettlementListModel{}
	}

	res = new(food.SettlementListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) SettlementView(ctx context.Context, req *food.SettlementViewReq) (res *food.SettlementViewRes, err error) {
	data, err := service.FoodSettlement().View(ctx, &req.FoodSettlementViewInp)
	if err != nil {
		return
	}

	res = new(food.SettlementViewRes)
	res.FoodSettlementViewModel = data
	return
}
func (c *ControllerFood) SettlementEdit(ctx context.Context, req *food.SettlementEditReq) (res *food.SettlementEditRes, err error) {
	err = service.FoodSettlement().Edit(ctx, &req.FoodSettlementEditInp)
	return
}
func (c *ControllerFood) SettlementDelete(ctx context.Context, req *food.SettlementDeleteReq) (res *food.SettlementDeleteRes, err error) {
	err = service.FoodSettlement().Delete(ctx, &req.FoodSettlementDeleteInp)
	return
}
func (c *ControllerFood) SettlementStatus(ctx context.Context, req *food.SettlementStatusReq) (res *food.SettlementStatusRes, err error) {
	err = service.FoodSettlement().Status(ctx, &req.FoodSettlementStatusInp)
	return
}
