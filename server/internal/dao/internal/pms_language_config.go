// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsLanguageConfigDao is the data access object for the table hg_pms_language_config.
type PmsLanguageConfigDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns PmsLanguageConfigColumns // columns contains all the column names of Table for convenient usage.
}

// PmsLanguageConfigColumns defines and stores column names for the table hg_pms_language_config.
type PmsLanguageConfigColumns struct {
	Id       string //
	Tag      string // 语言标签
	Name     string // 语言名称
	BaiduTag string // 百度翻译语种
	Flag     string // 国旗
	CreateAt string //
	UpdateAt string //
}

// pmsLanguageConfigColumns holds the columns for the table hg_pms_language_config.
var pmsLanguageConfigColumns = PmsLanguageConfigColumns{
	Id:       "id",
	Tag:      "tag",
	Name:     "name",
	BaiduTag: "baidu_tag",
	Flag:     "flag",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewPmsLanguageConfigDao creates and returns a new DAO object for table data access.
func NewPmsLanguageConfigDao() *PmsLanguageConfigDao {
	return &PmsLanguageConfigDao{
		group:   "default",
		table:   "hg_pms_language_config",
		columns: pmsLanguageConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsLanguageConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsLanguageConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsLanguageConfigDao) Columns() PmsLanguageConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsLanguageConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsLanguageConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsLanguageConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
