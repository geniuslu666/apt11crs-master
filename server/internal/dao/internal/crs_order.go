// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CrsOrderDao is the data access object for the table hg_crs_order.
type CrsOrderDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CrsOrderColumns // columns contains all the column names of Table for convenient usage.
}

// CrsOrderColumns defines and stores column names for the table hg_crs_order.
type CrsOrderColumns struct {
	Id           string // 主键
	OrderSn      string // 订单号
	OrderAmount  string // 订单金额
	OutOrderSn   string // 外部订单号
	OrderStatus  string // 订单状态
	RefundStatus string // 退款状态
	RefundAmount string // 退款金额
	ExpiredTime  string // 过期时间
	CreatedAt    string // 创建时间
	UpdatedAt    string // 修改时间
}

// crsOrderColumns holds the columns for the table hg_crs_order.
var crsOrderColumns = CrsOrderColumns{
	Id:           "id",
	OrderSn:      "order_sn",
	OrderAmount:  "order_amount",
	OutOrderSn:   "out_order_sn",
	OrderStatus:  "order_status",
	RefundStatus: "refund_status",
	RefundAmount: "refund_amount",
	ExpiredTime:  "expired_time",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewCrsOrderDao creates and returns a new DAO object for table data access.
func NewCrsOrderDao() *CrsOrderDao {
	return &CrsOrderDao{
		group:   "default",
		table:   "hg_crs_order",
		columns: crsOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CrsOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CrsOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CrsOrderDao) Columns() CrsOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CrsOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CrsOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CrsOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
