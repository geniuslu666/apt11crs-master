// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberLogDao is the data access object for the table hg_pms_member_log.
type PmsMemberLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsMemberLogColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberLogColumns defines and stores column names for the table hg_pms_member_log.
type PmsMemberLogColumns struct {
	Id        string //
	MemberId  string // 会员ID
	LoginTime string // 登录时间
	LoginType string // 登录方式
	LoginIp   string // 登录IP
	MdCode    string // 登录设备码
	MpModel   string // 登录设备型号
	Token     string // 登录token
	ExpirTime string // 过期时间
	CreatedAt string //
	UpdatedAt string //
}

// pmsMemberLogColumns holds the columns for the table hg_pms_member_log.
var pmsMemberLogColumns = PmsMemberLogColumns{
	Id:        "id",
	MemberId:  "member_id",
	LoginTime: "login_time",
	LoginType: "login_type",
	LoginIp:   "login_ip",
	MdCode:    "md_code",
	MpModel:   "mp_model",
	Token:     "token",
	ExpirTime: "expir_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPmsMemberLogDao creates and returns a new DAO object for table data access.
func NewPmsMemberLogDao() *PmsMemberLogDao {
	return &PmsMemberLogDao{
		group:   "default",
		table:   "hg_pms_member_log",
		columns: pmsMemberLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberLogDao) Columns() PmsMemberLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
