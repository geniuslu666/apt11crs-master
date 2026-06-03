// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleThumb is the golang structure for table homepage_article_thumb.
type HomepageArticleThumb struct {
	Id        uint64      `json:"id"        orm:"id"         description:"关联ID"`
	ArticleId uint64      `json:"articleId" orm:"article_id" description:"活动ID"`
	MemberId  uint64      `json:"memberId"  orm:"member_id"  description:"会员ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
