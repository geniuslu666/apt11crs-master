// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VirtualProductDao is the data access object for the table virtual_product.
type VirtualProductDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns VirtualProductColumns // columns contains all the column names of Table for convenient usage.
}

// VirtualProductColumns defines and stores column names for the table virtual_product.
type VirtualProductColumns struct {
	Id              string //
	EntId           string // 企业ID
	KefuName        string // 客服账户
	ProductName     string //
	ProductCategory string //
	Payment         string // 支付方式，wechat 微信支付；nan66 南星码支付
	Description     string //
	Price           string // 金额
	ProductImg      string //
	ResourceLink    string //
	IsActive        string // 在线状态，1在售，2下架
	CreatedAt       string //
	UpdatedAt       string //
}

// virtualProductColumns holds the columns for the table virtual_product.
var virtualProductColumns = VirtualProductColumns{
	Id:              "id",
	EntId:           "ent_id",
	KefuName:        "kefu_name",
	ProductName:     "product_name",
	ProductCategory: "product_category",
	Payment:         "payment",
	Description:     "description",
	Price:           "price",
	ProductImg:      "product_img",
	ResourceLink:    "resource_link",
	IsActive:        "is_active",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewVirtualProductDao creates and returns a new DAO object for table data access.
func NewVirtualProductDao() *VirtualProductDao {
	return &VirtualProductDao{
		group:   "default",
		table:   "virtual_product",
		columns: virtualProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VirtualProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VirtualProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VirtualProductDao) Columns() VirtualProductColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VirtualProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VirtualProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VirtualProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
