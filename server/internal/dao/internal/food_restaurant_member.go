// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantMemberDao is the data access object for the table hg_food_restaurant_member.
type FoodRestaurantMemberDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns FoodRestaurantMemberColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantMemberColumns defines and stores column names for the table hg_food_restaurant_member.
type FoodRestaurantMemberColumns struct {
	RestaurantId string // 餐厅ID
	MemberId     string // 用户ID
}

// foodRestaurantMemberColumns holds the columns for the table hg_food_restaurant_member.
var foodRestaurantMemberColumns = FoodRestaurantMemberColumns{
	RestaurantId: "restaurant_id",
	MemberId:     "member_id",
}

// NewFoodRestaurantMemberDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantMemberDao() *FoodRestaurantMemberDao {
	return &FoodRestaurantMemberDao{
		group:   "default",
		table:   "hg_food_restaurant_member",
		columns: foodRestaurantMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantMemberDao) Columns() FoodRestaurantMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
