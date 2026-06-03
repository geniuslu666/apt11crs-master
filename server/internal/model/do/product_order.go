// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductOrder is the golang structure of table product_order for DAO operations like Where/Data.
type ProductOrder struct {
	g.Meta          `orm:"table:product_order, do:true"`
	Id              interface{} //
	EntId           interface{} // 企业ID
	KefuName        interface{} // 客服账户
	UserId          interface{} // 用户ID
	OrderSn         interface{} // 订单编号
	OrderDesc       interface{} // 订单描述
	OrderStatus     interface{} // 订单状态：pending,processing,completed,cancelled
	TotalAmount     interface{} // 订单金额
	PaymentMethod   interface{} // 支付方式：wechat,alipay,bank,other
	PaymentStatus   interface{} // 支付状态：paid,unpaid,refunded
	ShippingAddress interface{} // 收货地址
	Email           interface{} // 邮箱
	Contact         interface{} // 联系人
	Tel             interface{} // 手机号
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
