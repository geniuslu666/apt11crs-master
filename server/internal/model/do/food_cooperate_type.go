// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodCooperateType is the golang structure of table hg_food_cooperate_type for DAO operations like Where/Data.
type FoodCooperateType struct {
	g.Meta    `orm:"table:hg_food_cooperate_type, do:true"`
	Id        interface{} //
	TypeName  interface{} //
	Status    interface{} // 状态1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
