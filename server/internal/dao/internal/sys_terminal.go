// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTerminalDao is the data access object for the table hg_sys_terminal.
type SysTerminalDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysTerminalColumns // columns contains all the column names of Table for convenient usage.
}

// SysTerminalColumns defines and stores column names for the table hg_sys_terminal.
type SysTerminalColumns struct {
	Id           string //
	TerminalName string // 终端名称
	TerminalType string // 终端类型
	BrandModel   string // 品牌型号
	Sn           string // 终端编号
	OnlineStatus string // 在线状态 1-在线 2离线
	StoreId      string // 绑定门店
	RestaurantId string // 绑定餐厅
	PrintTimes   string // 打印联数(次数)
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
	DeletedAt    string //
}

// sysTerminalColumns holds the columns for the table hg_sys_terminal.
var sysTerminalColumns = SysTerminalColumns{
	Id:           "id",
	TerminalName: "terminal_name",
	TerminalType: "terminal_type",
	BrandModel:   "brand_model",
	Sn:           "sn",
	OnlineStatus: "online_status",
	StoreId:      "store_id",
	RestaurantId: "restaurant_id",
	PrintTimes:   "print_times",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeletedAt:    "deleted_at",
}

// NewSysTerminalDao creates and returns a new DAO object for table data access.
func NewSysTerminalDao() *SysTerminalDao {
	return &SysTerminalDao{
		group:   "default",
		table:   "hg_sys_terminal",
		columns: sysTerminalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTerminalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTerminalDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTerminalDao) Columns() SysTerminalColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTerminalDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTerminalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysTerminalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
