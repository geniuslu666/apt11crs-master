// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CarServiceDao is the data access object for the table hg_car_service.
type CarServiceDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns CarServiceColumns // columns contains all the column names of Table for convenient usage.
}

// CarServiceColumns defines and stores column names for the table hg_car_service.
type CarServiceColumns struct {
	Id               string //
	ServiceType      string // 服务类型
	ServiceName      string // 路线名称
	CarTypeId        string // 车型ID
	StartIds         string // 出发地ID  多选
	EndIds           string // 目的地ID 多选
	Distance         string // 路线长度 KM
	UseTime          string // 线路时长  分钟
	Price            string // 基础费用
	LinePrice        string // 划线价
	FreeWaitTime     string // 免费等待时长  分钟
	MaxWaitTime      string // 最大等待时长  分钟
	TimeoutPreTime   string // 超时每xx分钟
	TimeoutPrePrice  string // 超时价格
	Status           string // 状态1、启用 2、禁用
	Sort             string // 排序(越大越靠前)
	TotalOrderNum    string // 预约单总数量（包含退款）
	TotalOrderAmount string // 预约单总金额（包含退款）
	PayOrderNum      string // 预约单支付数量（不包含退款）
	PayOrderAmount   string // 预约单支付金额（不包含退款）
	CreateAt         string // 创建时间
	UpdateAt         string // 更新时间
	DeletedAt        string // 删除时间
}

// carServiceColumns holds the columns for the table hg_car_service.
var carServiceColumns = CarServiceColumns{
	Id:               "id",
	ServiceType:      "service_type",
	ServiceName:      "service_name",
	CarTypeId:        "car_type_id",
	StartIds:         "start_ids",
	EndIds:           "end_ids",
	Distance:         "distance",
	UseTime:          "use_time",
	Price:            "price",
	LinePrice:        "line_price",
	FreeWaitTime:     "free_wait_time",
	MaxWaitTime:      "max_wait_time",
	TimeoutPreTime:   "timeout_pre_time",
	TimeoutPrePrice:  "timeout_pre_price",
	Status:           "status",
	Sort:             "sort",
	TotalOrderNum:    "total_order_num",
	TotalOrderAmount: "total_order_amount",
	PayOrderNum:      "pay_order_num",
	PayOrderAmount:   "pay_order_amount",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
	DeletedAt:        "deleted_at",
}

// NewCarServiceDao creates and returns a new DAO object for table data access.
func NewCarServiceDao() *CarServiceDao {
	return &CarServiceDao{
		group:   "default",
		table:   "hg_car_service",
		columns: carServiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CarServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CarServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CarServiceDao) Columns() CarServiceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CarServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CarServiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CarServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
