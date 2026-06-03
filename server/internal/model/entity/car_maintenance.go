// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarMaintenance is the golang structure for table car_maintenance.
type CarMaintenance struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Type      uint        `json:"type"      orm:"type"       description:"类型 1-全部 2-接送机 3-包机"`
	Language  string      `json:"language"  orm:"language"   description:"语言"`
	Content   string      `json:"content"   orm:"content"    description:"内容"`
	IsDefault uint        `json:"isDefault" orm:"is_default" description:"是否默认  1是 2否"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
}
