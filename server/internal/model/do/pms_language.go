// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsLanguage is the golang structure of table hg_pms_language for DAO operations like Where/Data.
type PmsLanguage struct {
	g.Meta   `orm:"table:hg_pms_language, do:true"`
	Id       interface{} //
	Uuid     interface{} // 标签ID
	Tag      interface{} // 标签  type = table  、 数据库表名     type = 其他的话  是前端的属性名
	Type     interface{} // 类型  table/数据库表  manage/管理端  mobile/移动端
	Key      interface{} // 字段标识
	Language interface{} // 语言
	Content  interface{} // 语言内容
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
