// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomUnit is the golang structure of table hg_pms_room_unit for DAO operations like Where/Data.
type PmsRoomUnit struct {
	g.Meta    `orm:"table:hg_pms_room_unit, do:true"`
	Uid       interface{} // 第三方系统的ID
	Id        interface{} // 主键
	RtUid     interface{} // 房型ID
	RoomNo    interface{} // 房间号
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
