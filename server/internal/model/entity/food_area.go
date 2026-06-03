// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodArea is the golang structure for table food_area.
type FoodArea struct {
	Id         int64       `json:"id"         orm:"id"          description:""`
	Pid        int64       `json:"pid"        orm:"pid"         description:"上级ID"`
	Level      int         `json:"level"      orm:"level"       description:"区域级别"`
	Tree       string      `json:"tree"       orm:"tree"        description:"关系树"`
	AreaName   string      `json:"areaName"   orm:"area_name"   description:"区域名称"`
	AreaStatus uint        `json:"areaStatus" orm:"area_status" description:"1、启用 2、禁用"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:"更新时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:""`
}
