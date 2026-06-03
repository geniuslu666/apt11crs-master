// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSettlementAccount is the golang structure for table food_settlement_account.
type FoodSettlementAccount struct {
	Id            int64       `json:"id"            orm:"id"             description:""`
	RestaurantId  int         `json:"restaurantId"  orm:"restaurant_id"  description:"餐厅ID"`
	Type          string      `json:"type"          orm:"type"           description:"类型"`
	BankName      string      `json:"bankName"      orm:"bank_name"      description:"开户行"`
	BankUser      string      `json:"bankUser"      orm:"bank_user"      description:"开户人姓名"`
	BankCard      string      `json:"bankCard"      orm:"bank_card"      description:"银行账号"`
	WechatAccount string      `json:"wechatAccount" orm:"wechat_account" description:"微信名"`
	AlipayName    string      `json:"alipayName"    orm:"alipay_name"    description:"支付宝真实姓名"`
	AlipayAccount string      `json:"alipayAccount" orm:"alipay_account" description:"支付宝账户"`
	Status        uint        `json:"status"        orm:"status"         description:"1、启用 2、禁用"`
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:"创建时间"`
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:""`
}
