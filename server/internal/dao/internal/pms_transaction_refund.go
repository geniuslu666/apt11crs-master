// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsTransactionRefundDao is the data access object for the table hg_pms_transaction_refund.
type PmsTransactionRefundDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns PmsTransactionRefundColumns // columns contains all the column names of Table for convenient usage.
}

// PmsTransactionRefundColumns defines and stores column names for the table hg_pms_transaction_refund.
type PmsTransactionRefundColumns struct {
	Id            string // 主键
	OrderSn       string // 订单号
	ChangeOrderSn string // 变更订单号
	TransactionSn string // 支付流水号
	Scene         string // 场景值
	RefundChannel string // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
	RefundType    string // '支付方式   BAL 余额'
	RefundSn      string // 退款流水号
	TransNo       string // 退款交易号
	CancelOrderSn string // 取消订单号
	RefundAmount  string // 退款金额
	RefundTime    string // 退款时间
	RefundStatus  string // 退款状态
	CancelId      string // 取消政策ID
	OperateType   string // 操作员类型
	OperateId     string // 操作员ID
	Remark        string // 备注
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	IsFx          string // 是否是分销订单
}

// pmsTransactionRefundColumns holds the columns for the table hg_pms_transaction_refund.
var pmsTransactionRefundColumns = PmsTransactionRefundColumns{
	Id:            "id",
	OrderSn:       "order_sn",
	ChangeOrderSn: "change_order_sn",
	TransactionSn: "transaction_sn",
	Scene:         "scene",
	RefundChannel: "refund_channel",
	RefundType:    "refund_type",
	RefundSn:      "refund_sn",
	TransNo:       "trans_no",
	CancelOrderSn: "cancel_order_sn",
	RefundAmount:  "refund_amount",
	RefundTime:    "refund_time",
	RefundStatus:  "refund_status",
	CancelId:      "cancel_id",
	OperateType:   "operate_type",
	OperateId:     "operate_id",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	IsFx:          "is_fx",
}

// NewPmsTransactionRefundDao creates and returns a new DAO object for table data access.
func NewPmsTransactionRefundDao() *PmsTransactionRefundDao {
	return &PmsTransactionRefundDao{
		group:   "default",
		table:   "hg_pms_transaction_refund",
		columns: pmsTransactionRefundColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsTransactionRefundDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsTransactionRefundDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsTransactionRefundDao) Columns() PmsTransactionRefundColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsTransactionRefundDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsTransactionRefundDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsTransactionRefundDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
