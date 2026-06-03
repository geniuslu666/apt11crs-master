// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaTechnicianBalanceChangeDao is the data access object for the table hg_spa_technician_balance_change.
type SpaTechnicianBalanceChangeDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of the current DAO.
	columns SpaTechnicianBalanceChangeColumns // columns contains all the column names of Table for convenient usage.
}

// SpaTechnicianBalanceChangeColumns defines and stores column names for the table hg_spa_technician_balance_change.
type SpaTechnicianBalanceChangeColumns struct {
	Id                string // 主键
	TechnicianId      string // 技师ID
	Type              string // 金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整
	ChangePrice       string // 变更金额
	SettlementOrderId string // 结算单ID
	WithdrawOrderId   string // 提现单ID
	Des               string // 描述
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// spaTechnicianBalanceChangeColumns holds the columns for the table hg_spa_technician_balance_change.
var spaTechnicianBalanceChangeColumns = SpaTechnicianBalanceChangeColumns{
	Id:                "id",
	TechnicianId:      "technician_id",
	Type:              "type",
	ChangePrice:       "change_price",
	SettlementOrderId: "settlement_order_id",
	WithdrawOrderId:   "withdraw_order_id",
	Des:               "des",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewSpaTechnicianBalanceChangeDao creates and returns a new DAO object for table data access.
func NewSpaTechnicianBalanceChangeDao() *SpaTechnicianBalanceChangeDao {
	return &SpaTechnicianBalanceChangeDao{
		group:   "default",
		table:   "hg_spa_technician_balance_change",
		columns: spaTechnicianBalanceChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaTechnicianBalanceChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaTechnicianBalanceChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaTechnicianBalanceChangeDao) Columns() SpaTechnicianBalanceChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaTechnicianBalanceChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaTechnicianBalanceChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaTechnicianBalanceChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
