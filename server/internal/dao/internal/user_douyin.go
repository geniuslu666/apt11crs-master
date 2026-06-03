// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDouyinDao is the data access object for the table user_douyin.
type UserDouyinDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns UserDouyinColumns // columns contains all the column names of Table for convenient usage.
}

// UserDouyinColumns defines and stores column names for the table user_douyin.
type UserDouyinColumns struct {
	Id                 string //
	KefuName           string // 客服账户
	Nickname           string // 抖音昵称
	Avatar             string // 抖音头像
	OpenId             string // 抖音OpenId
	UnionId            string // 抖音union_id
	AccessToken        string // 抖音AccessToken
	ExpiresIn          string // 抖音AccessToken过期时间
	RefreshToken       string // 抖音refresh_token
	RefreshExpiresIn   string // 抖音refresh_token过期时间
	ClientToken        string // 抖音client_token
	ClientTokenExpires string // 抖音client_token过期时间
	EntId              string // 企业ID
	CreatedAt          string // 创建时间
}

// userDouyinColumns holds the columns for the table user_douyin.
var userDouyinColumns = UserDouyinColumns{
	Id:                 "id",
	KefuName:           "kefu_name",
	Nickname:           "nickname",
	Avatar:             "avatar",
	OpenId:             "open_id",
	UnionId:            "union_id",
	AccessToken:        "access_token",
	ExpiresIn:          "expires_in",
	RefreshToken:       "refresh_token",
	RefreshExpiresIn:   "refresh_expires_in",
	ClientToken:        "client_token",
	ClientTokenExpires: "client_token_expires",
	EntId:              "ent_id",
	CreatedAt:          "created_at",
}

// NewUserDouyinDao creates and returns a new DAO object for table data access.
func NewUserDouyinDao() *UserDouyinDao {
	return &UserDouyinDao{
		group:   "default",
		table:   "user_douyin",
		columns: userDouyinColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDouyinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDouyinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDouyinDao) Columns() UserDouyinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDouyinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDouyinDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserDouyinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
