// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsMemberLevelDao is the data access object for table hg_pms_member_level.
type PmsMemberLevelDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns PmsMemberLevelColumns // columns contains all the column names of Table for convenient usage.
}

// PmsMemberLevelColumns defines and stores column names for table hg_pms_member_level.
type PmsMemberLevelColumns struct {
	Id             string //
	LevelName      string // 会员等级名称
	Exp            string // 达到等级所需经验值
	HotelGetRate   string // 酒店场景获取积分倍率
	FoodGetRate    string // 餐饮场景获取积分倍率
	SpaGetRate     string // 按摩场景获取积分倍率
	CarGetRate     string // 接送机/包车场景获取积分倍率
	CabinetGetRate string // 储物柜场景获取积分倍率
	TravelGetRate  string // 一日游场景获取积分倍率
	Desc           string // 等级说明
	WordColor      string // 等级字体颜色
	LevelBadge     string // 等级徽章（单图）
	LevelCard      string // 等级卡片（单图）
	LevelBigPic    string // 等级大图（单图）
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
}

// pmsMemberLevelColumns holds the columns for table hg_pms_member_level.
var pmsMemberLevelColumns = PmsMemberLevelColumns{
	Id:             "id",
	LevelName:      "level_name",
	Exp:            "exp",
	HotelGetRate:   "hotel_get_rate",
	FoodGetRate:    "food_get_rate",
	SpaGetRate:     "spa_get_rate",
	CarGetRate:     "car_get_rate",
	CabinetGetRate: "cabinet_get_rate",
	TravelGetRate:  "travel_get_rate",
	Desc:           "desc",
	WordColor:      "word_color",
	LevelBadge:     "level_badge",
	LevelCard:      "level_card",
	LevelBigPic:    "level_big_pic",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewPmsMemberLevelDao creates and returns a new DAO object for table data access.
func NewPmsMemberLevelDao() *PmsMemberLevelDao {
	return &PmsMemberLevelDao{
		group:   "default",
		table:   "hg_pms_member_level",
		columns: pmsMemberLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmsMemberLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmsMemberLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmsMemberLevelDao) Columns() PmsMemberLevelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmsMemberLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmsMemberLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmsMemberLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
