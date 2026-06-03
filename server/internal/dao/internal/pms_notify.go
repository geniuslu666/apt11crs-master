// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsNotifyDao is the data access object for the table hg_pms_notify.
type PmsNotifyDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns PmsNotifyColumns // columns contains all the column names of Table for convenient usage.
}

// PmsNotifyColumns defines and stores column names for the table hg_pms_notify.
type PmsNotifyColumns struct {
	Id            string //
	MemberId      string // 会员ID
	NotifyTitle   string // 通知标题
	NotifyType    string // 通知类型 SYSTEM、系统消息    BOOKING、预定消息
	NotifyContent string // 通知内容
	NotifyData    string // 通知数据
	IsRead        string // Y 已读  N 未读
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// pmsNotifyColumns holds the columns for the table hg_pms_notify.
var pmsNotifyColumns = PmsNotifyColumns{
	Id:            "id",
	MemberId:      "member_id",
	NotifyTitle:   "notify_title",
	NotifyType:    "notify_type",
	NotifyContent: "notify_content",
	NotifyData:    "notify_data",
	IsRead:        "is_read",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPmsNotifyDao creates and returns a new DAO object for table data access.
func NewPmsNotifyDao() *PmsNotifyDao {
	return &PmsNotifyDao{
		group:   "default",
		table:   "hg_pms_notify",
		columns: pmsNotifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsNotifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsNotifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsNotifyDao) Columns() PmsNotifyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsNotifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsNotifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsNotifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
