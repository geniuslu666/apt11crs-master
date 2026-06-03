// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarOrder is the golang structure of table hg_car_order for DAO operations like Where/Data.
type CarOrder struct {
	g.Meta                  `orm:"table:hg_car_order, do:true"`
	Id                      interface{} //
	OrderType               interface{} // 订单类型
	OrderSn                 interface{} // 订单编号
	OutOrderSn              interface{} // 三方订单号
	ServiceType             interface{} // 服务类型
	ConfirmType             interface{} // 1手动确认  2自动确认
	DispatchType            interface{} // 1手动派单  2自动派单
	MemberId                interface{} // 用户ID
	BookDate                interface{} // 预定日期
	BookTime                interface{} // 预定时间
	AdultNum                interface{} // 成人数
	ChildNum                interface{} // 儿童数
	BookStartTime           *gtime.Time // 预定开始日期时间
	BookEndTime             *gtime.Time // 预定结束日期时间
	DriverGoTime            *gtime.Time // 司机出发时间
	ActualStartTime         *gtime.Time // 实际开始日期时间
	ActualEndTime           *gtime.Time // 实际结束日期时间
	BookingName             interface{} // 预定人姓名
	PhoneArea               interface{} // 手机区号
	BookingMobile           interface{} // 预定人手机
	BookingEmail            interface{} // 预定人邮箱
	StartAddressId          interface{} // 出发地ID
	EndAddressId            interface{} // 目的地ID
	ServiceId               interface{} // 服务ID
	DriverId                interface{} // 司机ID
	CarId                   interface{} // 车辆ID
	IsReturn                interface{} // 是否退单中 1是  2否
	ReturnDriverId          interface{} // 退单司机ID
	ReturnCarId             interface{} // 退单车辆ID
	PickUpSign              interface{} // 是否选择举牌接机
	PickUpSignAmount        interface{} // 举牌接机价格
	ChildSeatAddNum         interface{} // 婴儿座椅数量
	ChildSeatAddAmount      interface{} // 婴儿座椅总价
	OrderAmount             interface{} // 订单金额
	CouponAmount            interface{} // 优惠券抵扣金额
	NightAmount             interface{} // 深夜费
	BalAmount               interface{} // 积分抵扣金额
	PayModel                interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 *gtime.Time // 支付时间
	PayStatus               interface{} // 订单付款状态
	FlightNumber            interface{} // 航班号
	OrderStatus             interface{} // 订单状态
	ConfirmTime             *gtime.Time // 订单确认时间
	ConfirmRefuseReason     interface{} // 审核拒绝原因
	DispatchStatus          interface{} // 订单调度状态
	DispatchTime            *gtime.Time // 订单调度时间
	DispatchDesc            interface{} // 订单调度备注
	DispatchOperatorId      interface{} // 调度操作人ID
	ExpirationTime          interface{} // 订单过期时间
	CancelTime              *gtime.Time // 取消时间
	DriverLanguage          interface{} // 司机语言
	EmergencyName           interface{} // 紧急联系人
	EmergencyPhoneArea      interface{} // 紧急联系人手机区号
	EmergencyMobile         interface{} // 紧急联系人手机
	StartServeImages        interface{} // 开始服务图集
	EndServeImages          interface{} // 服务结束图集
	SettlementRate          interface{} // 结算比例
	SettlementStatus        interface{} // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount        interface{} // 结算金额
	SettlementTime          *gtime.Time // 结算时间
	SettlementOrderId       interface{} // 结算单ID
	SettlementType          interface{} // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle         interface{} // 结算周期  1每日结算  2每周结算  3每月结算
	MemberMessage           interface{} // 购买人留言信息
	MemberMessageJa         interface{} // 购买人留言信息日语
	RefundStatus            interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee               interface{} // 退款手续费
	RefundTime              *gtime.Time // 退款时间
	RefundAmount            interface{} // 已退款总金额
	RefundBalAmount         interface{} // 已退款积分
	RefundCouponAmount      interface{} // 已退款优惠券
	RefundReason            interface{} // 退款原因
	Referrer                interface{} // 推荐人
	RebateRate              interface{} // 分佣比例
	RebateStatus            interface{} // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            interface{} // 分佣结算金额
	RebateTime              *gtime.Time // 分佣结算时间
	IsGetOpen               interface{} // 是否开启积分获取
	IsPayOpen               interface{} // 是否开启积分抵扣
	CarGetRateVip           interface{} // 接送机结算积分比例
	CarGetRateScene         interface{} // 场景结算积分比例
	CarGetScoreStatus       interface{} // 'WAIT','SUCCESS','FAIL'
	CarGetAmount            interface{} // 结算积分金额
	InnnLuggageFreeNum      interface{} // innn免费行李数
	InnnCarPreAmount        interface{} // innn车型单价
	InnnCarTotalAmount      interface{} // innn车型总价
	InnnLuggagePreAmount    interface{} // innn行李额单价
	InnnLuggageNum          interface{} // innn额外行李数
	InnnLuggageTotalAmount  interface{} // innn行李额总价
	InnnOrderId             interface{} // innn订单ID
	InnnOrderNo             interface{} // innn订单号
	InnnLuggageOrderId      interface{} // innn行李订单ID
	InnnLuggageOrderNo      interface{} // innn行李订单号
	InnnQrcode              interface{} // innn二维码内容
	InnnQrcodeExpireTime    interface{} // innn二维码过期时间戳
	InnnOrderCancelSuccess  interface{} // innn取消接口是否请求成功  0无请求   1请求成功  2请求失败
	InnnCancelRule          interface{} // innn取消政策
	InnnVerifyStatus        interface{} // innn核销状态
	ExpValue                interface{} // 结算的经验值
	ExpTime                 *gtime.Time // 经验结算时间
	AdminRefundAmount       interface{} // 后台已退款总金额
	AdminRefundBalAmount    interface{} // 后台已退款积分
	AdminRefundCouponAmount interface{} // 后台已退款优惠券
	AdminCancelReason       interface{} // 后台取消原因
	AdminCancelNum          interface{} // 后台取消次数
	AbnormalStatus          interface{} // 异常单状态（1：不是异常单  2异常待处理  3异常已处理）
	AbnormalReason          interface{} // 异常处理原因
	AbnormalOperatorId      interface{} // 异常处理操作人ID
	AbnormalTime            *gtime.Time // 异常处理时间
	CreatedAt               *gtime.Time // 创建时间
	UpdatedAt               *gtime.Time // 更新时间
	IsFx                    interface{} // 是否是分销订单
}
