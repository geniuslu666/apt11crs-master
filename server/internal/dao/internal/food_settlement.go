// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodSettlementDao is the data access object for the table hg_food_settlement.
type FoodSettlementDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns FoodSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// FoodSettlementColumns defines and stores column names for the table hg_food_settlement.
type FoodSettlementColumns struct {
	Id              string //
	Name            string // 名称
	Type            string // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	Cycle           string // 结算周期  1每日结算  2每周结算  3每月结算
	Rate            string // 结算比例
	Cost            string // 结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金
	IsAuditWithdraw string // 是否提现审核  1开启  2关闭
	Status          string // 1、启用 2、禁用
	CreateAt        string // 创建时间
	UpdateAt        string // 更新时间
	DeletedAt       string //
}

// foodSettlementColumns holds the columns for the table hg_food_settlement.
var foodSettlementColumns = FoodSettlementColumns{
	Id:              "id",
	Name:            "name",
	Type:            "type",
	Cycle:           "cycle",
	Rate:            "rate",
	Cost:            "cost",
	IsAuditWithdraw: "is_audit_withdraw",
	Status:          "status",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
	DeletedAt:       "deleted_at",
}

// NewFoodSettlementDao creates and returns a new DAO object for table data access.
func NewFoodSettlementDao() *FoodSettlementDao {
	return &FoodSettlementDao{
		group:   "default",
		table:   "hg_food_settlement",
		columns: foodSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodSettlementDao) Columns() FoodSettlementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
