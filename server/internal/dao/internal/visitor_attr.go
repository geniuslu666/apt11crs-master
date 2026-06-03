// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitorAttrDao is the data access object for the table visitor_attr.
type VisitorAttrDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns VisitorAttrColumns // columns contains all the column names of Table for convenient usage.
}

// VisitorAttrColumns defines and stores column names for the table visitor_attr.
type VisitorAttrColumns struct {
	Id            string //
	VisitorId     string // 访客ID
	RealName      string // 访客真实姓名
	Tel           string // 访客手机号
	Email         string // 访客邮箱
	Qq            string // 访客QQ
	Wechat        string // 访客微信
	Comment       string // 访客备注
	CreatedAt     string // 创建时间
	EntId         string // 对接企业ID
	MaxMessageNum string // 访客最大消息数
}

// visitorAttrColumns holds the columns for the table visitor_attr.
var visitorAttrColumns = VisitorAttrColumns{
	Id:            "id",
	VisitorId:     "visitor_id",
	RealName:      "real_name",
	Tel:           "tel",
	Email:         "email",
	Qq:            "qq",
	Wechat:        "wechat",
	Comment:       "comment",
	CreatedAt:     "created_at",
	EntId:         "ent_id",
	MaxMessageNum: "max_message_num",
}

// NewVisitorAttrDao creates and returns a new DAO object for table data access.
func NewVisitorAttrDao() *VisitorAttrDao {
	return &VisitorAttrDao{
		group:   "default",
		table:   "visitor_attr",
		columns: visitorAttrColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VisitorAttrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VisitorAttrDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VisitorAttrDao) Columns() VisitorAttrColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VisitorAttrDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VisitorAttrDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VisitorAttrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
