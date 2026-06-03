// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderExportDao is the data access object for the table hg_order_export.
type OrderExportDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns OrderExportColumns // columns contains all the column names of Table for convenient usage.
}

// OrderExportColumns defines and stores column names for the table hg_order_export.
type OrderExportColumns struct {
	Id        string //
	Scene     string // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车 5-储物柜
	Condition string // 查询条件
	Status    string // 0：导出中   1：导出成功  2：导出失败
	Path      string // 路径
	CreateAt  string // 创建时间
	UpdateAt  string // 修改时间
}

// orderExportColumns holds the columns for the table hg_order_export.
var orderExportColumns = OrderExportColumns{
	Id:        "id",
	Scene:     "scene",
	Condition: "condition",
	Status:    "status",
	Path:      "path",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// NewOrderExportDao creates and returns a new DAO object for table data access.
func NewOrderExportDao() *OrderExportDao {
	return &OrderExportDao{
		group:   "default",
		table:   "hg_order_export",
		columns: orderExportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OrderExportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OrderExportDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OrderExportDao) Columns() OrderExportColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OrderExportDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OrderExportDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OrderExportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
