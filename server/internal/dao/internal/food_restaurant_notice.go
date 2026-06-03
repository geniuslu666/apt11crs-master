// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodRestaurantNoticeDao is the data access object for the table hg_food_restaurant_notice.
type FoodRestaurantNoticeDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns FoodRestaurantNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// FoodRestaurantNoticeColumns defines and stores column names for the table hg_food_restaurant_notice.
type FoodRestaurantNoticeColumns struct {
	Id              string //
	RestaurantId    string // 餐厅ID
	Title           string // 通知标题
	Content         string // 公告内容
	Sort            string // 排序(越大越靠前)
	Status          string // 状态1、启用 2、禁用
	NeedUserConfirm string // 是否需要用户确认 1、需要  2、不需要
	CreateAt        string // 创建时间
	UpdateAt        string // 更新时间
	DeletedAt       string // 删除时间
}

// foodRestaurantNoticeColumns holds the columns for the table hg_food_restaurant_notice.
var foodRestaurantNoticeColumns = FoodRestaurantNoticeColumns{
	Id:              "id",
	RestaurantId:    "restaurant_id",
	Title:           "title",
	Content:         "content",
	Sort:            "sort",
	Status:          "status",
	NeedUserConfirm: "need_user_confirm",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
	DeletedAt:       "deleted_at",
}

// NewFoodRestaurantNoticeDao creates and returns a new DAO object for table data access.
func NewFoodRestaurantNoticeDao() *FoodRestaurantNoticeDao {
	return &FoodRestaurantNoticeDao{
		group:   "default",
		table:   "hg_food_restaurant_notice",
		columns: foodRestaurantNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodRestaurantNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodRestaurantNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodRestaurantNoticeDao) Columns() FoodRestaurantNoticeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodRestaurantNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodRestaurantNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodRestaurantNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
