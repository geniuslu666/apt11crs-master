// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelOrder is the golang structure of table hg_travel_order for DAO operations like Where/Data.
type TravelOrder struct {
	g.Meta                  `orm:"table:hg_travel_order, do:true"`
	Id                      interface{} //
	OrderSn                 interface{} // 预约单号
	ProductId               interface{} // 产品ID
	SkuId                   interface{} // SKUID
	MemberId                interface{} // 会员ID
	BookingName             interface{} // 预订人姓名
	FirstName               interface{} // 订单预定人名
	LastName                interface{} // 订单预定人姓
	PhoneArea               interface{} // 手机区号
	BookingMobile           interface{} // 预订人电话
	BookingEmail            interface{} // 预定人邮箱
	BookingNum              interface{} // 预约人数
	BookDate                *gtime.Time // 预约日期
	OrderAmount             interface{} // 订单金额（元）
	CouponAmount            interface{} // 优惠券抵扣金额
	BalAmount               interface{} // 积分抵扣金额
	OrderStatus             interface{} // 订单状态
	PayModel                interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	PayStatus               interface{} // 订单付款状态
	PayTime                 *gtime.Time // 支付时间
	VerifyStaffId           interface{} // 核销人员ID
	VerifyTime              *gtime.Time // 核销时间
	ExpirationTime          interface{} // 订单过期时间
	CancelTime              *gtime.Time // 取消时间
	CancelFee               interface{} // 取消手续费（元）
	RefundStatus            interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundAmount            interface{} // 退款金额（元）
	RefundTime              *gtime.Time // 退款时间
	RefundBalAmount         interface{} // 已退款积分
	RefundCouponAmount      interface{} // 已退款优惠券
	RefundReason            interface{} // 退款原因
	CreatedAt               *gtime.Time // 创建时间
	UpdatedAt               *gtime.Time // 更新时间
	Referrer                interface{} // 推荐人
	RebateRate              interface{} // 分佣比例
	RebateStatus            interface{} // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            interface{} // 分佣结算金额
	RebateTime              *gtime.Time // 分佣结算时间
	IsGetOpen               interface{} // 是否开启积分获取
	IsPayOpen               interface{} // 是否开启积分抵扣
	TravelGetRateVip        interface{} // 结算积分比例
	TravelGetRateScene      interface{} // 场景结算积分比例
	TravelGetScoreStatus    interface{} // 'WAIT','SUCCESS','FAIL'
	TravelGetAmount         interface{} // 结算积分金额
	ExpValue                interface{} // 结算的经验值
	ExpTime                 *gtime.Time // 经验结算时间
	AdminRefundAmount       interface{} // 后台已退款总金额
	AdminRefundBalAmount    interface{} // 后台已退款积分
	AdminRefundCouponAmount interface{} // 后台已退款优惠券
	AdminCancelReason       interface{} // 后台取消原因
	AdminCancelNum          interface{} // 后台取消次数
	IsFx                    interface{} // 是否是分销订单
}
