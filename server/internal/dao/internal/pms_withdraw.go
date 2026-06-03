// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsWithdrawDao is the data access object for the table hg_pms_withdraw.
type PmsWithdrawDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PmsWithdrawColumns // columns contains all the column names of Table for convenient usage.
}

// PmsWithdrawColumns defines and stores column names for the table hg_pms_withdraw.
type PmsWithdrawColumns struct {
	Id             string //
	Type           string // 类型
	StaffId        string // 员工ID
	ChannelId      string // 渠道ID
	WithdrawSn     string // 提现单号
	WithdrawStatus string // 提现状态
	WithdrawAmount string // 提现金额
	ArrivalAmount  string // 到账金额
	ServiceCharge  string // 提现手续费比例
	Transfer       string // 1、未转账 2、已转账
	ApplyRemark    string // 审核备注
	ApplyAt        string // 审核时间
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// pmsWithdrawColumns holds the columns for the table hg_pms_withdraw.
var pmsWithdrawColumns = PmsWithdrawColumns{
	Id:             "id",
	Type:           "type",
	StaffId:        "staff_id",
	ChannelId:      "channel_id",
	WithdrawSn:     "withdraw_sn",
	WithdrawStatus: "withdraw_status",
	WithdrawAmount: "withdraw_amount",
	ArrivalAmount:  "arrival_amount",
	ServiceCharge:  "service_charge",
	Transfer:       "transfer",
	ApplyRemark:    "apply_remark",
	ApplyAt:        "apply_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPmsWithdrawDao creates and returns a new DAO object for table data access.
func NewPmsWithdrawDao() *PmsWithdrawDao {
	return &PmsWithdrawDao{
		group:   "default",
		table:   "hg_pms_withdraw",
		columns: pmsWithdrawColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsWithdrawDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsWithdrawDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsWithdrawDao) Columns() PmsWithdrawColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsWithdrawDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsWithdrawDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsWithdrawDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
