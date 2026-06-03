// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelVerifyStaffScopeDao is the data access object for table hg_travel_verify_staff_scope.
type TravelVerifyStaffScopeDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns TravelVerifyStaffScopeColumns // columns contains all the column names of Table for convenient usage.
}

// TravelVerifyStaffScopeColumns defines and stores column names for table hg_travel_verify_staff_scope.
type TravelVerifyStaffScopeColumns struct {
	Id        string //
	StaffId   string // 核销人员ID
	ProductId string // 产品ID
	SkuId     string // SKU ID（0=产品全部SKU）
	IsAll     string // 是否全部可核销（1是 0否）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 软删除时间（NULL=正常）
}

// travelVerifyStaffScopeColumns holds the columns for table hg_travel_verify_staff_scope.
var travelVerifyStaffScopeColumns = TravelVerifyStaffScopeColumns{
	Id:        "id",
	StaffId:   "staff_id",
	ProductId: "product_id",
	SkuId:     "sku_id",
	IsAll:     "is_all",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewTravelVerifyStaffScopeDao creates and returns a new DAO object for table data access.
func NewTravelVerifyStaffScopeDao() *TravelVerifyStaffScopeDao {
	return &TravelVerifyStaffScopeDao{
		group:   "default",
		table:   "hg_travel_verify_staff_scope",
		columns: travelVerifyStaffScopeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelVerifyStaffScopeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelVerifyStaffScopeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelVerifyStaffScopeDao) Columns() TravelVerifyStaffScopeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelVerifyStaffScopeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelVerifyStaffScopeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelVerifyStaffScopeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
