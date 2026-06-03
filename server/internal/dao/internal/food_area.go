// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodAreaDao is the data access object for the table hg_food_area.
type FoodAreaDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns FoodAreaColumns // columns contains all the column names of Table for convenient usage.
}

// FoodAreaColumns defines and stores column names for the table hg_food_area.
type FoodAreaColumns struct {
	Id         string //
	Pid        string // 上级ID
	Level      string // 区域级别
	Tree       string // 关系树
	AreaName   string // 区域名称
	AreaStatus string // 1、启用 2、禁用
	CreateAt   string // 创建时间
	UpdateAt   string // 更新时间
	DeletedAt  string //
}

// foodAreaColumns holds the columns for the table hg_food_area.
var foodAreaColumns = FoodAreaColumns{
	Id:         "id",
	Pid:        "pid",
	Level:      "level",
	Tree:       "tree",
	AreaName:   "area_name",
	AreaStatus: "area_status",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeletedAt:  "deleted_at",
}

// NewFoodAreaDao creates and returns a new DAO object for table data access.
func NewFoodAreaDao() *FoodAreaDao {
	return &FoodAreaDao{
		group:   "default",
		table:   "hg_food_area",
		columns: foodAreaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodAreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodAreaDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodAreaDao) Columns() FoodAreaColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodAreaDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodAreaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodAreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
