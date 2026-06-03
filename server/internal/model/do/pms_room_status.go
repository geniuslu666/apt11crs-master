// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomStatus is the golang structure of table hg_pms_room_status for DAO operations like Where/Data.
type PmsRoomStatus struct {
	g.Meta    `orm:"table:hg_pms_room_status, do:true"`
	Id        interface{} // 主键
	Puid      interface{} // 物业ID
	Tuid      interface{} // 房型ID
	Ruid      interface{} // 房间ID
	RoomNo    interface{} // 房间号
	Date      *gtime.Time // 日期
	Status    interface{} // WAIT、等待预定  CONFIRM、确认预定  CHECK_IN、入住中  CHECK_OUT、退房  CHECK_IN_EXPIRE、签到到期  CHECK_OUT_EXPIRE、退房到期   LOCK、锁定
	ReserveId interface{} // 预定单ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
