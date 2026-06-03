// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarCarDao is the data access object for the table hg_car_car.
type CarCarDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of the current DAO.
	columns CarCarColumns // columns contains all the column names of Table for convenient usage.
}

// CarCarColumns defines and stores column names for the table hg_car_car.
type CarCarColumns struct {
	Id                  string //
	CarName             string // 车辆名称
	TypeId              string // 车型ID
	Brand               string // 车辆品牌型号
	LicenseNo           string // 车牌号码
	LicenseColor        string // 牌照颜色
	SeatNum             string // 座位数(废弃)
	PassengerNum        string // 建议乘员人数(废弃)
	MaxPackageNum       string // 最大容纳行李件数(废弃)
	ChildrenSeatSupport string // 是否支持儿童座椅(废弃)
	ChildrenSeatSpace   string // 儿童座椅占几个作为(废弃)
	WorkStatus          string // 工作状态
	Status              string // 状态1、启用 2、禁用
	Sort                string // 排序(越大越靠前)
	QualityMaterials    string // 资质信息(多图)
	TotalOrderNum       string // 预约单总数量（包含退款）
	TotalOrderAmount    string // 预约单总金额（包含退款）
	PayOrderNum         string // 预约单支付数量（不包含退款）
	PayOrderAmount      string // 预约单支付金额（不包含退款）
	CreateAt            string // 创建时间
	UpdateAt            string // 更新时间
	DeletedAt           string //
}

// carCarColumns holds the columns for the table hg_car_car.
var carCarColumns = CarCarColumns{
	Id:                  "id",
	CarName:             "car_name",
	TypeId:              "type_id",
	Brand:               "brand",
	LicenseNo:           "license_no",
	LicenseColor:        "license_color",
	SeatNum:             "seat_num",
	PassengerNum:        "passenger_num",
	MaxPackageNum:       "max_package_num",
	ChildrenSeatSupport: "children_seat_support",
	ChildrenSeatSpace:   "children_seat_space",
	WorkStatus:          "work_status",
	Status:              "status",
	Sort:                "sort",
	QualityMaterials:    "quality_materials",
	TotalOrderNum:       "total_order_num",
	TotalOrderAmount:    "total_order_amount",
	PayOrderNum:         "pay_order_num",
	PayOrderAmount:      "pay_order_amount",
	CreateAt:            "create_at",
	UpdateAt:            "update_at",
	DeletedAt:           "deleted_at",
}

// NewCarCarDao creates and returns a new DAO object for table data access.
func NewCarCarDao() *CarCarDao {
	return &CarCarDao{
		group:   "default",
		table:   "hg_car_car",
		columns: carCarColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarCarDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarCarDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarCarDao) Columns() CarCarColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarCarDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarCarDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarCarDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
