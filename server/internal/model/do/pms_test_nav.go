// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTestNav is the golang structure of table hg_pms_test_nav for DAO operations like Where/Data.
type PmsTestNav struct {
	g.Meta    `orm:"table:hg_pms_test_nav, do:true"`
	Id        interface{} //
	Name      interface{} // 名称
	Image     interface{} // 图标
	AppLink   interface{} // app跳转链接
	WxLink    interface{} // 微信跳转链接
	Sort      interface{} // 排序(越大越靠前)
	Status    interface{} // 状态1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time //
}
