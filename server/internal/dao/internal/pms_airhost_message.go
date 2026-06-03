// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAirhostMessageDao is the data access object for the table hg_pms_airhost_message.
type PmsAirhostMessageDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns PmsAirhostMessageColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAirhostMessageColumns defines and stores column names for the table hg_pms_airhost_message.
type PmsAirhostMessageColumns struct {
	Id               string //
	ObjectType       string // 消息来源类型
	MessageId        string // 消息 ID
	MessageEventCode string // 消息类型
	MessageContent   string // 消息内容
	IsHandle         string // 1、已处理  2、未处理
	CreatedAt        string //
	DeletedAt        string //
}

// pmsAirhostMessageColumns holds the columns for the table hg_pms_airhost_message.
var pmsAirhostMessageColumns = PmsAirhostMessageColumns{
	Id:               "id",
	ObjectType:       "object_type",
	MessageId:        "message_id",
	MessageEventCode: "message_event_code",
	MessageContent:   "message_content",
	IsHandle:         "is_handle",
	CreatedAt:        "created_at",
	DeletedAt:        "deleted_at",
}

// NewPmsAirhostMessageDao creates and returns a new DAO object for table data access.
func NewPmsAirhostMessageDao() *PmsAirhostMessageDao {
	return &PmsAirhostMessageDao{
		group:   "default",
		table:   "hg_pms_airhost_message",
		columns: pmsAirhostMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsAirhostMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsAirhostMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsAirhostMessageDao) Columns() PmsAirhostMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsAirhostMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsAirhostMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsAirhostMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
