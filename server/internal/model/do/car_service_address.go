// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CarServiceAddress is the golang structure of table hg_car_service_address for DAO operations like Where/Data.
type CarServiceAddress struct {
	g.Meta    `orm:"table:hg_car_service_address, do:true"`
	ServiceId interface{} // 服务ID
	AddressId interface{} // 地址ID
	Type      interface{} // 1 出发地  2 目的地
}
