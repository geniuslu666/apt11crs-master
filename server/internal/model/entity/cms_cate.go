// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsCate is the golang structure for table cms_cate.
type CmsCate struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	CatName   string      `json:"catName"   orm:"cat_name"   description:"分类名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
