// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CabinetOrderLog is the golang structure for table cabinet_order_log.
type CabinetOrderLog struct {
	Id          int64       `json:"id"          orm:"id"           description:"变动ID"`
	OrderId     int         `json:"orderId"     orm:"order_id"     description:"订单ID"`
	OrderStatus string      `json:"orderStatus" orm:"order_status" description:"订单状态"`
	ActionWay   string      `json:"actionWay"   orm:"action_way"   description:"操作名"`
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`
	Images      string      `json:"images"      orm:"images"       description:"图集"`
	OperateType string      `json:"operateType" orm:"operate_type" description:"操作员类型"`
	OperateId   int         `json:"operateId"   orm:"operate_id"   description:"操作员ID"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
}
