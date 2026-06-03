// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Qrcode is the golang structure for table qrcode.
type Qrcode struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	EntId     string      `json:"entId"     orm:"ent_id"     description:"客服企业ID"`
	KefuName  string      `json:"kefuName"  orm:"kefu_name"  description:"客服账户"`
	Uuid      string      `json:"uuid"      orm:"uuid"       description:"唯一ID"`
	Url       string      `json:"url"       orm:"url"        description:"跳转的URL"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
