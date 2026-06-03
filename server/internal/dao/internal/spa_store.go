// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaStoreDao is the data access object for the table hg_spa_store.
type SpaStoreDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns SpaStoreColumns // columns contains all the column names of Table for convenient usage.
}

// SpaStoreColumns defines and stores column names for the table hg_spa_store.
type SpaStoreColumns struct {
	Id            string //
	Name          string // 名称
	HeadName      string // 负责人
	PhoneArea     string // 区号
	Phone         string // 联系电话
	OpenTime      string // 营业时间
	AreaPid       string // 地区省级ID
	AreaId        string // 地区市级ID
	DetailAddress string // 详细地址
	GgLat         string // 谷歌纬度
	GgLng         string // 谷歌经度
	Lat           string // 纬度
	Lng           string // 经度
	CreateAt      string // 创建时间
	UpdateAt      string // 更新时间
	DeletedAt     string //
}

// spaStoreColumns holds the columns for the table hg_spa_store.
var spaStoreColumns = SpaStoreColumns{
	Id:            "id",
	Name:          "name",
	HeadName:      "head_name",
	PhoneArea:     "phone_area",
	Phone:         "phone",
	OpenTime:      "open_time",
	AreaPid:       "area_pid",
	AreaId:        "area_id",
	DetailAddress: "detail_address",
	GgLat:         "gg_lat",
	GgLng:         "gg_lng",
	Lat:           "lat",
	Lng:           "lng",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
	DeletedAt:     "deleted_at",
}

// NewSpaStoreDao creates and returns a new DAO object for table data access.
func NewSpaStoreDao() *SpaStoreDao {
	return &SpaStoreDao{
		group:   "default",
		table:   "hg_spa_store",
		columns: spaStoreColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpaStoreDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpaStoreDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpaStoreDao) Columns() SpaStoreColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpaStoreDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpaStoreDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SpaStoreDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
