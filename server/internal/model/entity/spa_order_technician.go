// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrderTechnician is the golang structure for table spa_order_technician.
type SpaOrderTechnician struct {
	Id                int64       `json:"id"                orm:"id"                  description:""`
	IspId             int         `json:"ispId"             orm:"isp_id"              description:"服务商ID"`
	OrderId           int64       `json:"orderId"           orm:"order_id"            description:"订单ID"`
	OrderGoodsId      int64       `json:"orderGoodsId"      orm:"order_goods_id"      description:"子订单ID"`
	TechnicianId      int64       `json:"technicianId"      orm:"technician_id"       description:"技师ID"`
	SettlementObject  string      `json:"settlementObject"  orm:"settlement_object"   description:"结算对象"`
	SettlementRate    float64     `json:"settlementRate"    orm:"settlement_rate"     description:"结算比例"`
	SettlementStatus  string      `json:"settlementStatus"  orm:"settlement_status"   description:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount  float64     `json:"settlementAmount"  orm:"settlement_amount"   description:"结算金额"`
	SettlementTime    *gtime.Time `json:"settlementTime"    orm:"settlement_time"     description:"结算时间"`
	SettlementOrderId uint        `json:"settlementOrderId" orm:"settlement_order_id" description:"结算单ID"`
	SettlementType    int         `json:"settlementType"    orm:"settlement_type"     description:"结算方式  1无需结算 2按周期自动结算  3手动申请结算"`
	SettlementCycle   int         `json:"settlementCycle"   orm:"settlement_cycle"    description:"结算周期  1每日结算  2每周结算  3每月结算"`
	SettlementCost    string      `json:"settlementCost"    orm:"settlement_cost"     description:"结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金"`
	OrderAmount       float64     `json:"orderAmount"       orm:"order_amount"        description:"单笔订单金额"`
	CouponAmount      float64     `json:"couponAmount"      orm:"coupon_amount"       description:"优惠券抵扣金额"`
	BalAmount         float64     `json:"balAmount"         orm:"bal_amount"          description:"积分抵扣金额"`
	ActualEndTime     *gtime.Time `json:"actualEndTime"     orm:"actual_end_time"     description:"实际结束日期时间"`
}
