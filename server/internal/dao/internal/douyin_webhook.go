// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DouyinWebhookDao is the data access object for the table douyin_webhook.
type DouyinWebhookDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns DouyinWebhookColumns // columns contains all the column names of Table for convenient usage.
}

// DouyinWebhookColumns defines and stores column names for the table douyin_webhook.
type DouyinWebhookColumns struct {
	Id         string //
	KefuName   string // 客服账户
	Event      string // event
	FromUserId string // from_user_id
	ToUserId   string // to_user_id
	ClientKey  string // client_key
	Content    string // content
	EntId      string // 企业ID
	CreatedAt  string // 创建时间
}

// douyinWebhookColumns holds the columns for the table douyin_webhook.
var douyinWebhookColumns = DouyinWebhookColumns{
	Id:         "id",
	KefuName:   "kefu_name",
	Event:      "event",
	FromUserId: "from_user_id",
	ToUserId:   "to_user_id",
	ClientKey:  "client_key",
	Content:    "content",
	EntId:      "ent_id",
	CreatedAt:  "created_at",
}

// NewDouyinWebhookDao creates and returns a new DAO object for table data access.
func NewDouyinWebhookDao() *DouyinWebhookDao {
	return &DouyinWebhookDao{
		group:   "default",
		table:   "douyin_webhook",
		columns: douyinWebhookColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DouyinWebhookDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DouyinWebhookDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DouyinWebhookDao) Columns() DouyinWebhookColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DouyinWebhookDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DouyinWebhookDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DouyinWebhookDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
