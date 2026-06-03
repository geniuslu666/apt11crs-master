// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThCouponCategory is the golang structure for table th_coupon_category.
type ThCouponCategory struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序(越大越靠前)"`
	Status    uint        `json:"status"    orm:"status"     description:"1、启用 2、禁用"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
