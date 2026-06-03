// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeDao is the data access object for the table hg_employee.
type EmployeeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns EmployeeColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeColumns defines and stores column names for the table hg_employee.
type EmployeeColumns struct {
	Id           string // 员工ID
	Name         string // 员工姓名
	PhoneArea    string // 电话区号
	Phone        string // 电话号码
	MemberId     string // 绑定用户ID
	DepartmentId string // 员工部门ID
	EmployeeNo   string // 员工编号
	Status       string // 状态：1-正常 2-禁用
	Remark       string // 备注
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// employeeColumns holds the columns for the table hg_employee.
var employeeColumns = EmployeeColumns{
	Id:           "id",
	Name:         "name",
	PhoneArea:    "phone_area",
	Phone:        "phone",
	MemberId:     "member_id",
	DepartmentId: "department_id",
	EmployeeNo:   "employee_no",
	Status:       "status",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewEmployeeDao creates and returns a new DAO object for table data access.
func NewEmployeeDao() *EmployeeDao {
	return &EmployeeDao{
		group:   "default",
		table:   "hg_employee",
		columns: employeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmployeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmployeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmployeeDao) Columns() EmployeeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmployeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmployeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EmployeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
