// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomepageArticleThumbDao is the data access object for the table hg_homepage_article_thumb.
type HomepageArticleThumbDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns HomepageArticleThumbColumns // columns contains all the column names of Table for convenient usage.
}

// HomepageArticleThumbColumns defines and stores column names for the table hg_homepage_article_thumb.
type HomepageArticleThumbColumns struct {
	Id        string // 关联ID
	ArticleId string // 活动ID
	MemberId  string // 会员ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// homepageArticleThumbColumns holds the columns for the table hg_homepage_article_thumb.
var homepageArticleThumbColumns = HomepageArticleThumbColumns{
	Id:        "id",
	ArticleId: "article_id",
	MemberId:  "member_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewHomepageArticleThumbDao creates and returns a new DAO object for table data access.
func NewHomepageArticleThumbDao() *HomepageArticleThumbDao {
	return &HomepageArticleThumbDao{
		group:   "default",
		table:   "hg_homepage_article_thumb",
		columns: homepageArticleThumbColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HomepageArticleThumbDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HomepageArticleThumbDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HomepageArticleThumbDao) Columns() HomepageArticleThumbColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HomepageArticleThumbDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HomepageArticleThumbDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *HomepageArticleThumbDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
