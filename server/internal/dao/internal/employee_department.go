// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeDepartmentDao is the data access object for the table hg_employee_department.
type EmployeeDepartmentDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns EmployeeDepartmentColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeDepartmentColumns defines and stores column names for the table hg_employee_department.
type EmployeeDepartmentColumns struct {
	Id          string // 部门ID
	Name        string // 部门名称
	ParentId    string // 上级部门ID（NULL表示顶级部门）
	Level       string // 部门层级（1表示顶级）
	Path        string // 部门路径（如：1,2,3）
	ManagerId   string // 部门负责人ID
	Description string // 部门描述
	Sort        string // 排序
	Status      string // 状态：1-正常 2-禁用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// employeeDepartmentColumns holds the columns for the table hg_employee_department.
var employeeDepartmentColumns = EmployeeDepartmentColumns{
	Id:          "id",
	Name:        "name",
	ParentId:    "parent_id",
	Level:       "level",
	Path:        "path",
	ManagerId:   "manager_id",
	Description: "description",
	Sort:        "sort",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewEmployeeDepartmentDao creates and returns a new DAO object for table data access.
func NewEmployeeDepartmentDao() *EmployeeDepartmentDao {
	return &EmployeeDepartmentDao{
		group:   "default",
		table:   "hg_employee_department",
		columns: employeeDepartmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmployeeDepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmployeeDepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmployeeDepartmentDao) Columns() EmployeeDepartmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmployeeDepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmployeeDepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EmployeeDepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
