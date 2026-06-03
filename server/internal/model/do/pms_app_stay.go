// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppStay is the golang structure of table hg_pms_app_stay for DAO operations like Where/Data.
type PmsAppStay struct {
	g.Meta              `orm:"table:hg_pms_app_stay, do:true"`
	Id                  interface{} // APP订单主键
	Uid                 interface{} // 我方系统 ID
	Uuid                interface{} // 三方系统 ID
	Puid                interface{} // 物业ID
	Source              interface{} // 订单来源
	MemberId            interface{} // 用户ID
	OrderSn             interface{} // 订单号
	OutOrderSn          interface{} // 三方订单号
	Booker              interface{} // 预定人
	OrderAmount         interface{} // 订单金额
	PayModel            interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	OrderStatus         interface{} // 订单付款状态
	ExpirationTime      interface{} // 订单过期时间
	CancelTime          *gtime.Time // 未支付取消时间
	RefundStatus        interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundTime          *gtime.Time // 退款时间
	RefundAmount        interface{} // 已退款金额
	CleanFee            interface{} // 取消费用
	CheckInDate         interface{} // 入住时间
	CheckOutDate        interface{} // 退房时间
	CancelRate          interface{} // 退款政策
	Referrer            interface{} // 推荐人
	RebateRate          interface{} // 分佣比例
	RebateStatus        interface{} // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount        interface{} // 分佣结算金额
	RebateTime          *gtime.Time // 分佣结算时间
	IsGetOpen           interface{} // 是否开启积分获取
	IsPayOpen           interface{} // 是否开启积分抵扣
	HotelGetRateVip     interface{} // 酒店结算积分比例
	HotelGetRateScene   interface{} // 场景结算积分比例
	HotelGetScoreStatus interface{} // 'WAIT','SUCCESS','FAIL'
	HotelGetAmount      interface{} // 结算积分金额
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	PricePercent        interface{} // 全局溢价比例
	TotalAmount         interface{} // 订单总价
	ChangeAmount        interface{} // 优惠金额
	IsFx                interface{} // 是否是分销订单
}
