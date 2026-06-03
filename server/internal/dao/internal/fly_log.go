// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FlyLogDao is the data access object for the table fly_log.
type FlyLogDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns FlyLogColumns // columns contains all the column names of Table for convenient usage.
}

// FlyLogColumns defines and stores column names for the table fly_log.
type FlyLogColumns struct {
	Id         string //
	EntId      string // 企业ID
	LogType    string // 日志类型
	LogContent string // 日志内容
	IpAddress  string // 用户IP地址
	CreatedAt  string // 创建时间
}

// flyLogColumns holds the columns for the table fly_log.
var flyLogColumns = FlyLogColumns{
	Id:         "id",
	EntId:      "ent_id",
	LogType:    "log_type",
	LogContent: "log_content",
	IpAddress:  "ip_address",
	CreatedAt:  "created_at",
}

// NewFlyLogDao creates and returns a new DAO object for table data access.
func NewFlyLogDao() *FlyLogDao {
	return &FlyLogDao{
		group:   "default",
		table:   "fly_log",
		columns: flyLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FlyLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FlyLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FlyLogDao) Columns() FlyLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FlyLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FlyLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FlyLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
