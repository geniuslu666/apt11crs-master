// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberSceneDao is the data access object for the table hg_pms_member_scene.
type PmsMemberSceneDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns PmsMemberSceneColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberSceneColumns defines and stores column names for the table hg_pms_member_scene.
type PmsMemberSceneColumns struct {
	Id                  string //
	SceneName           string // 场景名称
	IsGetOpen           string // 是否允许积分获取
	IsPayOpen           string // 是否允许积分抵扣
	LimitMoney          string // 限制金额
	GetRate             string // 获得积分的抵扣比例
	PayRate             string // 能够使用的总金额比例积分
	IsOpenReward        string // 1、启用 2、禁用
	RewardType          string // 奖励类型（balance：积分 | coupon：优惠券 | thcoupon：礼品券）
	RewardCouponTypeIds string // 奖励优惠券ID，逗号分隔
	RewardThCouponIds   string // 奖励礼品券ID，逗号分隔
	CreatedAt           string //
	UpdatedAt           string //
	DeletedAt           string //
}

// pmsMemberSceneColumns holds the columns for the table hg_pms_member_scene.
var pmsMemberSceneColumns = PmsMemberSceneColumns{
	Id:                  "id",
	SceneName:           "scene_name",
	IsGetOpen:           "is_get_open",
	IsPayOpen:           "is_pay_open",
	LimitMoney:          "limit_money",
	GetRate:             "get_rate",
	PayRate:             "pay_rate",
	IsOpenReward:        "is_open_reward",
	RewardType:          "reward_type",
	RewardCouponTypeIds: "reward_coupon_type_ids",
	RewardThCouponIds:   "reward_th_coupon_ids",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// NewPmsMemberSceneDao creates and returns a new DAO object for table data access.
func NewPmsMemberSceneDao() *PmsMemberSceneDao {
	return &PmsMemberSceneDao{
		group:   "default",
		table:   "hg_pms_member_scene",
		columns: pmsMemberSceneColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsMemberSceneDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsMemberSceneDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsMemberSceneDao) Columns() PmsMemberSceneColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsMemberSceneDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsMemberSceneDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsMemberSceneDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
