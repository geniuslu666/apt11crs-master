// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CsFastContentDao is the data access object for the table hg_cs_fast_content.
type CsFastContentDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns CsFastContentColumns // columns contains all the column names of Table for convenient usage.
}

// CsFastContentColumns defines and stores column names for the table hg_cs_fast_content.
type CsFastContentColumns struct {
	Id          string //
	Type        string // 类型    welcome    欢迎语    option
	ZhContent   string // 中文简体内容
	ZhCnContent string // 中文繁体内容
	EnContent   string // 英文内容
	JaContent   string // 日文内容
	KoContent   string // 韩文内容
	Sort        string // 排序  从大到小排序
	Status      string // 启禁用
	CreateAt    string //
	UpdateAt    string //
}

// csFastContentColumns holds the columns for the table hg_cs_fast_content.
var csFastContentColumns = CsFastContentColumns{
	Id:          "id",
	Type:        "type",
	ZhContent:   "zh_content",
	ZhCnContent: "zh_cn_content",
	EnContent:   "en_content",
	JaContent:   "ja_content",
	KoContent:   "ko_content",
	Sort:        "sort",
	Status:      "status",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewCsFastContentDao creates and returns a new DAO object for table data access.
func NewCsFastContentDao() *CsFastContentDao {
	return &CsFastContentDao{
		group:   "default",
		table:   "hg_cs_fast_content",
		columns: csFastContentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CsFastContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CsFastContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CsFastContentDao) Columns() CsFastContentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CsFastContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CsFastContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CsFastContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
