// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarServiceAddressDao is the data access object for the table hg_car_service_address.
type CarServiceAddressDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns CarServiceAddressColumns // columns contains all the column names of Table for convenient usage.
}

// CarServiceAddressColumns defines and stores column names for the table hg_car_service_address.
type CarServiceAddressColumns struct {
	ServiceId string // 服务ID
	AddressId string // 地址ID
	Type      string // 1 出发地  2 目的地
}

// carServiceAddressColumns holds the columns for the table hg_car_service_address.
var carServiceAddressColumns = CarServiceAddressColumns{
	ServiceId: "service_id",
	AddressId: "address_id",
	Type:      "type",
}

// NewCarServiceAddressDao creates and returns a new DAO object for table data access.
func NewCarServiceAddressDao() *CarServiceAddressDao {
	return &CarServiceAddressDao{
		group:   "default",
		table:   "hg_car_service_address",
		columns: carServiceAddressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarServiceAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarServiceAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarServiceAddressDao) Columns() CarServiceAddressColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarServiceAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarServiceAddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarServiceAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
