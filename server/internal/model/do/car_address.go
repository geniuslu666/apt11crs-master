// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarAddress is the golang structure of table hg_car_address for DAO operations like Where/Data.
type CarAddress struct {
	g.Meta        `orm:"table:hg_car_address, do:true"`
	Id            interface{} //
	TypeId        interface{} // 地点类型id
	Name          interface{} // 地点名称（后台）
	SubName       interface{} // 地点名称（app-多语）
	AirportCode   interface{} // 机场代码
	TerminalName  interface{} // 机场航站楼名称
	DetailAddress interface{} // 详细地址-多语
	GgLat         interface{} // 谷歌纬度
	GgLng         interface{} // 谷歌经度
	Lat           interface{} // 纬度
	Lng           interface{} // 经度
	Status        interface{} // 状态1、启用 2、禁用
	PropertyId    interface{} // 物业id
	CreateAt      *gtime.Time // 创建时间
	UpdateAt      *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
