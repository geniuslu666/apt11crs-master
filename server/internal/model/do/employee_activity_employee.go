// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeActivityEmployee is the golang structure of table hg_employee_activity_employee for DAO operations like Where/Data.
type EmployeeActivityEmployee struct {
	g.Meta     `orm:"table:hg_employee_activity_employee, do:true"`
	Id         interface{} // 关联ID
	ActivityId interface{} // 活动ID
	EmployeeId interface{} // 员工ID
	CreatedAt  *gtime.Time // 创建时间
}
