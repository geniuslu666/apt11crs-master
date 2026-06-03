// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTerminalVerifyDao is the data access object for the table hg_sys_terminal_verify.
type SysTerminalVerifyDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns SysTerminalVerifyColumns // columns contains all the column names of Table for convenient usage.
}

// SysTerminalVerifyColumns defines and stores column names for the table hg_sys_terminal_verify.
type SysTerminalVerifyColumns struct {
	Id             string //
	TerminalId     string // 终端ID
	VerifyType     string // 核销类型
	MchId          string // 商户ID
	VerifyMemberId string // 核销用户ID
	StoreId        string // 门店ID
	RestaurantId   string // 餐厅ID
	FoodOrderId    string // 餐厅订单ID
	MemberCouponId string // 用户礼品券ID
	CouponMchName  string // 核销商品名
	VerifyTime     string // 核销时间
	CreateAt       string // 创建时间
	UpdateAt       string // 更新时间
	DeletedAt      string //
}

// sysTerminalVerifyColumns holds the columns for the table hg_sys_terminal_verify.
var sysTerminalVerifyColumns = SysTerminalVerifyColumns{
	Id:             "id",
	TerminalId:     "terminal_id",
	VerifyType:     "verify_type",
	MchId:          "mch_id",
	VerifyMemberId: "verify_member_id",
	StoreId:        "store_id",
	RestaurantId:   "restaurant_id",
	FoodOrderId:    "food_order_id",
	MemberCouponId: "member_coupon_id",
	CouponMchName:  "coupon_mch_name",
	VerifyTime:     "verify_time",
	CreateAt:       "create_at",
	UpdateAt:       "update_at",
	DeletedAt:      "deleted_at",
}

// NewSysTerminalVerifyDao creates and returns a new DAO object for table data access.
func NewSysTerminalVerifyDao() *SysTerminalVerifyDao {
	return &SysTerminalVerifyDao{
		group:   "default",
		table:   "hg_sys_terminal_verify",
		columns: sysTerminalVerifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysTerminalVerifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysTerminalVerifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysTerminalVerifyDao) Columns() SysTerminalVerifyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysTerminalVerifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysTerminalVerifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysTerminalVerifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
