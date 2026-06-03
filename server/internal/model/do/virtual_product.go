// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VirtualProduct is the golang structure of table virtual_product for DAO operations like Where/Data.
type VirtualProduct struct {
	g.Meta          `orm:"table:virtual_product, do:true"`
	Id              interface{} //
	EntId           interface{} // 企业ID
	KefuName        interface{} // 客服账户
	ProductName     interface{} //
	ProductCategory interface{} //
	Payment         interface{} // 支付方式，wechat 微信支付；nan66 南星码支付
	Description     interface{} //
	Price           interface{} // 金额
	ProductImg      interface{} //
	ResourceLink    interface{} //
	IsActive        interface{} // 在线状态，1在售，2下架
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
