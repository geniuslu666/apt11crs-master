// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodLabel is the golang structure of table hg_food_label for DAO operations like Where/Data.
type FoodLabel struct {
	g.Meta    `orm:"table:hg_food_label, do:true"`
	Id        interface{} //
	Type      interface{} // 1 餐厅  2 套餐
	Name      interface{} //
	Status    interface{} // 状态1、启用 2、禁用
	Sort      interface{} // 排序(越大越靠前)
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
