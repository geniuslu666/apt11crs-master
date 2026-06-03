// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarHelpDao is the data access object for the table hg_car_help.
type CarHelpDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns CarHelpColumns // columns contains all the column names of Table for convenient usage.
}

// CarHelpColumns defines and stores column names for the table hg_car_help.
type CarHelpColumns struct {
	Id       string //
	Language string // 语言
	Image    string // 主图
	Content  string // 内容
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// carHelpColumns holds the columns for the table hg_car_help.
var carHelpColumns = CarHelpColumns{
	Id:       "id",
	Language: "language",
	Image:    "image",
	Content:  "content",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewCarHelpDao creates and returns a new DAO object for table data access.
func NewCarHelpDao() *CarHelpDao {
	return &CarHelpDao{
		group:   "default",
		table:   "hg_car_help",
		columns: carHelpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarHelpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarHelpDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarHelpDao) Columns() CarHelpColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarHelpDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarHelpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarHelpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
