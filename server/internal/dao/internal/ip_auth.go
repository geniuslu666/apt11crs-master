// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IpAuthDao is the data access object for the table ip_auth.
type IpAuthDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns IpAuthColumns // columns contains all the column names of Table for convenient usage.
}

// IpAuthColumns defines and stores column names for the table ip_auth.
type IpAuthColumns struct {
	Id         string //
	Content    string // 备注
	IpAddress  string // IP地址
	ExpireTime string // 过期时间，未启用
	Phone      string // 客服账户手机号
	CreatedAt  string // 创建时间
	Status     string // 开启状态，1正常，2关闭
	Code       string // 授权码
}

// ipAuthColumns holds the columns for the table ip_auth.
var ipAuthColumns = IpAuthColumns{
	Id:         "id",
	Content:    "content",
	IpAddress:  "ip_address",
	ExpireTime: "expire_time",
	Phone:      "phone",
	CreatedAt:  "created_at",
	Status:     "status",
	Code:       "code",
}

// NewIpAuthDao creates and returns a new DAO object for table data access.
func NewIpAuthDao() *IpAuthDao {
	return &IpAuthDao{
		group:   "default",
		table:   "ip_auth",
		columns: ipAuthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IpAuthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IpAuthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IpAuthDao) Columns() IpAuthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IpAuthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IpAuthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *IpAuthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
