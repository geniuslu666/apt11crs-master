// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorBlack is the golang structure for table visitor_black.
type VisitorBlack struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	VisitorId string      `json:"visitorId" orm:"visitor_id" description:"访客ID"`
	Name      string      `json:"name"      orm:"name"       description:"访客名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	EntId     string      `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
	KefuName  string      `json:"kefuName"  orm:"kefu_name"  description:"客服账户"`
}
