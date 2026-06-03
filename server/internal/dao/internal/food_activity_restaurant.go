// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodActivityRestaurantDao is the data access object for the table hg_food_activity_restaurant.
type FoodActivityRestaurantDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of the current DAO.
	columns FoodActivityRestaurantColumns // columns contains all the column names of Table for convenient usage.
}

// FoodActivityRestaurantColumns defines and stores column names for the table hg_food_activity_restaurant.
type FoodActivityRestaurantColumns struct {
	ActivityId   string // 活动ID
	RestaurantId string // 餐厅ID
}

// foodActivityRestaurantColumns holds the columns for the table hg_food_activity_restaurant.
var foodActivityRestaurantColumns = FoodActivityRestaurantColumns{
	ActivityId:   "activity_id",
	RestaurantId: "restaurant_id",
}

// NewFoodActivityRestaurantDao creates and returns a new DAO object for table data access.
func NewFoodActivityRestaurantDao() *FoodActivityRestaurantDao {
	return &FoodActivityRestaurantDao{
		group:   "default",
		table:   "hg_food_activity_restaurant",
		columns: foodActivityRestaurantColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodActivityRestaurantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodActivityRestaurantDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodActivityRestaurantDao) Columns() FoodActivityRestaurantColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodActivityRestaurantDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodActivityRestaurantDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodActivityRestaurantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
