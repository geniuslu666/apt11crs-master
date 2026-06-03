// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"context"

	"APT/api/app/basics"
	"APT/api/app/cabinet"
	"APT/api/app/employee"
	"APT/api/app/fx"
	"APT/api/app/hotel"
	"APT/api/app/member"
	"APT/api/app/pay"
	"APT/api/app/terminal"
	"APT/api/app/travel"
	"APT/api/app/verify"
)

type IAppBasics interface {
	AdSpaceAfterPayBannerList(ctx context.Context, req *basics.AdSpaceAfterPayBannerListReq) (res *basics.AdSpaceAfterPayBannerListRes, err error)
	BrokerageList(ctx context.Context, req *basics.BrokerageListReq) (res *basics.BrokerageListRes, err error)
	BrokerageDetail(ctx context.Context, req *basics.BrokerageDetailReq) (res *basics.BrokerageDetailRes, err error)
	BrokerageWithdrawBase(ctx context.Context, req *basics.BrokerageWithdrawBaseReq) (res *basics.BrokerageWithdrawBaseRes, err error)
	BrokerageWithdraw(ctx context.Context, req *basics.BrokerageWithdrawReq) (res *basics.BrokerageWithdrawRes, err error)
	BrokerageWithdrawList(ctx context.Context, req *basics.BrokerageWithdrawListReq) (res *basics.BrokerageWithdrawListRes, err error)
	BrokerageWithdrawDetail(ctx context.Context, req *basics.BrokerageWithdrawDetailReq) (res *basics.BrokerageWithdrawDetailRes, err error)
	ChannelFxCenter(ctx context.Context, req *basics.ChannelFxCenterReq) (res *basics.ChannelFxCenterRes, err error)
	Chat(ctx context.Context, req *basics.ChatReq) (res *basics.ChatRes, err error)
	CollectAdd(ctx context.Context, req *basics.CollectAddReq) (res *basics.CollectAddRes, err error)
	CollectList(ctx context.Context, req *basics.CollectListReq) (res *basics.CollectListRes, err error)
	CollectDelete(ctx context.Context, req *basics.CollectDeleteReq) (res *basics.CollectDeleteRes, err error)
	GetAppConfig(ctx context.Context, req *basics.GetAppConfigReq) (res *basics.GetAppConfigRes, err error)
	GetAppLanguageConfig(ctx context.Context, req *basics.GetAppLanguageConfigReq) (res *basics.GetAppLanguageConfigRes, err error)
	GetPayModeConfig(ctx context.Context, req *basics.GetPayModeConfigReq) (res *basics.GetPayModeConfigRes, err error)
	GetContactConfig(ctx context.Context, req *basics.GetContactConfigReq) (res *basics.GetContactConfigRes, err error)
	GetProfileButtonConfig(ctx context.Context, req *basics.GetProfileButtonConfigReq) (res *basics.GetProfileButtonConfigRes, err error)
	CouponList(ctx context.Context, req *basics.CouponListReq) (res *basics.CouponListRes, err error)
	CouponView(ctx context.Context, req *basics.CouponViewReq) (res *basics.CouponViewRes, err error)
	CouponOrderList(ctx context.Context, req *basics.CouponOrderListReq) (res *basics.CouponOrderListRes, err error)
	CouponTypeView(ctx context.Context, req *basics.CouponTypeViewReq) (res *basics.CouponTypeViewRes, err error)
	CouponTypeReceive(ctx context.Context, req *basics.CouponTypeReceiveReq) (res *basics.CouponTypeReceiveRes, err error)
	Dashboard(ctx context.Context, req *basics.DashboardReq) (res *basics.DashboardRes, err error)
	HelpCenterCategoryList(ctx context.Context, req *basics.HelpCenterCategoryListReq) (res *basics.HelpCenterCategoryListRes, err error)
	HelpCenterList(ctx context.Context, req *basics.HelpCenterListReq) (res *basics.HelpCenterListRes, err error)
	HelpCenterView(ctx context.Context, req *basics.HelpCenterViewReq) (res *basics.HelpCenterViewRes, err error)
	HomeDateLoad(ctx context.Context, req *basics.HomeDateLoadReq) (res *basics.HomeDateLoadRes, err error)
	HomeDateConfig(ctx context.Context, req *basics.HomeDateConfigReq) (res *basics.HomeDateConfigRes, err error)
	ReadAlertMessage(ctx context.Context, req *basics.ReadAlertMessageReq) (res *basics.ReadAlertMessageRes, err error)
	HomeDeviceInfo(ctx context.Context, req *basics.HomeDeviceInfoReq) (res *basics.HomeDeviceInfoRes, err error)
	ConfigGet(ctx context.Context, req *basics.ConfigGetReq) (res *basics.ConfigGetRes, err error)
	HomepageArticleList(ctx context.Context, req *basics.HomepageArticleListReq) (res *basics.HomepageArticleListRes, err error)
	HomepageArticleThumb(ctx context.Context, req *basics.HomepageArticleThumbReq) (res *basics.HomepageArticleThumbRes, err error)
	HomepageArticleView(ctx context.Context, req *basics.HomepageArticleViewReq) (res *basics.HomepageArticleViewRes, err error)
	Link(ctx context.Context, req *basics.LinkReq) (res *basics.LinkRes, err error)
	WxReturnUrl(ctx context.Context, req *basics.WxReturnUrlReq) (res *basics.WxReturnUrlRes, err error)
	NotifyList(ctx context.Context, req *basics.NotifyListReq) (res *basics.NotifyListRes, err error)
	NotifyView(ctx context.Context, req *basics.NotifyViewReq) (res *basics.NotifyViewRes, err error)
	NotifyEdit(ctx context.Context, req *basics.NotifyEditReq) (res *basics.NotifyEditRes, err error)
	NotifyDelete(ctx context.Context, req *basics.NotifyDeleteReq) (res *basics.NotifyDeleteRes, err error)
	ReferrerInfo(ctx context.Context, req *basics.ReferrerInfoReq) (res *basics.ReferrerInfoRes, err error)
	ReferrerList(ctx context.Context, req *basics.ReferrerListReq) (res *basics.ReferrerListRes, err error)
	ShareInfo(ctx context.Context, req *basics.ShareInfoReq) (res *basics.ShareInfoRes, err error)
	ShareInvitePageInfo(ctx context.Context, req *basics.ShareInvitePageInfoReq) (res *basics.ShareInvitePageInfoRes, err error)
	StaffFxCenter(ctx context.Context, req *basics.StaffFxCenterReq) (res *basics.StaffFxCenterRes, err error)
	AppLatestMessage(ctx context.Context, req *basics.AppLatestMessageReq) (res *basics.AppLatestMessageRes, err error)
	AppMessageList(ctx context.Context, req *basics.AppMessageListReq) (res *basics.AppMessageListRes, err error)
	AppMessageRead(ctx context.Context, req *basics.AppMessageReadReq) (res *basics.AppMessageReadRes, err error)
	TestNavList(ctx context.Context, req *basics.TestNavListReq) (res *basics.TestNavListRes, err error)
	ThCouponList(ctx context.Context, req *basics.ThCouponListReq) (res *basics.ThCouponListRes, err error)
	ThCouponView(ctx context.Context, req *basics.ThCouponViewReq) (res *basics.ThCouponViewRes, err error)
	ThMchView(ctx context.Context, req *basics.ThMchViewReq) (res *basics.ThMchViewRes, err error)
	ThCouponWaitView(ctx context.Context, req *basics.ThCouponWaitViewReq) (res *basics.ThCouponWaitViewRes, err error)
	ThCouponReceive(ctx context.Context, req *basics.ThCouponReceiveReq) (res *basics.ThCouponReceiveRes, err error)
	UploadImage(ctx context.Context, req *basics.UploadImageReq) (res *basics.UploadImageRes, err error)
}

type IAppCabinet interface {
	CabinetList(ctx context.Context, req *cabinet.CabinetListReq) (res *cabinet.CabinetListRes, err error)
	OrderList(ctx context.Context, req *cabinet.OrderListReq) (res *cabinet.OrderListRes, err error)
	OrderView(ctx context.Context, req *cabinet.OrderViewReq) (res *cabinet.OrderViewRes, err error)
	PayOvertimeInfo(ctx context.Context, req *cabinet.PayOvertimeInfoReq) (res *cabinet.PayOvertimeInfoRes, err error)
	PayOvertime(ctx context.Context, req *cabinet.PayOvertimeReq) (res *cabinet.PayOvertimeRes, err error)
	CancelOrderNoPay(ctx context.Context, req *cabinet.CancelOrderNoPayReq) (res *cabinet.CancelOrderNoPayRes, err error)
	PreOrderCreate(ctx context.Context, req *cabinet.PreOrderCreateReq) (res *cabinet.PreOrderCreateRes, err error)
	PreOrderCreateDetail(ctx context.Context, req *cabinet.PreOrderCreateDetailReq) (res *cabinet.PreOrderCreateDetailRes, err error)
	PrePayInfo(ctx context.Context, req *cabinet.PrePayInfoReq) (res *cabinet.PrePayInfoRes, err error)
	CreateOrder(ctx context.Context, req *cabinet.CreateOrderReq) (res *cabinet.CreateOrderRes, err error)
}

type IAppEmployee interface {
	EmployeeActivityList(ctx context.Context, req *employee.EmployeeActivityListReq) (res *employee.EmployeeActivityListRes, err error)
	EmployeeActivityView(ctx context.Context, req *employee.EmployeeActivityViewReq) (res *employee.EmployeeActivityViewRes, err error)
	EmployeeActivityReceiveCoupon(ctx context.Context, req *employee.EmployeeActivityReceiveCouponReq) (res *employee.EmployeeActivityReceiveCouponRes, err error)
	EmployeeCheckAuth(ctx context.Context, req *employee.EmployeeCheckAuthReq) (res *employee.EmployeeCheckAuthRes, err error)
}

type IAppFx interface {
	Login(ctx context.Context, req *fx.LoginReq) (res *fx.LoginRes, err error)
}

type IAppHotel interface {
	CancelOrderNoPay(ctx context.Context, req *hotel.CancelOrderNoPayReq) (res *hotel.CancelOrderNoPayRes, err error)
	OrderChangeOptions(ctx context.Context, req *hotel.OrderChangeOptionsReq) (res *hotel.OrderChangeOptionsRes, err error)
	OrderChangeGuestEdit(ctx context.Context, req *hotel.OrderChangeGuestEditReq) (res *hotel.OrderChangeGuestEditRes, err error)
	OrderChangeGuestInfo(ctx context.Context, req *hotel.OrderChangeGuestInfoReq) (res *hotel.OrderChangeGuestInfoRes, err error)
	OrderChangeBookingPeopleInfo(ctx context.Context, req *hotel.OrderChangeBookingPeopleInfoReq) (res *hotel.OrderChangeBookingPeopleInfoRes, err error)
	OrderChangeBookingPeople(ctx context.Context, req *hotel.OrderChangeBookingPeopleReq) (res *hotel.OrderChangeBookingPeopleRes, err error)
	OrderChangeBookingPeopleSubmit(ctx context.Context, req *hotel.OrderChangeBookingPeopleSubmitReq) (res *hotel.OrderChangeBookingPeopleSubmitRes, err error)
	OrderChangeDatePreInfo(ctx context.Context, req *hotel.OrderChangeDatePreInfoReq) (res *hotel.OrderChangeDatePreInfoRes, err error)
	OrderChangeDatePreInfoSubmit(ctx context.Context, req *hotel.OrderChangeDatePreInfoSubmitReq) (res *hotel.OrderChangeDatePreInfoSubmitRes, err error)
	OrderChangeStayOnInfo(ctx context.Context, req *hotel.OrderChangeStayOnInfoReq) (res *hotel.OrderChangeStayOnInfoRes, err error)
	OrderChangeInfoDetail(ctx context.Context, req *hotel.OrderChangeInfoDetailReq) (res *hotel.OrderChangeInfoDetailRes, err error)
	PreOrderCreate(ctx context.Context, req *hotel.PreOrderCreateReq) (res *hotel.PreOrderCreateRes, err error)
	PreOrderCreateDetail(ctx context.Context, req *hotel.PreOrderCreateDetailReq) (res *hotel.PreOrderCreateDetailRes, err error)
	PrePayInfo(ctx context.Context, req *hotel.PrePayInfoReq) (res *hotel.PrePayInfoRes, err error)
	CreateOrder(ctx context.Context, req *hotel.CreateOrderReq) (res *hotel.CreateOrderRes, err error)
	OrderList(ctx context.Context, req *hotel.OrderListReq) (res *hotel.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *hotel.OrderDetailReq) (res *hotel.OrderDetailRes, err error)
	PreOrderRefundDetail(ctx context.Context, req *hotel.PreOrderRefundDetailReq) (res *hotel.PreOrderRefundDetailRes, err error)
	RefundOrder(ctx context.Context, req *hotel.RefundOrderReq) (res *hotel.RefundOrderRes, err error)
	OrderRefundDetail(ctx context.Context, req *hotel.OrderRefundDetailReq) (res *hotel.OrderRefundDetailRes, err error)
	ReservationList(ctx context.Context, req *hotel.ReservationListReq) (res *hotel.ReservationListRes, err error)
	ReservationDetail(ctx context.Context, req *hotel.ReservationDetailReq) (res *hotel.ReservationDetailRes, err error)
	FindReservationRoomNo(ctx context.Context, req *hotel.FindReservationRoomNoReq) (res *hotel.FindReservationRoomNoRes, err error)
	PropertySearch(ctx context.Context, req *hotel.PropertySearchReq) (res *hotel.PropertySearchRes, err error)
	PropertyView(ctx context.Context, req *hotel.PropertyViewReq) (res *hotel.PropertyViewRes, err error)
	RoomTypeList(ctx context.Context, req *hotel.RoomTypeListReq) (res *hotel.RoomTypeListRes, err error)
	RoomTypeView(ctx context.Context, req *hotel.RoomTypeViewReq) (res *hotel.RoomTypeViewRes, err error)
}

type IAppMember interface {
	BalanceChange(ctx context.Context, req *member.BalanceChangeReq) (res *member.BalanceChangeRes, err error)
	Cancel(ctx context.Context, req *member.CancelReq) (res *member.CancelRes, err error)
	MemberIntentionEdit(ctx context.Context, req *member.MemberIntentionEditReq) (res *member.MemberIntentionEditRes, err error)
	GetMemberIntentionConfig(ctx context.Context, req *member.GetMemberIntentionConfigReq) (res *member.GetMemberIntentionConfigRes, err error)
	Login(ctx context.Context, req *member.LoginReq) (res *member.LoginRes, err error)
	AuthIdBindLogin(ctx context.Context, req *member.AuthIdBindLoginReq) (res *member.AuthIdBindLoginRes, err error)
	GetUserInfo(ctx context.Context, req *member.GetUserInfoReq) (res *member.GetUserInfoRes, err error)
	BindEmail(ctx context.Context, req *member.BindEmailReq) (res *member.BindEmailRes, err error)
	BindPhone(ctx context.Context, req *member.BindPhoneReq) (res *member.BindPhoneRes, err error)
	StaffInfo(ctx context.Context, req *member.StaffInfoReq) (res *member.StaffInfoRes, err error)
	ChannelInfo(ctx context.Context, req *member.ChannelInfoReq) (res *member.ChannelInfoRes, err error)
	UserLevelRule(ctx context.Context, req *member.UserLevelRuleReq) (res *member.UserLevelRuleRes, err error)
	EditUserInfo(ctx context.Context, req *member.EditUserInfoReq) (res *member.EditUserInfoRes, err error)
	EmailSendCode(ctx context.Context, req *member.EmailSendCodeReq) (res *member.EmailSendCodeRes, err error)
	SmsSendCode(ctx context.Context, req *member.SmsSendCodeReq) (res *member.SmsSendCodeRes, err error)
	OauthLogin(ctx context.Context, req *member.OauthLoginReq) (res *member.OauthLoginRes, err error)
	PhoneNumber(ctx context.Context, req *member.PhoneNumberReq) (res *member.PhoneNumberRes, err error)
	FxAuthLogin(ctx context.Context, req *member.FxAuthLoginReq) (res *member.FxAuthLoginRes, err error)
	GetMemberQrCode(ctx context.Context, req *member.GetMemberQrCodeReq) (res *member.GetMemberQrCodeRes, err error)
}

type IAppPay interface {
	ThirdPay(ctx context.Context, req *pay.ThirdPayReq) (res *pay.ThirdPayRes, err error)
	InnerPay(ctx context.Context, req *pay.InnerPayReq) (res *pay.InnerPayRes, err error)
	QueryStatus(ctx context.Context, req *pay.QueryStatusReq) (res *pay.QueryStatusRes, err error)
}

type IAppTerminal interface {
	TerminalLogin(ctx context.Context, req *terminal.TerminalLoginReq) (res *terminal.TerminalLoginRes, err error)
	TerminalLogout(ctx context.Context, req *terminal.TerminalLogoutReq) (res *terminal.TerminalLogoutRes, err error)
	VerifyLog(ctx context.Context, req *terminal.VerifyLogReq) (res *terminal.VerifyLogRes, err error)
	VerifyLogView(ctx context.Context, req *terminal.VerifyLogViewReq) (res *terminal.VerifyLogViewRes, err error)
	CodeView(ctx context.Context, req *terminal.CodeViewReq) (res *terminal.CodeViewRes, err error)
	CodeVerify(ctx context.Context, req *terminal.CodeVerifyReq) (res *terminal.CodeVerifyRes, err error)
}

type IAppTravel interface {
	CancelOrderNoPay(ctx context.Context, req *travel.CancelOrderNoPayReq) (res *travel.CancelOrderNoPayRes, err error)
	PreOrderCreate(ctx context.Context, req *travel.PreOrderCreateReq) (res *travel.PreOrderCreateRes, err error)
	PreOrderCreateDetail(ctx context.Context, req *travel.PreOrderCreateDetailReq) (res *travel.PreOrderCreateDetailRes, err error)
	PrePayInfo(ctx context.Context, req *travel.PrePayInfoReq) (res *travel.PrePayInfoRes, err error)
	CreateOrder(ctx context.Context, req *travel.CreateOrderReq) (res *travel.CreateOrderRes, err error)
	OrderList(ctx context.Context, req *travel.OrderListReq) (res *travel.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *travel.OrderDetailReq) (res *travel.OrderDetailRes, err error)
	MemberVerifyCodeRefresh(ctx context.Context, req *travel.MemberVerifyCodeRefreshReq) (res *travel.MemberVerifyCodeRefreshRes, err error)
	ApplyOrderRefundDetail(ctx context.Context, req *travel.ApplyOrderRefundDetailReq) (res *travel.ApplyOrderRefundDetailRes, err error)
	RefundOrder(ctx context.Context, req *travel.RefundOrderReq) (res *travel.RefundOrderRes, err error)
	OrderRefundDetail(ctx context.Context, req *travel.OrderRefundDetailReq) (res *travel.OrderRefundDetailRes, err error)
	ProductList(ctx context.Context, req *travel.ProductListReq) (res *travel.ProductListRes, err error)
	ProductView(ctx context.Context, req *travel.ProductViewReq) (res *travel.ProductViewRes, err error)
	ProductSkuStock(ctx context.Context, req *travel.ProductSkuStockReq) (res *travel.ProductSkuStockRes, err error)
	StaffConfig(ctx context.Context, req *travel.StaffConfigReq) (res *travel.StaffConfigRes, err error)
	StaffLogin(ctx context.Context, req *travel.StaffLoginReq) (res *travel.StaffLoginRes, err error)
	StaffLogout(ctx context.Context, req *travel.StaffLogoutReq) (res *travel.StaffLogoutRes, err error)
	VerifyLog(ctx context.Context, req *travel.VerifyLogReq) (res *travel.VerifyLogRes, err error)
	VerifyLogView(ctx context.Context, req *travel.VerifyLogViewReq) (res *travel.VerifyLogViewRes, err error)
	CodeView(ctx context.Context, req *travel.CodeViewReq) (res *travel.CodeViewRes, err error)
	CodeVerify(ctx context.Context, req *travel.CodeVerifyReq) (res *travel.CodeVerifyRes, err error)
}

type IAppVerify interface {
	MemberVerify(ctx context.Context, req *verify.MemberVerifyReq) (res *verify.MemberVerifyRes, err error)
	MemberVerifyCodeRefresh(ctx context.Context, req *verify.MemberVerifyCodeRefreshReq) (res *verify.MemberVerifyCodeRefreshRes, err error)
}
