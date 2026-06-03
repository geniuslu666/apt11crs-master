// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyRecord is the golang structure for table travel_verify_record.
type TravelVerifyRecord struct {
	Id            uint64      `json:"id"            orm:"id"              description:""`
	OrderId       uint64      `json:"orderId"       orm:"order_id"        description:"订单ID"`
	OrderSn       string      `json:"orderSn"       orm:"order_sn"        description:"预约单号"`
	ProductId     uint64      `json:"productId"     orm:"product_id"      description:"产品ID"`
	MemberId      uint64      `json:"memberId"      orm:"member_id"       description:"会员ID"`
	BookDate      *gtime.Time `json:"bookDate"      orm:"book_date"       description:"预约日期"`
	VerifyStaffId uint64      `json:"verifyStaffId" orm:"verify_staff_id" description:"核销人员ID"`
	VerifyTime    *gtime.Time `json:"verifyTime"    orm:"verify_time"     description:"核销时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`
}
