// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsPropertyRegionDao is the data access object for the table hg_pms_property_region.
type PmsPropertyRegionDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns PmsPropertyRegionColumns // columns contains all the column names of Table for convenient usage.
}

// PmsPropertyRegionColumns defines and stores column names for the table hg_pms_property_region.
type PmsPropertyRegionColumns struct {
	Id        string // 主键
	Name      string // 区域名称（多语言）
	Status    string // 状态1、启用 2、禁用
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// pmsPropertyRegionColumns holds the columns for the table hg_pms_property_region.
var pmsPropertyRegionColumns = PmsPropertyRegionColumns{
	Id:        "id",
	Name:      "name",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewPmsPropertyRegionDao creates and returns a new DAO object for table data access.
func NewPmsPropertyRegionDao() *PmsPropertyRegionDao {
	return &PmsPropertyRegionDao{
		group:   "default",
		table:   "hg_pms_property_region",
		columns: pmsPropertyRegionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsPropertyRegionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsPropertyRegionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsPropertyRegionDao) Columns() PmsPropertyRegionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsPropertyRegionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsPropertyRegionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsPropertyRegionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
