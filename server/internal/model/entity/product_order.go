// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductOrder is the golang structure for table product_order.
type ProductOrder struct {
	Id              int         `json:"id"              orm:"id"               description:""`
	EntId           string      `json:"entId"           orm:"ent_id"           description:"企业ID"`
	KefuName        string      `json:"kefuName"        orm:"kefu_name"        description:"客服账户"`
	UserId          string      `json:"userId"          orm:"user_id"          description:"用户ID"`
	OrderSn         string      `json:"orderSn"         orm:"order_sn"         description:"订单编号"`
	OrderDesc       string      `json:"orderDesc"       orm:"order_desc"       description:"订单描述"`
	OrderStatus     string      `json:"orderStatus"     orm:"order_status"     description:"订单状态：pending,processing,completed,cancelled"`
	TotalAmount     int         `json:"totalAmount"     orm:"total_amount"     description:"订单金额"`
	PaymentMethod   string      `json:"paymentMethod"   orm:"payment_method"   description:"支付方式：wechat,alipay,bank,other"`
	PaymentStatus   string      `json:"paymentStatus"   orm:"payment_status"   description:"支付状态：paid,unpaid,refunded"`
	ShippingAddress string      `json:"shippingAddress" orm:"shipping_address" description:"收货地址"`
	Email           string      `json:"email"           orm:"email"            description:"邮箱"`
	Contact         string      `json:"contact"         orm:"contact"          description:"联系人"`
	Tel             string      `json:"tel"             orm:"tel"              description:"手机号"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`
}
