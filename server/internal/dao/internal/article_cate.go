// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleCateDao is the data access object for the table article_cate.
type ArticleCateDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns ArticleCateColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleCateColumns defines and stores column names for the table article_cate.
type ArticleCateColumns struct {
	Id      string //
	CatName string // 自动回复分类名称
	UserId  string // 客服账户
	EntId   string // 客服企业ID
	IsTop   string // 是否置顶展示，1置顶
}

// articleCateColumns holds the columns for the table article_cate.
var articleCateColumns = ArticleCateColumns{
	Id:      "id",
	CatName: "cat_name",
	UserId:  "user_id",
	EntId:   "ent_id",
	IsTop:   "is_top",
}

// NewArticleCateDao creates and returns a new DAO object for table data access.
func NewArticleCateDao() *ArticleCateDao {
	return &ArticleCateDao{
		group:   "default",
		table:   "article_cate",
		columns: articleCateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ArticleCateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ArticleCateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ArticleCateDao) Columns() ArticleCateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ArticleCateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ArticleCateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ArticleCateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
