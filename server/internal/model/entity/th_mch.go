// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMch is the golang structure for table th_mch.
type ThMch struct {
	Id          int         `json:"id"          orm:"id"            description:""`
	CategoryId  uint        `json:"categoryId"  orm:"category_id"   description:"分类ID"`
	Name        string      `json:"name"        orm:"name"          description:"名称"`
	Logo        string      `json:"logo"        orm:"logo"          description:"LOGO"`
	ContactInfo string      `json:"contactInfo" orm:"contact_info"  description:"联系信息"`
	Sort        int         `json:"sort"        orm:"sort"          description:"排序(越大越靠前)"`
	Status      uint        `json:"status"      orm:"status"        description:"1、启用 2、禁用"`
	StoreOnNum  int         `json:"storeOnNum"  orm:"store_on_num"  description:"启用中门店数量"`
	StoreOffNum int         `json:"storeOffNum" orm:"store_off_num" description:"禁用中门店数量"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:""`
}
