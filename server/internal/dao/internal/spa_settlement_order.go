// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaSettlementOrderDao is the data access object for the table hg_spa_settlement_order.
type SpaSettlementOrderDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns SpaSettlementOrderColumns // columns contains all the column names of Table for convenient usage.
}

// SpaSettlementOrderColumns defines and stores column names for the table hg_spa_settlement_order.
type SpaSettlementOrderColumns struct {
	Id               string //
	SettlementObject string // 结算对象
	IspId            string // 服务商ID
	OrderSn          string // 结算单编号
	TechnicianId     string // 技师ID
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

// spaSettlementOrderColumns holds the columns for the table hg_spa_settlement_order.
var spaSettlementOrderColumns = SpaSettlementOrderColumns{
	Id:               "id",
	SettlementObject: "settlement_object",
	IspId:            "isp_id",
	OrderSn:          "order_sn",
	TechnicianId:     "technician_id",
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

// NewSpaSettlementOrderDao creates and returns a new DAO object for table data access.
func NewSpaSettlementOrderDao() *SpaSettlementOrderDao {
	return &SpaSettlementOrderDao{
		group:   "default",
		table:   "hg_spa_settlement_order",
		columns: spaSettlementOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaSettlementOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaSettlementOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaSettlementOrderDao) Columns() SpaSettlementOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaSettlementOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaSettlementOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaSettlementOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
