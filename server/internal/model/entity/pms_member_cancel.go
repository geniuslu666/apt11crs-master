// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberCancel is the golang structure for table pms_member_cancel.
type PmsMemberCancel struct {
	Id           int         `json:"id"           orm:"id"            description:"主键"`
	MemberId     uint        `json:"memberId"     orm:"member_id"     description:"用户ID"`
	MemberNo     string      `json:"memberNo"     orm:"member_no"     description:"会员号"`
	Phone        string      `json:"phone"        orm:"phone"         description:"手机号"`
	PhoneArea    string      `json:"phoneArea"    orm:"phone_area"    description:"手机区号"`
	Mail         string      `json:"mail"         orm:"mail"          description:"邮箱"`
	CancelReason string      `json:"cancelReason" orm:"cancel_reason" description:"注销原因"`
	OperatorId   int         `json:"operatorId"   orm:"operator_id"   description:"操作员ID"`
	AuditStatus  int         `json:"auditStatus"  orm:"audit_status"  description:"审核状态 1-待审核 2-审核通过 3-审核拒绝"`
	AuditReason  string      `json:"auditReason"  orm:"audit_reason"  description:"审核通过/拒绝原因"`
	AuditTime    *gtime.Time `json:"auditTime"    orm:"audit_time"    description:"审核时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
}
