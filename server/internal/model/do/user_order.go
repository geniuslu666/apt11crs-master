// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserOrder is the golang structure of table user_order for DAO operations like Where/Data.
type UserOrder struct {
	g.Meta        `orm:"table:user_order, do:true"`
	Id            interface{} //
	OrderSn       interface{} // 订单号
	Money         interface{} // 金额
	NewExpireTime interface{} // 新的到期时间
	Payment       interface{} // alipay支付宝，wechat微信，bank网银
	Type          interface{} // 1未支付，2已支付，3已取消
	Comment       interface{} // 备注
	CreatedAt     *gtime.Time // 创建时间
	Operator      interface{} // 操作人
	KefuName      interface{} // 客服账户
	EntId         interface{} // 客服企业ID
}
