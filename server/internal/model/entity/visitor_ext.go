// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitorExt is the golang structure for table visitor_ext.
type VisitorExt struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	VisitorId string      `json:"visitorId" orm:"visitor_id" description:"访客ID"`
	EntId     uint        `json:"entId"     orm:"ent_id"     description:"对接的企业ID"`
	Ua        string      `json:"ua"        orm:"ua"         description:"访客浏览器UserAgent"`
	Title     string      `json:"title"     orm:"title"      description:"页面标题"`
	Url       string      `json:"url"       orm:"url"        description:"页面地址"`
	Refer     string      `json:"refer"     orm:"refer"      description:"页面来源"`
	ReferUrl  string      `json:"referUrl"  orm:"refer_url"  description:"页面来源地址"`
	ClientIp  string      `json:"clientIp"  orm:"client_ip"  description:"访客IP"`
	City      string      `json:"city"      orm:"city"       description:"访客城市"`
	Language  string      `json:"language"  orm:"language"   description:"浏览器语言"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
