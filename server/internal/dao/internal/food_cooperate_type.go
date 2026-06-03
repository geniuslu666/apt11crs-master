// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodCooperateTypeDao is the data access object for the table hg_food_cooperate_type.
type FoodCooperateTypeDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns FoodCooperateTypeColumns // columns contains all the column names of Table for convenient usage.
}

// FoodCooperateTypeColumns defines and stores column names for the table hg_food_cooperate_type.
type FoodCooperateTypeColumns struct {
	Id        string //
	TypeName  string //
	Status    string // 状态1、启用 2、禁用
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// foodCooperateTypeColumns holds the columns for the table hg_food_cooperate_type.
var foodCooperateTypeColumns = FoodCooperateTypeColumns{
	Id:        "id",
	TypeName:  "type_name",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewFoodCooperateTypeDao creates and returns a new DAO object for table data access.
func NewFoodCooperateTypeDao() *FoodCooperateTypeDao {
	return &FoodCooperateTypeDao{
		group:   "default",
		table:   "hg_food_cooperate_type",
		columns: foodCooperateTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodCooperateTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodCooperateTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodCooperateTypeDao) Columns() FoodCooperateTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodCooperateTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodCooperateTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodCooperateTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
