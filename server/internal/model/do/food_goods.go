// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodGoods is the golang structure of table hg_food_goods for DAO operations like Where/Data.
type FoodGoods struct {
	g.Meta              `orm:"table:hg_food_goods, do:true"`
	Id                  interface{} //
	RestaurantId        interface{} // 餐厅ID
	ToretaCourseId      interface{} // Toreta的课程ID
	GoodsName           interface{} // 套餐名称
	Introduction        interface{} // 促销语
	LabelIds            interface{} // 标签（多选）
	Images              interface{} // 图集
	Price               interface{} // 套餐售价
	MarketPrice         interface{} // 套餐原价
	MaxOrderNum         interface{} // 单次最大可预定数
	TimeDuration        interface{} // 用餐停留时间
	GoodsContent        interface{} // 套餐详情
	Notice              interface{} // 注意事项
	GoodsState          interface{} // 状态（1.正常2下架）
	TotalOrderNum       interface{} // 预约单购买套餐的总数量（包含退款）
	TotalOrderAmount    interface{} // 预约单总金额（包含退款）
	PayOrderNum         interface{} // 预约单支付购买套餐的数量（不包含退款）
	PayOrderAmount      interface{} // 预约单支付金额（不包含退款）
	IsThCouponExclusive interface{} // 是否是礼品券兑换专属：0-否，1-是
	ThCouponId          interface{} // 绑定的礼品券ID，用于礼品券兑换专属商品
	IsNoPay             interface{} // 是否无需支付：0-否，1-是
	MaxTimeOrderOpen    interface{} // 是否开启时间限制 1-开始  2-关闭
	OrderTimeForm       interface{} // 限制时段
	CreateAt            *gtime.Time // 创建时间
	UpdateAt            *gtime.Time // 更新时间
	DeletedAt           *gtime.Time //
}
