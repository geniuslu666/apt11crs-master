// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberCancel is the golang structure of table hg_pms_member_cancel for DAO operations like Where/Data.
type PmsMemberCancel struct {
	g.Meta       `orm:"table:hg_pms_member_cancel, do:true"`
	Id           interface{} // 主键
	MemberId     interface{} // 用户ID
	MemberNo     interface{} // 会员号
	Phone        interface{} // 手机号
	PhoneArea    interface{} // 手机区号
	Mail         interface{} // 邮箱
	CancelReason interface{} // 注销原因
	OperatorId   interface{} // 操作员ID
	AuditStatus  interface{} // 审核状态 1-待审核 2-审核通过 3-审核拒绝
	AuditReason  interface{} // 审核通过/拒绝原因
	AuditTime    *gtime.Time // 审核时间
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
