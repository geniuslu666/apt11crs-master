// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaOrderTechnicianDao is the data access object for the table hg_spa_order_technician.
type SpaOrderTechnicianDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns SpaOrderTechnicianColumns // columns contains all the column names of Table for convenient usage.
}

// SpaOrderTechnicianColumns defines and stores column names for the table hg_spa_order_technician.
type SpaOrderTechnicianColumns struct {
	Id                string //
	IspId             string // 服务商ID
	OrderId           string // 订单ID
	OrderGoodsId      string // 子订单ID
	TechnicianId      string // 技师ID
	SettlementObject  string // 结算对象
	SettlementRate    string // 结算比例
	SettlementStatus  string // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount  string // 结算金额
	SettlementTime    string // 结算时间
	SettlementOrderId string // 结算单ID
	SettlementType    string // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle   string // 结算周期  1每日结算  2每周结算  3每月结算
	SettlementCost    string // 结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金
	OrderAmount       string // 单笔订单金额
	CouponAmount      string // 优惠券抵扣金额
	BalAmount         string // 积分抵扣金额
	ActualEndTime     string // 实际结束日期时间
}

// spaOrderTechnicianColumns holds the columns for the table hg_spa_order_technician.
var spaOrderTechnicianColumns = SpaOrderTechnicianColumns{
	Id:                "id",
	IspId:             "isp_id",
	OrderId:           "order_id",
	OrderGoodsId:      "order_goods_id",
	TechnicianId:      "technician_id",
	SettlementObject:  "settlement_object",
	SettlementRate:    "settlement_rate",
	SettlementStatus:  "settlement_status",
	SettlementAmount:  "settlement_amount",
	SettlementTime:    "settlement_time",
	SettlementOrderId: "settlement_order_id",
	SettlementType:    "settlement_type",
	SettlementCycle:   "settlement_cycle",
	SettlementCost:    "settlement_cost",
	OrderAmount:       "order_amount",
	CouponAmount:      "coupon_amount",
	BalAmount:         "bal_amount",
	ActualEndTime:     "actual_end_time",
}

// NewSpaOrderTechnicianDao creates and returns a new DAO object for table data access.
func NewSpaOrderTechnicianDao() *SpaOrderTechnicianDao {
	return &SpaOrderTechnicianDao{
		group:   "default",
		table:   "hg_spa_order_technician",
		columns: spaOrderTechnicianColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaOrderTechnicianDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaOrderTechnicianDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaOrderTechnicianDao) Columns() SpaOrderTechnicianColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaOrderTechnicianDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaOrderTechnicianDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaOrderTechnicianDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
