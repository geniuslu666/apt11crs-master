// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsPropertyDao is the data access object for the table hg_pms_property.
type PmsPropertyDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PmsPropertyColumns // columns contains all the column names of Table for convenient usage.
}

// PmsPropertyColumns defines and stores column names for the table hg_pms_property.
type PmsPropertyColumns struct {
	Id                   string // 主键
	Uid                  string // 在API合作伙伴系统中的物业ID
	Icon                 string // 图标
	GroupIds             string // 开放显示的会员分组IDS
	Name                 string // 物业名称   多语言
	Style                string // 物业类型
	Cover                string // 封面
	Currency             string // 物业的默认货币
	Language             string // 物业的默认自动消息语言
	TimeZone             string // 标准时区名称
	Address              string // 地址描述
	Lat                  string // 纬度
	Lng                  string // 经度
	AddressDetail        string // 百度详细地址
	GgAddressDetail      string // google详细地址
	GgLat                string // 谷歌纬度
	GgLng                string // 谷歌经度
	TagList              string // 标签多语言
	RoomDes              string // 房间描述
	Surroundings         string // 周边环境
	Description          string // 物业描述  多语言
	BusStation           string // 公交站
	Subway               string // 地铁站
	ContactName          string // 联系人
	Phone                string // 联系方式
	ContactEmail         string // 邮箱
	MinDaysNotice        string // 短租模式最小预定区间
	MaxDaysNotice        string // 最大预定区间
	MinutesAfterCheckout string // 退房后 分钟
	MinutesBeforeCheckin string // 入住前 分钟
	BookingLeadTimeLabel string // 预约期限
	TurnoverDays         string // 周转天数
	LinePrice            string // 单价
	Price                string // 单价
	CheckinAt            string // 入住时间
	CheckoutAt           string // 退房时间
	CancelPolicy         string // 取消政策 多语言
	GalleryImages        string // 画廊图片
	GalleryCover         string // 画廊封面
	RequiredBook         string //
	Close                string // 1、开启该物业  2、关闭该物业
	LeaseClose           string // 1、开启物业短租  2、关闭物业短租
	BookingClose         string // 1、开启预订模式  2、关闭预订模式
	Sort                 string // 排序(越大越靠前)
	CheckInGuide         string // 入住指南
	SpaCanOrder          string // 是否开放预定按摩服务  1开放  2关闭
	CreatedAt            string //
	UpdatedAt            string //
	DeletedAt            string // 删除时间
	BatchReservation     string // 是否开启多房型预定 （Y 开启 N  关闭）
	RegionId             string // 地区
	AccessPass           string // 门禁密码
}

// pmsPropertyColumns holds the columns for the table hg_pms_property.
var pmsPropertyColumns = PmsPropertyColumns{
	Id:                   "id",
	Uid:                  "uid",
	Icon:                 "icon",
	GroupIds:             "group_ids",
	Name:                 "name",
	Style:                "style",
	Cover:                "cover",
	Currency:             "currency",
	Language:             "language",
	TimeZone:             "time_zone",
	Address:              "address",
	Lat:                  "lat",
	Lng:                  "lng",
	AddressDetail:        "address_detail",
	GgAddressDetail:      "gg_address_detail",
	GgLat:                "gg_lat",
	GgLng:                "gg_lng",
	TagList:              "tag_list",
	RoomDes:              "room_des",
	Surroundings:         "surroundings",
	Description:          "description",
	BusStation:           "bus_station",
	Subway:               "subway",
	ContactName:          "contact_name",
	Phone:                "phone",
	ContactEmail:         "contact_email",
	MinDaysNotice:        "min_days_notice",
	MaxDaysNotice:        "max_days_notice",
	MinutesAfterCheckout: "minutes_after_checkout",
	MinutesBeforeCheckin: "minutes_before_checkin",
	BookingLeadTimeLabel: "booking_lead_time_label",
	TurnoverDays:         "turnover_days",
	LinePrice:            "line_price",
	Price:                "price",
	CheckinAt:            "checkin_at",
	CheckoutAt:           "checkout_at",
	CancelPolicy:         "cancel_policy",
	GalleryImages:        "gallery_images",
	GalleryCover:         "gallery_cover",
	RequiredBook:         "required_book",
	Close:                "close",
	LeaseClose:           "lease_close",
	BookingClose:         "booking_close",
	Sort:                 "sort",
	CheckInGuide:         "check_in_guide",
	SpaCanOrder:          "spa_can_order",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
	BatchReservation:     "batch_reservation",
	RegionId:             "region_id",
	AccessPass:           "access_pass",
}

// NewPmsPropertyDao creates and returns a new DAO object for table data access.
func NewPmsPropertyDao() *PmsPropertyDao {
	return &PmsPropertyDao{
		group:   "default",
		table:   "hg_pms_property",
		columns: pmsPropertyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsPropertyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsPropertyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsPropertyDao) Columns() PmsPropertyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsPropertyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsPropertyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsPropertyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
