// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleLinks is the golang structure for table homepage_article_links.
type HomepageArticleLinks struct {
	Id           int         `json:"id"           orm:"id"             description:""`
	ArticleId    uint64      `json:"articleId"    orm:"article_id"     description:"文章ID"`
	Title        string      `json:"title"        orm:"title"          description:"链接标题"`
	LinkText     string      `json:"linkText"     orm:"link_text"      description:"链接显示的文字"`
	AppLink      string      `json:"appLink"      orm:"app_link"       description:"APP跳转链接"`
	WxLink       string      `json:"wxLink"       orm:"wx_link"        description:"微信跳转链接"`
	UrlParam     *gjson.Json `json:"urlParam"     orm:"url_param"      description:"APP跳转链接参数，JSON格式存储"`
	LinkType     string      `json:"linkType"     orm:"link_type"      description:"coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页, outLink-外部链接"`
	LinkId       int         `json:"linkId"       orm:"link_id"        description:"对应跳转链接的ID"`
	LinkOpenType int         `json:"linkOpenType" orm:"link_open_type" description:"外部链接跳转方式 1-webview 2-浏览器"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:""`
}
