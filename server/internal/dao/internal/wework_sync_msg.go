// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WeworkSyncMsgDao is the data access object for the table wework_sync_msg.
type WeworkSyncMsgDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns WeworkSyncMsgColumns // columns contains all the column names of Table for convenient usage.
}

// WeworkSyncMsgColumns defines and stores column names for the table wework_sync_msg.
type WeworkSyncMsgColumns struct {
	Id         string //
	SyncCursor string //
	JsonTxt    string //
	VisitorId  string //
	KefuId     string //
	CreatedAt  string //
	EntId      string //
}

// weworkSyncMsgColumns holds the columns for the table wework_sync_msg.
var weworkSyncMsgColumns = WeworkSyncMsgColumns{
	Id:         "id",
	SyncCursor: "sync_cursor",
	JsonTxt:    "json_txt",
	VisitorId:  "visitor_id",
	KefuId:     "kefu_id",
	CreatedAt:  "created_at",
	EntId:      "ent_id",
}

// NewWeworkSyncMsgDao creates and returns a new DAO object for table data access.
func NewWeworkSyncMsgDao() *WeworkSyncMsgDao {
	return &WeworkSyncMsgDao{
		group:   "default",
		table:   "wework_sync_msg",
		columns: weworkSyncMsgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WeworkSyncMsgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WeworkSyncMsgDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WeworkSyncMsgDao) Columns() WeworkSyncMsgColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WeworkSyncMsgDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WeworkSyncMsgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WeworkSyncMsgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
