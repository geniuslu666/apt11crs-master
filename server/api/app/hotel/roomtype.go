package hotel

import (
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RoomTypeListReq struct {
	g.Meta       `path:"/roomType/list" method:"post" tags:"APP_HOTEL" summary:"房型_检索房型列表"`
	Puid         string `json:"puid" v:"required#please_enter_puid" dc:"物业id"`
	StartDate    string `json:"startDate" v:"required|date#please_enter_the_checkin_date|date_format_error" dc:"入住日期"`
	EndDate      string `json:"endDate" v:"required|date#please_enter_the_checkout_date|date_format_error" dc:"退房日期"`
	RoomNumber   int    `json:"roomNumber" v:"required#please_enter_the_number_of_rooms" dc:"房间数"`
	BookerNumber int    `json:"bookerNumber" v:"required#please_enter_the_number_of_occupants" dc:"入住人数"`
}

type RoomTypeListRes struct {
	List []*struct {
		Uid                    string      `json:"uid"          orm:"uid"           description:"第三方系统的ID"`
		Id                     int         `json:"id"           orm:"id"            description:"主键ID"`
		Puid                   string      `json:"puid"         orm:"puid"          description:"物业ID"`
		Cover                  string      `json:"cover"        orm:"cover"         description:"封面"`
		CoverList              *gjson.Json `json:"coverList"    orm:"cover_list"    description:"照片墙"`
		Name                   string      `json:"name"         orm:"name"          description:"房型名称"`
		BasePrice              float64     `json:"minPrice"     orm:"base_price"    description:"最低价格"`
		CheckinAt              string      `json:"checkinAt"    orm:"checkin_at"    description:"入住时间"`
		CheckoutAt             string      `json:"checkoutAt"   orm:"checkout_at"   description:"退房时间"`
		BookingStyle           string      `json:"bookingStyle" orm:"booking_style" description:"预订方式"`
		RoomStyle              string      `json:"roomStyle"    orm:"room_style"    description:"房间风格"`
		Occupancy              int         `json:"occupancy"    orm:"occupancy"     description:"占用"`
		Size                   string      `json:"size"         orm:"size"          description:"房间大小"`
		Bedrooms               string      `json:"bedrooms"     orm:"bedrooms"      description:"卧室"`
		Bathrooms              string      `json:"bathrooms"    orm:"bathrooms"     description:"浴室"`
		CleaningFee            float64     `json:"cleaningFee"  orm:"cleaning_fee"  description:"清理费"`
		RatePlanId             string      `json:"ratePlanId"   orm:"rate_plan_id"  description:"费率计划模板ID"`
		CreateAt               *gtime.Time `json:"createAt"     orm:"create_at"     description:""`
		UpdateAt               *gtime.Time `json:"updateAt"     orm:"update_at"     description:""`
		Inventory              int         `json:"inventory" dc:"库存量"`
		AdditionalGuestAmounts float64     `json:"additionalGuestAmounts" orm:"additional_guest_amounts" description:"额外客人金额"`
		OccupantsForBaseRate   int         `json:"occupantsForBaseRate"   orm:"occupants_for_base_rate"  description:"无需增加额外客人金额人数"`
	} `json:"property"`
}

type RoomTypeViewReq struct {
	g.Meta `path:"/roomType/view" method:"post" tags:"APP_HOTEL" summary:"房型_详情"`
	Uid    string `json:"uid" v:"required#please_enter_the_room_type_UID" dc:"房型UID"`
	Id     int    `json:"id" v:"required#please_enter_the_room_type_ID" dc:"房型ID"`
}

type RoomTypeViewRes struct {
	*input_hotel.PmsRoomTypeViewModel
}
