// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsBalanceChangeDao is the data access object for the table hg_pms_balance_change.
type PmsBalanceChangeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns PmsBalanceChangeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsBalanceChangeColumns defines and stores column names for the table hg_pms_balance_change.
type PmsBalanceChangeColumns struct {
	Id          string // 主键
	MemberId    string // 会员ID
	Scene       string // 场景值   HOTEL    酒店   SYSTEM 系统
	Type        string // 金额变动方式   CONSUME 消费   REFUND   退款    AWARD  奖励    BROKERAGE   佣金    SYS 系统调整
	ChangePrice string // 变更金额
	OrderSn     string // 订单号
	Des         string // 消费描述（后端展示）
	Reason      string // 原因（前端展示）
	OperatorId  string // 操作员ID
	MdCode      string // 注册设备码
	MpModel     string // 注册设备型号
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// pmsBalanceChangeColumns holds the columns for the table hg_pms_balance_change.
var pmsBalanceChangeColumns = PmsBalanceChangeColumns{
	Id:          "id",
	MemberId:    "member_id",
	Scene:       "scene",
	Type:        "type",
	ChangePrice: "change_price",
	OrderSn:     "order_sn",
	Des:         "des",
	Reason:      "reason",
	OperatorId:  "operator_id",
	MdCode:      "md_code",
	MpModel:     "mp_model",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewPmsBalanceChangeDao creates and returns a new DAO object for table data access.
func NewPmsBalanceChangeDao() *PmsBalanceChangeDao {
	return &PmsBalanceChangeDao{
		group:   "default",
		table:   "hg_pms_balance_change",
		columns: pmsBalanceChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsBalanceChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsBalanceChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsBalanceChangeDao) Columns() PmsBalanceChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsBalanceChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsBalanceChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsBalanceChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
