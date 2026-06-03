// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodLabelDao is the data access object for the table hg_food_label.
type FoodLabelDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns FoodLabelColumns // columns contains all the column names of Table for convenient usage.
}

// FoodLabelColumns defines and stores column names for the table hg_food_label.
type FoodLabelColumns struct {
	Id        string //
	Type      string // 1 餐厅  2 套餐
	Name      string //
	Status    string // 状态1、启用 2、禁用
	Sort      string // 排序(越大越靠前)
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// foodLabelColumns holds the columns for the table hg_food_label.
var foodLabelColumns = FoodLabelColumns{
	Id:        "id",
	Type:      "type",
	Name:      "name",
	Status:    "status",
	Sort:      "sort",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewFoodLabelDao creates and returns a new DAO object for table data access.
func NewFoodLabelDao() *FoodLabelDao {
	return &FoodLabelDao{
		group:   "default",
		table:   "hg_food_label",
		columns: foodLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodLabelDao) Columns() FoodLabelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
