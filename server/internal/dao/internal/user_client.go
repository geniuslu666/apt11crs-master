// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserClientDao is the data access object for the table user_client.
type UserClientDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns UserClientColumns // columns contains all the column names of Table for convenient usage.
}

// UserClientColumns defines and stores column names for the table user_client.
type UserClientColumns struct {
	Id        string //
	Kefu      string // 客服账户
	ClientId  string // 设备ID
	CreatedAt string // 创建时间
}

// userClientColumns holds the columns for the table user_client.
var userClientColumns = UserClientColumns{
	Id:        "id",
	Kefu:      "kefu",
	ClientId:  "client_id",
	CreatedAt: "created_at",
}

// NewUserClientDao creates and returns a new DAO object for table data access.
func NewUserClientDao() *UserClientDao {
	return &UserClientDao{
		group:   "default",
		table:   "user_client",
		columns: userClientColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserClientDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserClientDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserClientDao) Columns() UserClientColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserClientDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserClientDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserClientDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
