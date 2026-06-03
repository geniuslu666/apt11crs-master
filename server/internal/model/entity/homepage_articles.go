// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticles is the golang structure for table homepage_articles.
type HomepageArticles struct {
	Id                    int         `json:"id"                    orm:"id"                       description:""`
	Language              string      `json:"language"              orm:"language"                 description:""`
	No                    string      `json:"no"                    orm:"no"                       description:"活动编号"`
	ListPic               string      `json:"listPic"               orm:"list_pic"                 description:"列表图URL"`
	Title                 string      `json:"title"                 orm:"title"                    description:"标题"`
	Author                string      `json:"author"                orm:"author"                   description:"作者"`
	Content               string      `json:"content"               orm:"content"                  description:"内容"`
	Sort                  int         `json:"sort"                  orm:"sort"                     description:"排序(越大越靠前)"`
	Views                 int         `json:"views"                 orm:"views"                    description:"浏览量"`
	ThumbNum              int         `json:"thumbNum"              orm:"thumb_num"                description:"收藏数量"`
	RecommendTitle        string      `json:"recommendTitle"        orm:"recommend_title"          description:"推荐标题"`
	RecommendLinkType     string      `json:"recommendLinkType"     orm:"recommend_link_type"      description:"推荐链接类型 coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页"`
	CouponLimit           int         `json:"couponLimit"           orm:"coupon_limit"             description:"每人每张券可领取数量(废弃)"`
	HotelIds              string      `json:"hotelIds"              orm:"hotel_ids"                description:"民宿ids"`
	RestaurantIds         string      `json:"restaurantIds"         orm:"restaurant_ids"           description:"餐厅ids"`
	SpaServiceIds         string      `json:"spaServiceIds"         orm:"spa_service_ids"          description:"按摩ids"`
	StartTime             *gtime.Time `json:"startTime"             orm:"start_time"               description:"开始时间"`
	EndTime               *gtime.Time `json:"endTime"               orm:"end_time"                 description:"结束时间"`
	Status                int         `json:"status"                orm:"status"                   description:"状态：1-正常 2-禁用"`
	MinappStatus          int         `json:"minappStatus"          orm:"minapp_status"            description:"是否排除小程序显示 1-不排除 2-排除"`
	FoodIndexTitle        string      `json:"foodIndexTitle"        orm:"food_index_title"         description:"餐厅首页推荐标题"`
	SpaIndexTitle         string      `json:"spaIndexTitle"         orm:"spa_index_title"          description:"按摩首页推荐标题"`
	CarIndexTitle         string      `json:"carIndexTitle"         orm:"car_index_title"          description:"接送机首页推荐标题"`
	CouponLinkTitle       string      `json:"couponLinkTitle"       orm:"coupon_link_title"        description:"券链接区域标题"`
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"               description:""`
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"               description:""`
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"               description:""`
	CouponLinkSubTitle    string      `json:"couponLinkSubTitle"    orm:"coupon_link_sub_title"    description:"券链接区域副标题"`
	RecommendLinkTitle    string      `json:"recommendLinkTitle"    orm:"recommend_link_title"     description:"推荐链接区域标题"`
	RecommendLinkSubTitle string      `json:"recommendLinkSubTitle" orm:"recommend_link_sub_title" description:"推荐链接区域副标题"`
	ButtonTxt             string      `json:"buttonTxt"             orm:"button_txt"               description:"外链按钮文字"`
	OutLinkTitle          string      `json:"outLinkTitle"          orm:"out_link_title"           description:"外链推荐标题"`
	OutLinkContent        string      `json:"outLinkContent"        orm:"out_link_content"         description:"外链推荐链接"`
	Chain                 string      `json:"chain"                 orm:"chain"                    description:"活动内外链: IN 内链  OUT 外联"`
	Path                  string      `json:"path"                  orm:"path"                     description:"外链链接内容"`
	LinkOpenType          int         `json:"linkOpenType"          orm:"link_open_type"           description:"外部链接跳转方式 1-webview 2-浏览器"`
	LimitWeek             string      `json:"limitWeek"             orm:"limit_week"               description:"不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割"`
}
