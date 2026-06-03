// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCar is the golang structure of table hg_car_car for DAO operations like Where/Data.
type CarCar struct {
	g.Meta              `orm:"table:hg_car_car, do:true"`
	Id                  interface{} //
	CarName             interface{} // 车辆名称
	TypeId              interface{} // 车型ID
	Brand               interface{} // 车辆品牌型号
	LicenseNo           interface{} // 车牌号码
	LicenseColor        interface{} // 牌照颜色
	SeatNum             interface{} // 座位数(废弃)
	PassengerNum        interface{} // 建议乘员人数(废弃)
	MaxPackageNum       interface{} // 最大容纳行李件数(废弃)
	ChildrenSeatSupport interface{} // 是否支持儿童座椅(废弃)
	ChildrenSeatSpace   interface{} // 儿童座椅占几个作为(废弃)
	WorkStatus          interface{} // 工作状态
	Status              interface{} // 状态1、启用 2、禁用
	Sort                interface{} // 排序(越大越靠前)
	QualityMaterials    interface{} // 资质信息(多图)
	TotalOrderNum       interface{} // 预约单总数量（包含退款）
	TotalOrderAmount    interface{} // 预约单总金额（包含退款）
	PayOrderNum         interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount      interface{} // 预约单支付金额（不包含退款）
	CreateAt            *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 更新时间
	DeletedAt           *gtime.Time //
}
