package travel

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DashboardReq 一日游概况
type DashboardReq struct {
	g.Meta `path:"/travel/dashboard" method:"get" tags:"ADMIN_TRAVEL" summary:"一日游概况"`
	Type   string `json:"type"      dc:"类型(base:实时概况 todayOrder:今日订单)"`
}

type DashboardRes struct {
	Details struct {
		TodayOrderTotalNum     int     `json:"todayOrderTotalNum"    dc:"今日预约总数"`
		TodayEffectiveOrderNum int     `json:"todayEffectiveOrderNum" dc:"今日有效预约数（非取消/退款）"`
		TodayOrderAmount       float64 `json:"todayOrderAmount"      dc:"今日营业额"`
		TotalOrderAmount       float64 `json:"totalOrderAmount"      dc:"累计营业额"`
		TotalOrderNum          int     `json:"totalOrderNum"         dc:"累计预约数"`
		EffectiveOrderNum      int     `json:"effectiveOrderNum"     dc:"累计有效预约数"`
		TodayWaitVerifyNum     int     `json:"todayWaitVerifyNum"    dc:"今日待核销数"`
		TodayVerifiedNum       int     `json:"todayVerifiedNum"      dc:"今日已核销数"`
	} `json:"details" dc:"实时概况"`
	TodayOrder struct {
		WaitPayOrder    int `json:"waitPayOrder"    dc:"待支付订单数"`
		WaitVerifyOrder int `json:"waitVerifyOrder" dc:"待核销订单数"`
		DoneOrder       int `json:"doneOrder"       dc:"已完成订单数"`
		RefundOrder     int `json:"refundOrder"     dc:"已退款订单数"`
	} `json:"todayOrder" dc:"今日订单"`
}
