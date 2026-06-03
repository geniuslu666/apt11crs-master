// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsRoomUnitDao is the data access object for the table hg_pms_room_unit.
type PmsRoomUnitDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns PmsRoomUnitColumns // columns contains all the column names of Table for convenient usage.
}

// PmsRoomUnitColumns defines and stores column names for the table hg_pms_room_unit.
type PmsRoomUnitColumns struct {
	Uid       string // 第三方系统的ID
	Id        string // 主键
	RtUid     string // 房型ID
	RoomNo    string // 房间号
	CreatedAt string //
	UpdatedAt string //
}

// pmsRoomUnitColumns holds the columns for the table hg_pms_room_unit.
var pmsRoomUnitColumns = PmsRoomUnitColumns{
	Uid:       "uid",
	Id:        "id",
	RtUid:     "rt_uid",
	RoomNo:    "room_no",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPmsRoomUnitDao creates and returns a new DAO object for table data access.
func NewPmsRoomUnitDao() *PmsRoomUnitDao {
	return &PmsRoomUnitDao{
		group:   "default",
		table:   "hg_pms_room_unit",
		columns: pmsRoomUnitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsRoomUnitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsRoomUnitDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsRoomUnitDao) Columns() PmsRoomUnitColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsRoomUnitDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsRoomUnitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsRoomUnitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
