// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageDao is the data access object for the table message.
type MessageDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns MessageColumns // columns contains all the column names of Table for convenient usage.
}

// MessageColumns defines and stores column names for the table message.
type MessageColumns struct {
	Id        string //
	KefuId    string // 客服账户
	VisitorId string // 访客ID
	Content   string // 消息内容
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	MesType   string // 消息类型：kefu为客服发送，visitor为访客发送
	Status    string // 已读状态：read为已读，unread为未读
	EntId     string // 客服企业ID
}

// messageColumns holds the columns for the table message.
var messageColumns = MessageColumns{
	Id:        "id",
	KefuId:    "kefu_id",
	VisitorId: "visitor_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	MesType:   "mes_type",
	Status:    "status",
	EntId:     "ent_id",
}

// NewMessageDao creates and returns a new DAO object for table data access.
func NewMessageDao() *MessageDao {
	return &MessageDao{
		group:   "default",
		table:   "message",
		columns: messageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MessageDao) Columns() MessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
