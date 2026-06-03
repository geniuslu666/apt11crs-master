// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleCate is the golang structure of table article_cate for DAO operations like Where/Data.
type ArticleCate struct {
	g.Meta  `orm:"table:article_cate, do:true"`
	Id      interface{} //
	CatName interface{} // 自动回复分类名称
	UserId  interface{} // 客服账户
	EntId   interface{} // 客服企业ID
	IsTop   interface{} // 是否置顶展示，1置顶
}
