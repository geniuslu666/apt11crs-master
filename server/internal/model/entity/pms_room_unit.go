// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomUnit is the golang structure for table pms_room_unit.
type PmsRoomUnit struct {
	Uid       string      `json:"uid"       orm:"uid"        description:"第三方系统的ID"`
	Id        int         `json:"id"        orm:"id"         description:"主键"`
	RtUid     string      `json:"rtUid"     orm:"rt_uid"     description:"房型ID"`
	RoomNo    string      `json:"roomNo"    orm:"room_no"    description:"房间号"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
