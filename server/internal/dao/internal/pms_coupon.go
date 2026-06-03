// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsCouponDao is the data access object for the table hg_pms_coupon.
type PmsCouponDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns PmsCouponColumns // columns contains all the column names of Table for convenient usage.
}

// PmsCouponColumns defines and stores column names for the table hg_pms_coupon.
type PmsCouponColumns struct {
	Id              string // 主键ID
	Type            string // 优惠券类型 reward-满减 discount-折扣 random-随机
	CouponName      string // 优惠券名称
	CouponTypeId    string // 优惠券类型id
	MemberId        string // 领用人
	Scene           string // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车   5-储物柜
	PropertyIds     string //
	RestaurantIds   string // 餐厅ID
	ServiceIds      string // 按摩服务ID
	CarServiceTypes string // 汽车服务类型
	AtLeast         string // 满多少元使用 0代表无限制
	Money           string // 发放面额 当type为reward时需要添加
	Discount        string // 1 =< 折扣 <= 9.9 当type为discount时需要添加
	DiscountLimit   string // 最多折扣金额 当type为discount时可选择性添加
	State           string // 优惠券状态 1已领用（未使用） 2已使用 3已过期 4已关闭 5已回收
	FetchTime       string // 领取时间
	UseTime         string // 使用时间
	StartTime       string // 可使用的开始时间
	EndTime         string // 有效期结束时间
	Source          string // 来源：1-自主领取 2-系统发放 3-注册奖励 4-邀请奖励  5-首页活动
	SourceOrderId   string // 来源订单ID
	OperatorId      string // 操作员ID（回收）
	IndexActivityId string // 活动ID（首页活动）
	RecoveryTime    string // 回收时间
	CreateAt        string // 创建时间
	UpdateAt        string // 修改时间
	DeletedAt       string // 删除时间
}

// pmsCouponColumns holds the columns for the table hg_pms_coupon.
var pmsCouponColumns = PmsCouponColumns{
	Id:              "id",
	Type:            "type",
	CouponName:      "coupon_name",
	CouponTypeId:    "coupon_type_id",
	MemberId:        "member_id",
	Scene:           "scene",
	PropertyIds:     "property_ids",
	RestaurantIds:   "restaurant_ids",
	ServiceIds:      "service_ids",
	CarServiceTypes: "car_service_types",
	AtLeast:         "at_least",
	Money:           "money",
	Discount:        "discount",
	DiscountLimit:   "discount_limit",
	State:           "state",
	FetchTime:       "fetch_time",
	UseTime:         "use_time",
	StartTime:       "start_time",
	EndTime:         "end_time",
	Source:          "source",
	SourceOrderId:   "source_order_id",
	OperatorId:      "operator_id",
	IndexActivityId: "index_activity_id",
	RecoveryTime:    "recovery_time",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
	DeletedAt:       "deleted_at",
}

// NewPmsCouponDao creates and returns a new DAO object for table data access.
func NewPmsCouponDao() *PmsCouponDao {
	return &PmsCouponDao{
		group:   "default",
		table:   "hg_pms_coupon",
		columns: pmsCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsCouponDao) Columns() PmsCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
