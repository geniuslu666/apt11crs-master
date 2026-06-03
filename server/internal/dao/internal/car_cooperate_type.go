// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarCooperateTypeDao is the data access object for the table hg_car_cooperate_type.
type CarCooperateTypeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns CarCooperateTypeColumns // columns contains all the column names of Table for convenient usage.
}

// CarCooperateTypeColumns defines and stores column names for the table hg_car_cooperate_type.
type CarCooperateTypeColumns struct {
	Id        string //
	TypeName  string //
	Status    string // 状态1、启用 2、禁用
	IsThird   string // 是否第三方 1-是 2-否
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// carCooperateTypeColumns holds the columns for the table hg_car_cooperate_type.
var carCooperateTypeColumns = CarCooperateTypeColumns{
	Id:        "id",
	TypeName:  "type_name",
	Status:    "status",
	IsThird:   "is_third",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewCarCooperateTypeDao creates and returns a new DAO object for table data access.
func NewCarCooperateTypeDao() *CarCooperateTypeDao {
	return &CarCooperateTypeDao{
		group:   "default",
		table:   "hg_car_cooperate_type",
		columns: carCooperateTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarCooperateTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarCooperateTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarCooperateTypeDao) Columns() CarCooperateTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarCooperateTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarCooperateTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarCooperateTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
