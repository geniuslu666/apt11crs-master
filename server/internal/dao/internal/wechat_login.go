// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WechatLoginDao is the data access object for the table wechat_login.
type WechatLoginDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns WechatLoginColumns // columns contains all the column names of Table for convenient usage.
}

// WechatLoginColumns defines and stores column names for the table wechat_login.
type WechatLoginColumns struct {
	Id         string //
	KefuName   string // 客服账户
	OpenId     string // 微信公众号openid
	TempKefuId string // 临时客服ID
	Status     string // 当前状态
	EntId      string // 企业ID
	LoginIp    string // 登录IP
	CreatedAt  string // 创建时间
}

// wechatLoginColumns holds the columns for the table wechat_login.
var wechatLoginColumns = WechatLoginColumns{
	Id:         "id",
	KefuName:   "kefu_name",
	OpenId:     "open_id",
	TempKefuId: "temp_kefu_id",
	Status:     "status",
	EntId:      "ent_id",
	LoginIp:    "login_ip",
	CreatedAt:  "created_at",
}

// NewWechatLoginDao creates and returns a new DAO object for table data access.
func NewWechatLoginDao() *WechatLoginDao {
	return &WechatLoginDao{
		group:   "default",
		table:   "wechat_login",
		columns: wechatLoginColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WechatLoginDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WechatLoginDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WechatLoginDao) Columns() WechatLoginColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WechatLoginDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WechatLoginDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WechatLoginDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
