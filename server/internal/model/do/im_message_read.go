// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessageRead is the golang structure of table hg_im_message_read for DAO operations like Where/Data.
type ImMessageRead struct {
	g.Meta    `orm:"table:hg_im_message_read, do:true"`
	Id        interface{} //
	MemberId  interface{} //
	ImId      interface{} // 消息ID
	OrderSn   interface{} // 订单号
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
