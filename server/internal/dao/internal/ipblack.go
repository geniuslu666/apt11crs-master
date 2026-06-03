// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IpblackDao is the data access object for the table ipblack.
type IpblackDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns IpblackColumns // columns contains all the column names of Table for convenient usage.
}

// IpblackColumns defines and stores column names for the table ipblack.
type IpblackColumns struct {
	Id       string //
	Ip       string // IP地址
	Name     string // 名称
	CreateAt string // 创建时间
	KefuId   string // 操作的客服账户
	EntId    string // 客服企业ID
}

// ipblackColumns holds the columns for the table ipblack.
var ipblackColumns = IpblackColumns{
	Id:       "id",
	Ip:       "ip",
	Name:     "name",
	CreateAt: "create_at",
	KefuId:   "kefu_id",
	EntId:    "ent_id",
}

// NewIpblackDao creates and returns a new DAO object for table data access.
func NewIpblackDao() *IpblackDao {
	return &IpblackDao{
		group:   "default",
		table:   "ipblack",
		columns: ipblackColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IpblackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IpblackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IpblackDao) Columns() IpblackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IpblackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IpblackDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *IpblackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
