// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomStatus is the golang structure for table pms_room_status.
type PmsRoomStatus struct {
	Id        int         `json:"id"        orm:"id"         description:"主键"`
	Puid      string      `json:"puid"      orm:"puid"       description:"物业ID"`
	Tuid      string      `json:"tuid"      orm:"tuid"       description:"房型ID"`
	Ruid      string      `json:"ruid"      orm:"ruid"       description:"房间ID"`
	RoomNo    string      `json:"roomNo"    orm:"room_no"    description:"房间号"`
	Date      *gtime.Time `json:"date"      orm:"date"       description:"日期"`
	Status    string      `json:"status"    orm:"status"     description:"WAIT、等待预定  CONFIRM、确认预定  CHECK_IN、入住中  CHECK_OUT、退房  CHECK_IN_EXPIRE、签到到期  CHECK_OUT_EXPIRE、退房到期   LOCK、锁定"`
	ReserveId string      `json:"reserveId" orm:"reserve_id" description:"预定单ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
