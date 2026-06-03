// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Customer is the golang structure of table customer for DAO operations like Where/Data.
type Customer struct {
	g.Meta       `orm:"table:customer, do:true"`
	Id           interface{} //
	Name         interface{} // 会员名称
	Avatar       interface{} // 会员头像
	KefuName     interface{} // 客服账户
	Tel          interface{} // 会员手机号
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	Openid       interface{} // 会员ID
	AcountOpenid interface{} // 公众号会员ID
	Extra        interface{} // 会员扩展信息
	EntId        interface{} // 对接的企业ID
	Score        interface{} // 会员积分
}
