// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityDepartment is the golang structure for table employee_activity_department.
type EmployeeActivityDepartment struct {
	Id           uint64      `json:"id"           orm:"id"            description:"关联ID"`
	ActivityId   uint64      `json:"activityId"   orm:"activity_id"   description:"活动ID"`
	DepartmentId uint64      `json:"departmentId" orm:"department_id" description:"部门ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
}
