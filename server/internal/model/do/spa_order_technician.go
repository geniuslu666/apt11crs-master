// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrderTechnician is the golang structure of table hg_spa_order_technician for DAO operations like Where/Data.
type SpaOrderTechnician struct {
	g.Meta            `orm:"table:hg_spa_order_technician, do:true"`
	Id                interface{} //
	IspId             interface{} // 服务商ID
	OrderId           interface{} // 订单ID
	OrderGoodsId      interface{} // 子订单ID
	TechnicianId      interface{} // 技师ID
	SettlementObject  interface{} // 结算对象
	SettlementRate    interface{} // 结算比例
	SettlementStatus  interface{} // WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败
	SettlementAmount  interface{} // 结算金额
	SettlementTime    *gtime.Time // 结算时间
	SettlementOrderId interface{} // 结算单ID
	SettlementType    interface{} // 结算方式  1无需结算 2按周期自动结算  3手动申请结算
	SettlementCycle   interface{} // 结算周期  1每日结算  2每周结算  3每月结算
	SettlementCost    interface{} // 结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金
	OrderAmount       interface{} // 单笔订单金额
	CouponAmount      interface{} // 优惠券抵扣金额
	BalAmount         interface{} // 积分抵扣金额
	ActualEndTime     *gtime.Time // 实际结束日期时间
}
