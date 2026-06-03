// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodOrderLog is the golang structure of table hg_food_order_log for DAO operations like Where/Data.
type FoodOrderLog struct {
	g.Meta      `orm:"table:hg_food_order_log, do:true"`
	Id          interface{} // 变动ID
	OrderId     interface{} // 订单ID
	ActionWay   interface{} // 操作名
	Remark      interface{} // 备注
	Images      interface{} // 图集
	OperateType interface{} // 操作员类型
	OperateId   interface{} // 操作员ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
