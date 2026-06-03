// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsGuestProfileDao is the data access object for the table hg_pms_guest_profile.
type PmsGuestProfileDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns PmsGuestProfileColumns // columns contains all the column names of Table for convenient usage.
}

// PmsGuestProfileColumns defines and stores column names for the table hg_pms_guest_profile.
type PmsGuestProfileColumns struct {
	Id            string // 主键
	Uid           string // 三方系统 ID
	MemberId      string // 会员ID
	FirstName     string // 名
	LastName      string // 姓
	FirstNameKana string // 名的假名
	LastNameKana  string // 姓的假名
	FullName      string // 全名
	Language      string // 语言
	Email         string // 电子邮件
	Phone         string // 电话
	AreaNo        string // 电话国际区号
	Nationality   string // 国籍
	Address       string // 地址
	Password      string // 密码
	Register      string // Y 已注册  N 未注册
	RegisterAt    string // 注册时间
	CreatedAt     string //
	UpdatedAt     string //
}

// pmsGuestProfileColumns holds the columns for the table hg_pms_guest_profile.
var pmsGuestProfileColumns = PmsGuestProfileColumns{
	Id:            "id",
	Uid:           "uid",
	MemberId:      "member_id",
	FirstName:     "first_name",
	LastName:      "last_name",
	FirstNameKana: "first_name_kana",
	LastNameKana:  "last_name_kana",
	FullName:      "full_name",
	Language:      "language",
	Email:         "email",
	Phone:         "phone",
	AreaNo:        "area_no",
	Nationality:   "nationality",
	Address:       "address",
	Password:      "password",
	Register:      "register",
	RegisterAt:    "register_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewPmsGuestProfileDao creates and returns a new DAO object for table data access.
func NewPmsGuestProfileDao() *PmsGuestProfileDao {
	return &PmsGuestProfileDao{
		group:   "default",
		table:   "hg_pms_guest_profile",
		columns: pmsGuestProfileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsGuestProfileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsGuestProfileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsGuestProfileDao) Columns() PmsGuestProfileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsGuestProfileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsGuestProfileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsGuestProfileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
