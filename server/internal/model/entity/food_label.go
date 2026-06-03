// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodLabel is the golang structure for table food_label.
type FoodLabel struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Type      uint        `json:"type"      orm:"type"       description:"1 餐厅  2 套餐"`
	Name      string      `json:"name"      orm:"name"       description:""`
	Status    uint        `json:"status"    orm:"status"     description:"状态1、启用 2、禁用"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序(越大越靠前)"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
