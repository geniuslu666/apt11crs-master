package order

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderListReq struct {
	g.Meta    `path:"/order/list" method:"get" tags:"订单"  summary:"订单列表"`
	StartTime *gtime.Time `json:"startTime" v:"required#开始时间不能为空"`
	EndTime   *gtime.Time `json:"endTime" v:"required#结束时间不能为空"`
	NextId    int64       `json:"nextId" v:"required#下一个列表从哪个ID开始 第一页咳填入0"`
	Limit     int         `json:"limit" v:"required|max:1000#分页大小|最大为1000"`
}

type OrderListRes struct {
	List   []*OrderListItem `json:"list"`
	NextId int              `json:"nextId"`
}
type OrderListItem struct {
	Puid         string      `json:"puid"                   description:"物业ID"`
	SourceCode   string      `json:"sourceCode"             description:"来源渠道"`
	SourceName   string      `json:"sourceName"             description:"渠道名称"`
	PuidName     string      `json:"puidName"               description:"物业名称"`
	OrderSn      string      `json:"orderSn"                description:"系统订单号"`
	OrderIndex   int         `json:"orderIndex"             description:"订单索引"`
	OutOrderSn   string      `json:"outOrderSn"             description:"三方订单号"`
	RoomType     string      `json:"roomType"               description:"房型信息，参考房型uid"`
	RoomTypeName string      `json:"roomTypeName"           description:"房型名称"`
	BookingFee   float64     `json:"bookingFee"             description:"预订费"`
	CheckinDate  *gtime.Time `json:"checkinDate"           description:"入住时间"`
	CheckoutDate *gtime.Time `json:"checkoutDate"          description:"退房时间"`
	CreateTime   *gtime.Time `json:"createTime"           description:"创建时间"`
}
