// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaSettlementDao is the data access object for the table hg_spa_settlement.
type SpaSettlementDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SpaSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// SpaSettlementColumns defines and stores column names for the table hg_spa_settlement.
type SpaSettlementColumns struct {
	Id        string //
	Name      string // 名称
	Type      string // 结算方式  1无需结算 2按周期自动结算
	Cycle     string // 结算周期  1每日结算  2每周结算  3每月结算
	Rate      string // 结算比例
	Cost      string // 结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金
	Status    string // 1、启用 2、禁用
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string //
}

// spaSettlementColumns holds the columns for the table hg_spa_settlement.
var spaSettlementColumns = SpaSettlementColumns{
	Id:        "id",
	Name:      "name",
	Type:      "type",
	Cycle:     "cycle",
	Rate:      "rate",
	Cost:      "cost",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewSpaSettlementDao creates and returns a new DAO object for table data access.
func NewSpaSettlementDao() *SpaSettlementDao {
	return &SpaSettlementDao{
		group:   "default",
		table:   "hg_spa_settlement",
		columns: spaSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaSettlementDao) Columns() SpaSettlementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
