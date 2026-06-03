package app

import (
	"APT/api/app/travel"
	"APT/internal/service"
	"context"
)

func (c *ControllerTravel) PreOrderCreate(ctx context.Context, req *travel.PreOrderCreateReq) (res *travel.PreOrderCreateRes, err error) {
	res = new(travel.PreOrderCreateRes)
	if res.PreCreateOrderModel, err = service.TravelOrderCreatePreService().PreOrder(ctx, &req.PreCreateOrderInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) PreOrderCreateDetail(ctx context.Context, req *travel.PreOrderCreateDetailReq) (res *travel.PreOrderCreateDetailRes, err error) {
	res = new(travel.PreOrderCreateDetailRes)
	if res.PreOrderDetailModel, err = service.TravelOrderCreatePreService().PreOrderDetail(ctx, &req.PreOrderDetailInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) PrePayInfo(ctx context.Context, req *travel.PrePayInfoReq) (res *travel.PrePayInfoRes, err error) {
	res = new(travel.PrePayInfoRes)
	if res.PrePayInfoModel, err = service.TravelOrderCreatePreService().PrePayInfo(ctx, req.PrePayInfoInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) CreateOrder(ctx context.Context, req *travel.CreateOrderReq) (res *travel.CreateOrderRes, err error) {
	res = new(travel.CreateOrderRes)
	if res.CreateOrderModel, err = service.TravelOrderCreateService().CreateOrder(ctx, req.CreateOrderInp); err != nil {
		return
	}
	return
}
