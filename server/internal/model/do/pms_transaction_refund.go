// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransactionRefund is the golang structure of table hg_pms_transaction_refund for DAO operations like Where/Data.
type PmsTransactionRefund struct {
	g.Meta        `orm:"table:hg_pms_transaction_refund, do:true"`
	Id            interface{} // 主键
	OrderSn       interface{} // 订单号
	ChangeOrderSn interface{} // 变更订单号
	TransactionSn interface{} // 支付流水号
	Scene         interface{} // 场景值
	RefundChannel interface{} // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
	RefundType    interface{} // '支付方式   BAL 余额'
	RefundSn      interface{} // 退款流水号
	TransNo       interface{} // 退款交易号
	CancelOrderSn interface{} // 取消订单号
	RefundAmount  interface{} // 退款金额
	RefundTime    *gtime.Time // 退款时间
	RefundStatus  interface{} // 退款状态
	CancelId      interface{} // 取消政策ID
	OperateType   interface{} // 操作员类型
	OperateId     interface{} // 操作员ID
	Remark        interface{} // 备注
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	IsFx          interface{} // 是否是分销订单
}
