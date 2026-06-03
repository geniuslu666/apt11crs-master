// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelProductSkuDao is the data access object for the table hg_travel_product_sku.
type TravelProductSkuDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TravelProductSkuColumns // columns contains all the column names of Table for convenient usage.
}

// TravelProductSkuColumns defines and stores column names for the table hg_travel_product_sku.
type TravelProductSkuColumns struct {
	Id               string //
	ProductId        string // 产品ID
	Name             string // 车型名称（默认语言；多语言存 hg_pms_language）
	Price            string // 售价（元）
	DailyCapacity    string // 每日最大接待人数
	Status           string // 状态（1启用 2禁用）
	Sort             string // 排序（越大越靠前）
	SalesNum         string // 已售
	ContactMobile    string // 联系电话
	MeetingPlace     string // 集合地点
	MeetingTime      string // 集合时间（格式：HH:MM）
	GgLat            string // 谷歌纬度
	GgLng            string // 谷歌经度
	DeletedAt        string // 软删除时间（NULL=正常）
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	AllowCancel      string // 是否允许取消
	FreeCancelHours  string // 几小时前免费
	CancelFeePercent string // 取消费率%
}

// travelProductSkuColumns holds the columns for the table hg_travel_product_sku.
var travelProductSkuColumns = TravelProductSkuColumns{
	Id:               "id",
	ProductId:        "product_id",
	Name:             "name",
	Price:            "price",
	DailyCapacity:    "daily_capacity",
	Status:           "status",
	Sort:             "sort",
	SalesNum:         "sales_num",
	ContactMobile:    "contact_mobile",
	MeetingPlace:     "meeting_place",
	MeetingTime:      "meeting_time",
	GgLat:            "gg_lat",
	GgLng:            "gg_lng",
	DeletedAt:        "deleted_at",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	AllowCancel:      "allow_cancel",
	FreeCancelHours:  "free_cancel_hours",
	CancelFeePercent: "cancel_fee_percent",
}

// NewTravelProductSkuDao creates and returns a new DAO object for table data access.
func NewTravelProductSkuDao() *TravelProductSkuDao {
	return &TravelProductSkuDao{
		group:   "default",
		table:   "hg_travel_product_sku",
		columns: travelProductSkuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TravelProductSkuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TravelProductSkuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TravelProductSkuDao) Columns() TravelProductSkuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TravelProductSkuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TravelProductSkuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TravelProductSkuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
