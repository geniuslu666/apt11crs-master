// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WeworkSyncMsg is the golang structure for table wework_sync_msg.
type WeworkSyncMsg struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	SyncCursor string      `json:"syncCursor" orm:"sync_cursor" description:""`
	JsonTxt    string      `json:"jsonTxt"    orm:"json_txt"    description:""`
	VisitorId  string      `json:"visitorId"  orm:"visitor_id"  description:""`
	KefuId     string      `json:"kefuId"     orm:"kefu_id"     description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	EntId      int         `json:"entId"      orm:"ent_id"      description:""`
}
