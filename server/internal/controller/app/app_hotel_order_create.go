package app

import (
	"APT/internal/service"
	"context"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) PreOrderCreate(ctx context.Context, req *hotel.PreOrderCreateReq) (res *hotel.PreOrderCreateRes, err error) {
	res = new(hotel.PreOrderCreateRes)
	if res.PreCreateOrderModel, err = service.HotelService().PreOrder(ctx, &req.PreCreateOrderInp); err != nil {
		return
	}
	return
}
func (c *ControllerHotel) PreOrderCreateDetail(ctx context.Context, req *hotel.PreOrderCreateDetailReq) (res *hotel.PreOrderCreateDetailRes, err error) {
	res = new(hotel.PreOrderCreateDetailRes)
	if res.PreOrderDetailModel, err = service.HotelService().PreOrderDetail(ctx, &req.PreOrderDetailInp); err != nil {
		return
	}
	return
}
func (c *ControllerHotel) PrePayInfo(ctx context.Context, req *hotel.PrePayInfoReq) (res *hotel.PrePayInfoRes, err error) {
	res = new(hotel.PrePayInfoRes)
	if res.PrePayInfoModel, err = service.HotelService().PrePayInfo(ctx, req.PrePayInfoInp); err != nil {
		return
	}
	return
}
func (c *ControllerHotel) CreateOrder(ctx context.Context, req *hotel.CreateOrderReq) (res *hotel.CreateOrderRes, err error) {
	res = new(hotel.CreateOrderRes)
	if res.CreateOrderModel, err = service.HotelService().CreateHotelOrder(ctx, req.CreateOrderInp); err != nil {
		return
	}
	return
}
