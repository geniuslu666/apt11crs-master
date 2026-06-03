// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarAddressDao is the data access object for the table hg_car_address.
type CarAddressDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns CarAddressColumns // columns contains all the column names of Table for convenient usage.
}

// CarAddressColumns defines and stores column names for the table hg_car_address.
type CarAddressColumns struct {
	Id            string //
	TypeId        string // 地点类型id
	Name          string // 地点名称（后台）
	SubName       string // 地点名称（app-多语）
	AirportCode   string // 机场代码
	TerminalName  string // 机场航站楼名称
	DetailAddress string // 详细地址-多语
	GgLat         string // 谷歌纬度
	GgLng         string // 谷歌经度
	Lat           string // 纬度
	Lng           string // 经度
	Status        string // 状态1、启用 2、禁用
	PropertyId    string // 物业id
	CreateAt      string // 创建时间
	UpdateAt      string // 更新时间
	DeletedAt     string // 删除时间
}

// carAddressColumns holds the columns for the table hg_car_address.
var carAddressColumns = CarAddressColumns{
	Id:            "id",
	TypeId:        "type_id",
	Name:          "name",
	SubName:       "sub_name",
	AirportCode:   "airport_code",
	TerminalName:  "terminal_name",
	DetailAddress: "detail_address",
	GgLat:         "gg_lat",
	GgLng:         "gg_lng",
	Lat:           "lat",
	Lng:           "lng",
	Status:        "status",
	PropertyId:    "property_id",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
	DeletedAt:     "deleted_at",
}

// NewCarAddressDao creates and returns a new DAO object for table data access.
func NewCarAddressDao() *CarAddressDao {
	return &CarAddressDao{
		group:   "default",
		table:   "hg_car_address",
		columns: carAddressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarAddressDao) Columns() CarAddressColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarAddressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
