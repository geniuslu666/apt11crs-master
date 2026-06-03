// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for the table user.
type UserColumns struct {
	Id             string //
	Pid            string // 父级ID
	Name           string // 账户
	Password       string // 密码，md5加密
	Nickname       string // 显示名称
	CreatedAt      string // 创建时间
	ExpiredAt      string // 过期时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
	Avator         string // 客服头像
	RecNum         string // 正在接待数量
	OnlineStatus   string // 在线状态，1在线，2离线
	Status         string // 开启状态，0正常，1关闭
	AgentNum       string // 子账号个数
	Email          string // 绑定邮箱
	CompanyPic     string // 宣传图
	Tel            string // 绑定手机
	Uuid           string // 企业uuid
	ReceptionValue string // 接待固定值
	Reception      string // 接待模式
}

// userColumns holds the columns for the table user.
var userColumns = UserColumns{
	Id:             "id",
	Pid:            "pid",
	Name:           "name",
	Password:       "password",
	Nickname:       "nickname",
	CreatedAt:      "created_at",
	ExpiredAt:      "expired_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	Avator:         "avator",
	RecNum:         "rec_num",
	OnlineStatus:   "online_status",
	Status:         "status",
	AgentNum:       "agent_num",
	Email:          "email",
	CompanyPic:     "company_pic",
	Tel:            "tel",
	Uuid:           "uuid",
	ReceptionValue: "reception_value",
	Reception:      "reception",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
