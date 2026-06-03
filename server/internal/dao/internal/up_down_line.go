// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UpDownLineDao is the data access object for the table up_down_line.
type UpDownLineDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns UpDownLineColumns // columns contains all the column names of Table for convenient usage.
}

// UpDownLineColumns defines and stores column names for the table up_down_line.
type UpDownLineColumns struct {
	Id           string //
	KefuName     string // 客服账户
	EntId        string // 企业ID
	OnlineStatus string // 在线状态，1在线，2离线
	ClientIp     string // ip地址
	CreatedAt    string // 创建时间
}

// upDownLineColumns holds the columns for the table up_down_line.
var upDownLineColumns = UpDownLineColumns{
	Id:           "id",
	KefuName:     "kefu_name",
	EntId:        "ent_id",
	OnlineStatus: "online_status",
	ClientIp:     "client_ip",
	CreatedAt:    "created_at",
}

// NewUpDownLineDao creates and returns a new DAO object for table data access.
func NewUpDownLineDao() *UpDownLineDao {
	return &UpDownLineDao{
		group:   "default",
		table:   "up_down_line",
		columns: upDownLineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UpDownLineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UpDownLineDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UpDownLineDao) Columns() UpDownLineColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UpDownLineDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UpDownLineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UpDownLineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
