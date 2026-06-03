// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsBannerDao is the data access object for the table hg_pms_banner.
type PmsBannerDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns PmsBannerColumns // columns contains all the column names of Table for convenient usage.
}

// PmsBannerColumns defines and stores column names for the table hg_pms_banner.
type PmsBannerColumns struct {
	Id           string //
	Language     string // 语言
	BannerName   string // banner名称
	BannerImage  string // 轮播图
	Model        string // 模块   BANNER  横幅   EVENTS  事件
	Chain        string // IN 内链  OUT 外联
	Path         string // 链接内容
	BannerStatus string // 1、启用 2、禁用
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
}

// pmsBannerColumns holds the columns for the table hg_pms_banner.
var pmsBannerColumns = PmsBannerColumns{
	Id:           "id",
	Language:     "language",
	BannerName:   "banner_name",
	BannerImage:  "banner_image",
	Model:        "model",
	Chain:        "chain",
	Path:         "path",
	BannerStatus: "banner_status",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
}

// NewPmsBannerDao creates and returns a new DAO object for table data access.
func NewPmsBannerDao() *PmsBannerDao {
	return &PmsBannerDao{
		group:   "default",
		table:   "hg_pms_banner",
		columns: pmsBannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsBannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsBannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsBannerDao) Columns() PmsBannerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsBannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsBannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsBannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
