// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodActivity is the golang structure of table hg_food_activity for DAO operations like Where/Data.
type FoodActivity struct {
	g.Meta    `orm:"table:hg_food_activity, do:true"`
	Id        interface{} //
	Date      *gtime.Time // 活动日期
	Name      interface{} // 标题（多语言）
	SubName   interface{} // 副标题
	Pic       interface{} // 图片
	Status    interface{} // 状态1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
