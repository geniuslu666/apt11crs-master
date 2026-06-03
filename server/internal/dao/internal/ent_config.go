// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EntConfigDao is the data access object for the table ent_config.
type EntConfigDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns EntConfigColumns // columns contains all the column names of Table for convenient usage.
}

// EntConfigColumns defines and stores column names for the table ent_config.
type EntConfigColumns struct {
	Id        string //
	ConfName  string // 配置描述
	ConfKey   string // 配置key
	ConfValue string // 配置值
	Language  string // 语言
	EntId     string // 客服企业ID
}

// entConfigColumns holds the columns for the table ent_config.
var entConfigColumns = EntConfigColumns{
	Id:        "id",
	ConfName:  "conf_name",
	ConfKey:   "conf_key",
	ConfValue: "conf_value",
	Language:  "language",
	EntId:     "ent_id",
}

// NewEntConfigDao creates and returns a new DAO object for table data access.
func NewEntConfigDao() *EntConfigDao {
	return &EntConfigDao{
		group:   "default",
		table:   "ent_config",
		columns: entConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EntConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EntConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EntConfigDao) Columns() EntConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EntConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EntConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EntConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
