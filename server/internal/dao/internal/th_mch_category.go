// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThMchCategoryDao is the data access object for the table hg_th_mch_category.
type ThMchCategoryDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns ThMchCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// ThMchCategoryColumns defines and stores column names for the table hg_th_mch_category.
type ThMchCategoryColumns struct {
	Id        string //
	Name      string // 分类名称
	Sort      string // 排序(越大越靠前)
	Status    string // 1、启用 2、禁用
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// thMchCategoryColumns holds the columns for the table hg_th_mch_category.
var thMchCategoryColumns = ThMchCategoryColumns{
	Id:        "id",
	Name:      "name",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewThMchCategoryDao creates and returns a new DAO object for table data access.
func NewThMchCategoryDao() *ThMchCategoryDao {
	return &ThMchCategoryDao{
		group:   "default",
		table:   "hg_th_mch_category",
		columns: thMchCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThMchCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThMchCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThMchCategoryDao) Columns() ThMchCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThMchCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThMchCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThMchCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
