package car

import (
	"APT/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type DashboardReq struct {
	g.Meta          `path:"/car/dashboard" method:"post" tags:"ADMIN_CAR" summary:"APP数据看板_展示"`
	Type            string `json:"type"    dc:"类型(base:实时概况 todayOrder:今日订单  serviceLine:服务排行  driverLine:司机排行)"`
	ServiceLineType string `json:"service_line_type" dc:"服务排行类型  num:预定量  amount:预定金额"`
	DriverLineType  string `json:"driver_line_type" dc:"司机排行类型  num:服务量  amount:服务金额"`
}

type DashboardRes struct {
	Details struct {
		TodayOrderTotalNum             int     `json:"todayOrderTotalNum" dc:"今日预约总数"`
		TodayPickUpOrderTotalNum       int     `json:"todayPickUpOrderTotalNum" dc:"今日预约接机总数"`
		TodayDeliveryOrderTotalNum     int     `json:"todayDeliveryOrderTotalNum" dc:"今日预约接机送机总数"`
		TodayCharteredCarOrderTotalNum int     `json:"todayCharteredCarOrderTotalNum" dc:"今日预约包车送机总数"`
		TodayOrderAmount               float64 `json:"todayOrderAmount" dc:"今日营业额"`
		TodayOrderRefund               float64 `json:"todayOrderRefund" dc:"今日退款额"`
		TotalOrderTotalNum             int     `json:"totalOrderTotalNum" dc:"全部累计预约数"`
		EffectiveOrderTotalNum         int     `json:"effectiveOrderTotalNum" dc:"有效累计预约数"`
		TotalPickUpOrderTotalNum       int     `json:"totalPickUpOrderTotalNum" dc:"全部预约接机总数"`
		TotalDeliveryOrderTotalNum     int     `json:"totalDeliveryOrderTotalNum" dc:"全部预约送机总数"`
		TotalCharteredCarOrderTotalNum int     `json:"totalCharteredCarOrderTotalNum" dc:"全部预约包车总数"`
		TotalOrderAmount               float64 `json:"totalOrderAmount" dc:"总营业额"`
		TotalOrderRefund               float64 `json:"totalOrderRefund" dc:"总退款额"`
		TotalDriverNum                 int     `json:"totalDriverNum" dc:"司机总数"`
		OnDriverTotal                  int     `json:"onDriverTotal" dc:"司机正常总数"`
		OffDriverTotal                 int     `json:"offDriverTotal" dc:"停用司机数"`
		RestDriverTotal                int     `json:"restDriverTotal" dc:"休息中司机数"`
		TotalCarNum                    int     `json:"totalCarNum" dc:"车辆总数"`
		OnCarTotal                     int     `json:"onCarTotal" dc:"启用车辆数"`
		OffCarTotal                    int     `json:"offCarTotal" dc:"停用车辆数"`
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
	ServiceLine []*CarServiceLineModel `json:"serviceLine" dc:"服务排行"`
	DriverLine  []*entity.CarDriver    `json:"driverLine" dc:"司机排行"`
}

type CarServiceLineModel struct {
	ServiceType      string  `json:"serviceType"      dc:"服务类型"`
	TotalOrderNum    int     `json:"totalOrderNum"    dc:"预约单总数量（包含退款）"`
	TotalOrderAmount float64 `json:"totalOrderAmount" dc:"预约单总金额（包含退款）"`
}
