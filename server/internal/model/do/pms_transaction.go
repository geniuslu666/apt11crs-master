// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransaction is the golang structure of table hg_pms_transaction for DAO operations like Where/Data.
type PmsTransaction struct {
	g.Meta           `orm:"table:hg_pms_transaction, do:true"`
	Id               interface{} // 主键
	OrderSn          interface{} // 订单号
	ChangeOrderSn    interface{} // 变更订单号
	OrderType        interface{} //
	TransactionSn    interface{} // 支付流水号
	PaymentRequestId interface{} // 第三方支付流水号
	CaptureId        interface{} // paypal capture_id
	Scene            interface{} // 场景值
	PayChannel       interface{} // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
	PayType          interface{} // 支付方式   BAL 余额
	OpenId           interface{} // 用户ID
	Amount           interface{} // 总金额
	PayParams        interface{} // 支付参数
	PriceCurrency    interface{} // 币种
	PayAmount        interface{} // 支付金额
	PayCharge        interface{} // 支付税率
	PayStatus        interface{} // 支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付
	PayTime          *gtime.Time // 支付时间
	ExpiredTime      *gtime.Time // 过期时间
	RefundAmount     interface{} // 退款金额
	RefundStatus     interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	ScenePayRate     interface{} // 场景积分抵扣比例
	Level            interface{} // 等级
	ExchangeRate     interface{} // 积分汇率
	CouponId         interface{} // 优惠券ID
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
	Remark           interface{} // 备注
	IsFx             interface{} // 是否是分销订单
}
