// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurantNotice is the golang structure of table hg_food_restaurant_notice for DAO operations like Where/Data.
type FoodRestaurantNotice struct {
	g.Meta          `orm:"table:hg_food_restaurant_notice, do:true"`
	Id              interface{} //
	RestaurantId    interface{} // 餐厅ID
	Title           interface{} // 通知标题
	Content         interface{} // 公告内容
	Sort            interface{} // 排序(越大越靠前)
	Status          interface{} // 状态1、启用 2、禁用
	NeedUserConfirm interface{} // 是否需要用户确认 1、需要  2、不需要
	CreateAt        *gtime.Time // 创建时间
	UpdateAt        *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
