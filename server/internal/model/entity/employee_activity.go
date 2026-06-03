// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivity is the golang structure for table employee_activity.
type EmployeeActivity struct {
	Id              uint64      `json:"id"              orm:"id"               description:"活动ID"`
	Name            string      `json:"name"            orm:"name"             description:"活动名称（多语言）"`
	Cover           string      `json:"cover"           orm:"cover"            description:"活动封面"`
	Description     string      `json:"description"     orm:"description"      description:"活动描述（多语言）"`
	ValidityType    int         `json:"validityType"    orm:"validity_type"    description:"有效期类型：1-指定时间段 2-长期有效"`
	StartTime       *gtime.Time `json:"startTime"       orm:"start_time"       description:"开始时间"`
	EndTime         *gtime.Time `json:"endTime"         orm:"end_time"         description:"结束时间"`
	Status          int         `json:"status"          orm:"status"           description:"状态：1-未开始 2-进行中 3-已结束"`
	ManualClosed    int         `json:"manualClosed"    orm:"manual_closed"    description:"是否手动关闭：0-否 1-是"`
	RestrictionType int         `json:"restrictionType" orm:"restriction_type" description:"限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工"`
	Rule            string      `json:"rule"            orm:"rule"             description:"活动规则"`
	Sort            int         `json:"sort"            orm:"sort"             description:"排序"`
	Remark          string      `json:"remark"          orm:"remark"           description:"备注"`
	IsEnabled       int         `json:"isEnabled"       orm:"is_enabled"       description:"状态：1-正常 2-禁用"`
	CouponValidity  int         `json:"couponValidity"  orm:"coupon_validity"  description:"券有效期  1-跟随活动  2-自身有效期"`
	LimitWeek       string      `json:"limitWeek"       orm:"limit_week"       description:"不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:"删除时间"`
}
