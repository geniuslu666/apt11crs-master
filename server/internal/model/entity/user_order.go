// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserOrder is the golang structure for table user_order.
type UserOrder struct {
	Id            int         `json:"id"            orm:"id"              description:""`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"        description:"订单号"`
	Money         uint        `json:"money"         orm:"money"           description:"金额"`
	NewExpireTime string      `json:"newExpireTime" orm:"new_expire_time" description:"新的到期时间"`
	Payment       string      `json:"payment"       orm:"payment"         description:"alipay支付宝，wechat微信，bank网银"`
	Type          int         `json:"type"          orm:"type"            description:"1未支付，2已支付，3已取消"`
	Comment       string      `json:"comment"       orm:"comment"         description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
	Operator      string      `json:"operator"      orm:"operator"        description:"操作人"`
	KefuName      string      `json:"kefuName"      orm:"kefu_name"       description:"客服账户"`
	EntId         string      `json:"entId"         orm:"ent_id"          description:"客服企业ID"`
}
