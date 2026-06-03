// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurantNotice is the golang structure for table food_restaurant_notice.
type FoodRestaurantNotice struct {
	Id              int         `json:"id"              orm:"id"                description:""`
	RestaurantId    int         `json:"restaurantId"    orm:"restaurant_id"     description:"餐厅ID"`
	Title           string      `json:"title"           orm:"title"             description:"通知标题"`
	Content         string      `json:"content"         orm:"content"           description:"公告内容"`
	Sort            int         `json:"sort"            orm:"sort"              description:"排序(越大越靠前)"`
	Status          uint        `json:"status"          orm:"status"            description:"状态1、启用 2、禁用"`
	NeedUserConfirm uint        `json:"needUserConfirm" orm:"need_user_confirm" description:"是否需要用户确认 1、需要  2、不需要"`
	CreateAt        *gtime.Time `json:"createAt"        orm:"create_at"         description:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        orm:"update_at"         description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:"删除时间"`
}
