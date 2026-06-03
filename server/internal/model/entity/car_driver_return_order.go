// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriverReturnOrder is the golang structure for table car_driver_return_order.
type CarDriverReturnOrder struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	OrderId  uint        `json:"orderId"  orm:"order_id"  description:"订单ID"`
	DriverId int         `json:"driverId" orm:"driver_id" description:"司机ID"`
	CarId    int         `json:"carId"    orm:"car_id"    description:"车辆ID"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"更新时间"`
}
