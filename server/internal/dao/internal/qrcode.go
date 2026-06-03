// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QrcodeDao is the data access object for the table qrcode.
type QrcodeDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns QrcodeColumns // columns contains all the column names of Table for convenient usage.
}

// QrcodeColumns defines and stores column names for the table qrcode.
type QrcodeColumns struct {
	Id        string //
	EntId     string // 客服企业ID
	KefuName  string // 客服账户
	Uuid      string // 唯一ID
	Url       string // 跳转的URL
	CreatedAt string // 创建时间
}

// qrcodeColumns holds the columns for the table qrcode.
var qrcodeColumns = QrcodeColumns{
	Id:        "id",
	EntId:     "ent_id",
	KefuName:  "kefu_name",
	Uuid:      "uuid",
	Url:       "url",
	CreatedAt: "created_at",
}

// NewQrcodeDao creates and returns a new DAO object for table data access.
func NewQrcodeDao() *QrcodeDao {
	return &QrcodeDao{
		group:   "default",
		table:   "qrcode",
		columns: qrcodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *QrcodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *QrcodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *QrcodeDao) Columns() QrcodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *QrcodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *QrcodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *QrcodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
