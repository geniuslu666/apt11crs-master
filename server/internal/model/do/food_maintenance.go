// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodMaintenance is the golang structure of table hg_food_maintenance for DAO operations like Where/Data.
type FoodMaintenance struct {
	g.Meta    `orm:"table:hg_food_maintenance, do:true"`
	Id        interface{} //
	Language  interface{} // 语言
	Content   interface{} // 内容
	IsDefault interface{} // 是否默认  1是 2否
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
}
