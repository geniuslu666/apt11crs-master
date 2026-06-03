package appv2

import (
	"APT/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/appv2/hotel"
)

func (c *ControllerHotel) PreOrderCreate(ctx context.Context, req *hotel.PreOrderCreateReq) (res *hotel.PreOrderCreateRes, err error) {
	res = new(hotel.PreOrderCreateRes)
	if res.PreCreateOrderModel, err = service.HotelService().PrePricePlanOrder(ctx, &req.PreCreateOrderInp); err != nil {
		return
	}
	return
}
func (c *ControllerHotel) PreOrderCreateDetail(ctx context.Context, req *hotel.PreOrderCreateDetailReq) (res *hotel.PreOrderCreateDetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
func (c *ControllerHotel) PreMoreOrderCreate(ctx context.Context, req *hotel.PreMoreOrderCreateReq) (res *hotel.PreMoreOrderCreateRes, err error) {
	res = new(hotel.PreMoreOrderCreateRes)
	if res.PreCreateOrderModel, err = service.HotelService().PreMorePricePlanOrder(ctx, &req.PreMoreCreateOrderInp); err != nil {
		return
	}
	return
}
