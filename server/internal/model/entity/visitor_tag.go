// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorTag is the golang structure for table visitor_tag.
type VisitorTag struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	VisitorId string      `json:"visitorId" orm:"visitor_id" description:"访客ID"`
	TagId     int         `json:"tagId"     orm:"tag_id"     description:"标签ID"`
	EntId     int         `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Kefu      string      `json:"kefu"      orm:"kefu"       description:"客服账户"`
}
