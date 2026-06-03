// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaIspBalanceChangeDao is the data access object for the table hg_spa_isp_balance_change.
type SpaIspBalanceChangeDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns SpaIspBalanceChangeColumns // columns contains all the column names of Table for convenient usage.
}

// SpaIspBalanceChangeColumns defines and stores column names for the table hg_spa_isp_balance_change.
type SpaIspBalanceChangeColumns struct {
	Id                string // 主键
	IspId             string // 服务商ID
	Type              string // 金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整
	ChangePrice       string // 变更金额
	SettlementOrderId string // 结算单ID
	WithdrawOrderId   string // 提现单ID
	Des               string // 描述
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// spaIspBalanceChangeColumns holds the columns for the table hg_spa_isp_balance_change.
var spaIspBalanceChangeColumns = SpaIspBalanceChangeColumns{
	Id:                "id",
	IspId:             "isp_id",
	Type:              "type",
	ChangePrice:       "change_price",
	SettlementOrderId: "settlement_order_id",
	WithdrawOrderId:   "withdraw_order_id",
	Des:               "des",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewSpaIspBalanceChangeDao creates and returns a new DAO object for table data access.
func NewSpaIspBalanceChangeDao() *SpaIspBalanceChangeDao {
	return &SpaIspBalanceChangeDao{
		group:   "default",
		table:   "hg_spa_isp_balance_change",
		columns: spaIspBalanceChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaIspBalanceChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaIspBalanceChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaIspBalanceChangeDao) Columns() SpaIspBalanceChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaIspBalanceChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaIspBalanceChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaIspBalanceChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
