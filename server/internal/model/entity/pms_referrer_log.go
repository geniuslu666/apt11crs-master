// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsReferrerLog is the golang structure for table pms_referrer_log.
type PmsReferrerLog struct {
	Id           int         `json:"id"           orm:"id"            description:""`
	MemberId     int         `json:"memberId"     orm:"member_id"     description:"会员ID"`
	Referrer     int         `json:"referrer"     orm:"referrer"      description:"首次推荐人"`
	LastReferrer int         `json:"lastReferrer" orm:"last_referrer" description:"最后推荐人"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`
}
