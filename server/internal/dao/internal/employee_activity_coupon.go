// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmployeeActivityCouponDao is the data access object for table hg_employee_activity_coupon.
type EmployeeActivityCouponDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns EmployeeActivityCouponColumns // columns contains all the column names of Table for convenient usage.
}

// EmployeeActivityCouponColumns defines and stores column names for table hg_employee_activity_coupon.
type EmployeeActivityCouponColumns struct {
	Id                string // 关联ID
	ActivityId        string // 活动ID
	CouponId          string // 礼品券ID
	AvailableQuantity string // 可领取数量
	LimitDays         string // 限制天数
	PerDayAvailable   string // 每人N天可领取数量（0表示无限制）
	PerDayVerify      string // 每人每天可核销数量（0表示无限制）
	TotalReceived     string // 总领取数量
	TotalUsed         string // 已核销数量
	Status            string // 状态：1-正常 2-禁用
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// employeeActivityCouponColumns holds the columns for table hg_employee_activity_coupon.
var employeeActivityCouponColumns = EmployeeActivityCouponColumns{
	Id:                "id",
	ActivityId:        "activity_id",
	CouponId:          "coupon_id",
	AvailableQuantity: "available_quantity",
	LimitDays:         "limit_days",
	PerDayAvailable:   "per_day_available",
	PerDayVerify:      "per_day_verify",
	TotalReceived:     "total_received",
	TotalUsed:         "total_used",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewEmployeeActivityCouponDao creates and returns a new DAO object for table data access.
func NewEmployeeActivityCouponDao() *EmployeeActivityCouponDao {
	return &EmployeeActivityCouponDao{
		group:   "default",
		table:   "hg_employee_activity_coupon",
		columns: employeeActivityCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmployeeActivityCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmployeeActivityCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmployeeActivityCouponDao) Columns() EmployeeActivityCouponColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmployeeActivityCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmployeeActivityCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmployeeActivityCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
