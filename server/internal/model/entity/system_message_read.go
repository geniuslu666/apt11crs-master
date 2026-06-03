// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMessageRead is the golang structure for table system_message_read.
type SystemMessageRead struct {
	Id        uint64      `json:"id"        orm:"id"         description:"记录ID"`
	MessageId uint64      `json:"messageId" orm:"message_id" description:"消息ID"`
	MemberId  uint64      `json:"memberId"  orm:"member_id"  description:"会员ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"已读时间"`
}
