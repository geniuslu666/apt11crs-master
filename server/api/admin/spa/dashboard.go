package spa

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type DashboardReq struct {
	g.Meta             `path:"/spa/dashboard" method:"post" tags:"ADMIN_SPA" summary:"APP数据看板_展示"`
	Type               string `json:"type"    dc:"类型(base:实时概况 todayOrder:今日订单  serviceLine:服务排行  technicianLine:服务人员排行)"`
	ServiceLineType    string `json:"service_line_type" dc:"服务排行类型  num:预定量  amount:预定金额"`
	TechnicianLineType string `json:"technician_line_type" dc:"技师排行类型  num:服务量  amount:服务金额"`
}

type DashboardRes struct {
	Details struct {
		TodayOrderTotalNum          int     `json:"todayOrderTotalNum" dc:"今日预约总数"`
		TodayDdOrderTotalNum        int     `json:"todayDdOrderTotalNum" dc:"今日预约到店总数"`
		TodaySmOrderTotalNum        int     `json:"todaySmOrderTotalNum" dc:"今日预约上门总数"`
		TodayOrderAmount            float64 `json:"todayOrderAmount" dc:"今日营业额"`
		TotalOrderAmount            float64 `json:"totalOrderAmount" dc:"总营业额"`
		TotalOrderTotalNum          int     `json:"totalOrderTotalNum" dc:"全部预约数"`
		EffectiveTotalOrderTotalNum int     `json:"effectiveTotalOrderTotalNum" dc:"有效预约数"`
		CanOrderTechnicianTotal     int     `json:"canOrderTechnicianTotal" dc:"空闲中技师数"`
		WorkingTechnicianTotal      int     `json:"workingTechnicianTotal" dc:"工作中技师数"`
		RestTechnicianTotal         int     `json:"restTechnicianTotal" dc:"休息中技师数"`
	} `json:"details" dc:"实时概况"`
	TodayOrder struct {
		WaitPayOrder               int     `json:"waitPayOrder" dc:"待付款订单"`
		WaitPayOrderAmount         float64 `json:"waitPayOrderAmount" dc:"待付款订单金额"`
		WaitConfirmOrder           int     `json:"waitConfirmOrder" dc:"待确认订单"`
		WaitConfirmOrderAmount     float64 `json:"waitConfirmOrderAmount" dc:"待确认订单金额"`
		WaitServeOrder             int     `json:"waitServeOrder" dc:"待服务订单"`
		WaitServeOrderAmount       float64 `json:"waitServeOrderAmount" dc:"待服务订单金额"`
		ServiceDoingOrder          int     `json:"serviceDoingOrder" dc:"服务中订单"`
		ServiceDoingOrderAmount    float64 `json:"serviceDoingOrderAmount" dc:"服务中订单金额"`
		ServiceCompleteOrder       int     `json:"serviceCompleteOrder" dc:"服务完成订单"`
		ServiceCompleteOrderAmount float64 `json:"serviceCompleteOrderAmount" dc:"服务完成订单金额"`
	} `json:"todayOrder" dc:"今日订单"`
	ServiceLine    []*entity.SpaService    `json:"serviceLine" dc:"实时概况"`
	TechnicianLine []*entity.SpaTechnician `json:"technicianLine" dc:"实时概况"`
}
