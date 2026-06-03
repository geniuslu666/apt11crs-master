// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBedType is the golang structure for table pms_bed_type.
type PmsBedType struct {
	Id          int         `json:"id"          orm:"id"            description:"床型ID"`
	RoomTypeId  int         `json:"roomTypeId"  orm:"room_type_id"  description:"房型ID"`
	BedTypeName string      `json:"bedTypeName" orm:"bed_type_name" description:"床型名称"`
	BedWidth    string      `json:"bedWidth"    orm:"bed_width"     description:"床宽"`
	BedNum      int         `json:"bedNum"      orm:"bed_num"       description:"数量"`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"     description:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"     description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:"删除时间"`
}
