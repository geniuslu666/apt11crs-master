// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsCollectDao is the data access object for the table hg_pms_collect.
type PmsCollectDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns PmsCollectColumns // columns contains all the column names of Table for convenient usage.
}

// PmsCollectColumns defines and stores column names for the table hg_pms_collect.
type PmsCollectColumns struct {
	Id          string //
	MemberId    string //
	CollectType string // 收藏类型   HOST 、 酒店  REPAST 餐饮
	CollectId   string // 收藏数据ID
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// pmsCollectColumns holds the columns for the table hg_pms_collect.
var pmsCollectColumns = PmsCollectColumns{
	Id:          "id",
	MemberId:    "member_id",
	CollectType: "collect_type",
	CollectId:   "collect_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewPmsCollectDao creates and returns a new DAO object for table data access.
func NewPmsCollectDao() *PmsCollectDao {
	return &PmsCollectDao{
		group:   "default",
		table:   "hg_pms_collect",
		columns: pmsCollectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsCollectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsCollectDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsCollectDao) Columns() PmsCollectColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsCollectDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsCollectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsCollectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
