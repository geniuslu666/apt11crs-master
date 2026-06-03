// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberAuthDao is the data access object for the table hg_pms_member_auth.
type PmsMemberAuthDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsMemberAuthColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberAuthColumns defines and stores column names for the table hg_pms_member_auth.
type PmsMemberAuthColumns struct {
	Id        string //
	AuthId    string // 第三方用户ID标识
	WxUnionid string // 微信unionid
	Channel   string // WX、APPLE
	MemberId  string // 关联的会员ID
	Email     string // 邮箱
	Phone     string // 手机号
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	IsFxAuth  string // 是否是分销授权用户
}

// pmsMemberAuthColumns holds the columns for the table hg_pms_member_auth.
var pmsMemberAuthColumns = PmsMemberAuthColumns{
	Id:        "id",
	AuthId:    "auth_id",
	WxUnionid: "wx_unionid",
	Channel:   "channel",
	MemberId:  "member_id",
	Email:     "email",
	Phone:     "phone",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	IsFxAuth:  "is_fx_auth",
}

// NewPmsMemberAuthDao creates and returns a new DAO object for table data access.
func NewPmsMemberAuthDao() *PmsMemberAuthDao {
	return &PmsMemberAuthDao{
		group:   "default",
		table:   "hg_pms_member_auth",
		columns: pmsMemberAuthColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberAuthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberAuthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberAuthDao) Columns() PmsMemberAuthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberAuthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberAuthDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberAuthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
