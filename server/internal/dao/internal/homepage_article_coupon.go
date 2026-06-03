// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HomepageArticleCouponDao is the data access object for table hg_homepage_article_coupon.
type HomepageArticleCouponDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns HomepageArticleCouponColumns // columns contains all the column names of Table for convenient usage.
}

// HomepageArticleCouponColumns defines and stores column names for table hg_homepage_article_coupon.
type HomepageArticleCouponColumns struct {
	Id                string // 关联ID
	CouponType        string // coupon-优惠券，thCoupon-礼品券
	ArticleId         string // 活动ID
	CouponId          string // 礼品券ID
	AvailableQuantity string // 可领取数量
	PerDayAvailable   string // 每人N天可领取数量（0表示无限制）
	LimitDays         string // 限制天数
	TotalReceived     string // 总领取数量
	TotalUsed         string // 已核销数量
	Status            string // 状态：1-正常 2-禁用
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
	FixedTerm         string // 领取后几天内有效（0表示领取当日23:59:59）
}

// homepageArticleCouponColumns holds the columns for table hg_homepage_article_coupon.
var homepageArticleCouponColumns = HomepageArticleCouponColumns{
	Id:                "id",
	CouponType:        "coupon_type",
	ArticleId:         "article_id",
	CouponId:          "coupon_id",
	AvailableQuantity: "available_quantity",
	PerDayAvailable:   "per_day_available",
	LimitDays:         "limit_days",
	TotalReceived:     "total_received",
	TotalUsed:         "total_used",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	FixedTerm:         "fixed_term",
}

// NewHomepageArticleCouponDao creates and returns a new DAO object for table data access.
func NewHomepageArticleCouponDao() *HomepageArticleCouponDao {
	return &HomepageArticleCouponDao{
		group:   "default",
		table:   "hg_homepage_article_coupon",
		columns: homepageArticleCouponColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HomepageArticleCouponDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HomepageArticleCouponDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HomepageArticleCouponDao) Columns() HomepageArticleCouponColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HomepageArticleCouponDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HomepageArticleCouponDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HomepageArticleCouponDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
