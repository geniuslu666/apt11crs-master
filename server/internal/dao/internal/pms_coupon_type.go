// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsCouponTypeDao is the data access object for the table hg_pms_coupon_type.
type PmsCouponTypeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsCouponTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsCouponTypeColumns defines and stores column names for the table hg_pms_coupon_type.
type PmsCouponTypeColumns struct {
	Id                        string // 优惠券ID
	Type                      string // 优惠券类型 reward-满减 discount-折扣 random-随机
	CouponName                string // 优惠券名称
	Count                     string // 发放数量
	LeadCount                 string // 已领取数量
	UsedCount                 string // 已使用数量
	AtLeast                   string // 满多少元使用 0代表无限制
	Money                     string // 发放面额 当type为reward时需要添加
	Discount                  string // 1 =< 折扣 <= 9.9 当type为discount时需要添加
	DiscountLimit             string // 最多折扣金额 当type为discount时可选择性添加
	ValidityType              string // 过期类型1-古固定时间范围过期 2-领取之日固定日期后过期 3长期有效
	StartUseTime              string // 使用开始日期 过期类型0时必填
	EndUseTime                string // 使用结束日期 过期类型0时必填
	FixedTerm                 string // 当validity_type为2时需要添加 领取之日起或者次日N天内有效
	Sort                      string // 排序
	MaxFetch                  string // 每人最大领取个数
	IsShow                    string // 是否允许直接领取
	DiscountAppStayOrderMoney string // 住宿订单的优惠总金额
	AppStayOrderMoney         string // 住宿订单用券总成交额
	Status                    string // 状态（1进行中2已结束-1已关闭）
	Scene                     string // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车  5-储物柜
	PropertyIds               string //
	RestaurantIds             string // 餐厅ID
	ServiceIds                string // 按摩服务ID
	CarServiceTypes           string // 汽车服务类型
	Desc                      string // 优惠券使用说明
	CreateAt                  string // 创建时间
	UpdateAt                  string // 修改时间
	DeletedAt                 string // 删除时间
}

// pmsCouponTypeColumns holds the columns for the table hg_pms_coupon_type.
var pmsCouponTypeColumns = PmsCouponTypeColumns{
	Id:                        "id",
	Type:                      "type",
	CouponName:                "coupon_name",
	Count:                     "count",
	LeadCount:                 "lead_count",
	UsedCount:                 "used_count",
	AtLeast:                   "at_least",
	Money:                     "money",
	Discount:                  "discount",
	DiscountLimit:             "discount_limit",
	ValidityType:              "validity_type",
	StartUseTime:              "start_use_time",
	EndUseTime:                "end_use_time",
	FixedTerm:                 "fixed_term",
	Sort:                      "sort",
	MaxFetch:                  "max_fetch",
	IsShow:                    "is_show",
	DiscountAppStayOrderMoney: "discount_app_stay_order_money",
	AppStayOrderMoney:         "app_stay_order_money",
	Status:                    "status",
	Scene:                     "scene",
	PropertyIds:               "property_ids",
	RestaurantIds:             "restaurant_ids",
	ServiceIds:                "service_ids",
	CarServiceTypes:           "car_service_types",
	Desc:                      "desc",
	CreateAt:                  "create_at",
	UpdateAt:                  "update_at",
	DeletedAt:                 "deleted_at",
}

// NewPmsCouponTypeDao creates and returns a new DAO object for table data access.
func NewPmsCouponTypeDao() *PmsCouponTypeDao {
	return &PmsCouponTypeDao{
		group:   "default",
		table:   "hg_pms_coupon_type",
		columns: pmsCouponTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsCouponTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsCouponTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsCouponTypeDao) Columns() PmsCouponTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsCouponTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsCouponTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsCouponTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
