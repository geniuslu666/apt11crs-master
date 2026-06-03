// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThMchDao is the data access object for the table hg_th_mch.
type ThMchDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns ThMchColumns // columns contains all the column names of Table for convenient usage.
}

// ThMchColumns defines and stores column names for the table hg_th_mch.
type ThMchColumns struct {
	Id          string //
	CategoryId  string // 分类ID
	Name        string // 名称
	Logo        string // LOGO
	ContactInfo string // 联系信息
	Sort        string // 排序(越大越靠前)
	Status      string // 1、启用 2、禁用
	StoreOnNum  string // 启用中门店数量
	StoreOffNum string // 禁用中门店数量
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// thMchColumns holds the columns for the table hg_th_mch.
var thMchColumns = ThMchColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Name:        "name",
	Logo:        "logo",
	ContactInfo: "contact_info",
	Sort:        "sort",
	Status:      "status",
	StoreOnNum:  "store_on_num",
	StoreOffNum: "store_off_num",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewThMchDao creates and returns a new DAO object for table data access.
func NewThMchDao() *ThMchDao {
	return &ThMchDao{
		group:   "default",
		table:   "hg_th_mch",
		columns: thMchColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThMchDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThMchDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThMchDao) Columns() ThMchColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThMchDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThMchDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThMchDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
