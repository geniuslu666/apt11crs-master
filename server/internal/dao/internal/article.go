// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleDao is the data access object for the table article.
type ArticleDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns ArticleColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleColumns defines and stores column names for the table article.
type ArticleColumns struct {
	Id         string //
	Title      string // 自动回复关键词
	Content    string // 自动回复内容
	CatId      string // 自动回复分类ID
	UserId     string // 客服账户
	EntId      string // 客服企业ID
	ApiUrl     string // 第三方接口地址
	SearchType string // 1包含匹配,2精准匹配
	Score      string // 命中次数
}

// articleColumns holds the columns for the table article.
var articleColumns = ArticleColumns{
	Id:         "id",
	Title:      "title",
	Content:    "content",
	CatId:      "cat_id",
	UserId:     "user_id",
	EntId:      "ent_id",
	ApiUrl:     "api_url",
	SearchType: "search_type",
	Score:      "score",
}

// NewArticleDao creates and returns a new DAO object for table data access.
func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		group:   "default",
		table:   "article",
		columns: articleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ArticleDao) Columns() ArticleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ArticleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
