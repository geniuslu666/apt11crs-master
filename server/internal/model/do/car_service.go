// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarService is the golang structure of table hg_car_service for DAO operations like Where/Data.
type CarService struct {
	g.Meta           `orm:"table:hg_car_service, do:true"`
	Id               interface{} //
	ServiceType      interface{} // 服务类型
	ServiceName      interface{} // 路线名称
	CarTypeId        interface{} // 车型ID
	StartIds         interface{} // 出发地ID  多选
	EndIds           interface{} // 目的地ID 多选
	Distance         interface{} // 路线长度 KM
	UseTime          interface{} // 线路时长  分钟
	Price            interface{} // 基础费用
	LinePrice        interface{} // 划线价
	FreeWaitTime     interface{} // 免费等待时长  分钟
	MaxWaitTime      interface{} // 最大等待时长  分钟
	TimeoutPreTime   interface{} // 超时每xx分钟
	TimeoutPrePrice  interface{} // 超时价格
	Status           interface{} // 状态1、启用 2、禁用
	Sort             interface{} // 排序(越大越靠前)
	TotalOrderNum    interface{} // 预约单总数量（包含退款）
	TotalOrderAmount interface{} // 预约单总金额（包含退款）
	PayOrderNum      interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount   interface{} // 预约单支付金额（不包含退款）
	CreateAt         *gtime.Time // 创建时间
	UpdateAt         *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
