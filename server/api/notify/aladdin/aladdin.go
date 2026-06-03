package aladdin

import (
	"APT/internal/library/aladdinApi"
	"github.com/gogf/gf/v2/frame/g"
)

type OfferInfoReq struct {
	g.Meta `path:"/busOrder/offerInfo" method:"post" tags:"ALADDIN_HOOK" summary:"查询报价"`
	*aladdinApi.OfferInfoParams
}

type OfferInfoRes struct {
	aladdinApi.OfferInfoResponseItem
}

type BookingBusReq struct {
	g.Meta `path:"/busOrder/bookingBus" method:"post" tags:"ALADDIN_HOOK" summary:"巴士预约生单"`
	*aladdinApi.BookingBusParams
}

type BookingBusRes struct {
	aladdinApi.BookingBusResponseItem
}

type ConformOrderReq struct {
	g.Meta `path:"/busOrder/conformOrder" method:"post" tags:"ALADDIN_HOOK" summary:"确认订单"`
	*aladdinApi.ConformOrderParams
}

type ConformOrderRes struct {
	aladdinApi.ConformOrderResponseItem
}

type CancelOrderReq struct {
	g.Meta `path:"/busOrder/cancelOrder" method:"post" tags:"ALADDIN_HOOK" summary:"取消订单"`
	*aladdinApi.CancelOrderParams
}

type CancelOrderRes struct {
	aladdinApi.CancelOrderResponseItem
}

type UserOrderInfoReq struct {
	g.Meta `path:"/busOrder/userOrderInfo" method:"post" tags:"ALADDIN_HOOK" summary:"查询用户订单"`
	*aladdinApi.UserOrderInfoParams
}

type UserOrderInfoRes struct {
	aladdinApi.UserOrderInfoResponseItem
}

type BookingLuggageReq struct {
	g.Meta `path:"/busOrder/bookingLuggage" method:"post" tags:"ALADDIN_HOOK" summary:"行李预约"`
	*aladdinApi.BookingLuggageParams
}

type BookingLuggageRes struct {
	aladdinApi.BookingLuggageResponseItem
}

type ConformLuggageOrderReq struct {
	g.Meta `path:"/busOrder/conformLuggageOrder" method:"post" tags:"ALADDIN_HOOK" summary:"行李确认"`
	*aladdinApi.ConformLuggageOrderParams
}

type ConformLuggageOrderRes struct {
	aladdinApi.ConformLuggageOrderResponseItem
}

type CancelLuggageOrderReq struct {
	g.Meta `path:"/busOrder/cancelLuggageOrder" method:"post" tags:"ALADDIN_HOOK" summary:"行李取消"`
	*aladdinApi.CancelLuggageOrderParams
}

type CancelLuggageOrderRes struct {
	aladdinApi.CancelLuggageOrderResponseItem
}
