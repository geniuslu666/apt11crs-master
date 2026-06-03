// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DouyinWebhook is the golang structure for table douyin_webhook.
type DouyinWebhook struct {
	Id         int         `json:"id"         orm:"id"           description:""`
	KefuName   string      `json:"kefuName"   orm:"kefu_name"    description:"客服账户"`
	Event      string      `json:"event"      orm:"event"        description:"event"`
	FromUserId string      `json:"fromUserId" orm:"from_user_id" description:"from_user_id"`
	ToUserId   string      `json:"toUserId"   orm:"to_user_id"   description:"to_user_id"`
	ClientKey  string      `json:"clientKey"  orm:"client_key"   description:"client_key"`
	Content    string      `json:"content"    orm:"content"      description:"content"`
	EntId      string      `json:"entId"      orm:"ent_id"       description:"企业ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`
}
