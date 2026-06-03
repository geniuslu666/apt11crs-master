// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CrsOrder is the golang structure for table crs_order.
type CrsOrder struct {
	Id           int         `json:"id"           orm:"id"            description:"主键"`
	OrderSn      string      `json:"orderSn"      orm:"order_sn"      description:"订单号"`
	OrderAmount  float64     `json:"orderAmount"  orm:"order_amount"  description:"订单金额"`
	OutOrderSn   string      `json:"outOrderSn"   orm:"out_order_sn"  description:"外部订单号"`
	OrderStatus  string      `json:"orderStatus"  orm:"order_status"  description:"订单状态"`
	RefundStatus string      `json:"refundStatus" orm:"refund_status" description:"退款状态"`
	RefundAmount string      `json:"refundAmount" orm:"refund_amount" description:"退款金额"`
	ExpiredTime  int         `json:"expiredTime"  orm:"expired_time"  description:"过期时间"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"修改时间"`
}
