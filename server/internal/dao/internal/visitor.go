// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitorDao is the data access object for the table visitor.
type VisitorDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns VisitorColumns // columns contains all the column names of Table for convenient usage.
}

// VisitorColumns defines and stores column names for the table visitor.
type VisitorColumns struct {
	Id        string //
	Name      string // 访客显示名称
	RealName  string // 访客真实姓名
	Avator    string // 访客头像
	SourceIp  string // 访客来源IP
	ToId      string // 对接客服账户
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	VisitorId string // 访客唯一ID
	Status    string // 访客状态
	State     string // 访客状态位
	Refer     string // 访客来源
	City      string // 访客城市
	ClientIp  string // 访客IP
	Extra     string // 访客扩展信息
	EntId     string // 对接的企业ID
	VisitNum  string // 访客访问次数
}

// visitorColumns holds the columns for the table visitor.
var visitorColumns = VisitorColumns{
	Id:        "id",
	Name:      "name",
	RealName:  "real_name",
	Avator:    "avator",
	SourceIp:  "source_ip",
	ToId:      "to_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	VisitorId: "visitor_id",
	Status:    "status",
	State:     "state",
	Refer:     "refer",
	City:      "city",
	ClientIp:  "client_ip",
	Extra:     "extra",
	EntId:     "ent_id",
	VisitNum:  "visit_num",
}

// NewVisitorDao creates and returns a new DAO object for table data access.
func NewVisitorDao() *VisitorDao {
	return &VisitorDao{
		group:   "default",
		table:   "visitor",
		columns: visitorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *VisitorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *VisitorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *VisitorDao) Columns() VisitorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *VisitorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *VisitorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *VisitorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
