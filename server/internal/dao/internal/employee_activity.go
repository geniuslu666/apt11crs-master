// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeActivityDao is the data access object for the table hg_employee_activity.
type EmployeeActivityDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns EmployeeActivityColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeActivityColumns defines and stores column names for the table hg_employee_activity.
type EmployeeActivityColumns struct {
	Id              string // 活动ID
	Name            string // 活动名称（多语言）
	Cover           string // 活动封面
	Description     string // 活动描述（多语言）
	ValidityType    string // 有效期类型：1-指定时间段 2-长期有效
	StartTime       string // 开始时间
	EndTime         string // 结束时间
	Status          string // 状态：1-未开始 2-进行中 3-已结束
	ManualClosed    string // 是否手动关闭：0-否 1-是
	RestrictionType string // 限制类型：1-不做任何限制 2-限制指定部门 3-限制指定员工
	Rule            string // 活动规则
	Sort            string // 排序
	Remark          string // 备注
	IsEnabled       string // 状态：1-正常 2-禁用
	CouponValidity  string // 券有效期  1-跟随活动  2-自身有效期
	LimitWeek       string // 不可领取限制：6-是周六不可领， 7是周日不可领，以逗号分割
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
}

// employeeActivityColumns holds the columns for the table hg_employee_activity.
var employeeActivityColumns = EmployeeActivityColumns{
	Id:              "id",
	Name:            "name",
	Cover:           "cover",
	Description:     "description",
	ValidityType:    "validity_type",
	StartTime:       "start_time",
	EndTime:         "end_time",
	Status:          "status",
	ManualClosed:    "manual_closed",
	RestrictionType: "restriction_type",
	Rule:            "rule",
	Sort:            "sort",
	Remark:          "remark",
	IsEnabled:       "is_enabled",
	CouponValidity:  "coupon_validity",
	LimitWeek:       "limit_week",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewEmployeeActivityDao creates and returns a new DAO object for table data access.
func NewEmployeeActivityDao() *EmployeeActivityDao {
	return &EmployeeActivityDao{
		group:   "default",
		table:   "hg_employee_activity",
		columns: employeeActivityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmployeeActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmployeeActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmployeeActivityDao) Columns() EmployeeActivityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmployeeActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmployeeActivityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EmployeeActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
