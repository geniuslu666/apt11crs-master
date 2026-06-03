package admin

import (
	"APT/api/admin/travel"
	"APT/internal/model/input/input_travel"
	"APT/internal/service"
	"context"
)

func (c *ControllerTravel) OrderList(ctx context.Context, req *travel.OrderListReq) (res *travel.OrderListRes, err error) {
	list, totalCount, err := service.TravelOrder().List(ctx, &req.TravelOrderListInp)
	if err != nil {
		return
	}
	if list == nil {
		list = []*input_travel.TravelOrderListModel{}
	}
	res = new(travel.OrderListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

func (c *ControllerTravel) OrderView(ctx context.Context, req *travel.OrderViewReq) (res *travel.OrderViewRes, err error) {
	data, err := service.TravelOrder().View(ctx, &req.TravelOrderViewInp)
	if err != nil {
		return
	}
	res = new(travel.OrderViewRes)
	res.TravelOrderViewModel = data
	return
}

func (c *ControllerTravel) OrderRefund(ctx context.Context, req *travel.OrderRefundReq) (res *travel.OrderRefundRes, err error) {
	err = service.TravelOrder().Refund(ctx, &req.TravelOrderRefundInp)
	return
}
