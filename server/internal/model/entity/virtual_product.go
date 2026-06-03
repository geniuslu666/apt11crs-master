// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VirtualProduct is the golang structure for table virtual_product.
type VirtualProduct struct {
	Id              int         `json:"id"              orm:"id"               description:""`
	EntId           string      `json:"entId"           orm:"ent_id"           description:"企业ID"`
	KefuName        string      `json:"kefuName"        orm:"kefu_name"        description:"客服账户"`
	ProductName     string      `json:"productName"     orm:"product_name"     description:""`
	ProductCategory string      `json:"productCategory" orm:"product_category" description:""`
	Payment         string      `json:"payment"         orm:"payment"          description:"支付方式，wechat 微信支付；nan66 南星码支付"`
	Description     string      `json:"description"     orm:"description"      description:""`
	Price           int         `json:"price"           orm:"price"            description:"金额"`
	ProductImg      string      `json:"productImg"      orm:"product_img"      description:""`
	ResourceLink    string      `json:"resourceLink"    orm:"resource_link"    description:""`
	IsActive        int         `json:"isActive"        orm:"is_active"        description:"在线状态，1在售，2下架"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`
}
