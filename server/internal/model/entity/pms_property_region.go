// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsPropertyRegion is the golang structure for table pms_property_region.
type PmsPropertyRegion struct {
	Id        int         `json:"id"        orm:"id"         description:"主键"`
	Name      string      `json:"name"      orm:"name"       description:"区域名称（多语言）"`
	Status    uint        `json:"status"    orm:"status"     description:"状态1、启用 2、禁用"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
