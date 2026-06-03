// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsScanDeviceDao is the data access object for the table hg_pms_scan_device.
type PmsScanDeviceDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsScanDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// PmsScanDeviceColumns defines and stores column names for the table hg_pms_scan_device.
type PmsScanDeviceColumns struct {
	Id       string //
	DeviceSn string // 设备
}

// pmsScanDeviceColumns holds the columns for the table hg_pms_scan_device.
var pmsScanDeviceColumns = PmsScanDeviceColumns{
	Id:       "id",
	DeviceSn: "device_sn",
}

// NewPmsScanDeviceDao creates and returns a new DAO object for table data access.
func NewPmsScanDeviceDao() *PmsScanDeviceDao {
	return &PmsScanDeviceDao{
		group:   "default",
		table:   "hg_pms_scan_device",
		columns: pmsScanDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsScanDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsScanDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsScanDeviceDao) Columns() PmsScanDeviceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsScanDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsScanDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsScanDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
