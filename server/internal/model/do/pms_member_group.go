// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberGroup is the golang structure of table hg_pms_member_group for DAO operations like Where/Data.
type PmsMemberGroup struct {
	g.Meta      `orm:"table:hg_pms_member_group, do:true"`
	Id          interface{} // 会员分组ID
	MemberGroup interface{} // 会员分组名称
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
