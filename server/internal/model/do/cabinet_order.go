// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CabinetOrder is the golang structure of table hg_cabinet_order for DAO operations like Where/Data.
type CabinetOrder struct {
	g.Meta                  `orm:"table:hg_cabinet_order, do:true"`
	Id                      interface{} //
	OrderSn                 interface{} // 订单编号
	OutOrderSn              interface{} // 三方订单号
	MinHours                interface{} // 最小小时数
	CabinetId               interface{} // 储物柜ID
	CabinetName             interface{} // 储物柜名称
	CabinetNameJson         *gjson.Json // 储物柜名称多语言JSON
	CityId                  interface{} // 所属城市ID
	CityName                interface{} // 所属城市名
	CityNameJson            *gjson.Json // 所属城市名多语言JSON
	MchId                   interface{} // 运营商ID
	MchName                 interface{} // 运营商名称
	MchNameJson             *gjson.Json // 运营商名称多语言JSON
	MchBranchId             interface{} // 网点ID
	MchBranchName           interface{} // 网点名称
	MchBranchNameJson       *gjson.Json // 网点名称多语言JSON
	MchBranchLat            interface{} // 网点lat
	MchBranchLgt            interface{} // 网点lgt
	Address                 interface{} // 地址
	AddressJson             *gjson.Json // 地址多语言JSON
	BoxTypeJson             *gjson.Json // 格口类型列表数据
	BoxTypeId               interface{} // 格口类型ID
	BoxTypeName             interface{} // 格口类型名称
	BoxTypeNameJson         *gjson.Json // 格口类型名称多语言JSON
	BoxTypePrice            interface{} // 格口类型单价
	BoxId                   interface{} // 格口ID
	BoxNo                   interface{} // 格口编号
	BoxAlias                interface{} // 格口别名
	OrderFirstFeeRate       interface{} // 首次下单优惠
	BuyHours                interface{} // 购买的小时数
	Pin                     interface{} // 取件码(4位数字)
	MemberId                interface{} // 用户ID
	OrderAmount             interface{} // 订单金额
	BaseAmount              interface{} // 租赁时间内金额
	CouponAmount            interface{} // 优惠券抵扣金额
	BalAmount               interface{} // 积分抵扣金额
	OvertimeSecs            interface{} // 超时时长(秒)
	OvertimeHours           interface{} // 超时时长(小时)
	OvertimeFee             interface{} // 超时费用(日元)
	OvertimePayTime         *gtime.Time // 超时费支付时间
	OvertimePayStatus       interface{} // 超时费付款状态
	OvertimeBalAmount       interface{} // 超时费积分抵扣金额
	OvertimePayModel        interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	GraceSeconds            interface{} // 宽限期设置时长（秒）
	GraceEndTime            *gtime.Time // 宽限期结束时间
	PayStep                 interface{} // 支付流程
	PayModel                interface{} // 1、余额支付 2、组合支付 3、纯外部支付
	PayTime                 *gtime.Time // 支付时间
	PayStatus               interface{} // 订单付款状态
	OrderStatus             interface{} // 订单状态
	ExpirationTime          interface{} // 订单过期时间
	CancelTime              *gtime.Time // 取消时间
	StartTime               *gtime.Time // 订单开始时间
	EndTime                 *gtime.Time // 订单结束时间
	FinishTime              *gtime.Time // 完成时间
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
	CabinetGetRateVip       interface{} // 储物柜结算积分比例
	CabinetGetRateScene     interface{} // 场景结算积分比例
	CabinetGetScoreStatus   interface{} // 'WAIT','SUCCESS','FAIL'
	CabinetGetAmount        interface{} // 结算积分金额
	ExpValue                interface{} // 结算的经验值
	ExpTime                 *gtime.Time // 经验结算时间
	IsAbnormal              interface{} // 请求下单接口是否异常
	PayOvertimeAbnormal     interface{} // 请求支付超时费接口是否异常
	IsAdminComplete         interface{} // 是否是后台强制完成
	AdminCompleteOperatorId interface{} // 强制完成处理操作人ID
	CreatedAt               *gtime.Time //
	UpdatedAt               *gtime.Time //
	IsFx                    interface{} // 是否是分销订单
}
