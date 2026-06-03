// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarDriverReturnOrderDao is the data access object for the table hg_car_driver_return_order.
type CarDriverReturnOrderDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns CarDriverReturnOrderColumns // columns contains all the column names of Table for convenient usage.
}

// CarDriverReturnOrderColumns defines and stores column names for the table hg_car_driver_return_order.
type CarDriverReturnOrderColumns struct {
	Id       string //
	OrderId  string // 订单ID
	DriverId string // 司机ID
	CarId    string // 车辆ID
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// carDriverReturnOrderColumns holds the columns for the table hg_car_driver_return_order.
var carDriverReturnOrderColumns = CarDriverReturnOrderColumns{
	Id:       "id",
	OrderId:  "order_id",
	DriverId: "driver_id",
	CarId:    "car_id",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewCarDriverReturnOrderDao creates and returns a new DAO object for table data access.
func NewCarDriverReturnOrderDao() *CarDriverReturnOrderDao {
	return &CarDriverReturnOrderDao{
		group:   "default",
		table:   "hg_car_driver_return_order",
		columns: carDriverReturnOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarDriverReturnOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarDriverReturnOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarDriverReturnOrderDao) Columns() CarDriverReturnOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarDriverReturnOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarDriverReturnOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarDriverReturnOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
