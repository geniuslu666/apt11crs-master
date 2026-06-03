// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTestNav is the golang structure for table pms_test_nav.
type PmsTestNav struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"名称"`
	Image     string      `json:"image"     orm:"image"      description:"图标"`
	AppLink   string      `json:"appLink"   orm:"app_link"   description:"app跳转链接"`
	WxLink    string      `json:"wxLink"    orm:"wx_link"    description:"微信跳转链接"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序(越大越靠前)"`
	Status    uint        `json:"status"    orm:"status"     description:"状态1、启用 2、禁用"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
