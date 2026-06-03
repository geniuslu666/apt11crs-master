// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsTransactionDao is the data access object for the table hg_pms_transaction.
type PmsTransactionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns PmsTransactionColumns // columns contains all the column names of Table for convenient usage.
}

// PmsTransactionColumns defines and stores column names for the table hg_pms_transaction.
type PmsTransactionColumns struct {
	Id               string // 主键
	OrderSn          string // 订单号
	ChangeOrderSn    string // 变更订单号
	OrderType        string //
	TransactionSn    string // 支付流水号
	PaymentRequestId string // 第三方支付流水号
	CaptureId        string // paypal capture_id
	Scene            string // 场景值
	PayChannel       string // SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台
	PayType          string // 支付方式   BAL 余额
	OpenId           string // 用户ID
	Amount           string // 总金额
	PayParams        string // 支付参数
	PriceCurrency    string // 币种
	PayAmount        string // 支付金额
	PayCharge        string // 支付税率
	PayStatus        string // 支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付
	PayTime          string // 支付时间
	ExpiredTime      string // 过期时间
	RefundAmount     string // 退款金额
	RefundStatus     string // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	ScenePayRate     string // 场景积分抵扣比例
	Level            string // 等级
	ExchangeRate     string // 积分汇率
	CouponId         string // 优惠券ID
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	DeletedAt        string // 删除时间
	Remark           string // 备注
	IsFx             string // 是否是分销订单
}

// pmsTransactionColumns holds the columns for the table hg_pms_transaction.
var pmsTransactionColumns = PmsTransactionColumns{
	Id:               "id",
	OrderSn:          "order_sn",
	ChangeOrderSn:    "change_order_sn",
	OrderType:        "order_type",
	TransactionSn:    "transaction_sn",
	PaymentRequestId: "payment_request_id",
	CaptureId:        "capture_id",
	Scene:            "scene",
	PayChannel:       "pay_channel",
	PayType:          "pay_type",
	OpenId:           "open_id",
	Amount:           "amount",
	PayParams:        "pay_params",
	PriceCurrency:    "price_currency",
	PayAmount:        "pay_amount",
	PayCharge:        "pay_charge",
	PayStatus:        "pay_status",
	PayTime:          "pay_time",
	ExpiredTime:      "expired_time",
	RefundAmount:     "refund_amount",
	RefundStatus:     "refund_status",
	ScenePayRate:     "scene_pay_rate",
	Level:            "level",
	ExchangeRate:     "exchange_rate",
	CouponId:         "coupon_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	Remark:           "remark",
	IsFx:             "is_fx",
}

// NewPmsTransactionDao creates and returns a new DAO object for table data access.
func NewPmsTransactionDao() *PmsTransactionDao {
	return &PmsTransactionDao{
		group:   "default",
		table:   "hg_pms_transaction",
		columns: pmsTransactionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsTransactionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsTransactionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsTransactionDao) Columns() PmsTransactionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsTransactionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsTransactionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsTransactionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
