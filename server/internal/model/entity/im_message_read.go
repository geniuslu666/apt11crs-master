// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessageRead is the golang structure for table im_message_read.
type ImMessageRead struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	MemberId  int         `json:"memberId"  orm:"member_id"  description:""`
	ImId      int         `json:"imId"      orm:"im_id"      description:"消息ID"`
	OrderSn   string      `json:"orderSn"   orm:"order_sn"   description:"订单号"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
