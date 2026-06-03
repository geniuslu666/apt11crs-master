// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAirhostMessage is the golang structure for table pms_airhost_message.
type PmsAirhostMessage struct {
	Id               int         `json:"id"               orm:"id"                 description:""`
	ObjectType       string      `json:"objectType"       orm:"object_type"        description:"消息来源类型"`
	MessageId        string      `json:"messageId"        orm:"message_id"         description:"消息 ID"`
	MessageEventCode string      `json:"messageEventCode" orm:"message_event_code" description:"消息类型"`
	MessageContent   string      `json:"messageContent"   orm:"message_content"    description:"消息内容"`
	IsHandle         int         `json:"isHandle"         orm:"is_handle"          description:"1、已处理  2、未处理"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:""`
}
