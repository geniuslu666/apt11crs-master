// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransactionRefund is the golang structure for table pms_transaction_refund.
type PmsTransactionRefund struct {
	Id            int         `json:"id"            orm:"id"              description:"主键"`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"        description:"订单号"`
	ChangeOrderSn string      `json:"changeOrderSn" orm:"change_order_sn" description:"变更订单号"`
	TransactionSn string      `json:"transactionSn" orm:"transaction_sn"  description:"支付流水号"`
	Scene         string      `json:"scene"         orm:"scene"           description:"场景值"`
	RefundChannel string      `json:"refundChannel" orm:"refund_channel"  description:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
	RefundType    string      `json:"refundType"    orm:"refund_type"     description:"'支付方式   BAL 余额'"`
	RefundSn      string      `json:"refundSn"      orm:"refund_sn"       description:"退款流水号"`
	TransNo       string      `json:"transNo"       orm:"trans_no"        description:"退款交易号"`
	CancelOrderSn string      `json:"cancelOrderSn" orm:"cancel_order_sn" description:"取消订单号"`
	RefundAmount  float64     `json:"refundAmount"  orm:"refund_amount"   description:"退款金额"`
	RefundTime    *gtime.Time `json:"refundTime"    orm:"refund_time"     description:"退款时间"`
	RefundStatus  string      `json:"refundStatus"  orm:"refund_status"   description:"退款状态"`
	CancelId      string      `json:"cancelId"      orm:"cancel_id"       description:"取消政策ID"`
	OperateType   string      `json:"operateType"   orm:"operate_type"    description:"操作员类型"`
	OperateId     int         `json:"operateId"     orm:"operate_id"      description:"操作员ID"`
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`
	IsFx          string      `json:"isFx"          orm:"is_fx"           description:"是否是分销订单"`
}
