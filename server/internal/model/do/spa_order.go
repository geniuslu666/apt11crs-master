// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrder is the golang structure of table hg_spa_order for DAO operations like Where/Data.
type SpaOrder struct {
	g.Meta                  `orm:"table:hg_spa_order, do:true"`
	Id                      interface{} //
	IspId                   interface{} // 服务商ID
	OrderSn                 interface{} // 订单编号
	OutOrderSn              interface{} // 三方订单号
	ServiceType             interface{} // 1到店  2上门
	ConfirmType             interface{} // 1手动确认  2自动确认
	MemberId                interface{} // 用户ID
	BookDate                interface{} // 预定日期
	BookTime                interface{} // 预定时间
	BookStartTime           *gtime.Time // 预定开始日期时间
	BookEndTime             *gtime.Time // 预定结束日期时间
	TechnicianGoTime        *gtime.Time // 技师上门出发时间
	MemberArriveTime        *gtime.Time // 到店时间
	ActualStartTime         *gtime.Time // 实际开始日期时间
	ActualEndTime           *gtime.Time // 实际结束日期时间
	BookingName             interface{} // 预定人姓名
	PhoneArea               interface{} // 手机区号
	BookingMobile           interface{} // 预定人手机
	BookingEmail            interface{} // 预定人邮箱
	ServiceId               interface{} // 服务ID
	OrderAmount             interface{} // 订单金额
	CouponAmount            interface{} // 优惠券抵扣金额
	BalAmount               interface{} // 积分抵扣金额
	GoodsId                 interface{} // 项目ID
	GoodsNum                interface{} // 项目数量
	TechnicianIds           interface{} // 技师ID ,分隔
	IsReturn                interface{} // 是否退单中 1是  2否
	PayModel                interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 *gtime.Time // 支付时间
	PayStatus               interface{} // 订单付款状态
	StoreId                 interface{} // 到店门店ID
	PropertyId              interface{} // 上门物业ID
	RoomNo                  interface{} // 上门房间号
	OrderStatus             interface{} // 订单状态
	ConfirmTime             *gtime.Time // 订单确认时间
	ConfirmRefuseReason     interface{} // 审核拒绝原因
	DispatchStatus          interface{} // 订单调度状态
	DispatchTime            *gtime.Time // 订单调度时间
	DispatchDesc            interface{} // 订单调度备注
	DispatchOperatorId      interface{} // 调度操作人ID
	ExpirationTime          interface{} // 订单过期时间
	CancelTime              *gtime.Time // 取消时间
	StartServeImages        interface{} // 开始服务图集
	EndServeImages          interface{} // 服务结束图集
	SettlementRate          interface{} // 结算比例
	SettlementStatus        interface{} // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount        interface{} // 结算金额
	SettlementTime          *gtime.Time // 结算时间
	SettlementOrderId       interface{} // 结算单ID
	MemberMessage           interface{} // 购买人留言信息
	MemberMessageJa         interface{} // 购买人留言信息日语
	RefundStatus            interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee               interface{} // 退款手续费
	RefundTime              *gtime.Time // 退款时间
	RefundAmount            interface{} // 已退款总金额
	RefundBalAmount         interface{} // 已退款积分
	RefundCouponAmount      interface{} // 已退款优惠券
	AdminRefundAmount       interface{} // 后台已退款总金额
	AdminRefundBalAmount    interface{} // 后台已退款积分
	AdminRefundCouponAmount interface{} // 后台已退款优惠券
	RefundReason            interface{} // 退款原因
	Referrer                interface{} // 推荐人
	RebateRate              interface{} // 分佣比例
	RebateStatus            interface{} // WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败
	RebateAmount            interface{} // 分佣结算金额
	RebateTime              *gtime.Time // 分佣结算时间
	IsGetOpen               interface{} // 是否开启积分获取
	IsPayOpen               interface{} // 是否开启积分抵扣
	SpaGetRateVip           interface{} // 按摩结算积分比例
	SpaGetRateScene         interface{} // 场景结算积分比例
	SpaGetScoreStatus       interface{} // 'WAIT','SUCCESS','FAIL'
	SpaGetAmount            interface{} // 结算积分金额
	ExpValue                interface{} // 结算的经验值
	ExpTime                 *gtime.Time // 经验结算时间
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
