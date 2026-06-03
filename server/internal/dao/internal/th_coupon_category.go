// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThCouponCategoryDao is the data access object for the table hg_th_coupon_category.
type ThCouponCategoryDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns ThCouponCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// ThCouponCategoryColumns defines and stores column names for the table hg_th_coupon_category.
type ThCouponCategoryColumns struct {
	Id        string //
	Name      string // 分类名称
	Sort      string // 排序(越大越靠前)
	Status    string // 1、启用 2、禁用
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// thCouponCategoryColumns holds the columns for the table hg_th_coupon_category.
var thCouponCategoryColumns = ThCouponCategoryColumns{
	Id:        "id",
	Name:      "name",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewThCouponCategoryDao creates and returns a new DAO object for table data access.
func NewThCouponCategoryDao() *ThCouponCategoryDao {
	return &ThCouponCategoryDao{
		group:   "default",
		table:   "hg_th_coupon_category",
		columns: thCouponCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThCouponCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThCouponCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThCouponCategoryDao) Columns() ThCouponCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThCouponCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThCouponCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThCouponCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
