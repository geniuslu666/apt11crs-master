// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBanner is the golang structure for table pms_banner.
type PmsBanner struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	Language     string      `json:"language"     orm:"language"      description:"语言"`
	BannerName   string      `json:"bannerName"   orm:"banner_name"   description:"banner名称"`
	BannerImage  string      `json:"bannerImage"  orm:"banner_image"  description:"轮播图"`
	Model        string      `json:"model"        orm:"model"         description:"模块   BANNER  横幅   EVENTS  事件"`
	Chain        string      `json:"chain"        orm:"chain"         description:"IN 内链  OUT 外联"`
	Path         string      `json:"path"         orm:"path"          description:"链接内容"`
	BannerStatus uint        `json:"bannerStatus" orm:"banner_status" description:"1、启用 2、禁用"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"更新时间"`
}
