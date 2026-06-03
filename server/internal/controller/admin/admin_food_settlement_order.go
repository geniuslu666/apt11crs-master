package admin

import (
	"APT/api/admin/food"
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"
)

// SettlementOrderList 结算单列表
func (c *ControllerFood) SettlementOrderList(ctx context.Context, req *food.SettlementOrderListReq) (res *food.SettlementOrderListRes, err error) {
	list, totalCount, err := service.FoodSettlementOrder().List(ctx, &req.FoodSettlementOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodSettlementOrderListModel{}
	}

	res = new(food.SettlementOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// SettlementOrderStat 结算单概况
func (c *ControllerFood) SettlementOrderStat(ctx context.Context, req *food.SettlementOrderStatReq) (res *food.SettlementOrderStatRes, err error) {
	data, err := service.FoodSettlementOrder().SettlementOrderStat(ctx, &req.FoodSettlementOrderStatInp)

	res = new(food.SettlementOrderStatRes)
	res.FoodSettlementOrderStatModel = data
	return
}

// SettlementOrderView 结算单详情
func (c *ControllerFood) SettlementOrderView(ctx context.Context, req *food.SettlementOrderViewReq) (res *food.SettlementOrderViewRes, err error) {
	data, err := service.FoodSettlementOrder().View(ctx, &req.FoodSettlementOrderViewInp)
	if err != nil {
		return
	}

	res = new(food.SettlementOrderViewRes)
	res.FoodSettlementOrderViewModel = data
	return
}

// SettlementOrderVerify 结算单-核账
func (c *ControllerFood) SettlementOrderVerify(ctx context.Context, req *food.SettlementOrderVerifyReq) (res *food.SettlementOrderVerifyRes, err error) {
	err = service.FoodSettlementOrder().Verify(ctx, &req.FoodSettlementOrderVerifyInp)
	return
}
