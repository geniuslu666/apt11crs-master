package router

import (
	"APT/internal/controller/app"
	"APT/internal/controller/appv2"
	"APT/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

// App 路由
func App(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.ApiIsAuth)

		group.Bind(
			appv2.NewWsApi().SendWebsocketMessage,

			app.NewBasics().WxReturnUrl,

			app.NewBasics().HomeDateConfig,
			app.NewHotel().RoomTypeView,

			app.NewBasics().Chat,
			app.NewBasics().HomeDateLoad,
			app.NewBasics().HomeDeviceInfo,
			app.NewBasics().UploadImage,
			app.NewBasics().GetAppConfig,
			app.NewBasics().HelpCenterView,
			app.NewBasics().HelpCenterList,
			app.NewBasics().HelpCenterCategoryList,
			app.NewBasics().AdSpaceAfterPayBannerList,
			app.NewBasics().ConfigGet,
			app.NewBasics().GetContactConfig,
			app.NewBasics().GetProfileButtonConfig,

			app.NewMember().OauthLogin,
			app.NewMember().FxAuthLogin,
			app.NewMember().SmsSendCode,
			app.NewMember().EmailSendCode,
			app.NewMember().AuthIdBindLogin,
			app.NewMember().Login,
			app.NewMember().PhoneNumber,

			app.NewHotel().PropertySearch,
			app.NewHotel().PropertyView,
			app.NewHotel().RoomTypeList,
			appv2.NewHotel().RoomTypeList,
			app.NewHotel().PreOrderRefundDetail,
			app.NewHotel().RefundOrder,
			app.NewHotel().OrderRefundDetail,

			app.NewVerify().MemberVerify,

			app.NewBasics().Dashboard,
			app.NewBasics().GetAppLanguageConfig,
			app.NewBasics().TestNavList,
			app.NewBasics().HomepageArticleList,
			app.NewBasics().HomepageArticleView,
			app.NewTravel().ProductList,
			app.NewTravel().ProductView,
			app.NewTravel().ProductSkuStock,
		)
		group.Bind(
			appv2.NewWsApi().Ws,
			appv2.NewWsApi().ChatSendMessage,
			appv2.NewWsApi().ChatGetMessage,
			appv2.NewWsApi().TranslateMessage,
			appv2.NewWsApi().ChatGetMessageUnread,

			app.NewBasics().ReadAlertMessage,
			app.NewMember().MemberIntentionEdit,
			app.NewMember().GetMemberIntentionConfig,
			app.NewMember().BalanceChange,
			app.NewMember().GetUserInfo,
			app.NewMember().BindEmail,
			app.NewMember().BindPhone,
			app.NewMember().StaffInfo,
			app.NewMember().ChannelInfo,
			app.NewMember().UserLevelRule,
			app.NewMember().EditUserInfo,
			app.NewMember().Cancel,
			app.NewMember().GetMemberQrCode,

			app.NewBasics().CouponList,
			app.NewBasics().CouponView,
			app.NewBasics().CouponOrderList,
			app.NewBasics().ReferrerInfo,
			app.NewBasics().ReferrerList,
			app.NewBasics().NotifyList,
			app.NewBasics().NotifyView,
			app.NewBasics().NotifyEdit,
			app.NewBasics().NotifyDelete,
			app.NewBasics().CollectAdd,
			app.NewBasics().CollectList,
			app.NewBasics().CollectDelete,
			app.NewBasics().ShareInfo,
			app.NewBasics().ShareInvitePageInfo,
			app.NewBasics().ChannelFxCenter,
			app.NewBasics().GetPayModeConfig,
			app.NewBasics().BrokerageList,
			app.NewBasics().BrokerageDetail,
			app.NewBasics().BrokerageWithdrawBase,
			app.NewBasics().BrokerageWithdraw,
			app.NewBasics().BrokerageWithdrawList,
			app.NewBasics().BrokerageWithdrawDetail,
			app.NewBasics().StaffFxCenter,
			app.NewBasics().Link,
			app.NewBasics().ThCouponList,
			app.NewBasics().ThCouponView,
			app.NewBasics().ThMchView,
			app.NewBasics().ThCouponWaitView,
			app.NewBasics().ThCouponReceive,
			app.NewBasics().AppLatestMessage,
			app.NewBasics().AppMessageList,
			app.NewBasics().AppMessageRead,
			app.NewBasics().CouponTypeView,
			app.NewBasics().CouponTypeReceive,
			app.NewBasics().HomepageArticleThumb,

			app.NewEmployee().EmployeeActivityList,
			app.NewEmployee().EmployeeActivityView,
			app.NewEmployee().EmployeeActivityReceiveCoupon,
			app.NewEmployee().EmployeeCheckAuth,

			app.NewHotel().PreOrderCreate,
			app.NewHotel().PreOrderCreateDetail,
			app.NewHotel().PrePayInfo,
			app.NewHotel().CreateOrder,
			app.NewHotel().CancelOrderNoPay,
			app.NewHotel().OrderList,
			app.NewHotel().OrderDetail,
			app.NewHotel().OrderChangeOptions,
			app.NewHotel().OrderChangeGuestEdit,
			app.NewHotel().OrderChangeGuestInfo,
			app.NewHotel().OrderChangeDatePreInfo,
			app.NewHotel().OrderChangeBookingPeople,
			app.NewHotel().OrderChangeBookingPeopleSubmit,
			app.NewHotel().OrderChangeDatePreInfoSubmit,
			app.NewHotel().OrderChangeBookingPeopleInfo,
			app.NewHotel().OrderChangeStayOnInfo,
			app.NewHotel().OrderChangeInfoDetail,
			app.NewHotel().ReservationList,
			app.NewHotel().ReservationDetail,
			app.NewHotel().FindReservationRoomNo,
			appv2.NewHotel().PreOrderCreate,
			appv2.NewHotel().PreMoreOrderCreate,
			appv2.NewHotel().OrderReceipt,
			appv2.NewHotel().OrderReceiptView,

			app.NewPay().ThirdPay,
			app.NewPay().InnerPay,
			app.NewPay().QueryStatus,

			app.NewCabinet().PreOrderCreate,
			app.NewCabinet().PreOrderCreateDetail,
			app.NewCabinet().PrePayInfo,
			app.NewCabinet().CreateOrder,
			app.NewCabinet().OrderList,
			app.NewCabinet().OrderView,
			app.NewCabinet().PayOvertimeInfo,
			app.NewCabinet().PayOvertime,
			app.NewCabinet().CancelOrderNoPay,
			app.NewCabinet().CabinetList,

			app.NewTravel().PreOrderCreate,
			app.NewTravel().PreOrderCreateDetail,
			app.NewTravel().PrePayInfo,
			app.NewTravel().CreateOrder,
			app.NewTravel().CancelOrderNoPay,
			app.NewTravel().OrderList,
			app.NewTravel().OrderDetail,
			app.NewTravel().MemberVerifyCodeRefresh,
			app.NewTravel().ApplyOrderRefundDetail,
			app.NewTravel().RefundOrder,
			app.NewTravel().OrderRefundDetail,

			app.NewVerify().MemberVerifyCodeRefresh,
		).Middleware(middleware.ApiAuth).Middleware(middleware.ApiAuthCheckMember).Middleware()
	})
}
