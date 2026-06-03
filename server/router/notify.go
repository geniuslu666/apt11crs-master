package router

import (
	"APT/internal/controller/notify"
	"APT/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Notify 路由
func Notify(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api", func(group *ghttp.RouterGroup) {
		// airhost hook
		group.Bind(notify.NewHook().AirhostHook).Middleware(middleware.AirHostBasicAuthentication)

		group.Bind(
			notify.NewPaypay().DoPayPayCallback,

			notify.NewHook().VisitorMessageCallUrl,
			notify.NewHook().PaypelPayReturn,
			notify.NewHook().PaypelPayCancel,
			notify.NewHook().PaypelPayHook,
			notify.NewHook().StripeWebhook,
			notify.NewHook().StripeSuccess,
			notify.NewHook().StripeCancel,
			notify.NewHook().DoPayCloud,
			notify.NewHook().RefundCloud,
			notify.NewHook().InRefund,
			notify.NewHook().OrderPrint,
			notify.NewHook().Aladdin,
			notify.NewHook().MlilifePay,
			notify.NewHook().Cabinet,
			notify.NewHook().Toreta,
			notify.NewHook().PaySuccessNotify,
			notify.NewHook().RefundSuccessNotify,
			notify.NewAladdin().OfferInfo,
			notify.NewAladdin().UserOrderInfo,
			notify.NewAladdin().CancelOrder,
			notify.NewAladdin().BookingLuggage,
			notify.NewAladdin().ConformLuggageOrder,
			notify.NewAladdin().CancelLuggageOrder,
			notify.NewAladdin().BookingBus,
			notify.NewAladdin().ConformOrder,

			notify.NewKefu().AppReservationLanguageTemplate,
			notify.NewKefu().FindMemberInfos,
			notify.NewKefu().OrderList,
			notify.NewKefu().FindMemberInfo,
			notify.NewKefu().Welcome,

			notify.NewTest().PaycloudCardPay,

			notify.NewCabinet().CabinetInfo,
			notify.NewCabinet().CreateOrder,
			notify.NewCabinet().OrderQuery,
			notify.NewCabinet().PayOvertime,
			notify.NewCabinet().CabinetList,
			notify.NewCabinet().OrderComplete,

			notify.NewToreta().Restaurants,
			notify.NewToreta().RestaurantDetail,
			notify.NewToreta().Courses,
			notify.NewToreta().Slots,
			notify.NewToreta().CreateReservation,
			notify.NewToreta().CancelReservation,
			notify.NewToreta().ReservationDetail,
			notify.NewToreta().Notifications,
			notify.NewToreta().OARestaurants,
			notify.NewToreta().OARestaurantsBind,
			notify.NewToreta().RefreshOAToken,
		)
	})
}
