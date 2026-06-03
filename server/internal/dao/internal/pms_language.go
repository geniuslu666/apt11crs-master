// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsLanguageDao is the data access object for the table hg_pms_language.
type PmsLanguageDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PmsLanguageColumns // columns contains all the column names of Table for convenient usage.
}

// PmsLanguageColumns defines and stores column names for the table hg_pms_language.
type PmsLanguageColumns struct {
	Id       string //
	Uuid     string // 标签ID
	Tag      string // 标签  type = table  、 数据库表名     type = 其他的话  是前端的属性名
	Type     string // 类型  table/数据库表  manage/管理端  mobile/移动端
	Key      string // 字段标识
	Language string // 语言
	Content  string // 语言内容
	CreateAt string // 创建时间
	UpdateAt string // 更新时间
}

// pmsLanguageColumns holds the columns for the table hg_pms_language.
var pmsLanguageColumns = PmsLanguageColumns{
	Id:       "id",
	Uuid:     "uuid",
	Tag:      "tag",
	Type:     "type",
	Key:      "key",
	Language: "language",
	Content:  "content",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewPmsLanguageDao creates and returns a new DAO object for table data access.
func NewPmsLanguageDao() *PmsLanguageDao {
	return &PmsLanguageDao{
		group:   "default",
		table:   "hg_pms_language",
		columns: pmsLanguageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsLanguageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsLanguageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsLanguageDao) Columns() PmsLanguageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsLanguageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsLanguageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsLanguageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
