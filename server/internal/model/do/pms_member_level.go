// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLevel is the golang structure of table hg_pms_member_level for DAO operations like Where/Data.
type PmsMemberLevel struct {
	g.Meta         `orm:"table:hg_pms_member_level, do:true"`
	Id             interface{} //
	LevelName      interface{} // 会员等级名称
	Exp            interface{} // 达到等级所需经验值
	HotelGetRate   interface{} // 酒店场景获取积分倍率
	FoodGetRate    interface{} // 餐饮场景获取积分倍率
	SpaGetRate     interface{} // 按摩场景获取积分倍率
	CarGetRate     interface{} // 接送机/包车场景获取积分倍率
	CabinetGetRate interface{} // 储物柜场景获取积分倍率
	TravelGetRate  interface{} // 一日游场景获取积分倍率
	Desc           interface{} // 等级说明
	WordColor      interface{} // 等级字体颜色
	LevelBadge     interface{} // 等级徽章（单图）
	LevelCard      interface{} // 等级卡片（单图）
	LevelBigPic    interface{} // 等级大图（单图）
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
