// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarCarTypeDao is the data access object for the table hg_car_car_type.
type CarCarTypeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns CarCarTypeColumns // columns contains all the column names of Table for convenient usage.
}

// CarCarTypeColumns defines and stores column names for the table hg_car_car_type.
type CarCarTypeColumns struct {
	Id                  string //
	Name                string // 名称
	Desc                string // 简介
	Image               string // 图片
	SeatNum             string // 座位数
	PassengerNum        string // 建议乘员人数
	MaxPackageNum       string // 最大容纳行李件数
	ChildrenSeatSupport string // 是否支持儿童座椅
	ChildrenSeatSpace   string // 儿童座椅占几个作为
	Content             string // 内容
	Status              string // 状态1、启用 2、禁用
	Sort                string // 排序(越大越靠前)
	CreateAt            string // 创建时间
	UpdateAt            string // 更新时间
	DeletedAt           string // 删除时间
}

// carCarTypeColumns holds the columns for the table hg_car_car_type.
var carCarTypeColumns = CarCarTypeColumns{
	Id:                  "id",
	Name:                "name",
	Desc:                "desc",
	Image:               "image",
	SeatNum:             "seat_num",
	PassengerNum:        "passenger_num",
	MaxPackageNum:       "max_package_num",
	ChildrenSeatSupport: "children_seat_support",
	ChildrenSeatSpace:   "children_seat_space",
	Content:             "content",
	Status:              "status",
	Sort:                "sort",
	CreateAt:            "create_at",
	UpdateAt:            "update_at",
	DeletedAt:           "deleted_at",
}

// NewCarCarTypeDao creates and returns a new DAO object for table data access.
func NewCarCarTypeDao() *CarCarTypeDao {
	return &CarCarTypeDao{
		group:   "default",
		table:   "hg_car_car_type",
		columns: carCarTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarCarTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarCarTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarCarTypeDao) Columns() CarCarTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarCarTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarCarTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarCarTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
