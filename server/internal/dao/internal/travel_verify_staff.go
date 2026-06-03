// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TravelVerifyStaffDao is the data access object for the table hg_travel_verify_staff.
type TravelVerifyStaffDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns TravelVerifyStaffColumns // columns contains all the column names of Table for convenient usage.
}

// TravelVerifyStaffColumns defines and stores column names for the table hg_travel_verify_staff.
type TravelVerifyStaffColumns struct {
	Id                 string //
	Name               string // 姓名
	Mobile             string // 电话
	Username           string // 登录账号
	Status             string // 状态（1启用 2禁用）
	PasswordHash       string // 密码
	Salt               string // 密码盐
	PasswordResetToken string // 密码重置令牌
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
	DeletedAt          string // 软删除时间（NULL=正常）
}

// travelVerifyStaffColumns holds the columns for the table hg_travel_verify_staff.
var travelVerifyStaffColumns = TravelVerifyStaffColumns{
	Id:                 "id",
	Name:               "name",
	Mobile:             "mobile",
	Username:           "username",
	Status:             "status",
	PasswordHash:       "password_hash",
	Salt:               "salt",
	PasswordResetToken: "password_reset_token",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// NewTravelVerifyStaffDao creates and returns a new DAO object for table data access.
func NewTravelVerifyStaffDao() *TravelVerifyStaffDao {
	return &TravelVerifyStaffDao{
		group:   "default",
		table:   "hg_travel_verify_staff",
		columns: travelVerifyStaffColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TravelVerifyStaffDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TravelVerifyStaffDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TravelVerifyStaffDao) Columns() TravelVerifyStaffColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TravelVerifyStaffDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TravelVerifyStaffDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TravelVerifyStaffDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
