// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserAttrDao is the data access object for the table user_attr.
type UserAttrDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns UserAttrColumns // columns contains all the column names of Table for convenient usage.
}

// UserAttrColumns defines and stores column names for the table user_attr.
type UserAttrColumns struct {
	Id               string //
	KefuName         string // 客服账户
	AuthKey          string // authKey
	IdCard           string // 身份证号码
	EntId            string // 企业ID
	GoodNum          string // 评价好评数量
	NormalNum        string // 评价中评数量
	BadNum           string // 评价差评数量
	AigcSessionScore string // aigc积分
	Money            string // 总金额
	CreatedAt        string // 创建时间
}

// userAttrColumns holds the columns for the table user_attr.
var userAttrColumns = UserAttrColumns{
	Id:               "id",
	KefuName:         "kefu_name",
	AuthKey:          "auth_key",
	IdCard:           "id_card",
	EntId:            "ent_id",
	GoodNum:          "good_num",
	NormalNum:        "normal_num",
	BadNum:           "bad_num",
	AigcSessionScore: "aigc_session_score",
	Money:            "money",
	CreatedAt:        "created_at",
}

// NewUserAttrDao creates and returns a new DAO object for table data access.
func NewUserAttrDao() *UserAttrDao {
	return &UserAttrDao{
		group:   "default",
		table:   "user_attr",
		columns: userAttrColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserAttrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserAttrDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserAttrDao) Columns() UserAttrColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserAttrDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserAttrDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserAttrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
