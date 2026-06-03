// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomType is the golang structure for table pms_room_type.
type PmsRoomType struct {
	Uid                    string      `json:"uid"                    orm:"uid"                      description:"airhost房型ID"`
	Id                     int         `json:"id"                     orm:"id"                       description:"主键ID"`
	Puid                   string      `json:"puid"                   orm:"puid"                     description:"物业ID"`
	Cover                  string      `json:"cover"                  orm:"cover"                    description:"封面"`
	CoverList              *gjson.Json `json:"coverList"              orm:"cover_list"               description:"照片墙"`
	Name                   string      `json:"name"                   orm:"name"                     description:"房型多语言名称ID"`
	BasePrice              float64     `json:"basePrice"              orm:"base_price"               description:"最低价格"`
	CheckinAt              string      `json:"checkinAt"              orm:"checkin_at"               description:"入住时间"`
	CheckoutAt             string      `json:"checkoutAt"             orm:"checkout_at"              description:"退房时间"`
	BookingStyle           string      `json:"bookingStyle"           orm:"booking_style"            description:"预订方式"`
	RoomStyle              string      `json:"roomStyle"              orm:"room_style"               description:"房间的风格"`
	Occupancy              int         `json:"occupancy"              orm:"occupancy"                description:"占用"`
	Size                   string      `json:"size"                   orm:"size"                     description:"面积"`
	Bedrooms               string      `json:"bedrooms"               orm:"bedrooms"                 description:"卧室"`
	Bathrooms              string      `json:"bathrooms"              orm:"bathrooms"                description:"浴室"`
	CleaningFee            float64     `json:"cleaningFee"            orm:"cleaning_fee"             description:"清理费"`
	RatePlanId             string      `json:"ratePlanId"             orm:"rate_plan_id"             description:"费率ID"`
	AdditionalGuestAmounts float64     `json:"additionalGuestAmounts" orm:"additional_guest_amounts" description:"额外客人金额"`
	OccupantsForBaseRate   int         `json:"occupantsForBaseRate"   orm:"occupants_for_base_rate"  description:"无需增加额外客人金额人数"`
	RoomNum                int         `json:"roomNum"                orm:"room_num"                 description:"房间数"`
	IsShow                 int         `json:"isShow"                 orm:"Is_show"                  description:"1、显示 0 隐藏"`
	CreateAt               *gtime.Time `json:"createAt"               orm:"create_at"                description:""`
	UpdateAt               *gtime.Time `json:"updateAt"               orm:"update_at"                description:""`
}
