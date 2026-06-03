// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodActivity is the golang structure for table food_activity.
type FoodActivity struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Date      *gtime.Time `json:"date"      orm:"date"       description:"活动日期"`
	Name      string      `json:"name"      orm:"name"       description:"标题（多语言）"`
	SubName   string      `json:"subName"   orm:"sub_name"   description:"副标题"`
	Pic       string      `json:"pic"       orm:"pic"        description:"图片"`
	Status    uint        `json:"status"    orm:"status"     description:"状态1、启用 2、禁用"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
