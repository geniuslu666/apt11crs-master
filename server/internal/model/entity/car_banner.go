// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarBanner is the golang structure for table car_banner.
type CarBanner struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	Language    string      `json:"language"    orm:"language"     description:"语言"`
	BannerImage string      `json:"bannerImage" orm:"banner_image" description:"轮播图"`
	Status      uint        `json:"status"      orm:"status"       description:"状态1、启用 2、禁用"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序(越大越靠前)"`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    description:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    description:"更新时间"`
}
