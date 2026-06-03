// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenterCategory is the golang structure for table pms_helpcenter_category.
type PmsHelpcenterCategory struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`
	Language  string      `json:"language"  orm:"language"   description:""`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序(越大越靠前)"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
