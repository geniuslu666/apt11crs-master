// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaIspDao is the data access object for the table hg_spa_isp.
type SpaIspDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns SpaIspColumns // columns contains all the column names of Table for convenient usage.
}

// SpaIspColumns defines and stores column names for the table hg_spa_isp.
type SpaIspColumns struct {
	Id                    string //
	Name                  string // 服务商名称
	Status                string // 状态1、启用 2、禁用
	WorkStatus            string // 工作状态
	MemberId              string // 会员ID
	SettlementObject      string // 结算对象
	SettlementType        string // 服务分成类型 1跟随系统  2自定义
	SettlementId          string // 结算模式ID
	SettlementRate        string // 服务分成%
	SettlementOrderNum    string // 预约单已结算数量
	SettlementOrderAmount string // 已结算预约单订单金额
	TotalSettlementAmount string // 已结算金额
	Balance               string // 余额
	VerifyMoney           string // 已核账金额
	ApplyWithdrawBalance  string // 提现中余额
	WithdrawBalance       string // 已提现余额
	SendSmsPhone          string // 发送短信电话
	CreateAt              string // 创建时间
	UpdateAt              string // 更新时间
	DeletedAt             string //
}

// spaIspColumns holds the columns for the table hg_spa_isp.
var spaIspColumns = SpaIspColumns{
	Id:                    "id",
	Name:                  "name",
	Status:                "status",
	WorkStatus:            "work_status",
	MemberId:              "member_id",
	SettlementObject:      "settlement_object",
	SettlementType:        "settlement_type",
	SettlementId:          "settlement_id",
	SettlementRate:        "settlement_rate",
	SettlementOrderNum:    "settlement_order_num",
	SettlementOrderAmount: "settlement_order_amount",
	TotalSettlementAmount: "total_settlement_amount",
	Balance:               "balance",
	VerifyMoney:           "verify_money",
	ApplyWithdrawBalance:  "apply_withdraw_balance",
	WithdrawBalance:       "withdraw_balance",
	SendSmsPhone:          "send_sms_phone",
	CreateAt:              "create_at",
	UpdateAt:              "update_at",
	DeletedAt:             "deleted_at",
}

// NewSpaIspDao creates and returns a new DAO object for table data access.
func NewSpaIspDao() *SpaIspDao {
	return &SpaIspDao{
		group:   "default",
		table:   "hg_spa_isp",
		columns: spaIspColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaIspDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaIspDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaIspDao) Columns() SpaIspColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaIspDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaIspDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaIspDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
