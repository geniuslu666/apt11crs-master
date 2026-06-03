// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelOrderLog is the golang structure of table hg_travel_order_log for DAO operations like Where/Data.
type TravelOrderLog struct {
	g.Meta      `orm:"table:hg_travel_order_log, do:true"`
	Id          interface{} // 变动ID
	OrderId     interface{} // 订单ID
	OrderStatus interface{} // 订单状态
	ActionWay   interface{} // 操作名
	Remark      interface{} // 备注
	OperateType interface{} // 操作员类型
	OperateId   interface{} // 操作员ID
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
