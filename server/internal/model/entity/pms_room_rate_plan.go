// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomRatePlan is the golang structure for table pms_room_rate_plan.
type PmsRoomRatePlan struct {
	Id         int64       `json:"id"         orm:"id"           description:""`
	Tuid       string      `json:"tuid"       orm:"tuid"         description:"房型ID"`
	TName      string      `json:"tName"      orm:"t_name"       description:"房型名"`
	RateName   string      `json:"rateName"   orm:"rate_name"    description:"费率名称"`
	RatePlanId string      `json:"ratePlanId" orm:"rate_plan_id" description:"价格计划"`
	IsUse      string      `json:"isUse"      orm:"is_use"       description:"是否使用"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"更新时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"   description:"删除时间"`
}
