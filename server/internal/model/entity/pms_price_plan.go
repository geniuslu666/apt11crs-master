// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsPricePlan is the golang structure for table pms_price_plan.
type PmsPricePlan struct {
	Id              int         `json:"id"              orm:"id"                 description:""`
	PlanName        string      `json:"planName"        orm:"plan_name"          description:"价格plan名称"`
	PlanShowName    string      `json:"planShowName"    orm:"plan_show_name"     description:"展示标签 多语言"`
	PropertyId      string      `json:"propertyId"      orm:"property_id"        description:"物业ID"`
	RoomTypeId      string      `json:"roomTypeId"      orm:"room_type_id"       description:"房型ID"`
	BookingDays     int         `json:"bookingDays"     orm:"booking_days"       description:"预订天数"`
	MemberGroupId   string      `json:"memberGroupId"   orm:"member_group_id"    description:"用户组"`
	MemberLevelId   string      `json:"memberLevelId"   orm:"member_level_id"    description:"会员等级"`
	IsCancel        string      `json:"isCancel"        orm:"is_cancel"          description:"是否可取消   Y  是  N  否"`
	IsOpenPriceMode string      `json:"isOpenPriceMode" orm:"is_open_price_mode" description:"是否开启价格模式"`
	PriceMode       string      `json:"priceMode"       orm:"price_mode"         description:"模式    +  贵   - 便宜"`
	PriceStandard   string      `json:"priceStandard"   orm:"price_standard"     description:"基准  PERCENT 倍率  AMOUNT  金额"`
	PlanValue       int         `json:"planValue"       orm:"plan_value"         description:"价格基准值"`
	PlanTips        string      `json:"planTips"        orm:"plan_tips"          description:"价格计划提示"`
	PricePlanStatus string      `json:"pricePlanStatus" orm:"price_plan_status"  description:"Y 开启  N 关闭"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"         description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"         description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"         description:"删除时间"`
}
