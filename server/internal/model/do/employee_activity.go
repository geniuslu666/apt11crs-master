// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivity is the golang structure of table hg_employee_activity for DAO operations like Where/Data.
type EmployeeActivity struct {
	g.Meta          `orm:"table:hg_employee_activity, do:true"`
	Id              interface{} // 活动ID
	Name            interface{} // 活动名称（多语言）
	Cover           interface{} // 活动封面
	Description     interface{} // 活动描述（多语言）
	ValidityType    interface{} // 有效期类型：1-指定时间段 2-长期有效
	StartTime       *gtime.Time // 开始时间
	EndTime         *gtime.Time // 结束时间
	Status          interface{} // 状态：1-未开始 2-进行中 3-已结束
	ManualClosed    interface{} // 是否手动关闭：0-否 1-是
	RestrictionType interface{} // 限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工
	Rule            interface{} // 活动规则
	Sort            interface{} // 排序
	Remark          interface{} // 备注
	IsEnabled       interface{} // 状态：1-正常 2-禁用
	CouponValidity  interface{} // 券有效期  1-跟随活动  2-自身有效期
	LimitWeek       interface{} // 不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
