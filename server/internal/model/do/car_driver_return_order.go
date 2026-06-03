// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriverReturnOrder is the golang structure of table hg_car_driver_return_order for DAO operations like Where/Data.
type CarDriverReturnOrder struct {
	g.Meta   `orm:"table:hg_car_driver_return_order, do:true"`
	Id       interface{} //
	OrderId  interface{} // 订单ID
	DriverId interface{} // 司机ID
	CarId    interface{} // 车辆ID
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
