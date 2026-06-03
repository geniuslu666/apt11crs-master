// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PmsRoomStatusDao is the data access object for the table hg_pms_room_status.
type PmsRoomStatusDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns PmsRoomStatusColumns // columns contains all the column names of Table for convenient usage.
}

// PmsRoomStatusColumns defines and stores column names for the table hg_pms_room_status.
type PmsRoomStatusColumns struct {
	Id        string // 主键
	Puid      string // 物业ID
	Tuid      string // 房型ID
	Ruid      string // 房间ID
	RoomNo    string // 房间号
	Date      string // 日期
	Status    string // WAIT、等待预定  CONFIRM、确认预定  CHECK_IN、入住中  CHECK_OUT、退房  CHECK_IN_EXPIRE、签到到期  CHECK_OUT_EXPIRE、退房到期   LOCK、锁定
	ReserveId string // 预定单ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// pmsRoomStatusColumns holds the columns for the table hg_pms_room_status.
var pmsRoomStatusColumns = PmsRoomStatusColumns{
	Id:        "id",
	Puid:      "puid",
	Tuid:      "tuid",
	Ruid:      "ruid",
	RoomNo:    "room_no",
	Date:      "date",
	Status:    "status",
	ReserveId: "reserve_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPmsRoomStatusDao creates and returns a new DAO object for table data access.
func NewPmsRoomStatusDao() *PmsRoomStatusDao {
	return &PmsRoomStatusDao{
		group:   "default",
		table:   "hg_pms_room_status",
		columns: pmsRoomStatusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PmsRoomStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PmsRoomStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PmsRoomStatusDao) Columns() PmsRoomStatusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PmsRoomStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PmsRoomStatusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PmsRoomStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
