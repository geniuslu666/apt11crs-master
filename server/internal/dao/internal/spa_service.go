// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaServiceDao is the data access object for the table hg_spa_service.
type SpaServiceDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns SpaServiceColumns // columns contains all the column names of Table for convenient usage.
}

// SpaServiceColumns defines and stores column names for the table hg_spa_service.
type SpaServiceColumns struct {
	Id               string //
	IspId            string // 服务商ID
	Name             string // 服务名称(多语言)
	SubName          string // 副标题(多语言)
	LabelIds         string // 标签（多选）
	Images           string // 图集
	Channel          string // 服务渠道，1-到店和上门 2-仅上门 3-仅到店
	PropertyIds      string // 适用物业(以逗号分割)
	ServiceState     string // 状态（1-立即上架 2-放入仓库）
	Sort             string // 排序(越大越靠前)
	Content          string // 服务详情(多语言)
	TotalOrderNum    string // 预约单总数量（包含退款）
	TotalOrderAmount string // 预约单总金额（包含退款）
	PayOrderNum      string // 预约单支付数量（不包含退款）
	PayOrderAmount   string // 预约单支付金额（不包含退款）
	CreateAt         string // 创建时间
	UpdateAt         string // 更新时间
	DeletedAt        string //
}

// spaServiceColumns holds the columns for the table hg_spa_service.
var spaServiceColumns = SpaServiceColumns{
	Id:               "id",
	IspId:            "isp_id",
	Name:             "name",
	SubName:          "sub_name",
	LabelIds:         "label_ids",
	Images:           "images",
	Channel:          "channel",
	PropertyIds:      "property_ids",
	ServiceState:     "service_state",
	Sort:             "sort",
	Content:          "content",
	TotalOrderNum:    "total_order_num",
	TotalOrderAmount: "total_order_amount",
	PayOrderNum:      "pay_order_num",
	PayOrderAmount:   "pay_order_amount",
	CreateAt:         "create_at",
	UpdateAt:         "update_at",
	DeletedAt:        "deleted_at",
}

// NewSpaServiceDao creates and returns a new DAO object for table data access.
func NewSpaServiceDao() *SpaServiceDao {
	return &SpaServiceDao{
		group:   "default",
		table:   "hg_spa_service",
		columns: spaServiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaServiceDao) Columns() SpaServiceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaServiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
