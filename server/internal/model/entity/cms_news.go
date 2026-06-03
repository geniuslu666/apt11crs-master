// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsNews is the golang structure for table cms_news.
type CmsNews struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Title     string      `json:"title"     orm:"title"      description:"标题"`
	Content   string      `json:"content"   orm:"content"    description:"内容"`
	CatId     int         `json:"catId"     orm:"cat_id"     description:"分类ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
