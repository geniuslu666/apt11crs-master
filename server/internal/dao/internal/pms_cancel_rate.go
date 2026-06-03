// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsCancelRateDao is the data access object for the table hg_pms_cancel_rate.
type PmsCancelRateDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsCancelRateColumns // columns contains all the column names of Table for convenient usage.
}

// PmsCancelRateColumns defines and stores column names for the table hg_pms_cancel_rate.
type PmsCancelRateColumns struct {
	Id        string //
	Mode      string // 规则模式
	StartDays string // 开始天数
	EndDays   string // 结束天数
	Rate      string // 费率
	Name      string // 规则名
	Sort      string // 排序规则  从小到大
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// pmsCancelRateColumns holds the columns for the table hg_pms_cancel_rate.
var pmsCancelRateColumns = PmsCancelRateColumns{
	Id:        "id",
	Mode:      "mode",
	StartDays: "start_days",
	EndDays:   "end_days",
	Rate:      "rate",
	Name:      "name",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPmsCancelRateDao creates and returns a new DAO object for table data access.
func NewPmsCancelRateDao() *PmsCancelRateDao {
	return &PmsCancelRateDao{
		group:   "default",
		table:   "hg_pms_cancel_rate",
		columns: pmsCancelRateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsCancelRateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsCancelRateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsCancelRateDao) Columns() PmsCancelRateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsCancelRateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsCancelRateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsCancelRateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
