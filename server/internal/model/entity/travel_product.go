// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelProduct is the golang structure for table travel_product.
type TravelProduct struct {
	Id              uint64      `json:"id"              orm:"id"                description:""`
	Title           string      `json:"title"           orm:"title"             description:"标题（默认语言；多语言存 hg_pms_language）"`
	SubTitle        string      `json:"subTitle"        orm:"sub_title"         description:"副标题（默认语言；多语言存 hg_pms_language）"`
	ContentZh       string      `json:"contentZh"       orm:"content_zh"        description:"内容（中文 编辑器）"`
	ContentEn       string      `json:"contentEn"       orm:"content_en"        description:"内容（英文 编辑器）"`
	ContentKo       string      `json:"contentKo"       orm:"content_ko"        description:"内容（韩文 编辑器）"`
	ContentJa       string      `json:"contentJa"       orm:"content_ja"        description:"内容（日文 编辑器）"`
	ContentTw       string      `json:"contentTw"       orm:"content_tw"        description:"内容（繁文 编辑器）"`
	ListImage       string      `json:"listImage"       orm:"list_image"        description:"列表图（单图 URL）"`
	CarouselImages  string      `json:"carouselImages"  orm:"carousel_images"   description:"轮播图（多图，JSON）"`
	DailyCapacity   uint        `json:"dailyCapacity"   orm:"daily_capacity"    description:"每日最大接待人数"`
	Stock           uint        `json:"stock"           orm:"stock"             description:"总库存"`
	Price           float64     `json:"price"           orm:"price"             description:"售价（元）"`
	ContactMobile   string      `json:"contactMobile"   orm:"contact_mobile"    description:"联系电话"`
	MeetingPlace    string      `json:"meetingPlace"    orm:"meeting_place"     description:"集合地点"`
	MeetingTime     string      `json:"meetingTime"     orm:"meeting_time"      description:"集合时间（格式：HH:MM）"`
	GgLat           string      `json:"ggLat"           orm:"gg_lat"            description:"谷歌纬度"`
	GgLng           string      `json:"ggLng"           orm:"gg_lng"            description:"谷歌经度"`
	MaxBookDays     uint        `json:"maxBookDays"     orm:"max_book_days"     description:"最大可预约天数"`
	AdvanceBookDays uint        `json:"advanceBookDays" orm:"advance_book_days" description:"至少提前预订天数"`
	TripPlanningZh  string      `json:"tripPlanningZh"  orm:"trip_planning_zh"  description:"行程规划（中文 编辑器）"`
	TripPlanningEn  string      `json:"tripPlanningEn"  orm:"trip_planning_en"  description:"行程规划（英文 编辑器）"`
	TripPlanningJa  string      `json:"tripPlanningJa"  orm:"trip_planning_ja"  description:"行程规划（日文 编辑器）"`
	TripPlanningKo  string      `json:"tripPlanningKo"  orm:"trip_planning_ko"  description:"行程规划（韩文 编辑器）"`
	TripPlanningTw  string      `json:"tripPlanningTw"  orm:"trip_planning_tw"  description:"行程规划（繁文 编辑器）"`
	BookingNotesZh  string      `json:"bookingNotesZh"  orm:"booking_notes_zh"  description:"预约须知（中文 编辑器）"`
	BookingNotesEn  string      `json:"bookingNotesEn"  orm:"booking_notes_en"  description:"预约须知（英文 编辑器）"`
	BookingNotesJa  string      `json:"bookingNotesJa"  orm:"booking_notes_ja"  description:"预约须知（日文 编辑器）"`
	BookingNotesKo  string      `json:"bookingNotesKo"  orm:"booking_notes_ko"  description:"预约须知（韩文 编辑器）"`
	BookingNotesTw  string      `json:"bookingNotesTw"  orm:"booking_notes_tw"  description:"预约须知（繁文 编辑器）"`
	Status          int         `json:"status"          orm:"status"            description:"状态（1启用 2禁用）"`
	Sort            int         `json:"sort"            orm:"sort"              description:"排序（越大越靠前）"`
	SalesNum        uint        `json:"salesNum"        orm:"sales_num"         description:"已售"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"软删除时间（NULL=正常）"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"更新时间"`
}
