// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarAddressType is the golang structure of table hg_car_address_type for DAO operations like Where/Data.
type CarAddressType struct {
	g.Meta    `orm:"table:hg_car_address_type, do:true"`
	Id        interface{} //
	TypeName  interface{} //
	Sort      interface{} // 排序 越大越靠前
	Status    interface{} // 状态1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
