package basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_basics"
	"APT/internal/model/input/input_hotel"

	"github.com/gogf/gf/v2/frame/g"
)

type ImmovablesPathJsonData struct {
	Weapp string `json:"weapp"`
	App   string `json:"app"`
}

type HomeDateLoadReq struct {
	g.Meta `path:"/home/info" method:"post" tags:"APP_BASICS" summary:"[首页]加载信息"`
}

type HomeDateLoadRes struct {
	Banner                 []*input_basics.PmsBannerListModel      `json:"banner" dc:"banner数据"`
	Property               []*input_hotel.PmsPropertyAppIndexModel `json:"property" dc:"物业数据"`
	NavList                []*input_basics.PmsIndexNavApiListModel `json:"navList" dc:"导航栏"`
	ImmovablesCover        string                                  `json:"immovablesCover" dc:"不动产封面"`
	ImmovablesPath         string                                  `json:"immovablesPath" dc:"不动产跳转地址"`
	ImmovablesWxPath       string                                  `json:"immovablesWxPath" dc:"不动产微信跳转地址"`
	ImmovablesAppPath      string                                  `json:"immovablesAppPath" dc:"不动产APP跳转地址"`
	ImmovablesChain        string                                  `json:"immovablesChain" dc:"不动产内外联  IN 内链  OUT 外联"`
	AlertMessage           []*HomeDateAlertMessage                 `json:"alertMessage" dc:"首页弹窗信息"`
	MinDaysNotice          float64                                 `json:"min_days_notice" dc:"短租模式最小预定区间"`
	TripCardList           []*TripCardListItem                     `json:"tripCardList" dc:"行程卡列表"`
	UnReadMessageList      []*input_basics.MessageAppListModel     `json:"unReadMessageList" dc:"未读消息列表"`
	CarouselImageList      []*CarouselImageListItem                `json:"carouselImageList" dc:"首页广告位列表"`
	CarouselScrollInterval int                                     `json:"carouselScrollInterval" dc:"首页广告位滚动间隔(秒)"`
}

type CarouselImageListItem struct {
	Cover        string `json:"cover" dc:"封面"`
	Path         string `json:"path" dc:"跳转地址"`
	WxPath       string `json:"wxPath" dc:"微信跳转地址"`
	AppPath      string `json:"appPath" dc:"APP跳转地址"`
	Chain        string `json:"chain" dc:"内外联  IN 内链  OUT 外联"`
	LinkOpenType int    `json:"linkOpenType" dc:"外链打开方式  1-内部webview  2-外部浏览器"`
}

type TripCardListItem struct {
	OrderId         int                 `json:"orderId" dc:"订单ID"`
	OrderSn         string              `json:"orderSn" dc:"订单号"`
	OrderScene      string              `json:"orderScene" dc:"订单场景(hotel：酒店，food：餐厅，car：接送机，spa：按摩)"`
	BookingDateTime string              `json:"bookingDateTime" dc:"预订时间"`
	HotelOrderInfo  *HotelOrderInfoItem `json:"hotelOrderInfo" dc:"酒店订单信息"`
	FoodOrderInfo   *FoodOrderInfoItem  `json:"foodOrderInfo" dc:"餐厅订单信息"`
	CarOrderInfo    *CarOrderInfoItem   `json:"carOrderInfo" dc:"接送机订单信息"`
	SpaOrderInfo    *SpaOrderInfoItem   `json:"spaOrderInfo" dc:"按摩订单信息"`
}

type HotelOrderInfoItem struct {
	CheckinDate  string `json:"checkinDate"     dc:"入住日期"`
	CheckoutDate string `json:"checkoutDate"    dc:"退房日期"`
	CheckinTime  string `json:"checkinTime"     dc:"入住时间，24小时格式"`
	CheckoutTime string `json:"checkoutTime"    dc:"退房时间，24小时格式"`
	Nights       int    `json:"nights"            dc:"入住晚数"`
}
type FoodOrderInfoItem struct {
	BookDate       string `json:"bookDate"                 dc:"预定日期"`
	BookTime       string `json:"bookTime"                 dc:"预定时间"`
	RestaurantName string `json:"restaurantName"                 dc:"餐厅名称"`
	GoodsName      string `json:"goodsName"                 dc:"套餐名称"`
}

type CarOrderInfoItem struct {
	ServiceType      string `json:"serviceType"             dc:"服务类型（PICKUP：接机，DELIVERY：送机）"`
	BookDate         string `json:"bookDate"                 dc:"预定日期"`
	BookTime         string `json:"bookTime"                 dc:"预定时间"`
	StartAddressName string `json:"startAddressName"          dc:"出发地名称"`
	EndAddressName   string `json:"endAddressName"            dc:"目的地名称"`
	CarTypeName      string `json:"carTypeName"               dc:"车型名称"`
}

type SpaOrderInfoItem struct {
	ServiceType int    `json:"serviceType"             dc:"服务类型（1到店  2上门）"`
	BookDate    string `json:"bookDate"                 dc:"预定日期"`
	BookTime    string `json:"bookTime"                 dc:"预定时间"`
	ServiceName string `json:"serviceName"                 dc:"服务名称"`
	GoodsName   string `json:"goodsName"                 dc:"项目名称"`
}

type HomeDateAlertMessage struct {
	Id            int    `json:"id"`
	NotifyType    string `json:"notify_type"`
	NotifyContent string `json:"notify_content"`
}

type HomeDateConfigReq struct {
	g.Meta `path:"/home/config" method:"post" tags:"APP_BASICS" summary:"[首页]配置信息"`
}
type HomeDateConfigRes struct {
	ConfigList []*entity.PmsAppconfig `json:"config"`
}

type ReadAlertMessageReq struct {
	g.Meta `path:"/home/read_alert_message" method:"post" tags:"APP_BASICS" summary:"[首页]已读弹窗信息"`
	Id     int `json:"id" v:"required#id_unknown" dc:"id"`
}

type ReadAlertMessageRes struct{}

type HomeDeviceInfoReq struct {
	g.Meta `path:"/home/deviceInfo" method:"post" tags:"APP_BASICS" summary:"[首页]设备信息"`
}

type HomeDeviceInfoRes struct {
	DeviceSn int `json:"deviceSn" dc:"设备号"`
}

type ConfigGetReq struct {
	g.Meta `path:"/home/getConfig" method:"post" tags:"APP_BASICS" summary:"[首页]获取指定分组的配置"`
	input_basics.GetConfigInp
}

type ConfigGetRes struct {
	*input_basics.GetConfigModel
}
