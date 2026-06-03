// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsNews is the golang structure of table cms_news for DAO operations like Where/Data.
type CmsNews struct {
	g.Meta    `orm:"table:cms_news, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Content   interface{} // 内容
	CatId     interface{} // 分类ID
	CreatedAt *gtime.Time // 创建时间
}
