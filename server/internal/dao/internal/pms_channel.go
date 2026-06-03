// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsChannelDao is the data access object for the table hg_pms_channel.
type PmsChannelDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns PmsChannelColumns // columns contains all the column names of Table for convenient usage.
}

// PmsChannelColumns defines and stores column names for the table hg_pms_channel.
type PmsChannelColumns struct {
	Id                   string //
	Name                 string // 渠道姓名
	Phone                string // 手机号
	Email                string // 邮箱
	Rate                 string // 返佣比例
	Status               string // 状态   1、 开启   2、禁用
	Remark               string // 备注
	Balance              string // 可提现账户
	AllBalance           string // 总账户
	ApplyWithdrawBalance string // 提现中余额
	WithdrawBalance      string // 已提现余额
	MinWithdrawalAmount  string // 最低可提现额
	ServiceCharge        string // 手续费
	AfterDay             string // 预计提现时间周期
	CreatedAt            string //
	UpdatedAt            string //
	DeletedAt            string //
}

// pmsChannelColumns holds the columns for the table hg_pms_channel.
var pmsChannelColumns = PmsChannelColumns{
	Id:                   "id",
	Name:                 "name",
	Phone:                "phone",
	Email:                "email",
	Rate:                 "rate",
	Status:               "status",
	Remark:               "remark",
	Balance:              "balance",
	AllBalance:           "all_balance",
	ApplyWithdrawBalance: "apply_withdraw_balance",
	WithdrawBalance:      "withdraw_balance",
	MinWithdrawalAmount:  "min_withdrawal_amount",
	ServiceCharge:        "service_charge",
	AfterDay:             "after_day",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
}

// NewPmsChannelDao creates and returns a new DAO object for table data access.
func NewPmsChannelDao() *PmsChannelDao {
	return &PmsChannelDao{
		group:   "default",
		table:   "hg_pms_channel",
		columns: pmsChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsChannelDao) Columns() PmsChannelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
