// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"标签名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Kefu      string      `json:"kefu"      orm:"kefu"       description:"客服账户"`
	EntId     int         `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
}
