// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IndexBannerDao is the data access object for table hg_index_banner.
type IndexBannerDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns IndexBannerColumns // columns contains all the column names of Table for convenient usage.
}

// IndexBannerColumns defines and stores column names for table hg_index_banner.
type IndexBannerColumns struct {
	Id           string //
	Language     string // 语言
	BannerImage  string // 轮播图
	Model        string // 模块   BANNER  横幅   EVENTS  事件
	Chain        string // IN 内链  OUT 外联
	Path         string // 链接内容
	BannerStatus string // 1、启用 2、禁用
	MinappStatus string // 是否排除小程序显示 1-不排除 2-排除
	Sort         string // 排序(越大越靠前)
	LinkOpenType string // 外部链接跳转方式 1-webview 2-浏览器
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
}

// indexBannerColumns holds the columns for table hg_index_banner.
var indexBannerColumns = IndexBannerColumns{
	Id:           "id",
	Language:     "language",
	BannerImage:  "banner_image",
	Model:        "model",
	Chain:        "chain",
	Path:         "path",
	BannerStatus: "banner_status",
	MinappStatus: "minapp_status",
	Sort:         "sort",
	LinkOpenType: "link_open_type",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
}

// NewIndexBannerDao creates and returns a new DAO object for table data access.
func NewIndexBannerDao() *IndexBannerDao {
	return &IndexBannerDao{
		group:   "default",
		table:   "hg_index_banner",
		columns: indexBannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *IndexBannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *IndexBannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *IndexBannerDao) Columns() IndexBannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *IndexBannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *IndexBannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *IndexBannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
