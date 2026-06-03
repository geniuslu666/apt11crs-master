// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ToretaNotificationFailedLogDao is the data access object for the table hg_toreta_notification_failed_log.
type ToretaNotificationFailedLogDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of the current DAO.
	columns ToretaNotificationFailedLogColumns // columns contains all the column names of Table for convenient usage.
}

// ToretaNotificationFailedLogColumns defines and stores column names for the table hg_toreta_notification_failed_log.
type ToretaNotificationFailedLogColumns struct {
	Id             string // 主键ID
	ResourceKey    string // 资源密钥
	ResourceType   string // 资源类型
	ResourceAction string // 资源操作
	RestaurantKey  string // 餐厅密钥
	StatusCode     string // HTTP状态码
	RetryCount     string // 重试次数
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// toretaNotificationFailedLogColumns holds the columns for the table hg_toreta_notification_failed_log.
var toretaNotificationFailedLogColumns = ToretaNotificationFailedLogColumns{
	Id:             "id",
	ResourceKey:    "resource_key",
	ResourceType:   "resource_type",
	ResourceAction: "resource_action",
	RestaurantKey:  "restaurant_key",
	StatusCode:     "status_code",
	RetryCount:     "retry_count",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewToretaNotificationFailedLogDao creates and returns a new DAO object for table data access.
func NewToretaNotificationFailedLogDao() *ToretaNotificationFailedLogDao {
	return &ToretaNotificationFailedLogDao{
		group:   "default",
		table:   "hg_toreta_notification_failed_log",
		columns: toretaNotificationFailedLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ToretaNotificationFailedLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ToretaNotificationFailedLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ToretaNotificationFailedLogDao) Columns() ToretaNotificationFailedLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ToretaNotificationFailedLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ToretaNotificationFailedLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ToretaNotificationFailedLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
