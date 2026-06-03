// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPrinterDao is the data access object for the table hg_sys_printer.
type SysPrinterDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SysPrinterColumns // columns contains all the column names of Table for convenient usage.
}

// SysPrinterColumns defines and stores column names for the table hg_sys_printer.
type SysPrinterColumns struct {
	Id           string //
	PrinterType  string // 打印机类型
	PrinterName  string // 打印机名称
	ClientId     string // 易联云第三方应用ID
	ClientSecret string // 易联云第三方应用秘钥
	MachineCode  string // 打印机终端号
	MachineKey   string // 终端秘钥
	PrintTimes   string // 打印联数(次数)
	Sort         string // 排序 (数字越大越靠前)
	Status       string // 状态1、启用 2、禁用
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
	DeletedAt    string //
}

// sysPrinterColumns holds the columns for the table hg_sys_printer.
var sysPrinterColumns = SysPrinterColumns{
	Id:           "id",
	PrinterType:  "printer_type",
	PrinterName:  "printer_name",
	ClientId:     "client_id",
	ClientSecret: "client_secret",
	MachineCode:  "machine_code",
	MachineKey:   "machine_key",
	PrintTimes:   "print_times",
	Sort:         "sort",
	Status:       "status",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeletedAt:    "deleted_at",
}

// NewSysPrinterDao creates and returns a new DAO object for table data access.
func NewSysPrinterDao() *SysPrinterDao {
	return &SysPrinterDao{
		group:   "default",
		table:   "hg_sys_printer",
		columns: sysPrinterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPrinterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPrinterDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPrinterDao) Columns() SysPrinterColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPrinterDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPrinterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysPrinterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
