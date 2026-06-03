// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAppconfigDao is the data access object for the table hg_pms_appconfig.
type PmsAppconfigDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsAppconfigColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAppconfigColumns defines and stores column names for the table hg_pms_appconfig.
type PmsAppconfigColumns struct {
	Id        string //
	Name      string // 配置名称
	Key       string // 配置项
	Value     string // 配置值
	Language  string // 语言
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// pmsAppconfigColumns holds the columns for the table hg_pms_appconfig.
var pmsAppconfigColumns = PmsAppconfigColumns{
	Id:        "id",
	Name:      "name",
	Key:       "key",
	Value:     "value",
	Language:  "language",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPmsAppconfigDao creates and returns a new DAO object for table data access.
func NewPmsAppconfigDao() *PmsAppconfigDao {
	return &PmsAppconfigDao{
		group:   "default",
		table:   "hg_pms_appconfig",
		columns: pmsAppconfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsAppconfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsAppconfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsAppconfigDao) Columns() PmsAppconfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsAppconfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsAppconfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsAppconfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
