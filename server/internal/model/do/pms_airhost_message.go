// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAirhostMessage is the golang structure of table hg_pms_airhost_message for DAO operations like Where/Data.
type PmsAirhostMessage struct {
	g.Meta           `orm:"table:hg_pms_airhost_message, do:true"`
	Id               interface{} //
	ObjectType       interface{} // 消息来源类型
	MessageId        interface{} // 消息 ID
	MessageEventCode interface{} // 消息类型
	MessageContent   interface{} // 消息内容
	IsHandle         interface{} // 1、已处理  2、未处理
	CreatedAt        *gtime.Time //
	DeletedAt        *gtime.Time //
}
