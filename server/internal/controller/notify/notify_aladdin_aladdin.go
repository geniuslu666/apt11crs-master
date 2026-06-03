package notify

import (
	"APT/internal/library/aladdinApi"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/notify/aladdin"
)

func (c *ControllerAladdin) OfferInfo(ctx context.Context, req *aladdin.OfferInfoReq) (res *aladdin.OfferInfoRes, err error) {
	var (
		aladdinResponse *aladdinApi.OfferInfoResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).OfferInfo(ctx, req.OfferInfoParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.OfferInfoRes)
	res.OfferInfoResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) BookingBus(ctx context.Context, req *aladdin.BookingBusReq) (res *aladdin.BookingBusRes, err error) {
	var (
		aladdinResponse *aladdinApi.BookingBusResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).BookingBus(ctx, req.BookingBusParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.BookingBusRes)
	res.BookingBusResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) ConformOrder(ctx context.Context, req *aladdin.ConformOrderReq) (res *aladdin.ConformOrderRes, err error) {
	var (
		aladdinResponse *aladdinApi.ConformOrderResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).ConformOrder(ctx, req.ConformOrderParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.ConformOrderRes)
	res.ConformOrderResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) CancelOrder(ctx context.Context, req *aladdin.CancelOrderReq) (res *aladdin.CancelOrderRes, err error) {
	var (
		aladdinResponse *aladdinApi.CancelOrderResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).CancelOrder(ctx, req.CancelOrderParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.CancelOrderRes)
	res.CancelOrderResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) UserOrderInfo(ctx context.Context, req *aladdin.UserOrderInfoReq) (res *aladdin.UserOrderInfoRes, err error) {
	var (
		aladdinResponse *aladdinApi.UserOrderInfoResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).UserOrderInfo(ctx, req.UserOrderInfoParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.UserOrderInfoRes)
	res.UserOrderInfoResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) BookingLuggage(ctx context.Context, req *aladdin.BookingLuggageReq) (res *aladdin.BookingLuggageRes, err error) {
	var (
		aladdinResponse *aladdinApi.BookingLuggageResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).BookingLuggage(ctx, req.BookingLuggageParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.BookingLuggageRes)
	res.BookingLuggageResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) ConformLuggageOrder(ctx context.Context, req *aladdin.ConformLuggageOrderReq) (res *aladdin.ConformLuggageOrderRes, err error) {
	var (
		aladdinResponse *aladdinApi.ConformLuggageOrderResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).ConformLuggageOrder(ctx, req.ConformLuggageOrderParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.ConformLuggageOrderRes)
	res.ConformLuggageOrderResponseItem = aladdinResponse.Data
	return
}
func (c *ControllerAladdin) CancelLuggageOrder(ctx context.Context, req *aladdin.CancelLuggageOrderReq) (res *aladdin.CancelLuggageOrderRes, err error) {
	var (
		aladdinResponse *aladdinApi.CancelLuggageOrderResponse
	)
	if aladdinResponse, err = aladdinApi.NewClient(ctx).CancelLuggageOrder(ctx, req.CancelLuggageOrderParams); err != nil {
		return
	}
	if aladdinResponse.Status != 200 {
		err = gerror.New(aladdinResponse.Msg)
		return
	}
	res = new(aladdin.CancelLuggageOrderRes)
	res.CancelLuggageOrderResponseItem = aladdinResponse.Data
	return
}
