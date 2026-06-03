// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AigcSessionMessageDao is the data access object for the table aigc_session_message.
type AigcSessionMessageDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns AigcSessionMessageColumns // columns contains all the column names of Table for convenient usage.
}

// AigcSessionMessageColumns defines and stores column names for the table aigc_session_message.
type AigcSessionMessageColumns struct {
	Id         string //
	CollectId  string // 集合ID
	CreatedAt  string // 创建时间
	KefuAvatar string // 客服头像
	AiAvatar   string // AI头像
	Content    string // 内容
	KefuName   string // 客服名称
	EntId      string // 企业ID
	MsgType    string // 消息类型
}

// aigcSessionMessageColumns holds the columns for the table aigc_session_message.
var aigcSessionMessageColumns = AigcSessionMessageColumns{
	Id:         "id",
	CollectId:  "collect_id",
	CreatedAt:  "created_at",
	KefuAvatar: "kefu_avatar",
	AiAvatar:   "ai_avatar",
	Content:    "content",
	KefuName:   "kefu_name",
	EntId:      "ent_id",
	MsgType:    "msg_type",
}

// NewAigcSessionMessageDao creates and returns a new DAO object for table data access.
func NewAigcSessionMessageDao() *AigcSessionMessageDao {
	return &AigcSessionMessageDao{
		group:   "default",
		table:   "aigc_session_message",
		columns: aigcSessionMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AigcSessionMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AigcSessionMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AigcSessionMessageDao) Columns() AigcSessionMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AigcSessionMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AigcSessionMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AigcSessionMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
