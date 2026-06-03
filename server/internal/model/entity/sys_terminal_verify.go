// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTerminalVerify is the golang structure for table sys_terminal_verify.
type SysTerminalVerify struct {
	Id             int         `json:"id"             orm:"id"               description:""`
	TerminalId     int         `json:"terminalId"     orm:"terminal_id"      description:"终端ID"`
	VerifyType     string      `json:"verifyType"     orm:"verify_type"      description:"核销类型"`
	MchId          int         `json:"mchId"          orm:"mch_id"           description:"商户ID"`
	VerifyMemberId int         `json:"verifyMemberId" orm:"verify_member_id" description:"核销用户ID"`
	StoreId        int         `json:"storeId"        orm:"store_id"         description:"门店ID"`
	RestaurantId   int         `json:"restaurantId"   orm:"restaurant_id"    description:"餐厅ID"`
	FoodOrderId    int         `json:"foodOrderId"    orm:"food_order_id"    description:"餐厅订单ID"`
	MemberCouponId int         `json:"memberCouponId" orm:"member_coupon_id" description:"用户礼品券ID"`
	CouponMchName  string      `json:"couponMchName"  orm:"coupon_mch_name"  description:"核销商品名"`
	VerifyTime     *gtime.Time `json:"verifyTime"     orm:"verify_time"      description:"核销时间"`
	CreateAt       *gtime.Time `json:"createAt"       orm:"create_at"        description:"创建时间"`
	UpdateAt       *gtime.Time `json:"updateAt"       orm:"update_at"        description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       description:""`
}
