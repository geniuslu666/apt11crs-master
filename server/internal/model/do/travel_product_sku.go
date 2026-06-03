// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelProductSku is the golang structure of table hg_travel_product_sku for DAO operations like Where/Data.
type TravelProductSku struct {
	g.Meta           `orm:"table:hg_travel_product_sku, do:true"`
	Id               interface{} //
	ProductId        interface{} // 产品ID
	Name             interface{} // 车型名称（默认语言；多语言存 hg_pms_language）
	Price            interface{} // 售价（元）
	DailyCapacity    interface{} // 每日最大接待人数
	Status           interface{} // 状态（1启用 2禁用）
	Sort             interface{} // 排序（越大越靠前）
	SalesNum         interface{} // 已售
	ContactMobile    interface{} // 联系电话
	MeetingPlace     interface{} // 集合地点
	MeetingTime      interface{} // 集合时间（格式：HH:MM）
	GgLat            interface{} // 谷歌纬度
	GgLng            interface{} // 谷歌经度
	DeletedAt        *gtime.Time // 软删除时间（NULL=正常）
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	AllowCancel      interface{} // 是否允许取消
	FreeCancelHours  interface{} // 几小时前免费
	CancelFeePercent interface{} // 取消费率%
}
