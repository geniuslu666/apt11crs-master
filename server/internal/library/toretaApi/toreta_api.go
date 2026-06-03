package toretaApi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

// checkToretaError 检查 TORETA API 响应中的错误，返回响应体和错误
func checkToretaError(response *gclient.Response) (string, error) {
	responseBody := response.ReadAllString()

	// 如果 HTTP 状态码不是 2xx，检查是否是错误响应
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		// 尝试解析错误响应
		var errorResp ToretaErrorResponse
		if err := gjson.New(responseBody).Scan(&errorResp); err == nil && errorResp.Error.Status != 0 {
			return responseBody, gerror.Newf("TORETA API Error (Status: %d): %s", errorResp.Error.Status, errorResp.Error.Detail)
		}
		// 如果不是标准错误格式，返回 HTTP 错误
		return responseBody, gerror.Newf("TORETA API HTTP Error (Status: %d): %s", response.StatusCode, responseBody)
	}

	return responseBody, nil
}

// checkToretaOAError 检查 TORETA OAuth 响应中的错误，返回响应体和错误
func checkToretaOAError(response *gclient.Response) (string, error) {
	responseBody := response.ReadAllString()

	// 尝试解析错误响应
	var errorResp ToretaOAErrorResponse
	if err := gjson.New(responseBody).Scan(&errorResp); err == nil && !g.IsEmpty(errorResp.Error) {
		return responseBody, gerror.Newf("TORETA OAuth API Error (Error: %s, ErrorDescription: %s)", errorResp.Error, errorResp.ErrorDescription)
	}

	return responseBody, nil
}

// GetRestaurants 获取店铺一览
func (c *ToretaClient) GetRestaurants(ctx context.Context, params *RestaurantsParams) (resp *RestaurantsResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "GET", restaurantsEndpoint, false, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// GetRestaurantDetail 获取店铺详细信息
func (c *ToretaClient) GetRestaurantDetail(ctx context.Context, params *RestaurantDetailParams) (resp *RestaurantDetailResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
		url          string
	)
	// 构建 URL：/externals/public/restaurants/:key.json
	url = fmt.Sprintf("%s/%s.json", restaurantDetailEndpoint, params.Key)

	if response, err = c.DoRequest(ctx, "GET", url, false, nil); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// GetCourses 获取课程一览
func (c *ToretaClient) GetCourses(ctx context.Context, params *CoursesParams) (resp *CoursesResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
		url          string
	)
	// 构建 URL：/externals/public/restaurants/:restaurant_id/courses.json
	url = fmt.Sprintf("%s/%s/courses.json", coursesEndpoint, params.RestaurantId)

	if response, err = c.DoRequest(ctx, "GET", url, false, nil); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	// API返回的是直接的课程数组，需要包装到CoursesResponse结构中
	var courses []CourseItem
	if err = gjson.New(responseBody).Scan(&courses); err != nil {
		return
	}

	resp = &CoursesResponse{
		Courses: courses,
	}

	return
}

// GetSlots 获取单店铺空席时间段
func (c *ToretaClient) GetSlots(ctx context.Context, params *SlotsParams) (resp *SlotsResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "GET", slotsEndpoint, false, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// CreateReservation 创建 Web 预约
func (c *ToretaClient) CreateReservation(ctx context.Context, params *CreateReservationParams) (resp *CreateReservationResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "POST", createReservation, false, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// CancelReservation 取消预约
func (c *ToretaClient) CancelReservation(ctx context.Context, params *CancelReservationParams) (resp *CancelReservationResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
		url          string
	)
	// 构建 URL：/externals/public/web_reservations/:id.json
	url = fmt.Sprintf("%s/%s.json", cancelReservation, params.ReservationId)

	if response, err = c.DoRequest(ctx, "DELETE", url, false, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// ReservationDetail 预约详情
func (c *ToretaClient) ReservationDetail(ctx context.Context, params *ReservationDetailParams) (resp *ReservationDetailResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
		url          string
	)
	// 构建 URL：/externals/public/web_reservations/:id.json
	url = fmt.Sprintf("%s/%s.json", cancelReservation, params.ReservationId)

	if response, err = c.DoRequest(ctx, "GET", url, false, nil); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// Notifications 获取通知日志
func (c *ToretaClient) Notifications(ctx context.Context, params *NotificationsParams) (resp *NotificationsResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "GET", notificationsEndpoint, false, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	// 直接解析为数组
	resp = &NotificationsResponse{}
	if err = gjson.New(responseBody).Scan(resp); err != nil {
		return
	}

	return
}

// GetOARestaurants 获取店铺一览
func (c *ToretaClient) GetOARestaurants(ctx context.Context) (resp *OARestaurantsResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "GET", connections, true, nil); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// OARestaurantsBind 绑定店铺
func (c *ToretaClient) OARestaurantsBind(ctx context.Context, params *OARestaurantBindParams) (resp *OARestaurantBindResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
	)
	if response, err = c.DoRequest(ctx, "POST", connections, true, params); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}

// RefreshOAToken 刷新 OAuth Token
func (c *ToretaClient) RefreshOAToken(ctx context.Context, params *RefreshOATokenParams) (resp *RefreshOATokenResponse, err error) {
	var (
		response     *gclient.Response
		responseBody string
		apiParams    *RefreshOATokenApiParams
	)
	apiParams = &RefreshOATokenApiParams{
		GrantType:    "refresh_token",
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RefreshToken: params.RefreshToken,
	}
	if response, err = c.DoOARequest(ctx, "POST", oauth, apiParams); err != nil {
		return
	}

	// 检查错误响应
	if responseBody, err = checkToretaOAError(response); err != nil {
		return
	}

	if err = gjson.New(responseBody).Scan(&resp); err != nil {
		return
	}

	return
}
