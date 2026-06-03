// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsHelpcenterCategoryDao is the data access object for the table hg_pms_helpcenter_category.
type PmsHelpcenterCategoryDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of the current DAO.
	columns PmsHelpcenterCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// PmsHelpcenterCategoryColumns defines and stores column names for the table hg_pms_helpcenter_category.
type PmsHelpcenterCategoryColumns struct {
	Id        string //
	Name      string // 分类名称
	Language  string //
	Sort      string // 排序(越大越靠前)
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// pmsHelpcenterCategoryColumns holds the columns for the table hg_pms_helpcenter_category.
var pmsHelpcenterCategoryColumns = PmsHelpcenterCategoryColumns{
	Id:        "id",
	Name:      "name",
	Language:  "language",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPmsHelpcenterCategoryDao creates and returns a new DAO object for table data access.
func NewPmsHelpcenterCategoryDao() *PmsHelpcenterCategoryDao {
	return &PmsHelpcenterCategoryDao{
		group:   "default",
		table:   "hg_pms_helpcenter_category",
		columns: pmsHelpcenterCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsHelpcenterCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsHelpcenterCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsHelpcenterCategoryDao) Columns() PmsHelpcenterCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsHelpcenterCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsHelpcenterCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsHelpcenterCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
