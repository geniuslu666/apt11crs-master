package app

import (
	"APT/internal/service"
	"context"

	"APT/api/app/travel"
)

func (c *ControllerTravel) ApplyOrderRefundDetail(ctx context.Context, req *travel.ApplyOrderRefundDetailReq) (res *travel.ApplyOrderRefundDetailRes, err error) {
	res = new(travel.ApplyOrderRefundDetailRes)
	if res.TravelOrderApplyRefundDetailModel, err = service.TravelOrderRefundService().ApplyRefundDetail(ctx, &req.TravelOrderApplyRefundDetailInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) RefundOrder(ctx context.Context, req *travel.RefundOrderReq) (res *travel.RefundOrderRes, err error) {
	res = new(travel.RefundOrderRes)
	if err = service.TravelOrderRefundService().RefundOrder(ctx, &req.TravelOrderApplyRefundDetailInp); err != nil {
		return
	}
	return
}
func (c *ControllerTravel) OrderRefundDetail(ctx context.Context, req *travel.OrderRefundDetailReq) (res *travel.OrderRefundDetailRes, err error) {
	res = new(travel.OrderRefundDetailRes)
	if res.RefundDetailModel, err = service.TravelOrderRefundService().RefundOrderDetail(ctx, &req.TravelOrderApplyRefundDetailInp); err != nil {
		return
	}
	return
}
