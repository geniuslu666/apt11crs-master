// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderRefundLog is the golang structure for table order_refund_log.
type OrderRefundLog struct {
	Id           int         `json:"id"           orm:"id"            description:"主键"`
	OrderSn      string      `json:"orderSn"      orm:"order_sn"      description:"订单号"`
	Scene        string      `json:"scene"        orm:"scene"         description:"场景值"`
	RefundType   string      `json:"refundType"   orm:"refund_type"   description:"退款方式 BAL-积分退款 AMOUNT-金额退款"`
	RefundAmount float64     `json:"refundAmount" orm:"refund_amount" description:"退款金额"`
	RefundTime   *gtime.Time `json:"refundTime"   orm:"refund_time"   description:"退款时间"`
	RefundStatus string      `json:"refundStatus" orm:"refund_status" description:"退款状态"`
	OperateType  string      `json:"operateType"  orm:"operate_type"  description:"操作员类型"`
	OperateId    int         `json:"operateId"    orm:"operate_id"    description:"操作员ID"`
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
}
