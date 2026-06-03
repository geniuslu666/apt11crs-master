// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberCancelDao is the data access object for the table hg_pms_member_cancel.
type PmsMemberCancelDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns PmsMemberCancelColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberCancelColumns defines and stores column names for the table hg_pms_member_cancel.
type PmsMemberCancelColumns struct {
	Id           string // 主键
	MemberId     string // 用户ID
	MemberNo     string // 会员号
	Phone        string // 手机号
	PhoneArea    string // 手机区号
	Mail         string // 邮箱
	CancelReason string // 注销原因
	OperatorId   string // 操作员ID
	AuditStatus  string // 审核状态 1-待审核 2-审核通过 3-审核拒绝
	AuditReason  string // 审核通过/拒绝原因
	AuditTime    string // 审核时间
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// pmsMemberCancelColumns holds the columns for the table hg_pms_member_cancel.
var pmsMemberCancelColumns = PmsMemberCancelColumns{
	Id:           "id",
	MemberId:     "member_id",
	MemberNo:     "member_no",
	Phone:        "phone",
	PhoneArea:    "phone_area",
	Mail:         "mail",
	CancelReason: "cancel_reason",
	OperatorId:   "operator_id",
	AuditStatus:  "audit_status",
	AuditReason:  "audit_reason",
	AuditTime:    "audit_time",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewPmsMemberCancelDao creates and returns a new DAO object for table data access.
func NewPmsMemberCancelDao() *PmsMemberCancelDao {
	return &PmsMemberCancelDao{
		group:   "default",
		table:   "hg_pms_member_cancel",
		columns: pmsMemberCancelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberCancelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberCancelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberCancelDao) Columns() PmsMemberCancelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberCancelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberCancelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberCancelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
