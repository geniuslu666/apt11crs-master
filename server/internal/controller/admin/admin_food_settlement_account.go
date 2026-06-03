package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) SettlementAccountList(ctx context.Context, req *food.SettlementAccountListReq) (res *food.SettlementAccountListRes, err error) {
	list, totalCount, err := service.FoodSettlementAccount().List(ctx, &req.FoodSettlementAccountListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodSettlementAccountListModel{}
	}

	res = new(food.SettlementAccountListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) SettlementAccountView(ctx context.Context, req *food.SettlementAccountViewReq) (res *food.SettlementAccountViewRes, err error) {
	data, err := service.FoodSettlementAccount().View(ctx, &req.FoodSettlementAccountViewInp)
	if err != nil {
		return
	}

	res = new(food.SettlementAccountViewRes)
	res.FoodSettlementAccountViewModel = data
	return
}
func (c *ControllerFood) SettlementAccountEdit(ctx context.Context, req *food.SettlementAccountEditReq) (res *food.SettlementAccountEditRes, err error) {
	err = service.FoodSettlementAccount().Edit(ctx, &req.FoodSettlementAccountEditInp)
	return
}
func (c *ControllerFood) SettlementAccountDelete(ctx context.Context, req *food.SettlementAccountDeleteReq) (res *food.SettlementAccountDeleteRes, err error) {
	err = service.FoodSettlementAccount().Delete(ctx, &req.FoodSettlementAccountDeleteInp)
	return
}
func (c *ControllerFood) SettlementAccountStatus(ctx context.Context, req *food.SettlementAccountStatusReq) (res *food.SettlementAccountStatusRes, err error) {
	err = service.FoodSettlementAccount().Status(ctx, &req.FoodSettlementAccountStatusInp)
	return
}
