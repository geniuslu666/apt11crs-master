// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FoodGoodsLabelDao is the data access object for the table hg_food_goods_label.
type FoodGoodsLabelDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns FoodGoodsLabelColumns // columns contains all the column names of Table for convenient usage.
}

// FoodGoodsLabelColumns defines and stores column names for the table hg_food_goods_label.
type FoodGoodsLabelColumns struct {
	GoodsId string // 餐厅ID
	LabelId string // 标签ID
}

// foodGoodsLabelColumns holds the columns for the table hg_food_goods_label.
var foodGoodsLabelColumns = FoodGoodsLabelColumns{
	GoodsId: "goods_id",
	LabelId: "label_id",
}

// NewFoodGoodsLabelDao creates and returns a new DAO object for table data access.
func NewFoodGoodsLabelDao() *FoodGoodsLabelDao {
	return &FoodGoodsLabelDao{
		group:   "default",
		table:   "hg_food_goods_label",
		columns: foodGoodsLabelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FoodGoodsLabelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FoodGoodsLabelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FoodGoodsLabelDao) Columns() FoodGoodsLabelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FoodGoodsLabelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FoodGoodsLabelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FoodGoodsLabelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
