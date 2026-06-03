// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarSettlementOrderDao is the data access object for the table hg_car_settlement_order.
type CarSettlementOrderDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns CarSettlementOrderColumns // columns contains all the column names of Table for convenient usage.
}

// CarSettlementOrderColumns defines and stores column names for the table hg_car_settlement_order.
type CarSettlementOrderColumns struct {
	Id               string //
	OrderSn          string // 结算单编号
	DriverId         string // 司机ID
	OrderAmount      string // 订单总额
	SettlementAmount string // 结算总额
	Status           string // 结算状态
	StartTime        string // 账期开始时间
	EndTime          string // 账期结束时间
	VerifyStatus     string // 核账状态
	VerifyTime       string // 核帐时间
	VerifyImg        string // 核账凭证
	VerifyDesc       string // 核账说明
	VerifyOperateId  string // 核账操作人ID
	CreateAt         string // 创建时间
	UpdateAt         string // 更新时间
}

// carSettlementOrderColumns holds the columns for the table hg_car_settlement_order.
var carSettlementOrderColumns = CarSettlementOrderColumns{
	Id:               "id",
	OrderSn:          "order_sn",
	DriverId:         "driver_id",
	OrderAmount:      "order_amount",
	SettlementAmount: "settlement_amount",
	Status:           "status",
	StartTime:        "start_time",
	EndTime:          "end_time",
	VerifyStatus:     "verify_status",
	VerifyTime:       "verify_time",
	VerifyImg:        "verify_img",
	VerifyDesc:       "verify_desc",
	VerifyOperateId:  "verify_operate_id",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
}

// NewCarSettlementOrderDao creates and returns a new DAO object for table data access.
func NewCarSettlementOrderDao() *CarSettlementOrderDao {
	return &CarSettlementOrderDao{
		group:   "default",
		table:   "hg_car_settlement_order",
		columns: carSettlementOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarSettlementOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarSettlementOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarSettlementOrderDao) Columns() CarSettlementOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarSettlementOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarSettlementOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarSettlementOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
