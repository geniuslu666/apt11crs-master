// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThCouponDao is the data access object for the table hg_th_coupon.
type ThCouponDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns ThCouponColumns // columns contains all the column names of Table for convenient usage.
}

// ThCouponColumns defines and stores column names for the table hg_th_coupon.
type ThCouponColumns struct {
	Id                       string // ID
	CouponName               string // 提货券名称
	CouponSubName            string // 提货券副标题
	IdentityName             string // 券识别名称
	CouponNoPrefix           string // 编号前缀
	Logo                     string // LOGO
	Status                   string // 发放状态（1立即启用  2暂不启用）
	FixedTerm                string // 激活后几天内有效
	Desc                     string // 券说明
	Count                    string // 发放数量
	UseStatus                string // 使用状态（1开始使用  2停止使用）
	UseMode                  string // 使用模式
	CategoryId               string // 分类ID
	UsedCount                string // 已使用数量
	Sort                     string // 排序
	NeedReservation          string // 是否需要预约：0-不需要，1-需要
	ReservationRestaurantIds string // 需要预约的餐厅IDs，多个用逗号分隔，如：1,2,3
	CreateAt                 string // 创建时间
	UpdateAt                 string // 修改时间
	DeletedAt                string // 删除时间
}

// thCouponColumns holds the columns for the table hg_th_coupon.
var thCouponColumns = ThCouponColumns{
	Id:                       "id",
	CouponName:               "coupon_name",
	CouponSubName:            "coupon_sub_name",
	IdentityName:             "identity_name",
	CouponNoPrefix:           "coupon_no_prefix",
	Logo:                     "logo",
	Status:                   "status",
	FixedTerm:                "fixed_term",
	Desc:                     "desc",
	Count:                    "count",
	UseStatus:                "use_status",
	UseMode:                  "use_mode",
	CategoryId:               "category_id",
	UsedCount:                "used_count",
	Sort:                     "sort",
	NeedReservation:          "need_reservation",
	ReservationRestaurantIds: "reservation_restaurant_ids",
	CreateAt:                 "create_at",
	UpdateAt:                 "update_at",
	DeletedAt:                "deleted_at",
}

// NewThCouponDao creates and returns a new DAO object for table data access.
func NewThCouponDao() *ThCouponDao {
	return &ThCouponDao{
		group:   "default",
		table:   "hg_th_coupon",
		columns: thCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThCouponDao) Columns() ThCouponColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
