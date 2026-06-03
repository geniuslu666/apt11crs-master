// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsHelpcenterDao is the data access object for the table hg_pms_helpcenter.
type PmsHelpcenterDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsHelpcenterColumns // columns contains all the column names of Table for convenient usage.
}

// PmsHelpcenterColumns defines and stores column names for the table hg_pms_helpcenter.
type PmsHelpcenterColumns struct {
	Id         string //
	Language   string //
	CategoryId string // 分类ID
	Title      string // 标题
	Content    string // 内容
	Sort       string // 排序(越大越靠前)
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

// pmsHelpcenterColumns holds the columns for the table hg_pms_helpcenter.
var pmsHelpcenterColumns = PmsHelpcenterColumns{
	Id:         "id",
	Language:   "language",
	CategoryId: "category_id",
	Title:      "title",
	Content:    "content",
	Sort:       "sort",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewPmsHelpcenterDao creates and returns a new DAO object for table data access.
func NewPmsHelpcenterDao() *PmsHelpcenterDao {
	return &PmsHelpcenterDao{
		group:   "default",
		table:   "hg_pms_helpcenter",
		columns: pmsHelpcenterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsHelpcenterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsHelpcenterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsHelpcenterDao) Columns() PmsHelpcenterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsHelpcenterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsHelpcenterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsHelpcenterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
