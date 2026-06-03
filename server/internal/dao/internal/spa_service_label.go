// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServiceLabelDao is the data access object for the table hg_spa_service_label.
type SpaServiceLabelDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SpaServiceLabelColumns // columns contains all the column names of Table for convenient usage.
}

// SpaServiceLabelColumns defines and stores column names for the table hg_spa_service_label.
type SpaServiceLabelColumns struct {
	ServiceId string // 服务ID
	LabelId   string // 标签ID
}

// spaServiceLabelColumns holds the columns for the table hg_spa_service_label.
var spaServiceLabelColumns = SpaServiceLabelColumns{
	ServiceId: "service_id",
	LabelId:   "label_id",
}

// NewSpaServiceLabelDao creates and returns a new DAO object for table data access.
func NewSpaServiceLabelDao() *SpaServiceLabelDao {
	return &SpaServiceLabelDao{
		group:   "default",
		table:   "hg_spa_service_label",
		columns: spaServiceLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaServiceLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaServiceLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaServiceLabelDao) Columns() SpaServiceLabelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaServiceLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaServiceLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaServiceLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
