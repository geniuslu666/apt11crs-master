// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitorTagDao is the data access object for the table visitor_tag.
type VisitorTagDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns VisitorTagColumns // columns contains all the column names of Table for convenient usage.
}

// VisitorTagColumns defines and stores column names for the table visitor_tag.
type VisitorTagColumns struct {
	Id        string //
	VisitorId string // 访客ID
	TagId     string // 标签ID
	EntId     string // 客服企业ID
	CreatedAt string // 创建时间
	Kefu      string // 客服账户
}

// visitorTagColumns holds the columns for the table visitor_tag.
var visitorTagColumns = VisitorTagColumns{
	Id:        "id",
	VisitorId: "visitor_id",
	TagId:     "tag_id",
	EntId:     "ent_id",
	CreatedAt: "created_at",
	Kefu:      "kefu",
}

// NewVisitorTagDao creates and returns a new DAO object for table data access.
func NewVisitorTagDao() *VisitorTagDao {
	return &VisitorTagDao{
		group:   "default",
		table:   "visitor_tag",
		columns: visitorTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VisitorTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VisitorTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VisitorTagDao) Columns() VisitorTagColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VisitorTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VisitorTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VisitorTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
