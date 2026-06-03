// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// New is the golang structure for table new.
type New struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Tag       string      `json:"tag"       orm:"tag"        description:"公告标签"`
	Title     string      `json:"title"     orm:"title"      description:"公告标题"`
	Content   string      `json:"content"   orm:"content"    description:"公告内容"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	Status    int         `json:"status"    orm:"status"     description:"状态，未启用"`
}
