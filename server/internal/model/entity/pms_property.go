// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsProperty is the golang structure for table pms_property.
type PmsProperty struct {
	Id                   int         `json:"id"                   orm:"id"                      description:"主键"`
	Uid                  string      `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
	Icon                 string      `json:"icon"                 orm:"icon"                    description:"图标"`
	GroupIds             string      `json:"groupIds"             orm:"group_ids"               description:"开放显示的会员分组IDS"`
	Name                 string      `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	Style                string      `json:"style"                orm:"style"                   description:"物业类型"`
	Cover                string      `json:"cover"                orm:"cover"                   description:"封面"`
	Currency             string      `json:"currency"             orm:"currency"                description:"物业的默认货币"`
	Language             string      `json:"language"             orm:"language"                description:"物业的默认自动消息语言"`
	TimeZone             string      `json:"timeZone"             orm:"time_zone"               description:"标准时区名称"`
	Address              string      `json:"address"              orm:"address"                 description:"地址描述"`
	Lat                  string      `json:"lat"                  orm:"lat"                     description:"纬度"`
	Lng                  string      `json:"lng"                  orm:"lng"                     description:"经度"`
	AddressDetail        string      `json:"addressDetail"        orm:"address_detail"          description:"百度详细地址"`
	GgAddressDetail      string      `json:"ggAddressDetail"      orm:"gg_address_detail"       description:"google详细地址"`
	GgLat                string      `json:"ggLat"                orm:"gg_lat"                  description:"谷歌纬度"`
	GgLng                string      `json:"ggLng"                orm:"gg_lng"                  description:"谷歌经度"`
	TagList              string      `json:"tagList"              orm:"tag_list"                description:"标签多语言"`
	RoomDes              string      `json:"roomDes"              orm:"room_des"                description:"房间描述"`
	Surroundings         string      `json:"surroundings"         orm:"surroundings"            description:"周边环境"`
	Description          string      `json:"description"          orm:"description"             description:"物业描述  多语言"`
	BusStation           string      `json:"busStation"           orm:"bus_station"             description:"公交站"`
	Subway               string      `json:"subway"               orm:"subway"                  description:"地铁站"`
	ContactName          string      `json:"contactName"          orm:"contact_name"            description:"联系人"`
	Phone                string      `json:"phone"                orm:"phone"                   description:"联系方式"`
	ContactEmail         string      `json:"contactEmail"         orm:"contact_email"           description:"邮箱"`
	MinDaysNotice        int         `json:"minDaysNotice"        orm:"min_days_notice"         description:"短租模式最小预定区间"`
	MaxDaysNotice        int         `json:"maxDaysNotice"        orm:"max_days_notice"         description:"最大预定区间"`
	MinutesAfterCheckout int         `json:"minutesAfterCheckout" orm:"minutes_after_checkout"  description:"退房后 分钟"`
	MinutesBeforeCheckin int         `json:"minutesBeforeCheckin" orm:"minutes_before_checkin"  description:"入住前 分钟"`
	BookingLeadTimeLabel string      `json:"bookingLeadTimeLabel" orm:"booking_lead_time_label" description:"预约期限"`
	TurnoverDays         uint        `json:"turnoverDays"         orm:"turnover_days"           description:"周转天数"`
	LinePrice            float64     `json:"linePrice"            orm:"line_price"              description:"单价"`
	Price                float64     `json:"price"                orm:"price"                   description:"单价"`
	CheckinAt            string      `json:"checkinAt"            orm:"checkin_at"              description:"入住时间"`
	CheckoutAt           string      `json:"checkoutAt"           orm:"checkout_at"             description:"退房时间"`
	CancelPolicy         string      `json:"cancelPolicy"         orm:"cancel_policy"           description:"取消政策 多语言"`
	GalleryImages        string      `json:"galleryImages"        orm:"gallery_images"          description:"画廊图片"`
	GalleryCover         string      `json:"galleryCover"         orm:"gallery_cover"           description:"画廊封面"`
	RequiredBook         string      `json:"requiredBook"         orm:"required_book"           description:""`
	Close                int         `json:"close"                orm:"close"                   description:"1、开启该物业  2、关闭该物业"`
	LeaseClose           int         `json:"leaseClose"           orm:"lease_close"             description:"1、开启物业短租  2、关闭物业短租"`
	BookingClose         int         `json:"bookingClose"         orm:"booking_close"           description:"1、开启预订模式  2、关闭预订模式"`
	Sort                 int         `json:"sort"                 orm:"sort"                    description:"排序(越大越靠前)"`
	CheckInGuide         string      `json:"checkInGuide"         orm:"check_in_guide"          description:"入住指南"`
	SpaCanOrder          int         `json:"spaCanOrder"          orm:"spa_can_order"           description:"是否开放预定按摩服务  1开放  2关闭"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`
	DeletedAt            *gtime.Time `json:"deletedAt"            orm:"deleted_at"              description:"删除时间"`
	BatchReservation     string      `json:"batchReservation"     orm:"batch_reservation"       description:"是否开启多房型预定 （Y 开启 N  关闭）"`
	RegionId             uint        `json:"regionId"             orm:"region_id"               description:"地区"`
	AccessPass           string      `json:"accessPass"           orm:"access_pass"             description:"门禁密码"`
}
