// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantTerminalDao is the data access object for the table hg_food_restaurant_terminal.
type FoodRestaurantTerminalDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of the current DAO.
	columns FoodRestaurantTerminalColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantTerminalColumns defines and stores column names for the table hg_food_restaurant_terminal.
type FoodRestaurantTerminalColumns struct {
	RestaurantId string // 餐厅ID
	TerminalId   string // 终端ID
	PrintTimes   string // 打印联数(次数)
}

// foodRestaurantTerminalColumns holds the columns for the table hg_food_restaurant_terminal.
var foodRestaurantTerminalColumns = FoodRestaurantTerminalColumns{
	RestaurantId: "restaurant_id",
	TerminalId:   "terminal_id",
	PrintTimes:   "print_times",
}

// NewFoodRestaurantTerminalDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantTerminalDao() *FoodRestaurantTerminalDao {
	return &FoodRestaurantTerminalDao{
		group:   "default",
		table:   "hg_food_restaurant_terminal",
		columns: foodRestaurantTerminalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantTerminalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantTerminalDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantTerminalDao) Columns() FoodRestaurantTerminalColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantTerminalDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantTerminalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantTerminalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
