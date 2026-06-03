// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AigcSessionCollect is the golang structure for table aigc_session_collect.
type AigcSessionCollect struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Title     string      `json:"title"     orm:"title"      description:"集合标题"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	KefuName  string      `json:"kefuName"  orm:"kefu_name"  description:"客服名称"`
	EntId     string      `json:"entId"     orm:"ent_id"     description:"企业ID"`
}
