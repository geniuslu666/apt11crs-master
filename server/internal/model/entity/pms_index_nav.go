// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsIndexNav is the golang structure for table pms_index_nav.
type PmsIndexNav struct {
	Id           int64       `json:"id"           orm:"id"             description:""`
	Name         string      `json:"name"         orm:"name"           description:"导航名称(多语言)"`
	Tag          string      `json:"tag"          orm:"tag"            description:"标签"`
	Image        string      `json:"image"        orm:"image"          description:"图标"`
	AppLink      string      `json:"appLink"      orm:"app_link"       description:"app跳转链接"`
	WxLink       string      `json:"wxLink"       orm:"wx_link"        description:"微信跳转链接"`
	Sort         int         `json:"sort"         orm:"sort"           description:"排序(越大越靠前)"`
	Status       uint        `json:"status"       orm:"status"         description:"状态1、启用 2、禁用"`
	MinappStatus int         `json:"minappStatus" orm:"minapp_status"  description:"是否排除小程序显示 1-不排除 2-排除"`
	Chain        string      `json:"chain"        orm:"chain"          description:"IN 内链  OUT 外联"`
	LinkOpenType int         `json:"linkOpenType" orm:"link_open_type" description:"外部链接跳转方式 1-webview 2-浏览器"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"      description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"      description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"     description:""`
}
