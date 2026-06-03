// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyGroupDao is the data access object for the table reply_group.
type ReplyGroupDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns ReplyGroupColumns // columns contains all the column names of Table for convenient usage.
}

// ReplyGroupColumns defines and stores column names for the table reply_group.
type ReplyGroupColumns struct {
	Id        string //
	GroupName string // 组名
	UserId    string // 客服账户
	EntId     string // 客服企业ID
	IsTeam    string // 1个人,2团队
}

// replyGroupColumns holds the columns for the table reply_group.
var replyGroupColumns = ReplyGroupColumns{
	Id:        "id",
	GroupName: "group_name",
	UserId:    "user_id",
	EntId:     "ent_id",
	IsTeam:    "is_team",
}

// NewReplyGroupDao creates and returns a new DAO object for table data access.
func NewReplyGroupDao() *ReplyGroupDao {
	return &ReplyGroupDao{
		group:   "default",
		table:   "reply_group",
		columns: replyGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ReplyGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ReplyGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ReplyGroupDao) Columns() ReplyGroupColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ReplyGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ReplyGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ReplyGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
