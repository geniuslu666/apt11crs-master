// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarHelp is the golang structure for table car_help.
type CarHelp struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Language string      `json:"language" orm:"language"  description:"语言"`
	Image    string      `json:"image"    orm:"image"     description:"主图"`
	Content  string      `json:"content"  orm:"content"   description:"内容"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"更新时间"`
}
