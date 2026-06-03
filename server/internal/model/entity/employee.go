// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Employee is the golang structure for table employee.
type Employee struct {
	Id           uint64      `json:"id"           orm:"id"            description:"员工ID"`
	Name         string      `json:"name"         orm:"name"          description:"员工姓名"`
	PhoneArea    string      `json:"phoneArea"    orm:"phone_area"    description:"电话区号"`
	Phone        string      `json:"phone"        orm:"phone"         description:"电话号码"`
	MemberId     uint64      `json:"memberId"     orm:"member_id"     description:"绑定用户ID"`
	DepartmentId uint64      `json:"departmentId" orm:"department_id" description:"员工部门ID"`
	EmployeeNo   string      `json:"employeeNo"   orm:"employee_no"   description:"员工编号"`
	Status       int         `json:"status"       orm:"status"        description:"状态：1-正常 2-禁用"`
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`
}
