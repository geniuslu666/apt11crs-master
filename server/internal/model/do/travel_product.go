// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelProduct is the golang structure of table hg_travel_product for DAO operations like Where/Data.
type TravelProduct struct {
	g.Meta          `orm:"table:hg_travel_product, do:true"`
	Id              interface{} //
	Title           interface{} // 标题（默认语言；多语言存 hg_pms_language）
	SubTitle        interface{} // 副标题（默认语言；多语言存 hg_pms_language）
	ContentZh       interface{} // 内容（中文 编辑器）
	ContentEn       interface{} // 内容（英文 编辑器）
	ContentKo       interface{} // 内容（韩文 编辑器）
	ContentJa       interface{} // 内容（日文 编辑器）
	ContentTw       interface{} // 内容（繁文 编辑器）
	ListImage       interface{} // 列表图（单图 URL）
	CarouselImages  interface{} // 轮播图（多图，JSON）
	DailyCapacity   interface{} // 每日最大接待人数
	Stock           interface{} // 总库存
	Price           interface{} // 售价（元）
	ContactMobile   interface{} // 联系电话
	MeetingPlace    interface{} // 集合地点
	MeetingTime     interface{} // 集合时间（格式：HH:MM）
	GgLat           interface{} // 谷歌纬度
	GgLng           interface{} // 谷歌经度
	MaxBookDays     interface{} // 最大可预约天数
	AdvanceBookDays interface{} // 至少提前预订天数
	TripPlanningZh  interface{} // 行程规划（中文 编辑器）
	TripPlanningEn  interface{} // 行程规划（英文 编辑器）
	TripPlanningJa  interface{} // 行程规划（日文 编辑器）
	TripPlanningKo  interface{} // 行程规划（韩文 编辑器）
	TripPlanningTw  interface{} // 行程规划（繁文 编辑器）
	BookingNotesZh  interface{} // 预约须知（中文 编辑器）
	BookingNotesEn  interface{} // 预约须知（英文 编辑器）
	BookingNotesJa  interface{} // 预约须知（日文 编辑器）
	BookingNotesKo  interface{} // 预约须知（韩文 编辑器）
	BookingNotesTw  interface{} // 预约须知（繁文 编辑器）
	Status          interface{} // 状态（1启用 2禁用）
	Sort            interface{} // 排序（越大越靠前）
	SalesNum        interface{} // 已售
	DeletedAt       *gtime.Time // 软删除时间（NULL=正常）
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
