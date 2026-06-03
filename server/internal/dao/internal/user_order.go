// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserOrderDao is the data access object for the table user_order.
type UserOrderDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns UserOrderColumns // columns contains all the column names of Table for convenient usage.
}

// UserOrderColumns defines and stores column names for the table user_order.
type UserOrderColumns struct {
	Id            string //
	OrderSn       string // 订单号
	Money         string // 金额
	NewExpireTime string // 新的到期时间
	Payment       string // alipay支付宝，wechat微信，bank网银
	Type          string // 1未支付，2已支付，3已取消
	Comment       string // 备注
	CreatedAt     string // 创建时间
	Operator      string // 操作人
	KefuName      string // 客服账户
	EntId         string // 客服企业ID
}

// userOrderColumns holds the columns for the table user_order.
var userOrderColumns = UserOrderColumns{
	Id:            "id",
	OrderSn:       "order_sn",
	Money:         "money",
	NewExpireTime: "new_expire_time",
	Payment:       "payment",
	Type:          "type",
	Comment:       "comment",
	CreatedAt:     "created_at",
	Operator:      "operator",
	KefuName:      "kefu_name",
	EntId:         "ent_id",
}

// NewUserOrderDao creates and returns a new DAO object for table data access.
func NewUserOrderDao() *UserOrderDao {
	return &UserOrderDao{
		group:   "default",
		table:   "user_order",
		columns: userOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserOrderDao) Columns() UserOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
