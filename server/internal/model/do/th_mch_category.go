// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchCategory is the golang structure of table hg_th_mch_category for DAO operations like Where/Data.
type ThMchCategory struct {
	g.Meta    `orm:"table:hg_th_mch_category, do:true"`
	Id        interface{} //
	Name      interface{} // 分类名称
	Sort      interface{} // 排序(越大越靠前)
	Status    interface{} // 1、启用 2、禁用
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
