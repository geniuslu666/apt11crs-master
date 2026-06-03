package toreta

import (
	"APT/internal/library/toretaApi"

	"github.com/gogf/gf/v2/frame/g"
)

// RestaurantsReq 获取店铺一览请求
type RestaurantsReq struct {
	g.Meta `path:"/toretaNotify/restaurants" method:"post" tags:"TORETA_HOOK" summary:"获取店铺一览"`
	*toretaApi.RestaurantsParams
}

// RestaurantsRes 获取店铺一览响应
type RestaurantsRes struct {
	*toretaApi.RestaurantsResponse
}

// RestaurantDetailReq 获取店铺详细信息请求
type RestaurantDetailReq struct {
	g.Meta `path:"/toretaNotify/restaurantDetail" method:"post" tags:"TORETA_HOOK" summary:"获取店铺详细信息"`
	*toretaApi.RestaurantDetailParams
}

// RestaurantDetailRes 获取店铺详细信息响应
type RestaurantDetailRes struct {
	*toretaApi.RestaurantDetailResponse
}

// CoursesReq 获取课程一览请求
type CoursesReq struct {
	g.Meta `path:"/toretaNotify/courses" method:"post" tags:"TORETA_HOOK" summary:"获取课程一览"`
	*toretaApi.CoursesParams
}

// CoursesRes 获取课程一览响应
type CoursesRes struct {
	*toretaApi.CoursesResponse
}

// SlotsReq 获取空席时间段请求
type SlotsReq struct {
	g.Meta `path:"/toretaNotify/slots" method:"post" tags:"TORETA_HOOK" summary:"获取空席时间段"`
	*toretaApi.SlotsParams
}

// SlotsRes 获取空席时间段响应
type SlotsRes struct {
	*toretaApi.SlotsResponse
}

// CreateReservationReq 创建预约请求
type CreateReservationReq struct {
	g.Meta `path:"/toretaNotify/createReservation" method:"post" tags:"TORETA_HOOK" summary:"创建预约"`
	*toretaApi.CreateReservationParams
}

// CreateReservationRes 创建预约响应
type CreateReservationRes struct {
	*toretaApi.CreateReservationResponse
}

// CancelReservationReq 取消预约请求
type CancelReservationReq struct {
	g.Meta `path:"/toretaNotify/cancelReservation" method:"post" tags:"TORETA_HOOK" summary:"取消预约"`
	*toretaApi.CancelReservationParams
}

// CancelReservationRes 取消预约响应
type CancelReservationRes struct {
	*toretaApi.CancelReservationResponse
}

// ReservationDetailReq 预约详情请求
type ReservationDetailReq struct {
	g.Meta `path:"/toretaNotify/reservationDetail" method:"post" tags:"TORETA_HOOK" summary:"预约详情"`
	*toretaApi.ReservationDetailParams
}

// ReservationDetailRes 预约详情响应
type ReservationDetailRes struct {
	*toretaApi.ReservationDetailResponse
}

// NotificationsReq 获取通知日志请求
type NotificationsReq struct {
	g.Meta `path:"/toretaNotify/notifications" method:"post" tags:"TORETA_HOOK" summary:"获取通知日志"`
	*toretaApi.NotificationsParams
}

// NotificationsRes 获取通知日志响应
type NotificationsRes struct {
	*toretaApi.NotificationsResponse
}

// OARestaurantsReq 获取店铺一览请求
type OARestaurantsReq struct {
	g.Meta `path:"/toretaNotify/oARestaurants" method:"post" tags:"TORETA_HOOK" summary:"获取店铺一览"`
}

// OARestaurantsRes 获取店铺一览响应
type OARestaurantsRes struct {
	*toretaApi.OARestaurantsResponse
}

// OARestaurantsBindReq 店铺绑定请求
type OARestaurantsBindReq struct {
	g.Meta `path:"/toretaNotify/oARestaurantsBind" method:"post" tags:"TORETA_HOOK" summary:"店铺绑定"`
	*toretaApi.OARestaurantBindParams
}

// OARestaurantsBindRes 店铺绑定响应
type OARestaurantsBindRes struct {
	*toretaApi.OARestaurantBindResponse
}

// RefreshOATokenReq 刷新Token请求
type RefreshOATokenReq struct {
	g.Meta `path:"/toretaNotify/refreshOAToken" method:"post" tags:"TORETA_HOOK" summary:"刷新Token"`
	*toretaApi.RefreshOATokenParams
}

// RefreshOATokenRes 刷新Token响应
type RefreshOATokenRes struct {
	*toretaApi.RefreshOATokenResponse
}
