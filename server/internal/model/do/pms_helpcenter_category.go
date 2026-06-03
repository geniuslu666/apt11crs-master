// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenterCategory is the golang structure of table hg_pms_helpcenter_category for DAO operations like Where/Data.
type PmsHelpcenterCategory struct {
	g.Meta    `orm:"table:hg_pms_helpcenter_category, do:true"`
	Id        interface{} //
	Name      interface{} // 分类名称
	Language  interface{} //
	Sort      interface{} // 排序(越大越靠前)
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
