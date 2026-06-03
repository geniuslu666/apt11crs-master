// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta     `orm:"table:article, do:true"`
	Id         interface{} //
	Title      interface{} // 自动回复关键词
	Content    interface{} // 自动回复内容
	CatId      interface{} // 自动回复分类ID
	UserId     interface{} // 客服账户
	EntId      interface{} // 客服企业ID
	ApiUrl     interface{} // 第三方接口地址
	SearchType interface{} // 1包含匹配,2精准匹配
	Score      interface{} // 命中次数
}
