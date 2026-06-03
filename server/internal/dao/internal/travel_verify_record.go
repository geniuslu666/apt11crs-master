// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelVerifyRecordDao is the data access object for table hg_travel_verify_record.
type TravelVerifyRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns TravelVerifyRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TravelVerifyRecordColumns defines and stores column names for table hg_travel_verify_record.
type TravelVerifyRecordColumns struct {
	Id            string //
	OrderId       string // 订单ID
	OrderSn       string // 预约单号
	ProductId     string // 产品ID
	MemberId      string // 会员ID
	BookDate      string // 预约日期
	VerifyStaffId string // 核销人员ID
	VerifyTime    string // 核销时间
	CreatedAt     string // 创建时间
}

// travelVerifyRecordColumns holds the columns for table hg_travel_verify_record.
var travelVerifyRecordColumns = TravelVerifyRecordColumns{
	Id:            "id",
	OrderId:       "order_id",
	OrderSn:       "order_sn",
	ProductId:     "product_id",
	MemberId:      "member_id",
	BookDate:      "book_date",
	VerifyStaffId: "verify_staff_id",
	VerifyTime:    "verify_time",
	CreatedAt:     "created_at",
}

// NewTravelVerifyRecordDao creates and returns a new DAO object for table data access.
func NewTravelVerifyRecordDao() *TravelVerifyRecordDao {
	return &TravelVerifyRecordDao{
		group:   "default",
		table:   "hg_travel_verify_record",
		columns: travelVerifyRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelVerifyRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelVerifyRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelVerifyRecordDao) Columns() TravelVerifyRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelVerifyRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelVerifyRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelVerifyRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
