// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMessageRead is the golang structure of table hg_system_message_read for DAO operations like Where/Data.
type SystemMessageRead struct {
	g.Meta    `orm:"table:hg_system_message_read, do:true"`
	Id        interface{} // 记录ID
	MessageId interface{} // 消息ID
	MemberId  interface{} // 会员ID
	CreatedAt *gtime.Time // 已读时间
}
