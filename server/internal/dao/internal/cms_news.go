// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsNewsDao is the data access object for the table cms_news.
type CmsNewsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns CmsNewsColumns // columns contains all the column names of Table for convenient usage.
}

// CmsNewsColumns defines and stores column names for the table cms_news.
type CmsNewsColumns struct {
	Id        string //
	Title     string // 标题
	Content   string // 内容
	CatId     string // 分类ID
	CreatedAt string // 创建时间
}

// cmsNewsColumns holds the columns for the table cms_news.
var cmsNewsColumns = CmsNewsColumns{
	Id:        "id",
	Title:     "title",
	Content:   "content",
	CatId:     "cat_id",
	CreatedAt: "created_at",
}

// NewCmsNewsDao creates and returns a new DAO object for table data access.
func NewCmsNewsDao() *CmsNewsDao {
	return &CmsNewsDao{
		group:   "default",
		table:   "cms_news",
		columns: cmsNewsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CmsNewsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CmsNewsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CmsNewsDao) Columns() CmsNewsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CmsNewsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CmsNewsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CmsNewsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
