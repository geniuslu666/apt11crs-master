// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBedType is the golang structure of table hg_pms_bed_type for DAO operations like Where/Data.
type PmsBedType struct {
	g.Meta      `orm:"table:hg_pms_bed_type, do:true"`
	Id          interface{} // 床型ID
	RoomTypeId  interface{} // 房型ID
	BedTypeName interface{} // 床型名称
	BedWidth    interface{} // 床宽
	BedNum      interface{} // 数量
	CreateAt    *gtime.Time // 创建时间
	UpdateAt    *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
