// Package hook
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package hook

import (
	"context"
	"fmt"
)

// SystemMessageHelper 系统消息助手，提供常用的快捷方法
type SystemMessageHelper struct{}

// NewSystemMessageHelper 创建系统消息助手实例
func NewSystemMessageHelper() *SystemMessageHelper {
	return &SystemMessageHelper{}
}

// SendOrderMessage 发送订单相关消息的快捷方法
func (h *SystemMessageHelper) SendOrderMessage(ctx context.Context, scene string, memberId uint64, orderInfo OrderMessageInfo) (*SystemMessageHookResult, error) {
	params := &SystemMessageHookParams{
		Scene:       scene,
		Type:        "order",
		MemberId:    memberId,
		Title:       orderInfo.Title,
		Content:     orderInfo.Content,
		Image:       orderInfo.Image,
		AppLink:     orderInfo.AppLink,
		WxLink:      orderInfo.WxLink,
		UrlParam:    orderInfo.UrlParam,
		OperatorId:  orderInfo.OperatorId,
		EnablePush:  orderInfo.EnablePush,
		EnableSms:   orderInfo.EnableSms,
		PushTitle:   orderInfo.PushTitle,
		PushContent: orderInfo.PushContent,
		SmsTemplate: orderInfo.SmsTemplate,
		SmsParams:   orderInfo.SmsParams,
	}

	return ProcessSystemMessage(ctx, params)
}

// SendIMMessage 发送IM聊天消息的快捷方法
func (h *SystemMessageHelper) SendIMMessage(ctx context.Context, scene string, memberId uint64, imInfo IMMessageInfo) (*SystemMessageHookResult, error) {
	params := &SystemMessageHookParams{
		Scene:       scene,
		Type:        "im",
		MemberId:    memberId,
		Title:       imInfo.Title,
		Content:     imInfo.Content,
		Image:       imInfo.Image,
		AppLink:     imInfo.AppLink,
		WxLink:      imInfo.WxLink,
		UrlParam:    imInfo.UrlParam,
		OperatorId:  imInfo.OperatorId,
		EnablePush:  imInfo.EnablePush,
		PushTitle:   imInfo.PushTitle,
		PushContent: imInfo.PushContent,
		// IM消息通常不发送短信
		EnableSms: false,
	}

	return ProcessSystemMessage(ctx, params)
}

// OrderMessageInfo 订单消息信息结构体
type OrderMessageInfo struct {
	Title       map[string]string      `json:"title"`
	Content     map[string]string      `json:"content,omitempty"`
	Image       string                 `json:"image,omitempty"`
	AppLink     string                 `json:"appLink,omitempty"`
	WxLink      string                 `json:"wxLink,omitempty"`
	UrlParam    map[string]interface{} `json:"urlParam,omitempty"`
	OperatorId  uint64                 `json:"operatorId,omitempty"`
	EnablePush  bool                   `json:"enablePush"`
	EnableSms   bool                   `json:"enableSms"`
	PushTitle   string                 `json:"pushTitle,omitempty"`
	PushContent string                 `json:"pushContent,omitempty"`
	SmsTemplate string                 `json:"smsTemplate,omitempty"`
	SmsParams   map[string]string      `json:"smsParams,omitempty"`
}

// IMMessageInfo IM消息信息结构体
type IMMessageInfo struct {
	Title       map[string]string      `json:"title"`
	Content     map[string]string      `json:"content,omitempty"`
	Image       string                 `json:"image,omitempty"`
	AppLink     string                 `json:"appLink,omitempty"`
	WxLink      string                 `json:"wxLink,omitempty"`
	UrlParam    map[string]interface{} `json:"urlParam,omitempty"`
	OperatorId  uint64                 `json:"operatorId,omitempty"`
	EnablePush  bool                   `json:"enablePush"`
	PushTitle   string                 `json:"pushTitle,omitempty"`
	PushContent string                 `json:"pushContent,omitempty"`
}

// 预定义的常用消息模板

// CreateHotelOrderConfirmMessage 创建酒店订单确认消息
func (h *SystemMessageHelper) CreateHotelOrderConfirmMessage(hotelName, hotelNameEn, checkInDate, orderNo string, orderId, hotelId uint64) OrderMessageInfo {
	return OrderMessageInfo{
		Title: map[string]string{
			"zh": "酒店预订确认",
			"en": "Hotel Reservation Confirmed",
		},
		Content: map[string]string{
			"zh": fmt.Sprintf("您的%s预订已确认，入住时间：%s", hotelName, checkInDate),
			"en": fmt.Sprintf("Your reservation at %s is confirmed, check-in: %s", hotelNameEn, checkInDate),
		},
		AppLink: "/hotel/order/detail",
		UrlParam: map[string]interface{}{
			"order_id": orderId,
			"hotel_id": hotelId,
		},
		EnablePush:  true,
		EnableSms:   true,
		PushTitle:   "酒店预订确认",
		PushContent: fmt.Sprintf("您的%s预订已确认", hotelName),
		SmsTemplate: "HOTEL_ORDER_CONFIRM",
		SmsParams: map[string]string{
			"hotel_name": hotelName,
			"check_in":   checkInDate,
			"order_no":   orderNo,
		},
	}
}

// CreateFoodOrderConfirmMessage 创建餐厅订单确认消息
func (h *SystemMessageHelper) CreateFoodOrderConfirmMessage(restaurantName, restaurantNameEn, bookingTime, orderNo string, orderId, restaurantId uint64) OrderMessageInfo {
	return OrderMessageInfo{
		Title: map[string]string{
			"zh": "餐厅预订确认",
			"en": "Restaurant Reservation Confirmed",
		},
		Content: map[string]string{
			"zh": fmt.Sprintf("您的%s预订已确认，用餐时间：%s", restaurantName, bookingTime),
			"en": fmt.Sprintf("Your reservation at %s is confirmed, dining time: %s", restaurantNameEn, bookingTime),
		},
		AppLink: "/food/order/detail",
		UrlParam: map[string]interface{}{
			"order_id":      orderId,
			"restaurant_id": restaurantId,
		},
		EnablePush:  true,
		EnableSms:   true,
		PushTitle:   "餐厅预订确认",
		PushContent: fmt.Sprintf("您的%s预订已确认", restaurantName),
		SmsTemplate: "FOOD_ORDER_CONFIRM",
		SmsParams: map[string]string{
			"restaurant_name": restaurantName,
			"booking_time":    bookingTime,
			"order_no":        orderNo,
		},
	}
}

// CreateCarOrderConfirmMessage 创建接送机订单确认消息
func (h *SystemMessageHelper) CreateCarOrderConfirmMessage(serviceType, pickupTime, location, orderNo string, orderId uint64) OrderMessageInfo {
	var titleZh, titleEn, contentZh, contentEn string

	if serviceType == "pickup" {
		titleZh = "接机服务确认"
		titleEn = "Airport Pickup Confirmed"
		contentZh = fmt.Sprintf("您的接机服务已确认，时间：%s，地点：%s", pickupTime, location)
		contentEn = fmt.Sprintf("Your airport pickup is confirmed, time: %s, location: %s", pickupTime, location)
	} else {
		titleZh = "送机服务确认"
		titleEn = "Airport Drop-off Confirmed"
		contentZh = fmt.Sprintf("您的送机服务已确认，时间：%s，地点：%s", pickupTime, location)
		contentEn = fmt.Sprintf("Your airport drop-off is confirmed, time: %s, location: %s", pickupTime, location)
	}

	return OrderMessageInfo{
		Title: map[string]string{
			"zh": titleZh,
			"en": titleEn,
		},
		Content: map[string]string{
			"zh": contentZh,
			"en": contentEn,
		},
		AppLink: "/car/order/detail",
		UrlParam: map[string]interface{}{
			"order_id":     orderId,
			"service_type": serviceType,
		},
		EnablePush:  true,
		EnableSms:   true,
		PushTitle:   titleZh,
		PushContent: contentZh,
		SmsTemplate: "CAR_ORDER_CONFIRM",
		SmsParams: map[string]string{
			"service_type": serviceType,
			"pickup_time":  pickupTime,
			"location":     location,
			"order_no":     orderNo,
		},
	}
}

// CreateSpaOrderConfirmMessage 创建按摩订单确认消息
func (h *SystemMessageHelper) CreateSpaOrderConfirmMessage(serviceName, serviceNameEn, bookingTime, storeName, orderNo string, orderId, serviceId uint64) OrderMessageInfo {
	return OrderMessageInfo{
		Title: map[string]string{
			"zh": "按摩服务预约确认",
			"en": "Spa Service Booking Confirmed",
		},
		Content: map[string]string{
			"zh": fmt.Sprintf("您的%s预约已确认，时间：%s，地点：%s", serviceName, bookingTime, storeName),
			"en": fmt.Sprintf("Your %s booking is confirmed, time: %s, location: %s", serviceNameEn, bookingTime, storeName),
		},
		AppLink: "/spa/order/detail",
		UrlParam: map[string]interface{}{
			"order_id":   orderId,
			"service_id": serviceId,
		},
		EnablePush:  true,
		EnableSms:   true,
		PushTitle:   "按摩服务预约确认",
		PushContent: fmt.Sprintf("您的%s预约已确认", serviceName),
		SmsTemplate: "SPA_ORDER_CONFIRM",
		SmsParams: map[string]string{
			"service_name": serviceName,
			"booking_time": bookingTime,
			"store_name":   storeName,
			"order_no":     orderNo,
		},
	}
}

// CreateSystemNotificationMessage 创建系统通知消息
func (h *SystemMessageHelper) CreateSystemNotificationMessage(titleZh, titleEn, contentZh, contentEn string, enablePush bool) OrderMessageInfo {
	return OrderMessageInfo{
		Title: map[string]string{
			"zh": titleZh,
			"en": titleEn,
		},
		Content: map[string]string{
			"zh": contentZh,
			"en": contentEn,
		},
		EnablePush:  enablePush,
		EnableSms:   false, // 系统通知通常不发送短信
		PushTitle:   titleZh,
		PushContent: contentZh,
	}
}
