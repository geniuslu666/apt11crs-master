// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberGroup is the golang structure for table pms_member_group.
type PmsMemberGroup struct {
	Id          int         `json:"id"          orm:"id"           description:"会员分组ID"`
	MemberGroup string      `json:"memberGroup" orm:"member_group" description:"会员分组名称"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
