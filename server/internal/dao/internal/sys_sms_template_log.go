// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysSmsTemplateLogDao is the data access object for the table hg_sys_sms_template_log.
type SysSmsTemplateLogDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns SysSmsTemplateLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysSmsTemplateLogColumns defines and stores column names for the table hg_sys_sms_template_log.
type SysSmsTemplateLogColumns struct {
	Id             string // 主键
	TemplateId     string // 系统模版表中主键ID
	SendTemplateId string // 发送的模版ID
	Type           string // 1短信   2邮件   3推送
	To             string // 短信或邮件或推送头
	VipId          string // 会员ID
	Content        string // 发送内容
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string //
}

// sysSmsTemplateLogColumns holds the columns for the table hg_sys_sms_template_log.
var sysSmsTemplateLogColumns = SysSmsTemplateLogColumns{
	Id:             "id",
	TemplateId:     "template_id",
	SendTemplateId: "send_template_id",
	Type:           "type",
	To:             "to",
	VipId:          "vip_id",
	Content:        "content",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewSysSmsTemplateLogDao creates and returns a new DAO object for table data access.
func NewSysSmsTemplateLogDao() *SysSmsTemplateLogDao {
	return &SysSmsTemplateLogDao{
		group:   "default",
		table:   "hg_sys_sms_template_log",
		columns: sysSmsTemplateLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysSmsTemplateLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysSmsTemplateLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysSmsTemplateLogDao) Columns() SysSmsTemplateLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysSmsTemplateLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysSmsTemplateLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysSmsTemplateLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
