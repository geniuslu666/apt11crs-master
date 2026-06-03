// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodCuisine is the golang structure for table food_cuisine.
type FoodCuisine struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	CuisineName string      `json:"cuisineName" orm:"cuisine_name" description:""`
	Status      uint        `json:"status"      orm:"status"       description:"状态1、启用 2、禁用"`
	Pic         string      `json:"pic"         orm:"pic"          description:"图片"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序(越大越靠前)"`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    description:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
