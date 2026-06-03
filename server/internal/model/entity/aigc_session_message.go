// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AigcSessionMessage is the golang structure for table aigc_session_message.
type AigcSessionMessage struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	CollectId  int         `json:"collectId"  orm:"collect_id"  description:"集合ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	KefuAvatar string      `json:"kefuAvatar" orm:"kefu_avatar" description:"客服头像"`
	AiAvatar   string      `json:"aiAvatar"   orm:"ai_avatar"   description:"AI头像"`
	Content    string      `json:"content"    orm:"content"     description:"内容"`
	KefuName   string      `json:"kefuName"   orm:"kefu_name"   description:"客服名称"`
	EntId      string      `json:"entId"      orm:"ent_id"      description:"企业ID"`
	MsgType    string      `json:"msgType"    orm:"msg_type"    description:"消息类型"`
}
