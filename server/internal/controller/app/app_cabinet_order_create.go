package app

import (
	"APT/api/app/cabinet"
	"APT/internal/service"
	"context"
)

func (c *ControllerCabinet) PreOrderCreate(ctx context.Context, req *cabinet.PreOrderCreateReq) (res *cabinet.PreOrderCreateRes, err error) {
	res = new(cabinet.PreOrderCreateRes)
	if res.PreCreateOrderModel, err = service.CabinetService().PreOrder(ctx, &req.PreCreateOrderInp); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) PreOrderCreateDetail(ctx context.Context, req *cabinet.PreOrderCreateDetailReq) (res *cabinet.PreOrderCreateDetailRes, err error) {
	res = new(cabinet.PreOrderCreateDetailRes)
	if res.PreOrderDetailModel, err = service.CabinetService().PreOrderDetail(ctx, &req.PreOrderDetailInp); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) PrePayInfo(ctx context.Context, req *cabinet.PrePayInfoReq) (res *cabinet.PrePayInfoRes, err error) {
	res = new(cabinet.PrePayInfoRes)
	if res.OrderPayInfoModel, err = service.CabinetService().PrePayInfo(ctx, req.PrePayInfoInp); err != nil {
		return
	}
	return
}
func (c *ControllerCabinet) CreateOrder(ctx context.Context, req *cabinet.CreateOrderReq) (res *cabinet.CreateOrderRes, err error) {
	res = new(cabinet.CreateOrderRes)
	if res.CreateOrderModel, err = service.CabinetService().CreateOrder(ctx, req.CreateOrderInp); err != nil {
		return
	}
	return
}
