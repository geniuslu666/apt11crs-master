// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServicePropertyDao is the data access object for the table hg_spa_service_property.
type SpaServicePropertyDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns SpaServicePropertyColumns // columns contains all the column names of Table for convenient usage.
}

// SpaServicePropertyColumns defines and stores column names for the table hg_spa_service_property.
type SpaServicePropertyColumns struct {
	ServiceId  string // 服务ID
	PropertyId string // 物业ID
}

// spaServicePropertyColumns holds the columns for the table hg_spa_service_property.
var spaServicePropertyColumns = SpaServicePropertyColumns{
	ServiceId:  "service_id",
	PropertyId: "property_id",
}

// NewSpaServicePropertyDao creates and returns a new DAO object for table data access.
func NewSpaServicePropertyDao() *SpaServicePropertyDao {
	return &SpaServicePropertyDao{
		group:   "default",
		table:   "hg_spa_service_property",
		columns: spaServicePropertyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaServicePropertyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaServicePropertyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaServicePropertyDao) Columns() SpaServicePropertyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaServicePropertyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaServicePropertyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaServicePropertyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
