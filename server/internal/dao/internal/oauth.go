// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthDao is the data access object for the table oauth.
type OauthDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns OauthColumns // columns contains all the column names of Table for convenient usage.
}

// OauthColumns defines and stores column names for the table oauth.
type OauthColumns struct {
	Id        string //
	UserId    string // 访客/客服账户
	OauthId   string // 公众号OPENID
	CreatedAt string // 创建时间
	Status    string // 状态，未启用
}

// oauthColumns holds the columns for the table oauth.
var oauthColumns = OauthColumns{
	Id:        "id",
	UserId:    "user_id",
	OauthId:   "oauth_id",
	CreatedAt: "created_at",
	Status:    "status",
}

// NewOauthDao creates and returns a new DAO object for table data access.
func NewOauthDao() *OauthDao {
	return &OauthDao{
		group:   "default",
		table:   "oauth",
		columns: oauthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthDao) Columns() OauthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OauthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
