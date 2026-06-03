// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThStoreTerminalDao is the data access object for the table hg_th_store_terminal.
type ThStoreTerminalDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns ThStoreTerminalColumns // columns contains all the column names of Table for convenient usage.
}

// ThStoreTerminalColumns defines and stores column names for the table hg_th_store_terminal.
type ThStoreTerminalColumns struct {
	StoreId    string // 门店ID
	TerminalId string // 终端ID
	PrintTimes string // 打印联数(次数)
}

// thStoreTerminalColumns holds the columns for the table hg_th_store_terminal.
var thStoreTerminalColumns = ThStoreTerminalColumns{
	StoreId:    "store_id",
	TerminalId: "terminal_id",
	PrintTimes: "print_times",
}

// NewThStoreTerminalDao creates and returns a new DAO object for table data access.
func NewThStoreTerminalDao() *ThStoreTerminalDao {
	return &ThStoreTerminalDao{
		group:   "default",
		table:   "hg_th_store_terminal",
		columns: thStoreTerminalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThStoreTerminalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThStoreTerminalDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThStoreTerminalDao) Columns() ThStoreTerminalColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThStoreTerminalDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThStoreTerminalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThStoreTerminalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
