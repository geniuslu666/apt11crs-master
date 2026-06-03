// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsPricePlanDao is the data access object for the table hg_pms_price_plan.
type PmsPricePlanDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns PmsPricePlanColumns // columns contains all the column names of Table for convenient usage.
}

// PmsPricePlanColumns defines and stores column names for the table hg_pms_price_plan.
type PmsPricePlanColumns struct {
	Id              string //
	PlanName        string // 价格plan名称
	PlanShowName    string // 展示标签 多语言
	PropertyId      string // 物业ID
	RoomTypeId      string // 房型ID
	BookingDays     string // 预订天数
	MemberGroupId   string // 用户组
	MemberLevelId   string // 会员等级
	IsCancel        string // 是否可取消   Y  是  N  否
	IsOpenPriceMode string // 是否开启价格模式
	PriceMode       string // 模式    +  贵   - 便宜
	PriceStandard   string // 基准  PERCENT 倍率  AMOUNT  金额
	PlanValue       string // 价格基准值
	PlanTips        string // 价格计划提示
	PricePlanStatus string // Y 开启  N 关闭
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
}

// pmsPricePlanColumns holds the columns for the table hg_pms_price_plan.
var pmsPricePlanColumns = PmsPricePlanColumns{
	Id:              "id",
	PlanName:        "plan_name",
	PlanShowName:    "plan_show_name",
	PropertyId:      "property_id",
	RoomTypeId:      "room_type_id",
	BookingDays:     "booking_days",
	MemberGroupId:   "member_group_id",
	MemberLevelId:   "member_level_id",
	IsCancel:        "is_cancel",
	IsOpenPriceMode: "is_open_price_mode",
	PriceMode:       "price_mode",
	PriceStandard:   "price_standard",
	PlanValue:       "plan_value",
	PlanTips:        "plan_tips",
	PricePlanStatus: "price_plan_status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewPmsPricePlanDao creates and returns a new DAO object for table data access.
func NewPmsPricePlanDao() *PmsPricePlanDao {
	return &PmsPricePlanDao{
		group:   "default",
		table:   "hg_pms_price_plan",
		columns: pmsPricePlanColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsPricePlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsPricePlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsPricePlanDao) Columns() PmsPricePlanColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsPricePlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsPricePlanDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsPricePlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
