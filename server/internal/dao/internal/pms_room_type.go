// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsRoomTypeDao is the data access object for the table hg_pms_room_type.
type PmsRoomTypeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PmsRoomTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsRoomTypeColumns defines and stores column names for the table hg_pms_room_type.
type PmsRoomTypeColumns struct {
	Uid                    string // airhost房型ID
	Id                     string // 主键ID
	Puid                   string // 物业ID
	Cover                  string // 封面
	CoverList              string // 照片墙
	Name                   string // 房型多语言名称ID
	BasePrice              string // 最低价格
	CheckinAt              string // 入住时间
	CheckoutAt             string // 退房时间
	BookingStyle           string // 预订方式
	RoomStyle              string // 房间的风格
	Occupancy              string // 占用
	Size                   string // 面积
	Bedrooms               string // 卧室
	Bathrooms              string // 浴室
	CleaningFee            string // 清理费
	RatePlanId             string // 费率ID
	AdditionalGuestAmounts string // 额外客人金额
	OccupantsForBaseRate   string // 无需增加额外客人金额人数
	RoomNum                string // 房间数
	IsShow                 string // 1、显示 0 隐藏
	CreateAt               string //
	UpdateAt               string //
}

// pmsRoomTypeColumns holds the columns for the table hg_pms_room_type.
var pmsRoomTypeColumns = PmsRoomTypeColumns{
	Uid:                    "uid",
	Id:                     "id",
	Puid:                   "puid",
	Cover:                  "cover",
	CoverList:              "cover_list",
	Name:                   "name",
	BasePrice:              "base_price",
	CheckinAt:              "checkin_at",
	CheckoutAt:             "checkout_at",
	BookingStyle:           "booking_style",
	RoomStyle:              "room_style",
	Occupancy:              "occupancy",
	Size:                   "size",
	Bedrooms:               "bedrooms",
	Bathrooms:              "bathrooms",
	CleaningFee:            "cleaning_fee",
	RatePlanId:             "rate_plan_id",
	AdditionalGuestAmounts: "additional_guest_amounts",
	OccupantsForBaseRate:   "occupants_for_base_rate",
	RoomNum:                "room_num",
	IsShow:                 "Is_show",
	CreateAt:               "create_at",
	UpdateAt:               "update_at",
}

// NewPmsRoomTypeDao creates and returns a new DAO object for table data access.
func NewPmsRoomTypeDao() *PmsRoomTypeDao {
	return &PmsRoomTypeDao{
		group:   "default",
		table:   "hg_pms_room_type",
		columns: pmsRoomTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsRoomTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsRoomTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsRoomTypeDao) Columns() PmsRoomTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsRoomTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsRoomTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsRoomTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
