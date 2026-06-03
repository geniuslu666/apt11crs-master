// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityEmployee is the golang structure for table employee_activity_employee.
type EmployeeActivityEmployee struct {
	Id         uint64      `json:"id"         orm:"id"          description:"关联ID"`
	ActivityId uint64      `json:"activityId" orm:"activity_id" description:"活动ID"`
	EmployeeId uint64      `json:"employeeId" orm:"employee_id" description:"员工ID"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
}
