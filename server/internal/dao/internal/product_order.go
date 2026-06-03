// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProductOrderDao is the data access object for the table product_order.
type ProductOrderDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns ProductOrderColumns // columns contains all the column names of Table for convenient usage.
}

// ProductOrderColumns defines and stores column names for the table product_order.
type ProductOrderColumns struct {
	Id              string //
	EntId           string // 企业ID
	KefuName        string // 客服账户
	UserId          string // 用户ID
	OrderSn         string // 订单编号
	OrderDesc       string // 订单描述
	OrderStatus     string // 订单状态：pending,processing,completed,cancelled
	TotalAmount     string // 订单金额
	PaymentMethod   string // 支付方式：wechat,alipay,bank,other
	PaymentStatus   string // 支付状态：paid,unpaid,refunded
	ShippingAddress string // 收货地址
	Email           string // 邮箱
	Contact         string // 联系人
	Tel             string // 手机号
	CreatedAt       string //
	UpdatedAt       string //
}

// productOrderColumns holds the columns for the table product_order.
var productOrderColumns = ProductOrderColumns{
	Id:              "id",
	EntId:           "ent_id",
	KefuName:        "kefu_name",
	UserId:          "user_id",
	OrderSn:         "order_sn",
	OrderDesc:       "order_desc",
	OrderStatus:     "order_status",
	TotalAmount:     "total_amount",
	PaymentMethod:   "payment_method",
	PaymentStatus:   "payment_status",
	ShippingAddress: "shipping_address",
	Email:           "email",
	Contact:         "contact",
	Tel:             "tel",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewProductOrderDao creates and returns a new DAO object for table data access.
func NewProductOrderDao() *ProductOrderDao {
	return &ProductOrderDao{
		group:   "default",
		table:   "product_order",
		columns: productOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProductOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProductOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProductOrderDao) Columns() ProductOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProductOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProductOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ProductOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
