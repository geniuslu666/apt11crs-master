// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Receipt is the golang structure for table receipt.
type Receipt struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	UserName    string      `json:"userName"    orm:"user_name"    description:"发票抬头"`
	OrderSn     string      `json:"orderSn"     orm:"order_sn"     description:"订单号"`
	ReceiptPath string      `json:"receiptPath" orm:"receipt_path" description:"领售书文件地址"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
