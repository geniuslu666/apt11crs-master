// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleThumb is the golang structure of table hg_homepage_article_thumb for DAO operations like Where/Data.
type HomepageArticleThumb struct {
	g.Meta    `orm:"table:hg_homepage_article_thumb, do:true"`
	Id        interface{} // 关联ID
	ArticleId interface{} // 活动ID
	MemberId  interface{} // 会员ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
