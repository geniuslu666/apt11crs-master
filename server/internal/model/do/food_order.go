// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodOrder is the golang structure of table hg_food_order for DAO operations like Where/Data.
type FoodOrder struct {
	g.Meta                   `orm:"table:hg_food_order, do:true"`
	Id                       interface{} //
	OrderType                interface{} // 订单类型
	OrderSn                  interface{} // 订单编号
	OutOrderSn               interface{} // 三方订单号
	MemberId                 interface{} // 用户ID
	RestaurantId             interface{} // 餐厅ID
	OrderAmount              interface{} // 订单金额
	CouponAmount             interface{} // 优惠券抵扣金额
	BalAmount                interface{} // 积分抵扣金额
	GoodsId                  interface{} // 套餐ID
	BookingName              interface{} // 预定人姓名
	FirstName                interface{} // 订单预定人姓
	LastName                 interface{} // 订单预定人名
	GoodsNum                 interface{} // 套餐数量
	PhoneArea                interface{} // 手机区号
	BookingMobile            interface{} // 预定人手机
	BookingEmail             interface{} // 预定人邮箱
	BookDate                 interface{} // 预定日期
	BookTime                 interface{} // 预定时间
	BookDatetime             *gtime.Time // 预定日期时间
	SeatId                   interface{} // 座位ID
	PayModel                 interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                  *gtime.Time // 支付时间
	OrderStatus              interface{} // 订单付款状态
	BookingStatus            interface{} // 订单预定状态
	BookingTime              *gtime.Time // 订单确认时间
	ConfirmRefuseReason      interface{} // 审核拒绝原因
	VerifyStatus             interface{} // 订单核销状态
	VerifyCode               interface{} // 核销码
	VerifyTime               *gtime.Time // 核销时间
	ExpirationTime           interface{} // 订单过期时间
	SettlementRate           interface{} // 结算比例
	SettlementStatus         interface{} // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount         interface{} // 结算金额
	SettlementTime           *gtime.Time // 结算时间
	SettlementOrderId        interface{} // 结算单ID
	SettlementType           interface{} // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle          interface{} // 结算周期  1每日结算  2每周结算  3每月结算
	MemberMessage            interface{} // 购买人留言信息
	MemberMessageJa          interface{} // 购买人留言信息日语版
	RestaurantMessage        interface{} // 餐厅留言信息
	BookingCount             interface{} // 预定人数
	RefundStatus             interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundRate               interface{} // 退款比例（废弃）
	RefundFee                interface{} // 退款手续费
	RefundTime               *gtime.Time // 退款时间
	RefundAmount             interface{} // 已退款总金额
	RefundBalAmount          interface{} // 已退款积分
	RefundCouponAmount       interface{} // 已退款优惠券
	AdminRefundAmount        interface{} // 后台已退款总金额
	AdminRefundBalAmount     interface{} // 后台已退款积分
	AdminRefundCouponAmount  interface{} // 后台已退款优惠券
	AdminCancelReason        interface{} // 后台取消原因
	ActivityId               interface{} // 活动ID
	AdminCancelNum           interface{} // 后台取消次数
	Referrer                 interface{} // 推荐人
	RebateRate               interface{} // 分佣比例
	RebateStatus             interface{} // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount             interface{} // 分佣结算金额
	RebateTime               *gtime.Time // 分佣结算时间
	IsGetOpen                interface{} // 是否开启积分获取
	IsPayOpen                interface{} // 是否开启积分抵扣
	FoodGetRateVip           interface{} // 餐饮结算积分比例
	FoodGetRateScene         interface{} // 场景结算积分比例
	FoodGetScoreStatus       interface{} // 'WAIT','SUCCESS','FAIL'
	FoodGetAmount            interface{} // 结算积分金额
	ExpValue                 interface{} // 结算的经验值
	ExpTime                  *gtime.Time // 经验结算时间
	DepositRate              interface{} // 定金比例
	DepositAmount            interface{} // 定金
	DepositPayStatus         interface{} // 定金付款状态
	DepositPayTime           *gtime.Time // 定金支付时间
	DepositCancelTime        *gtime.Time // 定金取消时间
	DepositRefundTime        *gtime.Time // 定金退款时间
	DepositExpirationTime    interface{} // 定金支付过期时间
	RemainPayStatus          interface{} // 尾款付款状态
	RemainCancelTime         *gtime.Time // 尾款取消时间
	RemainRefundTime         *gtime.Time // 尾款退款时间
	DepositCancelSource      interface{} // 定金取消来源
	DepositCancelReason      interface{} // 定金取消原因
	RemainCancelSource       interface{} // 尾款取消来源
	RemainCancelReason       interface{} // 尾款取消原因
	PayStep                  interface{} //
	OldBookDate              interface{} // 原预定日期
	OldBookTime              interface{} // 原预定时间
	OldBookDatetime          *gtime.Time // 原预定日期时间
	OldBookingCount          interface{} // 原预定人数
	ToretaReservationNo      interface{} // Toreta预约号
	ToretaReservationId      interface{} // Toreta预约ID
	ToretaReservationStatus  interface{} // Toreta预约状态（0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置）
	ToretaReservationEndtime *gtime.Time // Toreta预约结束时间
	GoodsIsNoPay             interface{} // 套餐是否无需支付：0-否，1-是
	ToretaHasTimeLimit       interface{} // Toreta是否有时间限制  1-true  2-false
	ToretaEndTime            interface{} // Toreta用餐结束时间
	CreatedAt                *gtime.Time // 创建时间
	UpdatedAt                *gtime.Time // 更新时间
	IsFx                     interface{} // 是否是分销订单
}
