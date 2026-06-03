// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomType is the golang structure of table hg_pms_room_type for DAO operations like Where/Data.
type PmsRoomType struct {
	g.Meta                 `orm:"table:hg_pms_room_type, do:true"`
	Uid                    interface{} // airhost房型ID
	Id                     interface{} // 主键ID
	Puid                   interface{} // 物业ID
	Cover                  interface{} // 封面
	CoverList              *gjson.Json // 照片墙
	Name                   interface{} // 房型多语言名称ID
	BasePrice              interface{} // 最低价格
	CheckinAt              interface{} // 入住时间
	CheckoutAt             interface{} // 退房时间
	BookingStyle           interface{} // 预订方式
	RoomStyle              interface{} // 房间的风格
	Occupancy              interface{} // 占用
	Size                   interface{} // 面积
	Bedrooms               interface{} // 卧室
	Bathrooms              interface{} // 浴室
	CleaningFee            interface{} // 清理费
	RatePlanId             interface{} // 费率ID
	AdditionalGuestAmounts interface{} // 额外客人金额
	OccupantsForBaseRate   interface{} // 无需增加额外客人金额人数
	RoomNum                interface{} // 房间数
	IsShow                 interface{} // 1、显示 0 隐藏
	CreateAt               *gtime.Time //
	UpdateAt               *gtime.Time //
}
