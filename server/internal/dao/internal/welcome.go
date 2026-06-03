// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WelcomeDao is the data access object for the table welcome.
type WelcomeDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns WelcomeColumns // columns contains all the column names of Table for convenient usage.
}

// WelcomeColumns defines and stores column names for the table welcome.
type WelcomeColumns struct {
	Id          string //
	UserId      string // 客服账户
	Keyword     string // 关键词，welcome为默认欢迎；wechat为公众号欢迎语
	Content     string // 欢迎消息内容
	IsDefault   string // 是否默认，未启用
	DelaySecond string // 延迟秒数
	Ctime       string // 创建时间
}

// welcomeColumns holds the columns for the table welcome.
var welcomeColumns = WelcomeColumns{
	Id:          "id",
	UserId:      "user_id",
	Keyword:     "keyword",
	Content:     "content",
	IsDefault:   "is_default",
	DelaySecond: "delay_second",
	Ctime:       "ctime",
}

// NewWelcomeDao creates and returns a new DAO object for table data access.
func NewWelcomeDao() *WelcomeDao {
	return &WelcomeDao{
		group:   "default",
		table:   "welcome",
		columns: welcomeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WelcomeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WelcomeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WelcomeDao) Columns() WelcomeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WelcomeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WelcomeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WelcomeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
