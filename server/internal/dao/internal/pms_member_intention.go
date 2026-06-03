// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberIntentionDao is the data access object for the table hg_pms_member_intention.
type PmsMemberIntentionDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns PmsMemberIntentionColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberIntentionColumns defines and stores column names for the table hg_pms_member_intention.
type PmsMemberIntentionColumns struct {
	Id        string // 主键ID
	MemberId  string // 意向用户ID
	Name      string //
	Sex       string // 1、男 2、女
	Country   string //
	Phone     string //
	PhoneArea string //
	Language  string //
	Mail      string //
	Remark    string // 备注
	WechatNo  string // 微信号
	LineId    string // line_id
	CreateAt  string // 创建时间
	UpdateAt  string // 修改时间
	DeletedAt string // 删除时间
}

// pmsMemberIntentionColumns holds the columns for the table hg_pms_member_intention.
var pmsMemberIntentionColumns = PmsMemberIntentionColumns{
	Id:        "id",
	MemberId:  "member_id",
	Name:      "name",
	Sex:       "sex",
	Country:   "country",
	Phone:     "phone",
	PhoneArea: "phone_area",
	Language:  "language",
	Mail:      "mail",
	Remark:    "remark",
	WechatNo:  "wechat_no",
	LineId:    "line_id",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
	DeletedAt: "deleted_at",
}

// NewPmsMemberIntentionDao creates and returns a new DAO object for table data access.
func NewPmsMemberIntentionDao() *PmsMemberIntentionDao {
	return &PmsMemberIntentionDao{
		group:   "default",
		table:   "hg_pms_member_intention",
		columns: pmsMemberIntentionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberIntentionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberIntentionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberIntentionDao) Columns() PmsMemberIntentionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberIntentionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberIntentionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberIntentionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
