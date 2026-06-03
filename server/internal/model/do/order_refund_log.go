// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderRefundLog is the golang structure of table hg_order_refund_log for DAO operations like Where/Data.
type OrderRefundLog struct {
	g.Meta       `orm:"table:hg_order_refund_log, do:true"`
	Id           interface{} // 主键
	OrderSn      interface{} // 订单号
	Scene        interface{} // 场景值
	RefundType   interface{} // 退款方式 BAL-积分退款 AMOUNT-金额退款
	RefundAmount interface{} // 退款金额
	RefundTime   *gtime.Time // 退款时间
	RefundStatus interface{} // 退款状态
	OperateType  interface{} // 操作员类型
	OperateId    interface{} // 操作员ID
	Remark       interface{} // 备注
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
