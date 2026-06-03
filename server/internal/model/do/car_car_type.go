// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCarType is the golang structure of table hg_car_car_type for DAO operations like Where/Data.
type CarCarType struct {
	g.Meta              `orm:"table:hg_car_car_type, do:true"`
	Id                  interface{} //
	Name                interface{} // 名称
	Desc                interface{} // 简介
	Image               interface{} // 图片
	SeatNum             interface{} // 座位数
	PassengerNum        interface{} // 建议乘员人数
	MaxPackageNum       interface{} // 最大容纳行李件数
	ChildrenSeatSupport interface{} // 是否支持儿童座椅
	ChildrenSeatSpace   interface{} // 儿童座椅占几个作为
	Content             interface{} // 内容
	Status              interface{} // 状态1、启用 2、禁用
	Sort                interface{} // 排序(越大越靠前)
	CreateAt            *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 更新时间
	DeletedAt           *gtime.Time // 删除时间
}
