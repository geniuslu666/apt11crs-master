// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServiceGoodsDao is the data access object for the table hg_spa_service_goods.
type SpaServiceGoodsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns SpaServiceGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// SpaServiceGoodsColumns defines and stores column names for the table hg_spa_service_goods.
type SpaServiceGoodsColumns struct {
	Id               string //
	ServiceId        string // 服务ID
	GoodsName        string // 服务套餐名称(多语言)
	Image            string // 图片
	DetailImage      string // 详情图
	Price            string // 套餐售价
	Duration         string // 时长(分钟)
	TotalOrderNum    string // 预约单总数量（包含退款）
	TotalOrderAmount string // 预约单总金额（包含退款）
	PayOrderNum      string // 预约单支付数量（不包含退款）
	PayOrderAmount   string // 预约单支付金额（不包含退款）
	Status           string // 状态1、启用 2、禁用
	CreateAt         string // 创建时间
	UpdateAt         string // 更新时间
	DeletedAt        string //
}

// spaServiceGoodsColumns holds the columns for the table hg_spa_service_goods.
var spaServiceGoodsColumns = SpaServiceGoodsColumns{
	Id:               "id",
	ServiceId:        "service_id",
	GoodsName:        "goods_name",
	Image:            "image",
	DetailImage:      "detail_image",
	Price:            "price",
	Duration:         "duration",
	TotalOrderNum:    "total_order_num",
	TotalOrderAmount: "total_order_amount",
	PayOrderNum:      "pay_order_num",
	PayOrderAmount:   "pay_order_amount",
	Status:           "status",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
	DeletedAt:        "deleted_at",
}

// NewSpaServiceGoodsDao creates and returns a new DAO object for table data access.
func NewSpaServiceGoodsDao() *SpaServiceGoodsDao {
	return &SpaServiceGoodsDao{
		group:   "default",
		table:   "hg_spa_service_goods",
		columns: spaServiceGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaServiceGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaServiceGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaServiceGoodsDao) Columns() SpaServiceGoodsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaServiceGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaServiceGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaServiceGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
