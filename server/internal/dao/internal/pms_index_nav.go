// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsIndexNavDao is the data access object for table hg_pms_index_nav.
type PmsIndexNavDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PmsIndexNavColumns // columns contains all the column names of Table for convenient usage.
}

// PmsIndexNavColumns defines and stores column names for table hg_pms_index_nav.
type PmsIndexNavColumns struct {
	Id           string //
	Name         string // 导航名称(多语言)
	Tag          string // 标签
	Image        string // 图标
	AppLink      string // app跳转链接
	WxLink       string // 微信跳转链接
	Sort         string // 排序(越大越靠前)
	Status       string // 状态1、启用 2、禁用
	MinappStatus string // 是否排除小程序显示 1-不排除 2-排除
	Chain        string // IN 内链  OUT 外联
	LinkOpenType string // 外部链接跳转方式 1-webview 2-浏览器
	CreateAt     string // 创建时间
	UpdateAt     string // 更新时间
	DeletedAt    string //
}

// pmsIndexNavColumns holds the columns for table hg_pms_index_nav.
var pmsIndexNavColumns = PmsIndexNavColumns{
	Id:           "id",
	Name:         "name",
	Tag:          "tag",
	Image:        "image",
	AppLink:      "app_link",
	WxLink:       "wx_link",
	Sort:         "sort",
	Status:       "status",
	MinappStatus: "minapp_status",
	Chain:        "chain",
	LinkOpenType: "link_open_type",
	CreateAt:     "create_at",
	UpdateAt:     "update_at",
	DeletedAt:    "deleted_at",
}

// NewPmsIndexNavDao creates and returns a new DAO object for table data access.
func NewPmsIndexNavDao() *PmsIndexNavDao {
	return &PmsIndexNavDao{
		group:   "default",
		table:   "hg_pms_index_nav",
		columns: pmsIndexNavColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PmsIndexNavDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PmsIndexNavDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PmsIndexNavDao) Columns() PmsIndexNavColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PmsIndexNavDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PmsIndexNavDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PmsIndexNavDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
