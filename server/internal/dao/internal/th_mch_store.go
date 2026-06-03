// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThMchStoreDao is the data access object for the table hg_th_mch_store.
type ThMchStoreDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns ThMchStoreColumns // columns contains all the column names of Table for convenient usage.
}

// ThMchStoreColumns defines and stores column names for the table hg_th_mch_store.
type ThMchStoreColumns struct {
	Id                 string //
	MchId              string // 商户ID
	StoreName          string // 门店名称（多语）
	Images             string // 图集
	PhoneArea          string // 区号
	Phone              string // 电话
	DetailAddress      string // 详细地址
	GgLat              string // 谷歌纬度
	GgLng              string // 谷歌经度
	VerifyNum          string // 核销数量
	Account            string // 账号
	PasswordHash       string // 密码
	Salt               string // 密码盐
	PasswordResetToken string // 密码重置令牌
	Status             string // 1、启用 2、禁用
	CreateAt           string // 创建时间
	UpdateAt           string // 更新时间
	DeletedAt          string //
}

// thMchStoreColumns holds the columns for the table hg_th_mch_store.
var thMchStoreColumns = ThMchStoreColumns{
	Id:                 "id",
	MchId:              "mch_id",
	StoreName:          "store_name",
	Images:             "images",
	PhoneArea:          "phone_area",
	Phone:              "phone",
	DetailAddress:      "detail_address",
	GgLat:              "gg_lat",
	GgLng:              "gg_lng",
	VerifyNum:          "verify_num",
	Account:            "account",
	PasswordHash:       "password_hash",
	Salt:               "salt",
	PasswordResetToken: "password_reset_token",
	Status:             "status",
	CreateAt:           "create_at",
	UpdateAt:           "update_at",
	DeletedAt:          "deleted_at",
}

// NewThMchStoreDao creates and returns a new DAO object for table data access.
func NewThMchStoreDao() *ThMchStoreDao {
	return &ThMchStoreDao{
		group:   "default",
		table:   "hg_th_mch_store",
		columns: thMchStoreColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ThMchStoreDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ThMchStoreDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ThMchStoreDao) Columns() ThMchStoreColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ThMchStoreDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ThMchStoreDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ThMchStoreDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
