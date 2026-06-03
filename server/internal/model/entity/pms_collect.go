// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCollect is the golang structure for table pms_collect.
type PmsCollect struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	MemberId    int         `json:"memberId"    orm:"member_id"    description:""`
	CollectType string      `json:"collectType" orm:"collect_type" description:"收藏类型   HOST 、 酒店  REPAST 餐饮"`
	CollectId   int         `json:"collectId"   orm:"collect_id"   description:"收藏数据ID"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
