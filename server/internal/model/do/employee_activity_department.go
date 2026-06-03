// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityDepartment is the golang structure of table hg_employee_activity_department for DAO operations like Where/Data.
type EmployeeActivityDepartment struct {
	g.Meta       `orm:"table:hg_employee_activity_department, do:true"`
	Id           interface{} // 关联ID
	ActivityId   interface{} // 活动ID
	DepartmentId interface{} // 部门ID
	CreatedAt    *gtime.Time // 创建时间
}
