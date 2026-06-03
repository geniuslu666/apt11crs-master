// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSmsTemplateDao is the data access object for the table hg_sys_sms_template.
type SysSmsTemplateDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns SysSmsTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// SysSmsTemplateColumns defines and stores column names for the table hg_sys_sms_template.
type SysSmsTemplateColumns struct {
	Id              string // 主键
	Scene           string // 1会员  2住宿  3接送机  4按摩
	Event           string // 别名
	Title           string // 标题
	Content         string // 内容（多语言）
	UmsTemplate     string // 一信通模版json（id：模版id， content：模版内容）
	TencentTemplate string // 腾讯云模版json
	AliyunTemplate  string // 阿里云模版json
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string //
}

// sysSmsTemplateColumns holds the columns for the table hg_sys_sms_template.
var sysSmsTemplateColumns = SysSmsTemplateColumns{
	Id:              "id",
	Scene:           "scene",
	Event:           "event",
	Title:           "title",
	Content:         "content",
	UmsTemplate:     "ums_template",
	TencentTemplate: "tencent_template",
	AliyunTemplate:  "aliyun_template",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewSysSmsTemplateDao creates and returns a new DAO object for table data access.
func NewSysSmsTemplateDao() *SysSmsTemplateDao {
	return &SysSmsTemplateDao{
		group:   "default",
		table:   "hg_sys_sms_template",
		columns: sysSmsTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysSmsTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysSmsTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysSmsTemplateDao) Columns() SysSmsTemplateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysSmsTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysSmsTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysSmsTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
