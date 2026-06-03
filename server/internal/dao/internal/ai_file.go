// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AiFileDao is the data access object for the table ai_file.
type AiFileDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns AiFileColumns // columns contains all the column names of Table for convenient usage.
}

// AiFileColumns defines and stores column names for the table ai_file.
type AiFileColumns struct {
	Id          string //
	FileName    string // 文件名
	CreatedAt   string // 创建时间
	FileSize    string // 字符数
	CollectName string // 集合名称
}

// aiFileColumns holds the columns for the table ai_file.
var aiFileColumns = AiFileColumns{
	Id:          "id",
	FileName:    "file_name",
	CreatedAt:   "created_at",
	FileSize:    "file_size",
	CollectName: "collect_name",
}

// NewAiFileDao creates and returns a new DAO object for table data access.
func NewAiFileDao() *AiFileDao {
	return &AiFileDao{
		group:   "default",
		table:   "ai_file",
		columns: aiFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AiFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AiFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AiFileDao) Columns() AiFileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AiFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AiFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AiFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
