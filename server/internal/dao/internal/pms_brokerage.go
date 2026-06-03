// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsBrokerageDao is the data access object for the table hg_pms_brokerage.
type PmsBrokerageDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsBrokerageColumns // columns contains all the column names of Table for convenient usage.
}

// PmsBrokerageColumns defines and stores column names for the table hg_pms_brokerage.
type PmsBrokerageColumns struct {
	Id        string //
	Identity  string // 类型    CHANNEL   渠道   STAFF  员工
	Scene     string // 场景值   HOTEL 酒店  SYSTEM 系统
	Type      string // REBATE 返佣  WITHDRAW  提现  SYS 系统调整
	Balance   string // 变动金额
	StaffId   string // 员工ID
	ChannelId string // 渠道ID
	MemberId  string // 会员ID
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// pmsBrokerageColumns holds the columns for the table hg_pms_brokerage.
var pmsBrokerageColumns = PmsBrokerageColumns{
	Id:        "id",
	Identity:  "identity",
	Scene:     "scene",
	Type:      "type",
	Balance:   "balance",
	StaffId:   "staff_id",
	ChannelId: "channel_id",
	MemberId:  "member_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewPmsBrokerageDao creates and returns a new DAO object for table data access.
func NewPmsBrokerageDao() *PmsBrokerageDao {
	return &PmsBrokerageDao{
		group:   "default",
		table:   "hg_pms_brokerage",
		columns: pmsBrokerageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsBrokerageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsBrokerageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsBrokerageDao) Columns() PmsBrokerageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsBrokerageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsBrokerageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsBrokerageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
