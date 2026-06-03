// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodCuisine is the golang structure of table hg_food_cuisine for DAO operations like Where/Data.
type FoodCuisine struct {
	g.Meta      `orm:"table:hg_food_cuisine, do:true"`
	Id          interface{} //
	CuisineName interface{} //
	Status      interface{} // 状态1、启用 2、禁用
	Pic         interface{} // 图片
	Sort        interface{} // 排序(越大越靠前)
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
