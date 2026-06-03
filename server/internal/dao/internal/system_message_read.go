// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemMessageReadDao is the data access object for the table hg_system_message_read.
type SystemMessageReadDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns SystemMessageReadColumns // columns contains all the column names of Table for convenient usage.
}

// SystemMessageReadColumns defines and stores column names for the table hg_system_message_read.
type SystemMessageReadColumns struct {
	Id        string // 记录ID
	MessageId string // 消息ID
	MemberId  string // 会员ID
	CreatedAt string // 已读时间
}

// systemMessageReadColumns holds the columns for the table hg_system_message_read.
var systemMessageReadColumns = SystemMessageReadColumns{
	Id:        "id",
	MessageId: "message_id",
	MemberId:  "member_id",
	CreatedAt: "created_at",
}

// NewSystemMessageReadDao creates and returns a new DAO object for table data access.
func NewSystemMessageReadDao() *SystemMessageReadDao {
	return &SystemMessageReadDao{
		group:   "default",
		table:   "hg_system_message_read",
		columns: systemMessageReadColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemMessageReadDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemMessageReadDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemMessageReadDao) Columns() SystemMessageReadColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemMessageReadDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemMessageReadDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SystemMessageReadDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
