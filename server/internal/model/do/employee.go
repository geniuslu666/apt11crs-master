// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Employee is the golang structure of table hg_employee for DAO operations like Where/Data.
type Employee struct {
	g.Meta       `orm:"table:hg_employee, do:true"`
	Id           interface{} // 员工ID
	Name         interface{} // 员工姓名
	PhoneArea    interface{} // 电话区号
	Phone        interface{} // 电话号码
	MemberId     interface{} // 绑定用户ID
	DepartmentId interface{} // 员工部门ID
	EmployeeNo   interface{} // 员工编号
	Status       interface{} // 状态：1-正常 2-禁用
	Remark       interface{} // 备注
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
