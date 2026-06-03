// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppCancelOrder is the golang structure for table pms_app_cancel_order.
type PmsAppCancelOrder struct {
	Id            int         `json:"id"            orm:"id"              description:"ID"`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"        description:"订单号"`
	OutOrderSn    string      `json:"outOrderSn"    orm:"out_order_sn"    description:"外部订单号"`
	CancelOrderSn string      `json:"cancelOrderSn" orm:"cancel_order_sn" description:"取消订单号"`
	CancelAmount  float64     `json:"cancelAmount"  orm:"cancel_amount"   description:"取消费用"`
	CancelRate    *gjson.Json `json:"cancelRate"    orm:"cancel_rate"     description:"取消政策冗余"`
	CancelAt      *gtime.Time `json:"cancelAt"      orm:"cancel_at"       description:"取消时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"删除时间"`
}
