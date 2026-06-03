// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMch is the golang structure of table hg_th_mch for DAO operations like Where/Data.
type ThMch struct {
	g.Meta      `orm:"table:hg_th_mch, do:true"`
	Id          interface{} //
	CategoryId  interface{} // 分类ID
	Name        interface{} // 名称
	Logo        interface{} // LOGO
	ContactInfo interface{} // 联系信息
	Sort        interface{} // 排序(越大越靠前)
	Status      interface{} // 1、启用 2、禁用
	StoreOnNum  interface{} // 启用中门店数量
	StoreOffNum interface{} // 禁用中门店数量
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
