package cabinetApi

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"APT/internal/model"
)

var (
	baseUrl       string                      // 域名
	cabinetInfo   = "/mch/cabinet"            // 查询储物柜
	createOrder   = "/mch/order/create"       // 创建订单
	orderQuery    = "/mch/order/query"        // 订单查询
	payOvertime   = "/mch/order/pay-overtime" // 支付超时费
	cabinetList   = "/mch/cabinet/list"       // 支付超时费
	orderComplete = "/mch/order/complete"     // 订单完成
	Logger        = g.Log().Path("logs/SDK/CABINET_API")
)

type CabinetClient struct {
	Appid       string `json:"appid" dc:"应用ID"`
	Appkey      string `json:"appkey" dc:"应用秘钥"`
	ApiSecret   string `json:"apiSecret" dc:"API秘钥"`
	GHttpClient *gclient.Client
}

type CabinetInfoParams struct {
	CabinetId int `json:"cabinet_id" dc:"储物柜ID"`
}

type CabinetInfoResponse struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data CabinetInfoResponseItem `json:"data"`
}

type CabinetInfoResponseItem struct {
	ID                int         `json:"id" dc:"储物柜ID"`
	Name              string      `json:"name" dc:"储物柜名称"`
	NameZh            string      `json:"nameZh" dc:"储物柜名称_简体中文"`
	NameEn            string      `json:"nameEn" dc:"储物柜名称_英文"`
	NameJa            string      `json:"nameJa" dc:"储物柜名称_日语"`
	NameKo            string      `json:"nameKo" dc:"储物柜名称_韩语"`
	NameTw            string      `json:"nameTw" dc:"储物柜名称_繁体中文"`
	CityId            int         `json:"cityId" dc:"所属城市ID"`
	CityName          string      `json:"cityName" dc:"所属城市名称"`
	CityNameZh        string      `json:"cityNameZh" dc:"所属城市名称_简体中文"`
	CityNameEn        string      `json:"cityNameEn" dc:"所属城市名称_英文"`
	CityNameJa        string      `json:"cityNameJa" dc:"所属城市名称_日语"`
	CityNameKo        string      `json:"cityNameKo" dc:"所属城市名称_韩语"`
	CityNameTw        string      `json:"cityNameTw" dc:"所属城市名称_繁体中文"`
	MchId             int         `json:"mchId" dc:"所属运营商ID"`
	MchName           string      `json:"mchName" dc:"所属运营商名字"`
	MchNameZh         string      `json:"mchNameZh" dc:"所属运营商名字_简体中文"`
	MchNameEn         string      `json:"mchNameEn" dc:"所属运营商名字_英文"`
	MchNameJa         string      `json:"mchNameJa" dc:"所属运营商名字_日语"`
	MchNameKo         string      `json:"mchNameKo" dc:"所属运营商名字_韩语"`
	MchNameTw         string      `json:"mchNameTw" dc:"所属运营商名字_繁体中文"`
	MchBranchId       int         `json:"mchBranchId" dc:"所属网点ID"`
	MchBranchName     string      `json:"mchBranchName" dc:"所属网点名字"`
	MchBranchNameZh   string      `json:"mchBranchNameZh" dc:"所属网点名称_简体中文"`
	MchBranchNameEn   string      `json:"mchBranchNameEn" dc:"所属网点名称_英文"`
	MchBranchNameJa   string      `json:"mchBranchNameJa" dc:"所属网点名称_日语"`
	MchBranchNameKo   string      `json:"mchBranchNameKo" dc:"所属网点名称_韩语"`
	MchBranchNameTw   string      `json:"mchBranchNameTw" dc:"所属网点名称_繁体中文"`
	MchBranchLat      string      `json:"mchBranchLat" dc:"所属网点lat"`
	MchBranchLgt      string      `json:"mchBranchLgt" dc:"所属网点lgt"`
	BoxTypeA          BoxTypeItem `json:"boxTypeA" dc:"格口类型A信息"`
	BoxTypeB          BoxTypeItem `json:"boxTypeB" dc:"格口类型B信息"`
	BoxTypeC          BoxTypeItem `json:"boxTypeC" dc:"格口类型C信息"`
	MinHours          int         `json:"minHours" dc:"最小租赁时间"`
	OrderFirstFeeRate float64     `json:"orderFirstFeeRate" dc:"首次下单优惠比例"`
}

type BoxTypeItem struct {
	ID     int    `json:"id" dc:"格口类型ID"`
	Name   string `json:"name" dc:"格口类型名称"`
	NameZh string `json:"nameZh" dc:"格口类型名称_简体中文"`
	NameEn string `json:"nameEn" dc:"格口类型名称_英文"`
	NameJa string `json:"nameJa" dc:"格口类型名称_日语"`
	NameKo string `json:"nameKo" dc:"格口类型名称_韩语"`
	NameTw string `json:"nameTw" dc:"格口类型名称_繁体中文"`
	Size   string `json:"size" dc:"格口类型尺寸"`
	Price  int    `json:"price" dc:"格口类型单价（每小时）"`
	Num    int    `json:"num" dc:"格口类型库存"`
}

type CabinetCreateOrderParams struct {
	CabinetId     int    `json:"cabinet_id" dc:"储物柜ID"`
	BoxTypeId     int    `json:"box_type_id" dc:"格口类型ID"`
	Supercode     string `json:"supercode" dc:"超级码"`
	OutTradeNo    string `json:"out_trade_no" dc:"第三方订单号"`
	OutTradeVipid string `json:"out_trade_vipid" dc:"第三方会员ID"`
	BuyHours      int    `json:"buy_hours" dc:"购买时长"`
}

type CabinetCreateOrderResponse struct {
	Code int                            `json:"code"`
	Msg  string                         `json:"msg"`
	Data CabinetCreateOrderResponseItem `json:"data"`
}

type CabinetCreateOrderResponseItem struct {
	OrderId int    `json:"orderId"`
	OrderNo string `json:"orderNo"`
	BoxId   int    `json:"boxId"`
	BoxNo   string `json:"boxNo"`
	Pin     string `json:"pin"`
	//TotalFee      float64     `json:"totalFee"`
	//ExpireTime    *gtime.Time `json:"expireTime"`
	Status        int         `json:"status"`
	Color         string      `json:"color"`
	OutTradeNo    string      `json:"outTradeNo"`
	OutTradeVipid string      `json:"outTradeVipid"`
	CabinetId     int         `json:"cabinetId"`
	BoxAlias      string      `json:"boxAlias"`
	AddressZh     string      `json:"addressZh"`
	AddressEn     string      `json:"addressEn"`
	AddressJa     string      `json:"addressJa"`
	AddressKo     string      `json:"addressKo"`
	AddressTw     string      `json:"addressTw"`
	StartTime     *gtime.Time `json:"startTime"`
	EndTime       *gtime.Time `json:"endTime"`
	GraceSeconds  int         `json:"graceSeconds"`
}

type CabinetOrderQueryParams struct {
	OutTradeNo string `json:"out_trade_no" dc:"第三方订单号"`
}

type CabinetOrderQueryResponse struct {
	Code int                           `json:"code"`
	Msg  string                        `json:"msg"`
	Data CabinetOrderQueryResponseItem `json:"data"`
}

type CabinetOrderQueryResponseItem struct {
	Id            int         `json:"id"`
	OrderNo       string      `json:"orderNo"`
	OutTradeNo    string      `json:"outTradeNo"`
	OutTradeVipid string      `json:"outTradeVipid"`
	CabinetId     int         `json:"cabinetId"`
	CityId        int         `json:"cityId"`
	MchId         int         `json:"mchId"`
	MchBranchId   int         `json:"mchBranchId"`
	BoxId         int         `json:"boxId"`
	PricePlanId   int         `json:"pricePlanId"`
	BoxTypeId     int         `json:"boxTypeId"`
	Pin           string      `json:"pin"`
	Status        int         `json:"status"`
	BoxOpenCount  int         `json:"boxOpenCount"`
	BoxMiddleOpen int         `json:"boxMiddleOpen"`
	BuyHours      int         `json:"buyHours"`
	BasePrice     int         `json:"basePrice"`
	FirstFeeRate  string      `json:"firstFeeRate"`
	FirstFee      float64     `json:"firstFee"`
	OvertimeHours int         `json:"overtimeHours"`
	OvertimeFee   float64     `json:"overtimeFee"`
	TotalFee      float64     `json:"totalFee"`
	PaidFee       float64     `json:"paidFee"`
	StartTime     *gtime.Time `json:"startTime"`
	EndTime       *gtime.Time `json:"endTime"`
	GraceEndTime  *gtime.Time `json:"graceEndTime"`
	LastOpenTime  *gtime.Time `json:"lastOpenTime"`
	LastCheckTime *gtime.Time `json:"lastCheckTime"`
	CancelTime    *gtime.Time `json:"cancelTime"`
	FinishTime    *gtime.Time `json:"finishTime"`
	OpenLog       string      `json:"openLog"`
	MchBranchName string      `json:"mchBranchName"`
	MchBranchLat  string      `json:"mchBranchLat"`
	MchBranchLgt  string      `json:"mchBranchLgt"`
}

type CabinetPayOvertimeParams struct {
	OutTradeNo string `json:"out_trade_no" dc:"第三方订单号"`
}

type CabinetPayOvertimeResponse struct {
	Code int                            `json:"code"`
	Msg  string                         `json:"msg"`
	Data CabinetPayOvertimeResponseItem `json:"data"`
}

type CabinetPayOvertimeResponseItem struct {
	OrderId      int         `json:"orderId"`
	OrderNo      string      `json:"orderNo"`
	Status       int         `json:"status"`
	StatusText   string      `json:"statusText"`
	GraceEndTime *gtime.Time `json:"graceEndTime"`
	GraceSeconds int         `json:"graceSeconds"`
}

type CabinetListResponse struct {
	Code int                        `json:"code"`
	Msg  string                     `json:"msg"`
	Data []*CabinetListResponseItem `json:"data"`
}

type CabinetListResponseItem struct {
	Id                 int    `json:"id" dc:"储物柜ID"`
	Name               string `json:"name" dc:"储物柜名称"`
	NameZh             string `json:"nameZh" dc:"储物柜名称_简体中文"`
	NameEn             string `json:"nameEn" dc:"储物柜名称_英文"`
	NameJa             string `json:"nameJa" dc:"储物柜名称_日语"`
	NameKo             string `json:"nameKo" dc:"储物柜名称_韩语"`
	NameTw             string `json:"nameTw" dc:"储物柜名称_繁体中文"`
	No                 string `json:"no" dc:"储物柜编号"`
	MchBranchName      string `json:"mchBranchName" dc:"网点名称"`
	MchBranchNameZh    string `json:"mchBranchNameZh" dc:"网点名称_简体中文"`
	MchBranchNameEn    string `json:"mchBranchNameEn" dc:"网点名称_英文"`
	MchBranchNameJa    string `json:"mchBranchNameJa" dc:"网点名称_日语"`
	MchBranchNameKo    string `json:"mchBranchNameKo" dc:"网点名称_韩语"`
	MchBranchNameTw    string `json:"mchBranchNameTw" dc:"网点名称_繁体中文"`
	MchBranchAddress   string `json:"mchBranchAddress" dc:"网点地址"`
	MchBranchAddressZh string `json:"mchBranchAddressZh" dc:"网点地址_简体中文"`
	MchBranchAddressEn string `json:"mchBranchAddressEn" dc:"网点地址_英文"`
	MchBranchAddressJa string `json:"mchBranchAddressJa" dc:"网点地址_日语"`
	MchBranchAddressKo string `json:"mchBranchAddressKo" dc:"网点地址_韩语"`
	MchBranchAddressTw string `json:"mchBranchAddressTw" dc:"网点地址_繁体中文"`
	MchBranchLat       string `json:"mchBranchLat" dc:"网点纬度"`
	MchBranchLgt       string `json:"mchBranchLgt" dc:"网点经度"`
}

type OrderCompleteParams struct {
	OutTradeNo string `json:"out_trade_no" dc:"第三方订单号"`
}

type OrderCompleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewClient(ctx context.Context, config *model.CabinetApiConfig) *CabinetClient {
	var (
		appid     string
		appkey    string
		apiSecret string
	)
	// baseUrl = g.Cfg().MustGet(ctx, "cabinet.host").String()
	// appid = g.Cfg().MustGet(ctx, "cabinet.appid").String()
	// appkey = g.Cfg().MustGet(ctx, "cabinet.appkey").String()
	baseUrl = config.CabinetDomain + "/api"
	appid = config.CabinetAppid
	appkey = config.CabinetAppkey
	apiSecret = config.CabinetApiSecret

	return &CabinetClient{
		Appid:       appid,
		Appkey:      appkey,
		ApiSecret:   apiSecret,
		GHttpClient: g.Client(),
	}
}

func (c *CabinetClient) DoRequest(ctx context.Context, method string, url string, paramsReq interface{}) (response *gclient.Response, err error) {

	c.GHttpClient.SetHeader("Content-Type", "application/json")

	// apiSecret := g.Cfg().MustGet(ctx, "cabinet.apiSecret").String()

	apiSecret := c.ApiSecret

	// 1. 获取时间戳
	timestamp := gtime.Now().Timestamp() // 秒级，毫秒可用 TimestampMilli()

	// 2. 把原始参数转成 map
	paramMap := gconv.Map(paramsReq)
	if paramMap == nil {
		paramMap = make(map[string]interface{})
	}

	// 3. 加上通用参数
	paramMap["appid"] = c.Appid
	paramMap["appkey"] = c.Appkey
	paramMap["timestamp"] = timestamp

	// 4. 生成 sign（这里假设签名规则是 md5(appid+timestamp+appkey)）
	rawSignStr := fmt.Sprintf("%s%s%d%s", c.Appid, c.Appkey, timestamp, apiSecret)
	sign := gmd5.MustEncryptString(rawSignStr)
	paramMap["sign"] = sign

	if response, err = c.GHttpClient.DoRequest(ctx, method, baseUrl+url, paramMap); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())
	return
}
