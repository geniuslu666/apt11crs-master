package pms

import (
	"APT/internal/model/input/input_basics"
	"github.com/gogf/gf/v2/frame/g"
)

type DashboardReq struct {
	g.Meta    `path:"/pms/dashboard" method:"post" tags:"ADMIN_PMS" summary:"APP数据看板_展示"`
	Type      string `json:"type"    dc:"类型(base:房态统计  amount:营业额概况  echart:七天入住率)"`
	StartDate string `json:"start_date" dc:"开始日期"`
	EndDate   string `json:"end_date" dc:"结束日期"`
	Puid      string `json:"puid" dc:"物业ID"`
}

type DashboardRes struct {
	Details struct {
		ImportBookings   int `json:"import_bookings" dc:"预抵数"`
		CheckInBookings  int `json:"checkin_bookings" dc:"已抵数"`
		ExportBookings   int `json:"export_bookings" dc:"预离数"`
		CheckOutBookings int `json:"checkout_bookings" dc:"已离数"`
		CheckedBookings  int `json:"checked_bookings" dc:"在住客房数"`
		AllRoom          int `json:"all_room" dc:"客房总数"`
		ToDayStays       int `json:"today_stays" dc:"今日预定"`
		ToDayAppStays    int `json:"today_app_stays" dc:"app预定"`
	}
	AmountStat struct {
		TotalIncome    float64 `json:"totalIncome"           dc:"总收入额"`
		EffectIncome   float64 `json:"effectIncome"           dc:"有效收入"`
		PayCloudIncome float64 `json:"payCloudIncome"      dc:"payCloud支付"`
		PayCloudRefund float64 `json:"payCloudRefund"      dc:"payCloud退款"`
		StripIncome    float64 `json:"stripIncome"      dc:"stripCloud支付"`
		StripRefund    float64 `json:"stripRefund"      dc:"strip退款"`
		BalIncome      float64 `json:"balIncome"      dc:"积分抵扣支付"`
		BalRefund      float64 `json:"balRefund"      dc:"积分抵扣退款"`
		CouponIncome   float64 `json:"couponIncome"      dc:"优惠券抵扣"`
		CouponRefund   float64 `json:"couponRefund"      dc:"优惠券退款"`
		ExchangeRate   float64 `json:"exchangeRate"      dc:"全局积分汇率"`
	} `json:"amountStat" dc:"营业额概况"`
	CheckInTrend []*struct {
		Percent float64 `json:"percent"      dc:"入住率%"`
		Date    string  `json:"date"    description:"日期"`
	} `json:"CheckInTrend" dc:"近7天入住率"`
}

type DashboardAllReq struct {
	g.Meta `path:"/pms/dashboard/all" method:"post" tags:"ADMIN_PMS" summary:"APP数据看板_全量展示"`
	input_basics.BasicDashboardInp
	input_basics.RankingInp
}

type DashboardAllRes struct {
	*input_basics.BasicDashboardModel
	*input_basics.RankingModel
}
