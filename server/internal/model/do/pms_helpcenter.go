// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenter is the golang structure of table hg_pms_helpcenter for DAO operations like Where/Data.
type PmsHelpcenter struct {
	g.Meta     `orm:"table:hg_pms_helpcenter, do:true"`
	Id         interface{} //
	Language   interface{} //
	CategoryId interface{} // 分类ID
	Title      interface{} // 标题
	Content    interface{} // 内容
	Sort       interface{} // 排序(越大越靠前)
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
