// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAppReservationChangeDao is the data access object for the table hg_pms_app_reservation_change.
type PmsAppReservationChangeDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of the current DAO.
	columns PmsAppReservationChangeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAppReservationChangeColumns defines and stores column names for the table hg_pms_app_reservation_change.
type PmsAppReservationChangeColumns struct {
	Id              string // 主键
	ChangeOrderSn   string // 变更订单号
	OrderId         string // 订单ID
	OrderSn         string // 预订订单号
	OutOrderSn      string // 预订外部订单号
	ChangeType      string // 变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更
	OldCheckinDate  string // 入住日期
	OldCheckoutDate string // 退房日期
	NewCheckinDate  string // 入住日期
	NewCheckoutDate string // 退房日期
	OldMainGuest    string // 住宿人编号
	NewMainGuest    string // 住宿人编号
	OldAdultCount   string // 成人数量
	NewAdultCount   string // 成人数量
	OldChildCount   string // 儿童数量
	NewChildCount   string // 儿童数量
	OldInfantCount  string // 婴儿数量
	NewInfantCount  string // 婴儿数量
	ChangeStatus    string // 变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败
	ChangeAmount    string // 变动金额
	SubmitDate      string // 提交变更时间
	DoneDate        string // 变更成功时间
	OldOrderPrice   string // 原订单价格
	NewOrderPrice   string // 变更后订单价格
	ExpirationTime  string // 订单过期时间
	CreatedAt       string //
	UpdatedAt       string //
	PricePercent    string // 全局溢价比例
	IsFx            string // 是否是分销订单
}

// pmsAppReservationChangeColumns holds the columns for the table hg_pms_app_reservation_change.
var pmsAppReservationChangeColumns = PmsAppReservationChangeColumns{
	Id:              "id",
	ChangeOrderSn:   "change_order_sn",
	OrderId:         "order_id",
	OrderSn:         "order_sn",
	OutOrderSn:      "out_order_sn",
	ChangeType:      "change_type",
	OldCheckinDate:  "old_checkin_date",
	OldCheckoutDate: "old_checkout_date",
	NewCheckinDate:  "new_checkin_date",
	NewCheckoutDate: "new_checkout_date",
	OldMainGuest:    "old_main_guest",
	NewMainGuest:    "new_main_guest",
	OldAdultCount:   "old_adult_count",
	NewAdultCount:   "new_adult_count",
	OldChildCount:   "old_child_count",
	NewChildCount:   "new_child_count",
	OldInfantCount:  "old_infant_count",
	NewInfantCount:  "new_infant_count",
	ChangeStatus:    "change_status",
	ChangeAmount:    "change_amount",
	SubmitDate:      "submit_date",
	DoneDate:        "done_date",
	OldOrderPrice:   "old_order_price",
	NewOrderPrice:   "new_order_price",
	ExpirationTime:  "expiration_time",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	PricePercent:    "price_percent",
	IsFx:            "is_fx",
}

// NewPmsAppReservationChangeDao creates and returns a new DAO object for table data access.
func NewPmsAppReservationChangeDao() *PmsAppReservationChangeDao {
	return &PmsAppReservationChangeDao{
		group:   "default",
		table:   "hg_pms_app_reservation_change",
		columns: pmsAppReservationChangeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsAppReservationChangeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsAppReservationChangeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsAppReservationChangeDao) Columns() PmsAppReservationChangeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsAppReservationChangeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsAppReservationChangeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsAppReservationChangeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
