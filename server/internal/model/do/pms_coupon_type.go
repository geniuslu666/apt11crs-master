// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCouponType is the golang structure of table hg_pms_coupon_type for DAO operations like Where/Data.
type PmsCouponType struct {
	g.Meta                    `orm:"table:hg_pms_coupon_type, do:true"`
	Id                        interface{} // 优惠券ID
	Type                      interface{} // 优惠券类型 reward-满减 discount-折扣 random-随机
	CouponName                interface{} // 优惠券名称
	Count                     interface{} // 发放数量
	LeadCount                 interface{} // 已领取数量
	UsedCount                 interface{} // 已使用数量
	AtLeast                   interface{} // 满多少元使用 0代表无限制
	Money                     interface{} // 发放面额 当type为reward时需要添加
	Discount                  interface{} // 1 =< 折扣 <= 9.9 当type为discount时需要添加
	DiscountLimit             interface{} // 最多折扣金额 当type为discount时可选择性添加
	ValidityType              interface{} // 过期类型1-古固定时间范围过期 2-领取之日固定日期后过期 3长期有效
	StartUseTime              *gtime.Time // 使用开始日期 过期类型0时必填
	EndUseTime                *gtime.Time // 使用结束日期 过期类型0时必填
	FixedTerm                 interface{} // 当validity_type为2时需要添加 领取之日起或者次日N天内有效
	Sort                      interface{} // 排序
	MaxFetch                  interface{} // 每人最大领取个数
	IsShow                    interface{} // 是否允许直接领取
	DiscountAppStayOrderMoney interface{} // 住宿订单的优惠总金额
	AppStayOrderMoney         interface{} // 住宿订单用券总成交额
	Status                    interface{} // 状态（1进行中2已结束-1已关闭）
	Scene                     interface{} // 场景 1-住宿 2-餐饮  3-按摩 4-接送机/包车  5-储物柜
	PropertyIds               interface{} //
	RestaurantIds             interface{} // 餐厅ID
	ServiceIds                interface{} // 按摩服务ID
	CarServiceTypes           interface{} // 汽车服务类型
	Desc                      interface{} // 优惠券使用说明
	CreateAt                  *gtime.Time // 创建时间
	UpdateAt                  *gtime.Time // 修改时间
	DeletedAt                 *gtime.Time // 删除时间
}
