// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Rate is the golang structure for table rate.
type Rate struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	KefuName     string      `json:"kefuName"     orm:"kefu_name"     description:"客服账户"`
	KefuNickname string      `json:"kefuNickname" orm:"kefu_nickname" description:"客服昵称"`
	VisitorId    string      `json:"visitorId"    orm:"visitor_id"    description:"访客id"`
	Content      string      `json:"content"      orm:"content"       description:"评价内容"`
	EntId        string      `json:"entId"        orm:"ent_id"        description:"企业ID"`
	Score        uint        `json:"score"        orm:"score"         description:"评价分数"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}
