// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CustomerDao is the data access object for the table customer.
type CustomerDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns CustomerColumns // columns contains all the column names of Table for convenient usage.
}

// CustomerColumns defines and stores column names for the table customer.
type CustomerColumns struct {
	Id           string //
	Name         string // 会员名称
	Avatar       string // 会员头像
	KefuName     string // 客服账户
	Tel          string // 会员手机号
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	Openid       string // 会员ID
	AcountOpenid string // 公众号会员ID
	Extra        string // 会员扩展信息
	EntId        string // 对接的企业ID
	Score        string // 会员积分
}

// customerColumns holds the columns for the table customer.
var customerColumns = CustomerColumns{
	Id:           "id",
	Name:         "name",
	Avatar:       "avatar",
	KefuName:     "kefu_name",
	Tel:          "tel",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Openid:       "openid",
	AcountOpenid: "acount_openid",
	Extra:        "extra",
	EntId:        "ent_id",
	Score:        "score",
}

// NewCustomerDao creates and returns a new DAO object for table data access.
func NewCustomerDao() *CustomerDao {
	return &CustomerDao{
		group:   "default",
		table:   "customer",
		columns: customerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CustomerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CustomerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CustomerDao) Columns() CustomerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CustomerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CustomerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CustomerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
