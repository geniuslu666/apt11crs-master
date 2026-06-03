// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAttachmentKindDao is the data access object for the table hg_sys_attachment_kind.
type SysAttachmentKindDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns SysAttachmentKindColumns // columns contains all the column names of Table for convenient usage.
}

// SysAttachmentKindColumns defines and stores column names for the table hg_sys_attachment_kind.
type SysAttachmentKindColumns struct {
	Id    string //
	Label string // 附件分类名称
	Key   string // 分类key
	Value string // 分类value
	Icon  string // 分类图标
	Tag   string //
}

// sysAttachmentKindColumns holds the columns for the table hg_sys_attachment_kind.
var sysAttachmentKindColumns = SysAttachmentKindColumns{
	Id:    "id",
	Label: "label",
	Key:   "key",
	Value: "value",
	Icon:  "icon",
	Tag:   "tag",
}

// NewSysAttachmentKindDao creates and returns a new DAO object for table data access.
func NewSysAttachmentKindDao() *SysAttachmentKindDao {
	return &SysAttachmentKindDao{
		group:   "default",
		table:   "hg_sys_attachment_kind",
		columns: sysAttachmentKindColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAttachmentKindDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAttachmentKindDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAttachmentKindDao) Columns() SysAttachmentKindColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAttachmentKindDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAttachmentKindDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysAttachmentKindDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
