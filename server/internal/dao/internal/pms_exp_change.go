// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsExpChangeDao is the data access object for the table hg_pms_exp_change.
type PmsExpChangeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsExpChangeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsExpChangeColumns defines and stores column names for the table hg_pms_exp_change.
type PmsExpChangeColumns struct {
	Id         string //
	Scene      string // 场景值   HOTEL 酒店  SYSTEM 系统
	Exp        string // 变动金额
	OrderSn    string // 订单号
	Des        string // 描述
	MemberId   string // 会员ID
	OperatorId string // 操作员ID
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

// pmsExpChangeColumns holds the columns for the table hg_pms_exp_change.
var pmsExpChangeColumns = PmsExpChangeColumns{
	Id:         "id",
	Scene:      "scene",
	Exp:        "exp",
	OrderSn:    "order_sn",
	Des:        "des",
	MemberId:   "member_id",
	OperatorId: "operator_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewPmsExpChangeDao creates and returns a new DAO object for table data access.
func NewPmsExpChangeDao() *PmsExpChangeDao {
	return &PmsExpChangeDao{
		group:   "default",
		table:   "hg_pms_exp_change",
		columns: pmsExpChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsExpChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsExpChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsExpChangeDao) Columns() PmsExpChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsExpChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsExpChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsExpChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
