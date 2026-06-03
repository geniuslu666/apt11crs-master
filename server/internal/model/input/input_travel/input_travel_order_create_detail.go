package input_travel

import (
	"APT/internal/model/input/input_hotel"
)

// PreOrderDetailInp 预订单详情入参
type PreOrderDetailInp struct {
	PreOrderSn string `json:"preOrderSn" dc:"预定订单号"`
}

type PreOrderDetailModel struct {
	PreOrderSn string `json:"preOrderSn"     dc:"预订单号"`
	IsFx       string `json:"isFx"       dc:"是否是分销"`
	MemberId   uint64 `json:"memberId"       dc:"会员ID"`
	PreUser    struct {
		FullName  string `json:"fullName"      dc:"全名"`
		Phone     string `json:"phone"         dc:"手机号"`
		PhoneArea string `json:"phoneArea"     dc:"手机区号"`
		Mail      string `json:"mail"          dc:"邮箱"`
	} `json:"preUser" dc:"订单预定人信息"`
	ProductInfo struct {
		Id           uint64  `json:"id"              dc:""`
		Title        string  `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
		SubTitle     string  `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
		MeetingPlace string  `json:"meetingPlace"    dc:"集合地点"`
		MeetingTime  string  `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
		GgLat        string  `json:"ggLat"           dc:"谷歌纬度"`
		GgLng        string  `json:"ggLng"           dc:"谷歌经度"`
		Price        float64 `json:"price"           dc:"售价（JPY）"`
	} `json:"productInfo"  dc:"产品"`
	SkuInfo struct {
		Id            uint64  `json:"id"              dc:""`
		Name          string  `json:"name"           dc:"车型名称（默认语言；多语言存 hg_pms_language）"`
		Price         float64 `json:"price"           dc:"售价（JPY）"`
		MeetingPlace  string  `json:"meetingPlace"  dc:"集合地点"`
		MeetingTime   string  `json:"meetingTime"   dc:"集合时间（HH:MM）"`
		GgLat         string  `json:"ggLat"         dc:"谷歌纬度"`
		GgLng         string  `json:"ggLng"         dc:"谷歌经度"`
		ContactMobile string  `json:"contactMobile" dc:"联系电话"`
	} `json:"skuInfo"  dc:"车型"`
	BookingInfo struct {
		BookDate string `json:"bookDate"       dc:"预定日期"`
		BookNum  int    `json:"bookNum"     dc:"预定人数"`
	} `json:"bookingInfo"  dc:"预定信息"`
	PayInfo        *input_hotel.OrderPayInfoModel `json:"payInfo"     dc:"支付信息"`
	CancelFreeDate string                         `json:"cancelFreeDate"    dc:"免费取消日期"`
	CancelPolicy   string                         `json:"cancelPolicy"      dc:"取消政策"`
	BookingNotice  string                         `json:"bookingTips"       dc:"预定须知"`
}
