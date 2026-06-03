// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IndexBanner is the golang structure for table index_banner.
type IndexBanner struct {
	Id           int         `json:"id"           orm:"id"             description:""`
	Language     string      `json:"language"     orm:"language"       description:"语言"`
	BannerImage  string      `json:"bannerImage"  orm:"banner_image"   description:"轮播图"`
	Model        string      `json:"model"        orm:"model"          description:"模块   BANNER  横幅   EVENTS  事件"`
	Chain        string      `json:"chain"        orm:"chain"          description:"IN 内链  OUT 外联"`
	Path         string      `json:"path"         orm:"path"           description:"链接内容"`
	BannerStatus uint        `json:"bannerStatus" orm:"banner_status"  description:"1、启用 2、禁用"`
	MinappStatus int         `json:"minappStatus" orm:"minapp_status"  description:"是否排除小程序显示 1-不排除 2-排除"`
	Sort         int         `json:"sort"         orm:"sort"           description:"排序(越大越靠前)"`
	LinkOpenType int         `json:"linkOpenType" orm:"link_open_type" description:"外部链接跳转方式 1-webview 2-浏览器"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"      description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"      description:"更新时间"`
}
