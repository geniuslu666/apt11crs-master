// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodSettlementAccountDao is the data access object for the table hg_food_settlement_account.
type FoodSettlementAccountDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of the current DAO.
	columns FoodSettlementAccountColumns // columns contains all the column names of Table for convenient usage.
}

// FoodSettlementAccountColumns defines and stores column names for the table hg_food_settlement_account.
type FoodSettlementAccountColumns struct {
	Id            string //
	RestaurantId  string // 餐厅ID
	Type          string // 类型
	BankName      string // 开户行
	BankUser      string // 开户人姓名
	BankCard      string // 银行账号
	WechatAccount string // 微信名
	AlipayName    string // 支付宝真实姓名
	AlipayAccount string // 支付宝账户
	Status        string // 1、启用 2、禁用
	CreateAt      string // 创建时间
	UpdateAt      string // 更新时间
	DeletedAt     string //
}

// foodSettlementAccountColumns holds the columns for the table hg_food_settlement_account.
var foodSettlementAccountColumns = FoodSettlementAccountColumns{
	Id:            "id",
	RestaurantId:  "restaurant_id",
	Type:          "type",
	BankName:      "bank_name",
	BankUser:      "bank_user",
	BankCard:      "bank_card",
	WechatAccount: "wechat_account",
	AlipayName:    "alipay_name",
	AlipayAccount: "alipay_account",
	Status:        "status",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
	DeletedAt:     "deleted_at",
}

// NewFoodSettlementAccountDao creates and returns a new DAO object for table data access.
func NewFoodSettlementAccountDao() *FoodSettlementAccountDao {
	return &FoodSettlementAccountDao{
		group:   "default",
		table:   "hg_food_settlement_account",
		columns: foodSettlementAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodSettlementAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodSettlementAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodSettlementAccountDao) Columns() FoodSettlementAccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodSettlementAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodSettlementAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodSettlementAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
