// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodSeatDao is the data access object for the table hg_food_seat.
type FoodSeatDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns FoodSeatColumns // columns contains all the column names of Table for convenient usage.
}

// FoodSeatColumns defines and stores column names for the table hg_food_seat.
type FoodSeatColumns struct {
	Id        string //
	SeatName  string //
	Status    string // 状态1、启用 2、禁用
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// foodSeatColumns holds the columns for the table hg_food_seat.
var foodSeatColumns = FoodSeatColumns{
	Id:        "id",
	SeatName:  "seat_name",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewFoodSeatDao creates and returns a new DAO object for table data access.
func NewFoodSeatDao() *FoodSeatDao {
	return &FoodSeatDao{
		group:   "default",
		table:   "hg_food_seat",
		columns: foodSeatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodSeatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodSeatDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodSeatDao) Columns() FoodSeatColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodSeatDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodSeatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodSeatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
