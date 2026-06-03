package notify

import (
	"context"

	"APT/api/notify/toreta"
	"APT/internal/library/toretaApi"
	"APT/internal/model"
	"APT/internal/service"
)

func (c *ControllerToreta) Restaurants(ctx context.Context, req *toreta.RestaurantsReq) (res *toreta.RestaurantsRes, err error) {
	var (
		toretaResponse  *toretaApi.RestaurantsResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetRestaurants(ctx, req.RestaurantsParams); err != nil {
		return
	}
	res = &toreta.RestaurantsRes{
		RestaurantsResponse: toretaResponse,
	}
	return
}

func (c *ControllerToreta) RestaurantDetail(ctx context.Context, req *toreta.RestaurantDetailReq) (res *toreta.RestaurantDetailRes, err error) {
	var (
		toretaResponse  *toretaApi.RestaurantDetailResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetRestaurantDetail(ctx, req.RestaurantDetailParams); err != nil {
		return
	}
	res = &toreta.RestaurantDetailRes{
		RestaurantDetailResponse: toretaResponse,
	}
	return
}

func (c *ControllerToreta) Courses(ctx context.Context, req *toreta.CoursesReq) (res *toreta.CoursesRes, err error) {
	var (
		toretaResponse  *toretaApi.CoursesResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetCourses(ctx, req.CoursesParams); err != nil {
		return
	}
	res = &toreta.CoursesRes{
		CoursesResponse: toretaResponse,
	}
	return
}

func (c *ControllerToreta) Slots(ctx context.Context, req *toreta.SlotsReq) (res *toreta.SlotsRes, err error) {
	var (
		toretaResponse  *toretaApi.SlotsResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetSlots(ctx, req.SlotsParams); err != nil {
		return
	}
	res = &toreta.SlotsRes{
		SlotsResponse: toretaResponse,
	}
	return
}

func (c *ControllerToreta) CreateReservation(ctx context.Context, req *toreta.CreateReservationReq) (res *toreta.CreateReservationRes, err error) {
	var (
		toretaResponse  *toretaApi.CreateReservationResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).CreateReservation(ctx, req.CreateReservationParams); err != nil {
		return
	}
	res = &toreta.CreateReservationRes{
		CreateReservationResponse: toretaResponse,
	}
	return
}

func (c *ControllerToreta) CancelReservation(ctx context.Context, req *toreta.CancelReservationReq) (res *toreta.CancelReservationRes, err error) {
	var (
		toretaResponse  *toretaApi.CancelReservationResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).CancelReservation(ctx, req.CancelReservationParams); err != nil {
		return
	}
	res = &toreta.CancelReservationRes{
		CancelReservationResponse: toretaResponse,
	}
	return
}
func (c *ControllerToreta) ReservationDetail(ctx context.Context, req *toreta.ReservationDetailReq) (res *toreta.ReservationDetailRes, err error) {
	var (
		toretaResponse  *toretaApi.ReservationDetailResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).ReservationDetail(ctx, req.ReservationDetailParams); err != nil {
		return
	}
	res = &toreta.ReservationDetailRes{
		ReservationDetailResponse: toretaResponse,
	}
	return
}
func (c *ControllerToreta) Notifications(ctx context.Context, req *toreta.NotificationsReq) (res *toreta.NotificationsRes, err error) {
	var (
		toretaResponse  *toretaApi.NotificationsResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).Notifications(ctx, req.NotificationsParams); err != nil {
		return
	}
	res = &toreta.NotificationsRes{
		NotificationsResponse: toretaResponse,
	}
	return
}
func (c *ControllerToreta) OARestaurants(ctx context.Context, req *toreta.OARestaurantsReq) (res *toreta.OARestaurantsRes, err error) {
	var (
		toretaResponse  *toretaApi.OARestaurantsResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).GetOARestaurants(ctx); err != nil {
		return
	}
	res = &toreta.OARestaurantsRes{
		OARestaurantsResponse: toretaResponse,
	}
	return
}
func (c *ControllerToreta) OARestaurantsBind(ctx context.Context, req *toreta.OARestaurantsBindReq) (res *toreta.OARestaurantsBindRes, err error) {
	var (
		toretaResponse  *toretaApi.OARestaurantBindResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).OARestaurantsBind(ctx, req.OARestaurantBindParams); err != nil {
		return
	}
	res = &toreta.OARestaurantsBindRes{
		OARestaurantBindResponse: toretaResponse,
	}
	return
}
func (c *ControllerToreta) RefreshOAToken(ctx context.Context, req *toreta.RefreshOATokenReq) (res *toreta.RefreshOATokenRes, err error) {
	var (
		toretaResponse  *toretaApi.RefreshOATokenResponse
		ToretaApiConfig *model.ToretaApiConfig
	)
	if ToretaApiConfig, err = service.BasicsConfig().GetToretaApi(ctx); err != nil {
		return
	}
	if toretaResponse, err = toretaApi.NewClient(ctx, ToretaApiConfig).RefreshOAToken(ctx, req.RefreshOATokenParams); err != nil {
		return
	}
	res = &toreta.RefreshOATokenRes{
		RefreshOATokenResponse: toretaResponse,
	}
	return
}
