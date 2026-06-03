// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Learn is the golang structure for table learn.
type Learn struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Content   string      `json:"content"   orm:"content"    description:"问题内容"`
	Score     int         `json:"score"     orm:"score"      description:"次数"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	EntId     string      `json:"entId"     orm:"ent_id"     description:"企业ID"`
	KefuName  string      `json:"kefuName"  orm:"kefu_name"  description:"kefu名称"`
	Finshed   int         `json:"finshed"   orm:"finshed"    description:"是否解决，1未解决，2已解决"`
}
