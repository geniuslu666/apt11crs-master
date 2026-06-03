// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImMessageReadDao is the data access object for the table hg_im_message_read.
type ImMessageReadDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns ImMessageReadColumns // columns contains all the column names of Table for convenient usage.
}

// ImMessageReadColumns defines and stores column names for the table hg_im_message_read.
type ImMessageReadColumns struct {
	Id        string //
	MemberId  string //
	ImId      string // 消息ID
	OrderSn   string // 订单号
	CreatedAt string //
	UpdatedAt string //
}

// imMessageReadColumns holds the columns for the table hg_im_message_read.
var imMessageReadColumns = ImMessageReadColumns{
	Id:        "id",
	MemberId:  "member_id",
	ImId:      "im_id",
	OrderSn:   "order_sn",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewImMessageReadDao creates and returns a new DAO object for table data access.
func NewImMessageReadDao() *ImMessageReadDao {
	return &ImMessageReadDao{
		group:   "default",
		table:   "hg_im_message_read",
		columns: imMessageReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImMessageReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImMessageReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImMessageReadDao) Columns() ImMessageReadColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImMessageReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImMessageReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ImMessageReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
