// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyStaffScope is the golang structure for table travel_verify_staff_scope.
type TravelVerifyStaffScope struct {
	Id        uint64      `json:"id"        orm:"id"         description:""`
	StaffId   uint64      `json:"staffId"   orm:"staff_id"   description:"核销人员ID"`
	ProductId uint64      `json:"productId" orm:"product_id" description:"产品ID"`
	SkuId     uint64      `json:"skuId"     orm:"sku_id"     description:"SKU ID（0=产品全部SKU）"`
	IsAll     int         `json:"isAll"     orm:"is_all"     description:"是否全部可核销（1是 0否）"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"软删除时间（NULL=正常）"`
}
