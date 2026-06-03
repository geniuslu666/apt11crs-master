package input_basics

import "github.com/gogf/gf/v2/frame/g"

// BasicDashboardInp 基础统计
type BasicDashboardInp struct {
	Scene string `json:"scene" dc:"场景值"`
}

type BasicDashboardModel struct {
	MemberStat struct {
		TodayRegMemberNum      int     `json:"todayRegMemberNum"        dc:"今日新增会员数"`
		TodayRegGrewPer        float64 `json:"todayMemberNumGrewPer"    dc:"会员增长同比"`
		TotalMemberNum         int     `json:"totalMemberNum"           dc:"累计会员数"`
		TotalMemberGrewPer     float64 `json:"totalMemberGrewPer"    dc:"累计会员数增长同比"`
		TodayIncBal            float64 `json:"todayIncBal"           dc:"今日新增积分"`
		TodayIncBalGrewPer     float64 `json:"todayIncBalGrewPer"       dc:"今日积分较昨日增长同比"`
		TodayConsumeBal        float64 `json:"todayConsumeBal"          dc:"今日消耗积分"`
		TodayConsumeBalGrewPer float64 `json:"todayConsumeBalGrewPer"   dc:"今日消耗积分较昨日增长同比"`
	} `json:"memberStat" dc:"会员统计"`
	OrderStat struct {
		AllStat struct {
			TodayOrderNum           int     `json:"todayOrderNum"           dc:"今日订单"`
			TodayOrderNumGrewPer    float64 `json:"todayOrderNumGrewPer"            dc:"今日订单增长同比"`
			TotalOrderNum           int     `json:"totalOrderNum"           dc:"总订单数"`
			TotalOrderNumGrewPer    float64 `json:"totalOrderNumGrewPer"       dc:"总订单数增长同比"`
			TodayRefundNum          int     `json:"todayRefundNum"          dc:"今日退款单数"`
			TodayRefundGrewPer      float64 `json:"todayRefundGrewPer"      dc:"今日退款订单数增长同比"`
			TotalRefundNum          int     `json:"totalRefundNum"          dc:"全量退款单数"`
			TotalRefundGrewPer      float64 `json:"totalRefundGrewPer"      dc:"全量退款订单数增长同比"`
			TotalPropertyNum        int     `json:"totalPropertyNum"          dc:"全部物业数"`
			OnlinePropertyNum       int     `json:"onlinePropertyNum"          dc:"在线物业数"`
			TodayOrderMoney         float64 `json:"todayOrderMoney"          dc:"今日订单总额"`
			TodayOrderMoneyGrewPer  float64 `json:"todayOrderMoneyGrewPer"      dc:"今日订单总额增长同比"`
			TotalOrderMoney         float64 `json:"totalOrderMoney"          dc:"今日订单总额"`
			TotalOrderMoneyGrewPer  float64 `json:"totalOrderMoneyGrewPer"      dc:"今日订单总额增长同比"`
			TodayRefundMoney        float64 `json:"todayRefundMoney"          dc:"今日退款总额"`
			TodayRefundMoneyGrewPer float64 `json:"todayRefundMoneyGrewPer"      dc:"今日退款总额增长同比"`
			TotalRefundMoney        float64 `json:"totalRefundMoney"          dc:"今日退款总额"`
			TotalRefundMoneyGrewPer float64 `json:"totalRefundMoneyGrewPer"      dc:"今日退款总额增长同比"`
			TotalRoomTypeNum        int     `json:"totalRoomTypeNum"          dc:"全部房型数"`
			OnlineRoomTypeNum       int     `json:"onlineRoomTypeNum"          dc:"在线房型数"`
		} `json:"allStat" dc:"总订单概况"`
		HotelStat struct {
			TodayOrderNum           int     `json:"todayOrderNum"           dc:"今日订单"`
			TodayOrderNumGrewPer    float64 `json:"todayOrderNumGrewPer"            dc:"今日订单增长同比"`
			TotalOrderNum           int     `json:"totalOrderNum"           dc:"总订单数"`
			TotalOrderNumGrewPer    float64 `json:"totalOrderNumGrewPer"       dc:"总订单数增长同比"`
			TodayRefundNum          int     `json:"todayRefundNum"          dc:"今日退款单数"`
			TodayRefundGrewPer      float64 `json:"todayRefundGrewPer"      dc:"今日退款订单数增长同比"`
			TotalRefundNum          int     `json:"totalRefundNum"          dc:"全量退款单数"`
			TotalRefundGrewPer      float64 `json:"totalRefundGrewPer"      dc:"全量退款订单数增长同比"`
			TotalPropertyNum        int     `json:"totalPropertyNum"          dc:"全部物业数"`
			OnlinePropertyNum       int     `json:"onlinePropertyNum"          dc:"在线物业数"`
			TodayOrderMoney         float64 `json:"todayOrderMoney"          dc:"今日订单总额"`
			TodayOrderMoneyGrewPer  float64 `json:"todayOrderMoneyGrewPer"      dc:"今日订单总额增长同比"`
			TotalOrderMoney         float64 `json:"totalOrderMoney"          dc:"今日订单总额"`
			TotalOrderMoneyGrewPer  float64 `json:"totalOrderMoneyGrewPer"      dc:"今日订单总额增长同比"`
			TodayRefundMoney        float64 `json:"todayRefundMoney"          dc:"今日退款总额"`
			TodayRefundMoneyGrewPer float64 `json:"todayRefundMoneyGrewPer"      dc:"今日退款总额增长同比"`
			TotalRefundMoney        float64 `json:"totalRefundMoney"          dc:"今日退款总额"`
			TotalRefundMoneyGrewPer float64 `json:"totalRefundMoneyGrewPer"      dc:"今日退款总额增长同比"`
			TotalRoomTypeNum        int     `json:"totalRoomTypeNum"          dc:"全部房型数"`
			OnlineRoomTypeNum       int     `json:"onlineRoomTypeNum"          dc:"在线房型数"`
		} `json:"hotelStat" dc:"酒店订单概况"`
	} `json:"orderStat" dc:"订单概况"`
}

type RankingInp struct {
}

type RankingModel struct {
	NewOrderNumberList []*RankingOrderItem `json:"newOrderNumberList" dc:"30订单量趋势列表"`
	OrderNumberList    []*RankingOrderItem `json:"orderNumberList" dc:"7天订单量趋势列表"`
	PropertyList       []*PropertyItem     `json:"propertyList" dc:"物业预订量列表"`
	OTAChannelList     []*OTAChannelItem   `json:"oTAChannelList" dc:"渠道预订量列表"`
	NationalityList    []*NationalityItem  `json:"nationalityList" dc:"国家预订量列表"`
}

type RankingOrderItem struct {
	OrderNumber int64  `json:"orderNumber"`
	OrderDate   string `json:"orderDate"`
}

type PropertyItem struct {
	Puid           string `json:"puid" dc:"物业ID"`
	PropertyDetail *struct {
		g.Meta `orm:"table:hg_pms_property"`
		Uid    string `json:"puid" dc:"物业ID"`
		Name   string `json:"name" dc:"物业名称"`
	} `json:"propertyDetail" orm:"with:uid=puid" dc:"物业详情"`
	Rate        float64 `json:"rate" dc:"物业预订量占比"`
	OrderNumber int64   `json:"number" dc:"物业预订量占比"`
}

type OTAChannelItem struct {
	Name        string  `json:"name" dc:"OTA渠道名称"`
	Rate        float64 `json:"rate" dc:"OTA渠道预订量占比"`
	OrderNumber int64   `json:"number" dc:"OTA渠道预订量"`
}
type NationalityItem struct {
	Name        string  `json:"name" dc:"国家名称"`
	Rate        float64 `json:"rate" dc:"国家预订量占比"`
	OrderNumber int64   `json:"number" dc:"国家预订量"`
}
