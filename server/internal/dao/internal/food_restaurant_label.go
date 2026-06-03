// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantLabelDao is the data access object for the table hg_food_restaurant_label.
type FoodRestaurantLabelDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns FoodRestaurantLabelColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantLabelColumns defines and stores column names for the table hg_food_restaurant_label.
type FoodRestaurantLabelColumns struct {
	RestaurantId string // 餐厅ID
	LabelId      string // 标签ID
}

// foodRestaurantLabelColumns holds the columns for the table hg_food_restaurant_label.
var foodRestaurantLabelColumns = FoodRestaurantLabelColumns{
	RestaurantId: "restaurant_id",
	LabelId:      "label_id",
}

// NewFoodRestaurantLabelDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantLabelDao() *FoodRestaurantLabelDao {
	return &FoodRestaurantLabelDao{
		group:   "default",
		table:   "hg_food_restaurant_label",
		columns: foodRestaurantLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantLabelDao) Columns() FoodRestaurantLabelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
