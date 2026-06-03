// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeActivityDepartmentDao is the data access object for the table hg_employee_activity_department.
type EmployeeActivityDepartmentDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of the current DAO.
	columns EmployeeActivityDepartmentColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeActivityDepartmentColumns defines and stores column names for the table hg_employee_activity_department.
type EmployeeActivityDepartmentColumns struct {
	Id           string // 关联ID
	ActivityId   string // 活动ID
	DepartmentId string // 部门ID
	CreatedAt    string // 创建时间
}

// employeeActivityDepartmentColumns holds the columns for the table hg_employee_activity_department.
var employeeActivityDepartmentColumns = EmployeeActivityDepartmentColumns{
	Id:           "id",
	ActivityId:   "activity_id",
	DepartmentId: "department_id",
	CreatedAt:    "created_at",
}

// NewEmployeeActivityDepartmentDao creates and returns a new DAO object for table data access.
func NewEmployeeActivityDepartmentDao() *EmployeeActivityDepartmentDao {
	return &EmployeeActivityDepartmentDao{
		group:   "default",
		table:   "hg_employee_activity_department",
		columns: employeeActivityDepartmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmployeeActivityDepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmployeeActivityDepartmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmployeeActivityDepartmentDao) Columns() EmployeeActivityDepartmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmployeeActivityDepartmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmployeeActivityDepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EmployeeActivityDepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
