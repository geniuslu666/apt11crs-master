// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaCooperateTypeDao is the data access object for the table hg_spa_cooperate_type.
type SpaCooperateTypeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns SpaCooperateTypeColumns // columns contains all the column names of Table for convenient usage.
}

// SpaCooperateTypeColumns defines and stores column names for the table hg_spa_cooperate_type.
type SpaCooperateTypeColumns struct {
	Id        string //
	TypeName  string // 营业类型
	Status    string // 状态1、启用 2、禁用
	CreateAt  string // 创建时间
	UpdateAt  string // 更新时间
	DeletedAt string // 删除时间
}

// spaCooperateTypeColumns holds the columns for the table hg_spa_cooperate_type.
var spaCooperateTypeColumns = SpaCooperateTypeColumns{
	Id:        "id",
	TypeName:  "type_name",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewSpaCooperateTypeDao creates and returns a new DAO object for table data access.
func NewSpaCooperateTypeDao() *SpaCooperateTypeDao {
	return &SpaCooperateTypeDao{
		group:   "default",
		table:   "hg_spa_cooperate_type",
		columns: spaCooperateTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaCooperateTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaCooperateTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaCooperateTypeDao) Columns() SpaCooperateTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaCooperateTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaCooperateTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaCooperateTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
