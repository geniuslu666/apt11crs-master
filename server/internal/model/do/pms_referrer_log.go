// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsReferrerLog is the golang structure of table hg_pms_referrer_log for DAO operations like Where/Data.
type PmsReferrerLog struct {
	g.Meta       `orm:"table:hg_pms_referrer_log, do:true"`
	Id           interface{} //
	MemberId     interface{} // 会员ID
	Referrer     interface{} // 首次推荐人
	LastReferrer interface{} // 最后推荐人
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
