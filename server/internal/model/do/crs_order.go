// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CrsOrder is the golang structure of table hg_crs_order for DAO operations like Where/Data.
type CrsOrder struct {
	g.Meta       `orm:"table:hg_crs_order, do:true"`
	Id           interface{} // 主键
	OrderSn      interface{} // 订单号
	OrderAmount  interface{} // 订单金额
	OutOrderSn   interface{} // 外部订单号
	OrderStatus  interface{} // 订单状态
	RefundStatus interface{} // 退款状态
	RefundAmount interface{} // 退款金额
	ExpiredTime  interface{} // 过期时间
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 修改时间
}
