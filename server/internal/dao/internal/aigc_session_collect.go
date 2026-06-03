// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AigcSessionCollectDao is the data access object for the table aigc_session_collect.
type AigcSessionCollectDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns AigcSessionCollectColumns // columns contains all the column names of Table for convenient usage.
}

// AigcSessionCollectColumns defines and stores column names for the table aigc_session_collect.
type AigcSessionCollectColumns struct {
	Id        string //
	Title     string // 集合标题
	CreatedAt string // 创建时间
	KefuName  string // 客服名称
	EntId     string // 企业ID
}

// aigcSessionCollectColumns holds the columns for the table aigc_session_collect.
var aigcSessionCollectColumns = AigcSessionCollectColumns{
	Id:        "id",
	Title:     "title",
	CreatedAt: "created_at",
	KefuName:  "kefu_name",
	EntId:     "ent_id",
}

// NewAigcSessionCollectDao creates and returns a new DAO object for table data access.
func NewAigcSessionCollectDao() *AigcSessionCollectDao {
	return &AigcSessionCollectDao{
		group:   "default",
		table:   "aigc_session_collect",
		columns: aigcSessionCollectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AigcSessionCollectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AigcSessionCollectDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AigcSessionCollectDao) Columns() AigcSessionCollectColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AigcSessionCollectDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AigcSessionCollectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AigcSessionCollectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
