// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsCateDao is the data access object for the table cms_cate.
type CmsCateDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns CmsCateColumns // columns contains all the column names of Table for convenient usage.
}

// CmsCateColumns defines and stores column names for the table cms_cate.
type CmsCateColumns struct {
	Id        string //
	CatName   string // 分类名称
	CreatedAt string // 创建时间
}

// cmsCateColumns holds the columns for the table cms_cate.
var cmsCateColumns = CmsCateColumns{
	Id:        "id",
	CatName:   "cat_name",
	CreatedAt: "created_at",
}

// NewCmsCateDao creates and returns a new DAO object for table data access.
func NewCmsCateDao() *CmsCateDao {
	return &CmsCateDao{
		group:   "default",
		table:   "cms_cate",
		columns: cmsCateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CmsCateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CmsCateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CmsCateDao) Columns() CmsCateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CmsCateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CmsCateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CmsCateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
