// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurantOperateLog is the golang structure of table hg_food_restaurant_operate_log for DAO operations like Where/Data.
type FoodRestaurantOperateLog struct {
	g.Meta          `orm:"table:hg_food_restaurant_operate_log, do:true"`
	Id              interface{} //
	RestaurantId    interface{} // 餐厅ID
	OperateType     interface{} // 操作类型
	OperateMemberId interface{} // 后台操作人ID
	Remark          interface{} // 备注
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 修改时间
}
