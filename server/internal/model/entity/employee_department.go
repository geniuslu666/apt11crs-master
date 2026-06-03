// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmployeeDepartment is the golang structure for table employee_department.
type EmployeeDepartment struct {
	Id          uint64      `json:"id"          orm:"id"          description:"部门ID"`
	Name        string      `json:"name"        orm:"name"        description:"部门名称"`
	ParentId    uint64      `json:"parentId"    orm:"parent_id"   description:"上级部门ID（NULL表示顶级部门）"`
	Level       int         `json:"level"       orm:"level"       description:"部门层级（1表示顶级）"`
	Path        string      `json:"path"        orm:"path"        description:"部门路径（如：1,2,3）"`
	ManagerId   uint64      `json:"managerId"   orm:"manager_id"  description:"部门负责人ID"`
	Description string      `json:"description" orm:"description" description:"部门描述"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Status      int         `json:"status"      orm:"status"      description:"状态：1-正常 2-禁用"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
