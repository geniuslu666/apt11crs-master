// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSettlementAccount is the golang structure of table hg_food_settlement_account for DAO operations like Where/Data.
type FoodSettlementAccount struct {
	g.Meta        `orm:"table:hg_food_settlement_account, do:true"`
	Id            interface{} //
	RestaurantId  interface{} // 餐厅ID
	Type          interface{} // 类型
	BankName      interface{} // 开户行
	BankUser      interface{} // 开户人姓名
	BankCard      interface{} // 银行账号
	WechatAccount interface{} // 微信名
	AlipayName    interface{} // 支付宝真实姓名
	AlipayAccount interface{} // 支付宝账户
	Status        interface{} // 1、启用 2、禁用
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time //
}
