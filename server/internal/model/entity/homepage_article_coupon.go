// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleCoupon is the golang structure for table homepage_article_coupon.
type HomepageArticleCoupon struct {
	Id                uint64      `json:"id"                orm:"id"                 description:"关联ID"`
	CouponType        string      `json:"couponType"        orm:"coupon_type"        description:"coupon-优惠券，thCoupon-礼品券"`
	ArticleId         uint64      `json:"articleId"         orm:"article_id"         description:"活动ID"`
	CouponId          uint64      `json:"couponId"          orm:"coupon_id"          description:"礼品券ID"`
	AvailableQuantity int         `json:"availableQuantity" orm:"available_quantity" description:"可领取数量"`
	PerDayAvailable   int         `json:"perDayAvailable"   orm:"per_day_available"  description:"每人N天可领取数量（0表示无限制）"`
	LimitDays         int         `json:"limitDays"         orm:"limit_days"         description:"限制天数"`
	TotalReceived     int         `json:"totalReceived"     orm:"total_received"     description:"总领取数量"`
	TotalUsed         int         `json:"totalUsed"         orm:"total_used"         description:"已核销数量"`
	Status            int         `json:"status"            orm:"status"             description:"状态：1-正常 2-禁用"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"         description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"         description:"更新时间"`
	FixedTerm         int         `json:"fixedTerm"         orm:"fixed_term"         description:"领取后几天内有效（0表示领取当日23:59:59）"`
}
