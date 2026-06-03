// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitorExtDao is the data access object for the table visitor_ext.
type VisitorExtDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns VisitorExtColumns // columns contains all the column names of Table for convenient usage.
}

// VisitorExtColumns defines and stores column names for the table visitor_ext.
type VisitorExtColumns struct {
	Id        string //
	VisitorId string // 访客ID
	EntId     string // 对接的企业ID
	Ua        string // 访客浏览器UserAgent
	Title     string // 页面标题
	Url       string // 页面地址
	Refer     string // 页面来源
	ReferUrl  string // 页面来源地址
	ClientIp  string // 访客IP
	City      string // 访客城市
	Language  string // 浏览器语言
	CreatedAt string // 创建时间
}

// visitorExtColumns holds the columns for the table visitor_ext.
var visitorExtColumns = VisitorExtColumns{
	Id:        "id",
	VisitorId: "visitor_id",
	EntId:     "ent_id",
	Ua:        "ua",
	Title:     "title",
	Url:       "url",
	Refer:     "refer",
	ReferUrl:  "refer_url",
	ClientIp:  "client_ip",
	City:      "city",
	Language:  "language",
	CreatedAt: "created_at",
}

// NewVisitorExtDao creates and returns a new DAO object for table data access.
func NewVisitorExtDao() *VisitorExtDao {
	return &VisitorExtDao{
		group:   "default",
		table:   "visitor_ext",
		columns: visitorExtColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VisitorExtDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VisitorExtDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VisitorExtDao) Columns() VisitorExtColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VisitorExtDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VisitorExtDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VisitorExtDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
