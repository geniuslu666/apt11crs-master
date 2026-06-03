// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeDepartment is the golang structure of table hg_employee_department for DAO operations like Where/Data.
type EmployeeDepartment struct {
	g.Meta      `orm:"table:hg_employee_department, do:true"`
	Id          interface{} // 部门ID
	Name        interface{} // 部门名称
	ParentId    interface{} // 上级部门ID（NULL表示顶级部门）
	Level       interface{} // 部门层级（1表示顶级）
	Path        interface{} // 部门路径（如：1,2,3）
	ManagerId   interface{} // 部门负责人ID
	Description interface{} // 部门描述
	Sort        interface{} // 排序
	Status      interface{} // 状态：1-正常 2-禁用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
