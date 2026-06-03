// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsLanguage is the golang structure for table pms_language.
type PmsLanguage struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Uuid     string      `json:"uuid"     orm:"uuid"      description:"标签ID"`
	Tag      string      `json:"tag"      orm:"tag"       description:"标签  type = table  、 数据库表名     type = 其他的话  是前端的属性名"`
	Type     string      `json:"type"     orm:"type"      description:"类型  table/数据库表  manage/管理端  mobile/移动端"`
	Key      string      `json:"key"      orm:"key"       description:"字段标识"`
	Language string      `json:"language" orm:"language"  description:"语言"`
	Content  string      `json:"content"  orm:"content"   description:"语言内容"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"更新时间"`
}
