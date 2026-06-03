// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCancelRate is the golang structure for table pms_cancel_rate.
type PmsCancelRate struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Mode      string      `json:"mode"      orm:"mode"       description:"规则模式"`
	StartDays int         `json:"startDays" orm:"start_days" description:"开始天数"`
	EndDays   int         `json:"endDays"   orm:"end_days"   description:"结束天数"`
	Rate      int         `json:"rate"      orm:"rate"       description:"费率"`
	Name      string      `json:"name"      orm:"name"       description:"规则名"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序规则  从小到大"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
