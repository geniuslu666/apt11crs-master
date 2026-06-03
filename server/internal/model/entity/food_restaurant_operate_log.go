// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurantOperateLog is the golang structure for table food_restaurant_operate_log.
type FoodRestaurantOperateLog struct {
	Id              int64       `json:"id"              orm:"id"                description:""`
	RestaurantId    uint        `json:"restaurantId"    orm:"restaurant_id"     description:"餐厅ID"`
	OperateType     string      `json:"operateType"     orm:"operate_type"      description:"操作类型"`
	OperateMemberId uint        `json:"operateMemberId" orm:"operate_member_id" description:"后台操作人ID"`
	Remark          string      `json:"remark"          orm:"remark"            description:"备注"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:"修改时间"`
}
