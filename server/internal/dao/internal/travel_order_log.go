// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelOrderLogDao is the data access object for table hg_travel_order_log.
type TravelOrderLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TravelOrderLogColumns // columns contains all the column names of Table for convenient usage.
}

// TravelOrderLogColumns defines and stores column names for table hg_travel_order_log.
type TravelOrderLogColumns struct {
	Id          string // 变动ID
	OrderId     string // 订单ID
	OrderStatus string // 订单状态
	ActionWay   string // 操作名
	Remark      string // 备注
	OperateType string // 操作员类型
	OperateId   string // 操作员ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// travelOrderLogColumns holds the columns for table hg_travel_order_log.
var travelOrderLogColumns = TravelOrderLogColumns{
	Id:          "id",
	OrderId:     "order_id",
	OrderStatus: "order_status",
	ActionWay:   "action_way",
	Remark:      "remark",
	OperateType: "operate_type",
	OperateId:   "operate_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTravelOrderLogDao creates and returns a new DAO object for table data access.
func NewTravelOrderLogDao() *TravelOrderLogDao {
	return &TravelOrderLogDao{
		group:   "default",
		table:   "hg_travel_order_log",
		columns: travelOrderLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelOrderLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelOrderLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelOrderLogDao) Columns() TravelOrderLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelOrderLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelOrderLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelOrderLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
