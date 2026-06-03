// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CabinetOrderLogDao is the data access object for the table hg_cabinet_order_log.
type CabinetOrderLogDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns CabinetOrderLogColumns // columns contains all the column names of Table for convenient usage.
}

// CabinetOrderLogColumns defines and stores column names for the table hg_cabinet_order_log.
type CabinetOrderLogColumns struct {
	Id          string // 变动ID
	OrderId     string // 订单ID
	OrderStatus string // 订单状态
	ActionWay   string // 操作名
	Remark      string // 备注
	Images      string // 图集
	OperateType string // 操作员类型
	OperateId   string // 操作员ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// cabinetOrderLogColumns holds the columns for the table hg_cabinet_order_log.
var cabinetOrderLogColumns = CabinetOrderLogColumns{
	Id:          "id",
	OrderId:     "order_id",
	OrderStatus: "order_status",
	ActionWay:   "action_way",
	Remark:      "remark",
	Images:      "images",
	OperateType: "operate_type",
	OperateId:   "operate_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewCabinetOrderLogDao creates and returns a new DAO object for table data access.
func NewCabinetOrderLogDao() *CabinetOrderLogDao {
	return &CabinetOrderLogDao{
		group:   "default",
		table:   "hg_cabinet_order_log",
		columns: cabinetOrderLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CabinetOrderLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CabinetOrderLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CabinetOrderLogDao) Columns() CabinetOrderLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CabinetOrderLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CabinetOrderLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CabinetOrderLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
