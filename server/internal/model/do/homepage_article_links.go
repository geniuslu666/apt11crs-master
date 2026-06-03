// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HomepageArticleLinks is the golang structure of table hg_homepage_article_links for DAO operations like Where/Data.
type HomepageArticleLinks struct {
	g.Meta       `orm:"table:hg_homepage_article_links, do:true"`
	Id           interface{} //
	ArticleId    interface{} // 文章ID
	Title        interface{} // 链接标题
	LinkText     interface{} // 链接显示的文字
	AppLink      interface{} // APP跳转链接
	WxLink       interface{} // 微信跳转链接
	UrlParam     *gjson.Json // APP跳转链接参数，JSON格式存储
	LinkType     interface{} // coupon-优惠券 thCoupon-礼品券，hotelDetail-民宿详情，foodIndex-餐厅首页，foodDetail-餐厅详情，spaIndex-按摩首页，spaDetail-按摩详情，carIndex-接送机首页, outLink-外部链接
	LinkId       interface{} // 对应跳转链接的ID
	LinkOpenType interface{} // 外部链接跳转方式 1-webview 2-浏览器
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
