package aladdinApi

import (
	"APT/internal/library/contexts"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

var (
	//baseUrl       = "https://api.innnair.com"         // 域名
	baseUrl             string                                    // 域名
	offerInfo           = "/openApi/busOrder/offerInfo"           // 查询报价
	bookingBus          = "/openApi/busOrder/bookingBus"          // 巴士预约生成单
	conformOrder        = "/openApi/busOrder/conformOrder"        // 巴士预约确认单
	cancelOrder         = "/openApi/busOrder/cancelOrder"         // 巴士预约取消单
	userOrderInfo       = "/openApi/busOrder/userOrderInfo"       // 订单详情
	bookingLuggage      = "/openApi/busOrder/bookingLuggage"      // 行李生单
	conformLuggageOrder = "/openApi/busOrder/conformLuggageOrder" // 行李订单确认
	cancelLuggageOrder  = "/openApi/busOrder/cancelLuggageOrder"  // 行李订单取消
	Logger              = g.Log().Path("logs/SDK/ALADDIN_API")
)

type AladdinClient struct {
	Cid         string `json:"cid" dc:"接入应用类型 h5,web,android,ios"`
	AppOs       string `json:"app_os" dc:"商户id 接入方唯一标识"`
	Language    string `json:"language" dc:"日语-ja，中文-zh，英文-en，韩语-ko （默认响应为日语-ja）"`
	GHttpClient *gclient.Client
}

type OfferInfoParams struct {
	Type                 int    `json:"type" dc:"接送机类型:1接机2送机"`
	ArrivalAirportCode   string `json:"arrivalAirportCode,omitempty" dc:"到达机场三字码"`
	DepartureAirportCode string `json:"departureAirportCode,omitempty" dc:"出发机场三字码"`
	DateTime             string `json:"dateTime" dc:"yyyy-MM-dd HH:mm:ss"`
}

type OfferInfoResponse struct {
	Status int                   `json:"status"`
	Msg    string                `json:"msg"`
	Data   OfferInfoResponseItem `json:"data"`
}

type OfferInfoResponseItem struct {
	Session      string  `json:"session"`
	ID           *string `json:"id"`
	Type         int     `json:"type"`
	Currency     string  `json:"currency"`
	Price        int     `json:"price"`
	LuggagePrice int     `json:"luggagePrice"`
	Baggage      string  `json:"baggage"`
	Waiting      int     `json:"waiting"`
	MaxWaiting   int     `json:"maxWaiting"`
	CancelRule   string  `json:"cancelRule"`
	Service      string  `json:"Service"`
}

type BookingBusParams struct {
	Currency             string `json:"currency"`
	Type                 int    `json:"type"`
	Session              string `json:"session"`
	DepartureAirportCode string `json:"departureAirportCode"`
	DepartureTerminal    string `json:"departureTerminal"`
	ArrivalAirportCode   string `json:"arrivalAirportCode"`
	ArrivingTerminal     string `json:"arrivingTerminal"`
	Contacts             string `json:"contacts"`
	CountryCode          string `json:"countryCode"`
	ContactsPhone        string `json:"contactsPhone"`
	ContactsEmail        string `json:"contactsEmail"`
	DateTime             string `json:"dateTime"`
	PersonPrice          int    `json:"personPrice"`
	PersonNum            int    `json:"personNum"`
	OrderPrice           int    `json:"orderPrice"`
	Hotel                string `json:"hotel"`
	Address              string `json:"address"`
}

type BookingBusResponse struct {
	Status int                    `json:"status"`
	Msg    string                 `json:"msg"`
	Data   BookingBusResponseItem `json:"data"`
}

type BookingBusResponseItem struct {
	OrderId      string `json:"orderId"`
	OrderNo      string `json:"orderNo"`
	Contacts     string `json:"contacts"`
	OrderPrice   int    `json:"orderPrice"`
	CustomerInfo string `json:"customerInfo"`
}

type ConformOrderParams struct {
	Type       int    `json:"type"`
	Currency   string `json:"currency"`
	OrderId    string `json:"orderId"`
	OrderPrice int    `json:"orderPrice"`
}

type ConformOrderResponse struct {
	Status int                      `json:"status"`
	Msg    string                   `json:"msg"`
	Data   ConformOrderResponseItem `json:"data"`
}

type ConformOrderResponseItem struct {
	OrderId     string `json:"orderId"`
	OrderNo     string `json:"orderNo"`
	OrderQrCode string `json:"orderQrCode"`
	ExpireTime  string `json:"expireTime"`
}

type CancelOrderParams struct {
	OrderId string `json:"orderId" dc:"订单ID"`
}
type CancelOrderResponse struct {
	Status int                     `json:"status"`
	Msg    string                  `json:"msg"`
	Data   CancelOrderResponseItem `json:"data"`
}

type CancelOrderResponseItem struct {
	OrderId        string `json:"orderId"`
	LuggageOrderId string `json:"luggageOrderId"`
	LuggageOrderNo string `json:"luggageOrderNo"`
	OrderNo        string `json:"orderNo"`
}

type UserOrderInfoParams struct {
	OrderId string `json:"orderId" dc:"订单ID"`
}

// UserOrderInfoResponse 订单详情
type UserOrderInfoResponse struct {
	Status int                       `json:"status"`
	Msg    string                    `json:"msg"`
	Data   UserOrderInfoResponseItem `json:"data"`
}

type UserOrderInfoResponseItem struct {
	OrderId              string `json:"orderId"`
	OrderNo              string `json:"orderNo"`
	Type                 int    `json:"type"`
	FlightNumber         string `json:"flightNumber"`
	DepartureAirportCode string `json:"departureAirportCode"`
	DepartureTerminal    string `json:"departureTerminal"`
	ArrivalAirportCode   string `json:"arrivalAirportCode"`
	ArrivingTerminal     string `json:"arrivingTerminal"`
	Hotel                string `json:"hotel"`
	Address              string `json:"address"`
	DateTime             string `json:"dateTime"`
	BaggagePassColor     string `json:"baggagePassColor"`
	Contacts             string `json:"contacts"`
	CountryCode          string `json:"countryCode"`
	ContactsPhone        string `json:"contactsPhone"`
	ContactEmail         string `json:"contactEmail"`
	WechatNo             string `json:"wechatNo"`
	LineNo               string `json:"lineNo"`
	WhatsAppNo           string `json:"whatsAppNo"`
	Price                int    `json:"price"`
	PersonNum            int    `json:"personNum"`
	ChildSeatNum         int    `json:"childSeatNum"`
	ChildSeatCosts       int    `json:"childSeatCosts"`
	OrderPrice           int    `json:"orderPrice"`
	TradeStatus          int    `json:"tradeStatus"`
	RidingStatus         int    `json:"ridingStatus"`
	ReviewFlag           bool   `json:"reviewFlag"`
	OrderQrCode          string `json:"orderQrCode"`
	ExpireTime           int64  `json:"expireTime"`
	ReviewContentBO      []struct {
		ServiceQuality  int    `json:"serviceQuality"`
		PriceDiscount   int    `json:"priceDiscount"`
		ComfortableRide int    `json:"comfortableRide"`
		Content         string `json:"content"`
		ImgUrl          string `json:"imgUrl"`
	} `json:"reviewContentBO"`
	LuggageOrders []struct {
		OrderId        int64  `json:"orderId"`
		LuggageOrderId string `json:"luggageOrderId"`
		LuggageOrderNo string `json:"luggageOrderNo"`
		OrderPirce     int    `json:"orderPirce"`
		FlowStatus     int    `json:"flowStatus"`
	} `json:"luggageOrders"`
}

type BookingLuggageParams struct {
	Session      string `json:"session"`
	Currency     string `json:"currency"`
	BusOrderId   string `json:"busOrderId"`
	LuggagePrice int    `json:"luggagePrice"`
	LuggageNum   int    `json:"luggageNum"`
	OrderPrice   int    `json:"orderPrice"`
}

type BookingLuggageResponse struct {
	Status int                        `json:"status"`
	Msg    string                     `json:"msg"`
	Data   BookingLuggageResponseItem `json:"data"`
}

type BookingLuggageResponseItem struct {
	BusOrderId     string `json:"busOrderId"`
	LuggageOrderId string `json:"luggageOrderId"`
	LuggageOrderNo string `json:"luggageOrderNo"`
	OrderPirce     int    `json:"orderPirce"`
	FlowStatus     int    `json:"flowStatus"`
}

type ConformLuggageOrderParams struct {
	Currency       string `json:"currency"`
	LuggageOrderId string `json:"luggageOrderId"`
	OrderPrice     int    `json:"orderPrice"`
}

type ConformLuggageOrderResponse struct {
	Status int                             `json:"status"`
	Msg    string                          `json:"msg"`
	Data   ConformLuggageOrderResponseItem `json:"data"`
}

type ConformLuggageOrderResponseItem struct {
	Currency       string `json:"currency"`
	LuggageOrderId string `json:"luggageOrderId"`
	OrderPrice     int    `json:"orderPrice"`
}

type CancelLuggageOrderParams struct {
	LuggageOrderId string `json:"luggageOrderId"`
}

type CancelLuggageOrderResponse struct {
	Status int                            `json:"status"`
	Msg    string                         `json:"msg"`
	Data   CancelLuggageOrderResponseItem `json:"data"`
}

type CancelLuggageOrderResponseItem struct {
	LuggageOrderId string `json:"luggageOrderId"`
	LuggageOrderNo string `json:"luggageOrderNo"`
}

// VerifyOrderParams 核销回调接口入参
type VerifyOrderParams struct {
	OrderId string `json:"orderId"`
	OrderNo string `json:"orderNo"`
	Status  string `json:"status"`
}

func NewClient(ctx context.Context) *AladdinClient {
	// language 日语-ja，中文-zh，英文-en，韩语-ko （默认响应为日语-ja）
	var (
		language string
	)
	baseUrl = g.Cfg().MustGet(ctx, "aladdin.host").String()
	switch contexts.GetLanguage(ctx) {
	case "zh_CN":
		language = "zh"
	case "zh":
		language = "zh"
	case "ko":
		language = "ko"
	case "en":
		language = "en"
	default:
		language = "ja"
	}

	return &AladdinClient{
		Cid:         "HOTEL+ANDCLAN",
		AppOs:       "ios",
		Language:    language,
		GHttpClient: g.Client(),
	}
}

func (c *AladdinClient) DoRequest(ctx context.Context, method string, url string, paramsReq interface{}) (response *gclient.Response, err error) {
	c.GHttpClient.SetHeader("Content-Type", "application/json")
	c.GHttpClient.SetHeader("appOs", c.AppOs)
	c.GHttpClient.SetHeader("cid", c.Cid)
	c.GHttpClient.SetHeader("language", c.Language)
	if response, err = c.GHttpClient.DoRequest(ctx, method, baseUrl+url, paramsReq); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())
	return
}
