// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrderGoods is the golang structure of table hg_spa_order_goods for DAO operations like Where/Data.
type SpaOrderGoods struct {
	g.Meta             `orm:"table:hg_spa_order_goods, do:true"`
	Id                 interface{} //
	IspId              interface{} // 服务商ID
	OrderId            interface{} // 订单ID
	ServiceType        interface{} // 1到店  2上门
	ServiceId          interface{} // 服务ID
	TechnicianGoTime   *gtime.Time // 技师上门出发时间
	ActualStartTime    *gtime.Time // 实际开始日期时间
	ActualEndTime      *gtime.Time // 实际结束日期时间
	GoodsId            interface{} // 项目ID
	GoodsNum           interface{} // 项目数量
	OrderGoodsAmount   interface{} // 子订单金额
	CouponAmount       interface{} // 子订单优惠券抵扣金额
	BalAmount          interface{} // 子订单积分抵扣金额
	TechnicianIds      interface{} // 技师ID ,分隔
	PayStatus          interface{} // 订单付款状态
	StoreId            interface{} // 到店门店ID
	PropertyId         interface{} // 上门物业ID
	RoomNo             interface{} // 上门房间号
	OrderStatus        interface{} // 订单状态
	DispatchStatus     interface{} // 订单调度状态
	DispatchTime       *gtime.Time // 订单调度时间
	DispatchDesc       interface{} // 订单调度备注
	DispatchOperatorId interface{} // 调度操作人ID
	StartServeImages   interface{} // 开始服务图集
	EndServeImages     interface{} // 服务结束图集
	RefundStatus       interface{} // 退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款
	RefundFee          interface{} // 退款手续费
	RefundTime         *gtime.Time // 退款时间
	RefundAmount       interface{} // 已退款总金额
	RefundBalAmount    interface{} // 已退款积分
	RefundCouponAmount interface{} // 已退款优惠券
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
