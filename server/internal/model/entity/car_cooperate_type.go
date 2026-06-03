// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCooperateType is the golang structure for table car_cooperate_type.
type CarCooperateType struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	TypeName  string      `json:"typeName"  orm:"type_name"  description:""`
	Status    uint        `json:"status"    orm:"status"     description:"状态1、启用 2、禁用"`
	IsThird   uint        `json:"isThird"   orm:"is_third"   description:"是否第三方 1-是 2-否"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
