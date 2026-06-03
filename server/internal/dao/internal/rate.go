// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RateDao is the data access object for the table rate.
type RateDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns RateColumns // columns contains all the column names of Table for convenient usage.
}

// RateColumns defines and stores column names for the table rate.
type RateColumns struct {
	Id           string //
	KefuName     string // 客服账户
	KefuNickname string // 客服昵称
	VisitorId    string // 访客id
	Content      string // 评价内容
	EntId        string // 企业ID
	Score        string // 评价分数
	CreatedAt    string // 创建时间
}

// rateColumns holds the columns for the table rate.
var rateColumns = RateColumns{
	Id:           "id",
	KefuName:     "kefu_name",
	KefuNickname: "kefu_nickname",
	VisitorId:    "visitor_id",
	Content:      "content",
	EntId:        "ent_id",
	Score:        "score",
	CreatedAt:    "created_at",
}

// NewRateDao creates and returns a new DAO object for table data access.
func NewRateDao() *RateDao {
	return &RateDao{
		group:   "default",
		table:   "rate",
		columns: rateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RateDao) Columns() RateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
