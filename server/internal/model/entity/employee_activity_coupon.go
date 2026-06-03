// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityCoupon is the golang structure for table employee_activity_coupon.
type EmployeeActivityCoupon struct {
	Id                uint64      `json:"id"                orm:"id"                 description:"关联ID"`
	ActivityId        uint64      `json:"activityId"        orm:"activity_id"        description:"活动ID"`
	CouponId          uint64      `json:"couponId"          orm:"coupon_id"          description:"礼品券ID"`
	AvailableQuantity int         `json:"availableQuantity" orm:"available_quantity" description:"可领取数量"`
	LimitDays         int         `json:"limitDays"         orm:"limit_days"         description:"限制天数"`
	PerDayAvailable   int         `json:"perDayAvailable"   orm:"per_day_available"  description:"每人N天可领取数量（0表示无限制）"`
	PerDayVerify      int         `json:"perDayVerify"      orm:"per_day_verify"     description:"每人每天可核销数量（0表示无限制）"`
	TotalReceived     int         `json:"totalReceived"     orm:"total_received"     description:"总领取数量"`
	TotalUsed         int         `json:"totalUsed"         orm:"total_used"         description:"已核销数量"`
	Status            int         `json:"status"            orm:"status"             description:"状态：1-正常 2-禁用"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"         description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"         description:"更新时间"`
}
