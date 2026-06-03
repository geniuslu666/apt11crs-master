// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppCancelOrder is the golang structure of table hg_pms_app_cancel_order for DAO operations like Where/Data.
type PmsAppCancelOrder struct {
	g.Meta        `orm:"table:hg_pms_app_cancel_order, do:true"`
	Id            interface{} // ID
	OrderSn       interface{} // 订单号
	OutOrderSn    interface{} // 外部订单号
	CancelOrderSn interface{} // 取消订单号
	CancelAmount  interface{} // 取消费用
	CancelRate    *gjson.Json // 取消政策冗余
	CancelAt      *gtime.Time // 取消时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
