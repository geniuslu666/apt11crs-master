// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReplyItemDao is the data access object for the table reply_item.
type ReplyItemDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns ReplyItemColumns // columns contains all the column names of Table for convenient usage.
}

// ReplyItemColumns defines and stores column names for the table reply_item.
type ReplyItemColumns struct {
	Id       string //
	Content  string // 快捷回复内容
	GroupId  string // 快捷回复分组ID
	UserId   string // 客服账户
	ItemName string // 快捷回复标题
	EntId    string // 客服企业ID
	IsTeam   string // 1个人,2团队
}

// replyItemColumns holds the columns for the table reply_item.
var replyItemColumns = ReplyItemColumns{
	Id:       "id",
	Content:  "content",
	GroupId:  "group_id",
	UserId:   "user_id",
	ItemName: "item_name",
	EntId:    "ent_id",
	IsTeam:   "is_team",
}

// NewReplyItemDao creates and returns a new DAO object for table data access.
func NewReplyItemDao() *ReplyItemDao {
	return &ReplyItemDao{
		group:   "default",
		table:   "reply_item",
		columns: replyItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ReplyItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ReplyItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ReplyItemDao) Columns() ReplyItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ReplyItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ReplyItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ReplyItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
