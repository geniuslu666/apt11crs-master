// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelProductDao is the data access object for table hg_travel_product.
type TravelProductDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TravelProductColumns // columns contains all the column names of Table for convenient usage.
}

// TravelProductColumns defines and stores column names for table hg_travel_product.
type TravelProductColumns struct {
	Id              string //
	Title           string // 标题（默认语言；多语言存 hg_pms_language）
	SubTitle        string // 副标题（默认语言；多语言存 hg_pms_language）
	ContentZh       string // 内容（中文 编辑器）
	ContentEn       string // 内容（英文 编辑器）
	ContentKo       string // 内容（韩文 编辑器）
	ContentJa       string // 内容（日文 编辑器）
	ContentTw       string // 内容（繁文 编辑器）
	ListImage       string // 列表图（单图 URL）
	CarouselImages  string // 轮播图（多图，JSON）
	DailyCapacity   string // 每日最大接待人数
	Stock           string // 总库存
	Price           string // 售价（元）
	ContactMobile   string // 联系电话
	MeetingPlace    string // 集合地点
	MeetingTime     string // 集合时间（格式：HH:MM）
	GgLat           string // 谷歌纬度
	GgLng           string // 谷歌经度
	MaxBookDays     string // 最大可预约天数
	AdvanceBookDays string // 至少提前预订天数
	TripPlanningZh  string // 行程规划（中文 编辑器）
	TripPlanningEn  string // 行程规划（英文 编辑器）
	TripPlanningJa  string // 行程规划（日文 编辑器）
	TripPlanningKo  string // 行程规划（韩文 编辑器）
	TripPlanningTw  string // 行程规划（繁文 编辑器）
	BookingNotesZh  string // 预约须知（中文 编辑器）
	BookingNotesEn  string // 预约须知（英文 编辑器）
	BookingNotesJa  string // 预约须知（日文 编辑器）
	BookingNotesKo  string // 预约须知（韩文 编辑器）
	BookingNotesTw  string // 预约须知（繁文 编辑器）
	Status          string // 状态（1启用 2禁用）
	Sort            string // 排序（越大越靠前）
	SalesNum        string // 已售
	DeletedAt       string // 软删除时间（NULL=正常）
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// travelProductColumns holds the columns for table hg_travel_product.
var travelProductColumns = TravelProductColumns{
	Id:              "id",
	Title:           "title",
	SubTitle:        "sub_title",
	ContentZh:       "content_zh",
	ContentEn:       "content_en",
	ContentKo:       "content_ko",
	ContentJa:       "content_ja",
	ContentTw:       "content_tw",
	ListImage:       "list_image",
	CarouselImages:  "carousel_images",
	DailyCapacity:   "daily_capacity",
	Stock:           "stock",
	Price:           "price",
	ContactMobile:   "contact_mobile",
	MeetingPlace:    "meeting_place",
	MeetingTime:     "meeting_time",
	GgLat:           "gg_lat",
	GgLng:           "gg_lng",
	MaxBookDays:     "max_book_days",
	AdvanceBookDays: "advance_book_days",
	TripPlanningZh:  "trip_planning_zh",
	TripPlanningEn:  "trip_planning_en",
	TripPlanningJa:  "trip_planning_ja",
	TripPlanningKo:  "trip_planning_ko",
	TripPlanningTw:  "trip_planning_tw",
	BookingNotesZh:  "booking_notes_zh",
	BookingNotesEn:  "booking_notes_en",
	BookingNotesJa:  "booking_notes_ja",
	BookingNotesKo:  "booking_notes_ko",
	BookingNotesTw:  "booking_notes_tw",
	Status:          "status",
	Sort:            "sort",
	SalesNum:        "sales_num",
	DeletedAt:       "deleted_at",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewTravelProductDao creates and returns a new DAO object for table data access.
func NewTravelProductDao() *TravelProductDao {
	return &TravelProductDao{
		group:   "default",
		table:   "hg_travel_product",
		columns: travelProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TravelProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TravelProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TravelProductDao) Columns() TravelProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TravelProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TravelProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TravelProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
