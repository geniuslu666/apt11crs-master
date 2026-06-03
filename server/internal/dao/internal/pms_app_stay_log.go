// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsAppStayLogDao is the data access object for the table hg_pms_app_stay_log.
type PmsAppStayLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsAppStayLogColumns // columns contains all the column names of Table for convenient usage.
}

// PmsAppStayLogColumns defines and stores column names for the table hg_pms_app_stay_log.
type PmsAppStayLogColumns struct {
	Id                 string // 变动ID
	OrderId            string // 订单ID
	ReservationOrderId string // 房间订单ID
	ActionWay          string // 操作名
	Remark             string // 备注
	Images             string // 图集
	OperateType        string // 操作员类型
	OperateId          string // 操作员ID
	CreatedAt          string // 创建时间
	UpdatedAt          string // 修改时间
}

// pmsAppStayLogColumns holds the columns for the table hg_pms_app_stay_log.
var pmsAppStayLogColumns = PmsAppStayLogColumns{
	Id:                 "id",
	OrderId:            "order_id",
	ReservationOrderId: "reservation_order_id",
	ActionWay:          "action_way",
	Remark:             "remark",
	Images:             "images",
	OperateType:        "operate_type",
	OperateId:          "operate_id",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewPmsAppStayLogDao creates and returns a new DAO object for table data access.
func NewPmsAppStayLogDao() *PmsAppStayLogDao {
	return &PmsAppStayLogDao{
		group:   "default",
		table:   "hg_pms_app_stay_log",
		columns: pmsAppStayLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsAppStayLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsAppStayLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsAppStayLogDao) Columns() PmsAppStayLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsAppStayLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsAppStayLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsAppStayLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
