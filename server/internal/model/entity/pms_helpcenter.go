// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenter is the golang structure for table pms_helpcenter.
type PmsHelpcenter struct {
	Id         int         `json:"id"         orm:"id"          description:""`
	Language   string      `json:"language"   orm:"language"    description:""`
	CategoryId uint        `json:"categoryId" orm:"category_id" description:"分类ID"`
	Title      string      `json:"title"      orm:"title"       description:"标题"`
	Content    string      `json:"content"    orm:"content"     description:"内容"`
	Sort       int         `json:"sort"       orm:"sort"        description:"排序(越大越靠前)"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:""`
}
