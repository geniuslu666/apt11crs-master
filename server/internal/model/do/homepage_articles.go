// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticles is the golang structure of table hg_homepage_articles for DAO operations like Where/Data.
type HomepageArticles struct {
	g.Meta                `orm:"table:hg_homepage_articles, do:true"`
	Id                    interface{} //
	Language              interface{} //
	No                    interface{} // 活动编号
	ListPic               interface{} // 列表图URL
	Title                 interface{} // 标题
	Author                interface{} // 作者
	Content               interface{} // 内容
	Sort                  interface{} // 排序(越大越靠前)
	Views                 interface{} // 浏览量
	ThumbNum              interface{} // 收藏数量
	RecommendTitle        interface{} // 推荐标题
	RecommendLinkType     interface{} // 推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页
	CouponLimit           interface{} // 每人每张券可领取数量(废弃)
	HotelIds              interface{} // 民宿ids
	RestaurantIds         interface{} // 餐厅ids
	SpaServiceIds         interface{} // 按摩ids
	StartTime             *gtime.Time // 开始时间
	EndTime               *gtime.Time // 结束时间
	Status                interface{} // 状态：1-正常 2-禁用
	MinappStatus          interface{} // 是否排除小程序显示 1-不排除 2-排除
	FoodIndexTitle        interface{} // 餐厅首页推荐标题
	SpaIndexTitle         interface{} // 按摩首页推荐标题
	CarIndexTitle         interface{} // 接送机首页推荐标题
	CouponLinkTitle       interface{} // 券链接区域标题
	CreatedAt             *gtime.Time //
	UpdatedAt             *gtime.Time //
	DeletedAt             *gtime.Time //
	CouponLinkSubTitle    interface{} // 券链接区域副标题
	RecommendLinkTitle    interface{} // 推荐链接区域标题
	RecommendLinkSubTitle interface{} // 推荐链接区域副标题
	ButtonTxt             interface{} // 外链按钮文字
	OutLinkTitle          interface{} // 外链推荐标题
	OutLinkContent        interface{} // 外链推荐链接
	Chain                 interface{} // 活动内外链: IN 内链  OUT 外联
	Path                  interface{} // 外链链接内容
	LinkOpenType          interface{} // 外部链接跳转方式 1-webview 2-浏览器
	LimitWeek             interface{} // 不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割
}
