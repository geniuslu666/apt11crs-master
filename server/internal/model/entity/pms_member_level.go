// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLevel is the golang structure for table pms_member_level.
type PmsMemberLevel struct {
	Id             int         `json:"id"             orm:"id"               description:""`
	LevelName      string      `json:"levelName"      orm:"level_name"       description:"会员等级名称"`
	Exp            int         `json:"exp"            orm:"exp"              description:"达到等级所需经验值"`
	HotelGetRate   float64     `json:"hotelGetRate"   orm:"hotel_get_rate"   description:"酒店场景获取积分倍率"`
	FoodGetRate    float64     `json:"foodGetRate"    orm:"food_get_rate"    description:"餐饮场景获取积分倍率"`
	SpaGetRate     float64     `json:"spaGetRate"     orm:"spa_get_rate"     description:"按摩场景获取积分倍率"`
	CarGetRate     float64     `json:"carGetRate"     orm:"car_get_rate"     description:"接送机/包车场景获取积分倍率"`
	CabinetGetRate float64     `json:"cabinetGetRate" orm:"cabinet_get_rate" description:"储物柜场景获取积分倍率"`
	TravelGetRate  float64     `json:"travelGetRate"  orm:"travel_get_rate"  description:"一日游场景获取积分倍率"`
	Desc           string      `json:"desc"           orm:"desc"             description:"等级说明"`
	WordColor      string      `json:"wordColor"      orm:"word_color"       description:"等级字体颜色"`
	LevelBadge     string      `json:"levelBadge"     orm:"level_badge"      description:"等级徽章（单图）"`
	LevelCard      string      `json:"levelCard"      orm:"level_card"       description:"等级卡片（单图）"`
	LevelBigPic    string      `json:"levelBigPic"    orm:"level_big_pic"    description:"等级大图（单图）"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`
}
