// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaServiceGoods is the golang structure of table hg_spa_service_goods for DAO operations like Where/Data.
type SpaServiceGoods struct {
	g.Meta           `orm:"table:hg_spa_service_goods, do:true"`
	Id               interface{} //
	ServiceId        interface{} // 服务ID
	GoodsName        interface{} // 服务套餐名称(多语言)
	Image            interface{} // 图片
	DetailImage      interface{} // 详情图
	Price            interface{} // 套餐售价
	Duration         interface{} // 时长(分钟)
	TotalOrderNum    interface{} // 预约单总数量（包含退款）
	TotalOrderAmount interface{} // 预约单总金额（包含退款）
	PayOrderNum      interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount   interface{} // 预约单支付金额（不包含退款）
	Status           interface{} // 状态1、启用 2、禁用
	CreateAt         *gtime.Time // 创建时间
	UpdateAt         *gtime.Time // 更新时间
	DeletedAt        *gtime.Time //
}
