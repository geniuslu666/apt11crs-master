// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsumerDao is the data access object for the table consumer.
type ConsumerDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns ConsumerColumns // columns contains all the column names of Table for convenient usage.
}

// ConsumerColumns defines and stores column names for the table consumer.
type ConsumerColumns struct {
	Id         string //
	Company    string // 公司名
	Realname   string // 姓名
	Score      string // 级别
	ConsumerSn string // 客户编号
	EntId      string // 企业ID
	KefuName   string // kefu名称
	Tel        string // 手机
	Wechat     string // 微信
	Qq         string // qq
	Email      string // 邮箱
	Remark     string // 备注
	CreatedAt  string // 创建时间
}

// consumerColumns holds the columns for the table consumer.
var consumerColumns = ConsumerColumns{
	Id:         "id",
	Company:    "company",
	Realname:   "realname",
	Score:      "score",
	ConsumerSn: "consumer_sn",
	EntId:      "ent_id",
	KefuName:   "kefu_name",
	Tel:        "tel",
	Wechat:     "wechat",
	Qq:         "qq",
	Email:      "email",
	Remark:     "remark",
	CreatedAt:  "created_at",
}

// NewConsumerDao creates and returns a new DAO object for table data access.
func NewConsumerDao() *ConsumerDao {
	return &ConsumerDao{
		group:   "default",
		table:   "consumer",
		columns: consumerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ConsumerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ConsumerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ConsumerDao) Columns() ConsumerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ConsumerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ConsumerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ConsumerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
