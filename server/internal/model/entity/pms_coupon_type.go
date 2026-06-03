// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCouponType is the golang structure for table pms_coupon_type.
type PmsCouponType struct {
	Id                        int         `json:"id"                        orm:"id"                            description:"优惠券ID"`
	Type                      string      `json:"type"                      orm:"type"                          description:"优惠券类型 reward-满减 discount-折扣 random-随机"`
	CouponName                string      `json:"couponName"                orm:"coupon_name"                   description:"优惠券名称"`
	Count                     int         `json:"count"                     orm:"count"                         description:"发放数量"`
	LeadCount                 int         `json:"leadCount"                 orm:"lead_count"                    description:"已领取数量"`
	UsedCount                 int         `json:"usedCount"                 orm:"used_count"                    description:"已使用数量"`
	AtLeast                   float64     `json:"atLeast"                   orm:"at_least"                      description:"满多少元使用 0代表无限制"`
	Money                     float64     `json:"money"                     orm:"money"                         description:"发放面额 当type为reward时需要添加"`
	Discount                  float64     `json:"discount"                  orm:"discount"                      description:"1 =< 折扣 <= 9.9 当type为discount时需要添加"`
	DiscountLimit             float64     `json:"discountLimit"             orm:"discount_limit"                description:"最多折扣金额 当type为discount时可选择性添加"`
	ValidityType              int         `json:"validityType"              orm:"validity_type"                 description:"过期类型1-古固定时间范围过期 2-领取之日固定日期后过期 3长期有效"`
	StartUseTime              *gtime.Time `json:"startUseTime"              orm:"start_use_time"                description:"使用开始日期 过期类型0时必填"`
	EndUseTime                *gtime.Time `json:"endUseTime"                orm:"end_use_time"                  description:"使用结束日期 过期类型0时必填"`
	FixedTerm                 int         `json:"fixedTerm"                 orm:"fixed_term"                    description:"当validity_type为2时需要添加 领取之日起或者次日N天内有效"`
	Sort                      int         `json:"sort"                      orm:"sort"                          description:"排序"`
	MaxFetch                  int         `json:"maxFetch"                  orm:"max_fetch"                     description:"每人最大领取个数"`
	IsShow                    int         `json:"isShow"                    orm:"is_show"                       description:"是否允许直接领取"`
	DiscountAppStayOrderMoney float64     `json:"discountAppStayOrderMoney" orm:"discount_app_stay_order_money" description:"住宿订单的优惠总金额"`
	AppStayOrderMoney         float64     `json:"appStayOrderMoney"         orm:"app_stay_order_money"          description:"住宿订单用券总成交额"`
	Status                    int         `json:"status"                    orm:"status"                        description:"状态（1进行中2已结束-1已关闭）"`
	Scene                     int         `json:"scene"                     orm:"scene"                         description:"场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车  5-储物柜"`
	PropertyIds               string      `json:"propertyIds"               orm:"property_ids"                  description:""`
	RestaurantIds             string      `json:"restaurantIds"             orm:"restaurant_ids"                description:"餐厅ID"`
	ServiceIds                string      `json:"serviceIds"                orm:"service_ids"                   description:"按摩服务ID"`
	CarServiceTypes           string      `json:"carServiceTypes"           orm:"car_service_types"             description:"汽车服务类型"`
	Desc                      string      `json:"desc"                      orm:"desc"                          description:"优惠券使用说明"`
	CreateAt                  *gtime.Time `json:"createAt"                  orm:"create_at"                     description:"创建时间"`
	UpdateAt                  *gtime.Time `json:"updateAt"                  orm:"update_at"                     description:"修改时间"`
	DeletedAt                 *gtime.Time `json:"deletedAt"                 orm:"deleted_at"                    description:"删除时间"`
}
