// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThMemberCouponDao is the data access object for the table hg_th_member_coupon.
type ThMemberCouponDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns ThMemberCouponColumns // columns contains all the column names of Table for convenient usage.
}

// ThMemberCouponColumns defines and stores column names for the table hg_th_member_coupon.
type ThMemberCouponColumns struct {
	Id              string // 主键ID
	CouponNo        string // 券号
	CouponId        string // 券id
	MemberId        string // 领用人
	State           string // 状态 1待生效 2未使用 3已核销 4已过期  5已失效  6已回收
	InvalidTime     string // 失效时间
	VerifyTime      string // 核销时间
	VerifyMchId     string // 核销商户ID
	VerifyStoreId   string // 核销门店ID
	StartTime       string // 有效期开始时间
	EndTime         string // 有效期结束时间
	Source          string // 来源：1-手动发放 2-自动发放(下单奖励)  3-员工福利  4-首页活动
	SourceOrderId   string // 来源订单ID
	ActivityId      string // 活动ID（员工福利活动）
	EmployeeId      string // 员工ID（员工福利发放）
	IndexActivityId string // 活动ID（首页活动）
	CountDown       string // 倒计时
	OperatorId      string // 操作员ID（回收）
	RecoveryTime    string // 回收时间
	CreateAt        string // 创建时间
	UpdateAt        string // 修改时间
	DeletedAt       string // 删除时间
}

// thMemberCouponColumns holds the columns for the table hg_th_member_coupon.
var thMemberCouponColumns = ThMemberCouponColumns{
	Id:              "id",
	CouponNo:        "coupon_no",
	CouponId:        "coupon_id",
	MemberId:        "member_id",
	State:           "state",
	InvalidTime:     "invalid_time",
	VerifyTime:      "verify_time",
	VerifyMchId:     "verify_mch_id",
	VerifyStoreId:   "verify_store_id",
	StartTime:       "start_time",
	EndTime:         "end_time",
	Source:          "source",
	SourceOrderId:   "source_order_id",
	ActivityId:      "activity_id",
	EmployeeId:      "employee_id",
	IndexActivityId: "index_activity_id",
	CountDown:       "count_down",
	OperatorId:      "operator_id",
	RecoveryTime:    "recovery_time",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
	DeletedAt:       "deleted_at",
}

// NewThMemberCouponDao creates and returns a new DAO object for table data access.
func NewThMemberCouponDao() *ThMemberCouponDao {
	return &ThMemberCouponDao{
		group:   "default",
		table:   "hg_th_member_coupon",
		columns: thMemberCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThMemberCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThMemberCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThMemberCouponDao) Columns() ThMemberCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThMemberCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThMemberCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThMemberCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
