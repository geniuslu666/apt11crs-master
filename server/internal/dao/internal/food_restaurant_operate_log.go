// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantOperateLogDao is the data access object for the table hg_food_restaurant_operate_log.
type FoodRestaurantOperateLogDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of the current DAO.
	columns FoodRestaurantOperateLogColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantOperateLogColumns defines and stores column names for the table hg_food_restaurant_operate_log.
type FoodRestaurantOperateLogColumns struct {
	Id              string //
	RestaurantId    string // 餐厅ID
	OperateType     string // 操作类型
	OperateMemberId string // 后台操作人ID
	Remark          string // 备注
	CreatedAt       string // 创建时间
	UpdatedAt       string // 修改时间
}

// foodRestaurantOperateLogColumns holds the columns for the table hg_food_restaurant_operate_log.
var foodRestaurantOperateLogColumns = FoodRestaurantOperateLogColumns{
	Id:              "id",
	RestaurantId:    "restaurant_id",
	OperateType:     "operate_type",
	OperateMemberId: "operate_member_id",
	Remark:          "remark",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewFoodRestaurantOperateLogDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantOperateLogDao() *FoodRestaurantOperateLogDao {
	return &FoodRestaurantOperateLogDao{
		group:   "default",
		table:   "hg_food_restaurant_operate_log",
		columns: foodRestaurantOperateLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantOperateLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantOperateLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantOperateLogDao) Columns() FoodRestaurantOperateLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantOperateLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantOperateLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantOperateLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
