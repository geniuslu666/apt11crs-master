// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTerminalModelDao is the data access object for the table hg_sys_terminal_model.
type SysTerminalModelDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns SysTerminalModelColumns // columns contains all the column names of Table for convenient usage.
}

// SysTerminalModelColumns defines and stores column names for the table hg_sys_terminal_model.
type SysTerminalModelColumns struct {
	Id           string //
	BrandModel   string // 品牌型号
	ClientId     string // 开发者ID
	ClientSecret string // 开发者秘钥
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
	DeletedAt    string //
}

// sysTerminalModelColumns holds the columns for the table hg_sys_terminal_model.
var sysTerminalModelColumns = SysTerminalModelColumns{
	Id:           "id",
	BrandModel:   "brand_model",
	ClientId:     "client_id",
	ClientSecret: "client_secret",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeletedAt:    "deleted_at",
}

// NewSysTerminalModelDao creates and returns a new DAO object for table data access.
func NewSysTerminalModelDao() *SysTerminalModelDao {
	return &SysTerminalModelDao{
		group:   "default",
		table:   "hg_sys_terminal_model",
		columns: sysTerminalModelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTerminalModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTerminalModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTerminalModelDao) Columns() SysTerminalModelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTerminalModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTerminalModelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysTerminalModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
