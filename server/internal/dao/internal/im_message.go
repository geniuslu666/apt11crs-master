// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImMessageDao is the data access object for the table hg_im_message.
type ImMessageDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns ImMessageColumns // columns contains all the column names of Table for convenient usage.
}

// ImMessageColumns defines and stores column names for the table hg_im_message.
type ImMessageColumns struct {
	Id           string //
	Source       string // 消息发起方
	SourceId     string // 身份对应角色ID
	SourceName   string // 身份对应角色昵称
	SourcePhoto  string // 身份对应角色照片
	OrderSn      string // 订单号
	FromMemberId string // 发起消息方的用户ID
	ToMemberId   string // 接收消息方的用户ID
	Content      string // 消息内容
	CreatedAt    string // 创建消息时间
	UpdatedAt    string // 更新消息时间
}

// imMessageColumns holds the columns for the table hg_im_message.
var imMessageColumns = ImMessageColumns{
	Id:           "id",
	Source:       "source",
	SourceId:     "source_id",
	SourceName:   "source_name",
	SourcePhoto:  "source_photo",
	OrderSn:      "order_sn",
	FromMemberId: "from_member_id",
	ToMemberId:   "to_member_id",
	Content:      "content",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewImMessageDao creates and returns a new DAO object for table data access.
func NewImMessageDao() *ImMessageDao {
	return &ImMessageDao{
		group:   "default",
		table:   "hg_im_message",
		columns: imMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImMessageDao) Columns() ImMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ImMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
