// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarDriverBalanceChangeDao is the data access object for the table hg_car_driver_balance_change.
type CarDriverBalanceChangeDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of the current DAO.
	columns CarDriverBalanceChangeColumns // columns contains all the column names of Table for convenient usage.
}

// CarDriverBalanceChangeColumns defines and stores column names for the table hg_car_driver_balance_change.
type CarDriverBalanceChangeColumns struct {
	Id                string // 主键
	DriverId          string // 司机ID
	Type              string // 金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整
	ChangePrice       string // 变更金额
	SettlementOrderId string // 接送机结算单ID
	WithdrawOrderId   string // 提现单ID
	ReturnOrderId     string // 退单ID
	Des               string // 描述
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// carDriverBalanceChangeColumns holds the columns for the table hg_car_driver_balance_change.
var carDriverBalanceChangeColumns = CarDriverBalanceChangeColumns{
	Id:                "id",
	DriverId:          "driver_id",
	Type:              "type",
	ChangePrice:       "change_price",
	SettlementOrderId: "settlement_order_id",
	WithdrawOrderId:   "withdraw_order_id",
	ReturnOrderId:     "return_order_id",
	Des:               "des",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewCarDriverBalanceChangeDao creates and returns a new DAO object for table data access.
func NewCarDriverBalanceChangeDao() *CarDriverBalanceChangeDao {
	return &CarDriverBalanceChangeDao{
		group:   "default",
		table:   "hg_car_driver_balance_change",
		columns: carDriverBalanceChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarDriverBalanceChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarDriverBalanceChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarDriverBalanceChangeDao) Columns() CarDriverBalanceChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarDriverBalanceChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarDriverBalanceChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarDriverBalanceChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
