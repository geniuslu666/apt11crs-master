// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderRefundLogDao is the data access object for the table hg_order_refund_log.
type OrderRefundLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns OrderRefundLogColumns // columns contains all the column names of Table for convenient usage.
}

// OrderRefundLogColumns defines and stores column names for the table hg_order_refund_log.
type OrderRefundLogColumns struct {
	Id           string // 主键
	OrderSn      string // 订单号
	Scene        string // 场景值
	RefundType   string // 退款方式 BAL-积分退款 AMOUNT-金额退款
	RefundAmount string // 退款金额
	RefundTime   string // 退款时间
	RefundStatus string // 退款状态
	OperateType  string // 操作员类型
	OperateId    string // 操作员ID
	Remark       string // 备注
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// orderRefundLogColumns holds the columns for the table hg_order_refund_log.
var orderRefundLogColumns = OrderRefundLogColumns{
	Id:           "id",
	OrderSn:      "order_sn",
	Scene:        "scene",
	RefundType:   "refund_type",
	RefundAmount: "refund_amount",
	RefundTime:   "refund_time",
	RefundStatus: "refund_status",
	OperateType:  "operate_type",
	OperateId:    "operate_id",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewOrderRefundLogDao creates and returns a new DAO object for table data access.
func NewOrderRefundLogDao() *OrderRefundLogDao {
	return &OrderRefundLogDao{
		group:   "default",
		table:   "hg_order_refund_log",
		columns: orderRefundLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderRefundLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderRefundLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderRefundLogDao) Columns() OrderRefundLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderRefundLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderRefundLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderRefundLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
