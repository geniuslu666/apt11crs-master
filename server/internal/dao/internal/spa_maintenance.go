// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaMaintenanceDao is the data access object for the table hg_spa_maintenance.
type SpaMaintenanceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SpaMaintenanceColumns // columns contains all the column names of Table for convenient usage.
}

// SpaMaintenanceColumns defines and stores column names for the table hg_spa_maintenance.
type SpaMaintenanceColumns struct {
	Id        string //
	Language  string // 语言
	Content   string // 内容
	IsDefault string // 是否默认  1是 2否
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
}

// spaMaintenanceColumns holds the columns for the table hg_spa_maintenance.
var spaMaintenanceColumns = SpaMaintenanceColumns{
	Id:        "id",
	Language:  "language",
	Content:   "content",
	IsDefault: "is_default",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// NewSpaMaintenanceDao creates and returns a new DAO object for table data access.
func NewSpaMaintenanceDao() *SpaMaintenanceDao {
	return &SpaMaintenanceDao{
		group:   "default",
		table:   "hg_spa_maintenance",
		columns: spaMaintenanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaMaintenanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaMaintenanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaMaintenanceDao) Columns() SpaMaintenanceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaMaintenanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaMaintenanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaMaintenanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
