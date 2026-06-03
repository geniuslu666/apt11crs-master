// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsFolioDao is the data access object for the table hg_pms_folio.
type PmsFolioDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns PmsFolioColumns // columns contains all the column names of Table for convenient usage.
}

// PmsFolioColumns defines and stores column names for the table hg_pms_folio.
type PmsFolioColumns struct {
	Id                 string // 主键
	Uid                string // 第三方系统的ID
	Charges            string // 费用数组，参考Charges
	Payments           string // 付款数组，参考Payments
	TotalChargeAmount  string // 总收费金额
	TotalPaymentAmount string // 总付款金额
	OutstandingBalance string // 未清余额
	CreatedAt          string //
	UpdatedAt          string //
}

// pmsFolioColumns holds the columns for the table hg_pms_folio.
var pmsFolioColumns = PmsFolioColumns{
	Id:                 "id",
	Uid:                "uid",
	Charges:            "charges",
	Payments:           "payments",
	TotalChargeAmount:  "total_charge_amount",
	TotalPaymentAmount: "total_payment_amount",
	OutstandingBalance: "outstanding_balance",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewPmsFolioDao creates and returns a new DAO object for table data access.
func NewPmsFolioDao() *PmsFolioDao {
	return &PmsFolioDao{
		group:   "default",
		table:   "hg_pms_folio",
		columns: pmsFolioColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsFolioDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsFolioDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsFolioDao) Columns() PmsFolioColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsFolioDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsFolioDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsFolioDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
