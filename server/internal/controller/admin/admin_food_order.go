package admin

import (
	"APT/internal/model/input/input_food"
	"APT/internal/service"
	"context"

	"APT/api/admin/food"
)

func (c *ControllerFood) OrderList(ctx context.Context, req *food.OrderListReq) (res *food.OrderListRes, err error) {
	list, totalCount, err := service.FoodOrder().List(ctx, &req.FoodOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodOrderListModel{}
	}

	res = new(food.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) OrderView(ctx context.Context, req *food.OrderViewReq) (res *food.OrderViewRes, err error) {
	data, err := service.FoodOrder().View(ctx, &req.FoodOrderViewInp)
	if err != nil {
		return
	}

	res = new(food.OrderViewRes)
	res.FoodOrderViewModel = data
	return
}
func (c *ControllerFood) OrderConfirmAgree(ctx context.Context, req *food.OrderConfirmAgreeReq) (res *food.OrderConfirmAgreeRes, err error) {
	err = service.FoodOrder().ConfirmAgree(ctx, &req.FoodOrderConfirmAgreeInp)
	return
}
func (c *ControllerFood) OrderConfirmDisagree(ctx context.Context, req *food.OrderConfirmDisagreeReq) (res *food.OrderConfirmDisagreeRes, err error) {
	err = service.FoodOrder().ConfirmDisagree(ctx, &req.FoodOrderConfirmDisagreeInp)
	return
}

// SettleOrderList 结算订单列表
func (c *ControllerFood) SettleOrderList(ctx context.Context, req *food.SettleOrderListReq) (res *food.SettleOrderListRes, err error) {
	list, totalCount, err := service.FoodOrder().SettleOrderList(ctx, &req.SettleFoodOrderListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.SettleFoodOrderListModel{}
	}

	res = new(food.SettleOrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerFood) OrderCancelPay(ctx context.Context, req *food.OrderCancelPayReq) (res *food.OrderCancelPayRes, err error) {
	err = service.FoodOrder().CancelPay(ctx, &req.FoodOrderCancelPayInp)
	return
}
func (c *ControllerFood) OrderExport(ctx context.Context, req *food.OrderExportReq) (res *food.OrderExportRes, err error) {
	err = service.FoodOrder().ExportOrder(ctx, &req.FoodOrderExportInp)
	return
}
func (c *ControllerFood) OrderExportList(ctx context.Context, req *food.OrderExportListReq) (res *food.OrderExportListRes, err error) {
	list, totalCount, err := service.FoodOrder().ExportList(ctx, &req.FoodOrderExportListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_food.FoodOrderExportListModel{}
	}

	res = new(food.OrderExportListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
