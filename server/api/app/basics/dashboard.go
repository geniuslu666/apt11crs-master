package basics

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type DashboardReq struct {
	g.Meta `path:"/dashboard/view" method:"post" tags:"APP_BASICS" summary:"APP数据看板_展示"`
}

type DashboardRes struct {
	Permissions   []string `json:"permissions" dc:"数据看板权限列表: member-会员, hotel-民宿, spa-按摩, car-车队"`
	PrivateDetail struct {
		MemberTotal            int `json:"memberTotal"           dc:"会员总数"`
		TodayMemberAdd         int `json:"todayMemberAdd"           dc:"今日增加会员"`
		YesterdayMemberAdd     int `json:"yesterdayMemberAdd"           dc:"昨日增加会员"`
		TotalHotelOrderNum     int `json:"totalHotelOrderNum"         dc:"民宿下单总数"`
		TodayHotelOrderNum     int `json:"todayHotelOrderNum"         dc:"今日民宿下单数"`
		YesterdayHotelOrderNum int `json:"yesterdayHotelOrderNum"         dc:"昨日民宿下单数"`
	} `json:"privateDetail" dc:"私域数据"`
	RegSourceDetail struct {
		IOSNum     int `json:"iosNum"           dc:"IOS来源"`
		AndroidNum int `json:"androidNum"           dc:"android来源"`
		WeixinNum  int `json:"weixinNum"           dc:"Weixin来源"`
		H5Num      int `json:"h5Num"           dc:"H5来源"`
	} `json:"regSourceDetail" dc:"注册来源"`
	TotalDetail struct {
		TotalHotelOrderNum    int `json:"totalHotelOrderNum"         dc:"民宿下单总数"`
		TodayHotelOrderNum    int `json:"todayHotelOrderNum"         dc:"今日民宿下单总数"`
		TotalAPTHotelOrderNum int `json:"totalAPTHotelOrderNum"         dc:"APT民宿下单总数"`
		TodayAPTHotelOrderNum int `json:"todayAPTHotelOrderNum"         dc:"今日APT民宿下单总数"`
		TotalOTAHotelOrderNum int `json:"totalOTAHotelOrderNum"         dc:"OTA民宿下单总数"`
		TodayOTAHotelOrderNum int `json:"todayOTAHotelOrderNum"         dc:"今日OTA民宿下单总数"`
	} `json:"totalDetail" dc:"全量数据"`
	CarData struct {
		TotalCarOrderNum      int                `json:"totalCarOrderNum"         dc:"车队总订单数"`
		TotalOrderMoney       float64            `json:"totalOrderMoney"          dc:"累计营业额"`
		TodayOrderMoney       float64            `json:"todayOrderMoney"          dc:"今日营业额"`
		TodayCarOrderNum      int                `json:"todayCarBookOrderNum"     dc:"今日总预约数"`
		TodayPickUpOrderNum   int                `json:"todayPickUpOrderNum"      dc:"今日接机总数"`
		TodayDeliveryOrderNum int                `json:"todayDeliveryOrderNum"    dc:"今日送机总数"`
		RefuseOrderList       []*RefuseOrderItem `json:"refuseOrderList" dc:"拒单列表"`
	} `json:"carData" dc:"车队数据"`
	SpaData struct {
		TotalSpaOrderNum        int                `json:"totalSpaOrderNum"         dc:"总订单数"`
		TodaySpaOrderNum        int                `json:"todaySpaBookOrderNum"     dc:"今日总预约数"`
		TotalOrderMoney         float64            `json:"totalOrderMoney"          dc:"累计营业额"`
		TodayOrderMoney         float64            `json:"todayOrderMoney"          dc:"今日营业额"`
		TodayDoorToDoorOrderNum int                `json:"todayDoorToDoorOrderNum"  dc:"今日上门总数"`
		TodayInStoreOrderNum    int                `json:"todayInStoreOrderNum"     dc:"今日到店总数"`
		RefuseOrderList         []*RefuseOrderItem `json:"refuseOrderList" dc:"拒单列表"`
	} `json:"spaData" dc:"按摩数据"`
	UpdateTime *gtime.Time `json:"updateTime"     dc:"更新时间"`
}

type RefuseOrderItem struct {
	Id                  int    `json:"id"               dc:"订单ID"`
	OrderSn             string `json:"orderSn"               dc:"订单号"`
	ServiceType         string `json:"serviceType"           dc:"服务类型"`
	BookStartTime       string `json:"bookStartTime"         dc:"预定时间"`
	ConfirmRefuseReason string `json:"confirmRefuseReason"   dc:"确认拒绝原因"`
}
