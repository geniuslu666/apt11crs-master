// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsChargeDao is the data access object for table hg_pms_charge.
type PmsChargeDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PmsChargeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsChargeColumns defines and stores column names for table hg_pms_charge.
type PmsChargeColumns struct {
	Id             string // 主键
	Uid            string // 三方系统 ID
	AirUid         string // airhost uid
	Date           string // 费用日期，可以是发生日期或记账日期
	Name           string // 名称
	FeeType        string // 费用类型（预订费、餐饮费、清洁费、取消费、其他费用）
	Amount         string // 金额
	Currency       string // 货币
	OriginalAmount string // 原价
	Description    string // 费用详情
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// pmsChargeColumns holds the columns for table hg_pms_charge.
var pmsChargeColumns = PmsChargeColumns{
	Id:             "id",
	Uid:            "uid",
	AirUid:         "air_uid",
	Date:           "date",
	Name:           "name",
	FeeType:        "fee_type",
	Amount:         "amount",
	Currency:       "currency",
	OriginalAmount: "original_amount",
	Description:    "description",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPmsChargeDao creates and returns a new DAO object for table data access.
func NewPmsChargeDao() *PmsChargeDao {
	return &PmsChargeDao{
		group:   "default",
		table:   "hg_pms_charge",
		columns: pmsChargeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmsChargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmsChargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmsChargeDao) Columns() PmsChargeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmsChargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmsChargeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmsChargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
