// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsBedTypeDao is the data access object for the table hg_pms_bed_type.
type PmsBedTypeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns PmsBedTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PmsBedTypeColumns defines and stores column names for the table hg_pms_bed_type.
type PmsBedTypeColumns struct {
	Id          string // 床型ID
	RoomTypeId  string // 房型ID
	BedTypeName string // 床型名称
	BedWidth    string // 床宽
	BedNum      string // 数量
	CreateAt    string // 创建时间
	UpdateAt    string // 修改时间
	DeletedAt   string // 删除时间
}

// pmsBedTypeColumns holds the columns for the table hg_pms_bed_type.
var pmsBedTypeColumns = PmsBedTypeColumns{
	Id:          "id",
	RoomTypeId:  "room_type_id",
	BedTypeName: "bed_type_name",
	BedWidth:    "bed_width",
	BedNum:      "bed_num",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeletedAt:   "deleted_at",
}

// NewPmsBedTypeDao creates and returns a new DAO object for table data access.
func NewPmsBedTypeDao() *PmsBedTypeDao {
	return &PmsBedTypeDao{
		group:   "default",
		table:   "hg_pms_bed_type",
		columns: pmsBedTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsBedTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsBedTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsBedTypeDao) Columns() PmsBedTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsBedTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsBedTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsBedTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
