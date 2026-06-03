// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodGoods is the golang structure for table food_goods.
type FoodGoods struct {
	Id                  int64       `json:"id"                  orm:"id"                     description:""`
	RestaurantId        uint64      `json:"restaurantId"        orm:"restaurant_id"          description:"餐厅ID"`
	ToretaCourseId      string      `json:"toretaCourseId"      orm:"toreta_course_id"       description:"Toreta的课程ID"`
	GoodsName           string      `json:"goodsName"           orm:"goods_name"             description:"套餐名称"`
	Introduction        string      `json:"introduction"        orm:"introduction"           description:"促销语"`
	LabelIds            string      `json:"labelIds"            orm:"label_ids"              description:"标签（多选）"`
	Images              string      `json:"images"              orm:"images"                 description:"图集"`
	Price               float64     `json:"price"               orm:"price"                  description:"套餐售价"`
	MarketPrice         float64     `json:"marketPrice"         orm:"market_price"           description:"套餐原价"`
	MaxOrderNum         uint        `json:"maxOrderNum"         orm:"max_order_num"          description:"单次最大可预定数"`
	TimeDuration        string      `json:"timeDuration"        orm:"time_duration"          description:"用餐停留时间"`
	GoodsContent        string      `json:"goodsContent"        orm:"goods_content"          description:"套餐详情"`
	Notice              string      `json:"notice"              orm:"notice"                 description:"注意事项"`
	GoodsState          int         `json:"goodsState"          orm:"goods_state"            description:"状态（1.正常2下架）"`
	TotalOrderNum       int         `json:"totalOrderNum"       orm:"total_order_num"        description:"预约单购买套餐的总数量（包含退款）"`
	TotalOrderAmount    float64     `json:"totalOrderAmount"    orm:"total_order_amount"     description:"预约单总金额（包含退款）"`
	PayOrderNum         int         `json:"payOrderNum"         orm:"pay_order_num"          description:"预约单支付购买套餐的数量（不包含退款）"`
	PayOrderAmount      float64     `json:"payOrderAmount"      orm:"pay_order_amount"       description:"预约单支付金额（不包含退款）"`
	IsThCouponExclusive int         `json:"isThCouponExclusive" orm:"is_th_coupon_exclusive" description:"是否是礼品券兑换专属：0-否，1-是"`
	ThCouponId          int         `json:"thCouponId"          orm:"th_coupon_id"           description:"绑定的礼品券ID，用于礼品券兑换专属商品"`
	IsNoPay             int         `json:"isNoPay"             orm:"is_no_pay"              description:"是否无需支付：0-否，1-是"`
	MaxTimeOrderOpen    uint        `json:"maxTimeOrderOpen"    orm:"max_time_order_open"    description:"是否开启时间限制 1-开始  2-关闭"`
	OrderTimeForm       string      `json:"orderTimeForm"       orm:"order_time_form"        description:"限制时段"`
	CreateAt            *gtime.Time `json:"createAt"            orm:"create_at"              description:"创建时间"`
	UpdateAt            *gtime.Time `json:"updateAt"            orm:"update_at"              description:"更新时间"`
	DeletedAt           *gtime.Time `json:"deletedAt"           orm:"deleted_at"             description:""`
}
