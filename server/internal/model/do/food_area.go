// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodArea is the golang structure of table hg_food_area for DAO operations like Where/Data.
type FoodArea struct {
	g.Meta     `orm:"table:hg_food_area, do:true"`
	Id         interface{} //
	Pid        interface{} // 上级ID
	Level      interface{} // 区域级别
	Tree       interface{} // 关系树
	AreaName   interface{} // 区域名称
	AreaStatus interface{} // 1、启用 2、禁用
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 更新时间
	DeletedAt  *gtime.Time //
}
