// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notify

import (
	"context"

	"APT/api/notify/aladdin"
	"APT/api/notify/cabinet"
	"APT/api/notify/hook"
	"APT/api/notify/kefu"
	"APT/api/notify/paypay"
	"APT/api/notify/test"
	"APT/api/notify/toreta"
)

type INotifyAladdin interface {
	OfferInfo(ctx context.Context, req *aladdin.OfferInfoReq) (res *aladdin.OfferInfoRes, err error)
	BookingBus(ctx context.Context, req *aladdin.BookingBusReq) (res *aladdin.BookingBusRes, err error)
	ConformOrder(ctx context.Context, req *aladdin.ConformOrderReq) (res *aladdin.ConformOrderRes, err error)
	CancelOrder(ctx context.Context, req *aladdin.CancelOrderReq) (res *aladdin.CancelOrderRes, err error)
	UserOrderInfo(ctx context.Context, req *aladdin.UserOrderInfoReq) (res *aladdin.UserOrderInfoRes, err error)
	BookingLuggage(ctx context.Context, req *aladdin.BookingLuggageReq) (res *aladdin.BookingLuggageRes, err error)
	ConformLuggageOrder(ctx context.Context, req *aladdin.ConformLuggageOrderReq) (res *aladdin.ConformLuggageOrderRes, err error)
	CancelLuggageOrder(ctx context.Context, req *aladdin.CancelLuggageOrderReq) (res *aladdin.CancelLuggageOrderRes, err error)
}

type INotifyCabinet interface {
	CabinetInfo(ctx context.Context, req *cabinet.CabinetInfoReq) (res *cabinet.CabinetInfoRes, err error)
	CreateOrder(ctx context.Context, req *cabinet.CreateOrderReq) (res *cabinet.CreateOrderRes, err error)
	OrderQuery(ctx context.Context, req *cabinet.OrderQueryReq) (res *cabinet.OrderQueryRes, err error)
	PayOvertime(ctx context.Context, req *cabinet.PayOvertimeReq) (res *cabinet.PayOvertimeRes, err error)
	CabinetList(ctx context.Context, req *cabinet.CabinetListReq) (res *cabinet.CabinetListRes, err error)
	OrderComplete(ctx context.Context, req *cabinet.OrderCompleteReq) (res *cabinet.OrderCompleteRes, err error)
}

type INotifyHook interface {
	InRefund(ctx context.Context, req *hook.InRefundReq) (res *hook.InRefundRes, err error)
	OrderPrint(ctx context.Context, req *hook.OrderPrintReq) (res *hook.OrderPrintRes, err error)
	AirhostHook(ctx context.Context, req *hook.AirhostHookReq) (res *hook.AirhostHookRes, err error)
	Aladdin(ctx context.Context, req *hook.AladdinReq) (res *hook.AladdinRes, err error)
	Cabinet(ctx context.Context, req *hook.CabinetReq) (res *hook.CabinetRes, err error)
	VisitorMessageCallUrl(ctx context.Context, req *hook.VisitorMessageCallUrlReq) (res *hook.VisitorMessageCallUrlRes, err error)
	PaySuccessNotify(ctx context.Context, req *hook.PaySuccessNotifyReq) (res *hook.PaySuccessNotifyRes, err error)
	RefundSuccessNotify(ctx context.Context, req *hook.RefundSuccessNotifyReq) (res *hook.RefundSuccessNotifyRes, err error)
	MlilifePay(ctx context.Context, req *hook.MlilifePayReq) (res *hook.MlilifePayRes, err error)
	DoPayCloud(ctx context.Context, req *hook.DoPayCloudReq) (res *hook.DoPayCloudRes, err error)
	RefundCloud(ctx context.Context, req *hook.RefundCloudReq) (res *hook.RefundCloudRes, err error)
	PaypelPayReturn(ctx context.Context, req *hook.PaypelPayReturnReq) (res *hook.PaypelPayReturnRes, err error)
	PaypelPayCancel(ctx context.Context, req *hook.PaypelPayCancelReq) (res *hook.PaypelPayCancelRes, err error)
	PaypelPayHook(ctx context.Context, req *hook.PaypelPayHookReq) (res *hook.PaypelPayHookRes, err error)
	StripeWebhook(ctx context.Context, req *hook.StripeWebhookReq) (res *hook.StripeWebhookRes, err error)
	StripeSuccess(ctx context.Context, req *hook.StripeSuccessReq) (res *hook.StripeSuccessRes, err error)
	StripeCancel(ctx context.Context, req *hook.StripeCancelReq) (res *hook.StripeCancelRes, err error)
	Toreta(ctx context.Context, req *hook.ToretaReq) (res *hook.ToretaRes, err error)
}

type INotifyKefu interface {
	FindMemberInfo(ctx context.Context, req *kefu.FindMemberInfoReq) (res *kefu.FindMemberInfoRes, err error)
	FindMemberInfos(ctx context.Context, req *kefu.FindMemberInfosReq) (res *kefu.FindMemberInfosRes, err error)
	OrderList(ctx context.Context, req *kefu.OrderListReq) (res *kefu.OrderListRes, err error)
	AppReservationLanguageTemplate(ctx context.Context, req *kefu.AppReservationLanguageTemplateReq) (res *kefu.AppReservationLanguageTemplateRes, err error)
	Welcome(ctx context.Context, req *kefu.WelcomeReq) (res *kefu.WelcomeRes, err error)
}

type INotifyPaypay interface {
	DoPayPayCallback(ctx context.Context, req *paypay.DoPayPayCallbackReq) (res *paypay.DoPayPayCallbackRes, err error)
}

type INotifyTest interface {
	PaycloudCardPay(ctx context.Context, req *test.PaycloudCardPayReq) (res *test.PaycloudCardPayRes, err error)
}

type INotifyToreta interface {
	Restaurants(ctx context.Context, req *toreta.RestaurantsReq) (res *toreta.RestaurantsRes, err error)
	RestaurantDetail(ctx context.Context, req *toreta.RestaurantDetailReq) (res *toreta.RestaurantDetailRes, err error)
	Courses(ctx context.Context, req *toreta.CoursesReq) (res *toreta.CoursesRes, err error)
	Slots(ctx context.Context, req *toreta.SlotsReq) (res *toreta.SlotsRes, err error)
	CreateReservation(ctx context.Context, req *toreta.CreateReservationReq) (res *toreta.CreateReservationRes, err error)
	CancelReservation(ctx context.Context, req *toreta.CancelReservationReq) (res *toreta.CancelReservationRes, err error)
	ReservationDetail(ctx context.Context, req *toreta.ReservationDetailReq) (res *toreta.ReservationDetailRes, err error)
	Notifications(ctx context.Context, req *toreta.NotificationsReq) (res *toreta.NotificationsRes, err error)
	OARestaurants(ctx context.Context, req *toreta.OARestaurantsReq) (res *toreta.OARestaurantsRes, err error)
	OARestaurantsBind(ctx context.Context, req *toreta.OARestaurantsBindReq) (res *toreta.OARestaurantsBindRes, err error)
	RefreshOAToken(ctx context.Context, req *toreta.RefreshOATokenReq) (res *toreta.RefreshOATokenRes, err error)
}
