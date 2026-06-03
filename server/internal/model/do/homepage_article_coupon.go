// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleCoupon is the golang structure of table hg_homepage_article_coupon for DAO operations like Where/Data.
type HomepageArticleCoupon struct {
	g.Meta            `orm:"table:hg_homepage_article_coupon, do:true"`
	Id                interface{} // 关联ID
	CouponType        interface{} // coupon-优惠券，thCoupon-礼品券
	ArticleId         interface{} // 活动ID
	CouponId          interface{} // 礼品券ID
	AvailableQuantity interface{} // 可领取数量
	PerDayAvailable   interface{} // 每人N天可领取数量（0表示无限制）
	LimitDays         interface{} // 限制天数
	TotalReceived     interface{} // 总领取数量
	TotalUsed         interface{} // 已核销数量
	Status            interface{} // 状态：1-正常 2-禁用
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
	FixedTerm         interface{} // 领取后几天内有效（0表示领取当日23:59:59）
}
