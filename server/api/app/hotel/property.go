package hotel

import (
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type PropertySearchReq struct {
	g.Meta       `path:"/home/searchPropertyOnline" method:"post" tags:"APP_HOTEL" summary:"物业_检索列表"`
	StartDate    string `json:"startDate" v:"required|date#please_enter_the_checkin_date|date_format_error" dc:"入住日期"`
	EndDate      string `json:"endDate" v:"required|date#please_enter_the_checkout_date|date_format_error" dc:"退房日期"`
	RoomNumber   int    `json:"roomNumber" v:"required#please_enter_the_number_of_rooms" dc:"房间数"`
	BookerNumber int    `json:"bookerNumber" v:"required#please_enter_the_number_of_occupants" dc:"入住人数"`
	AdultNumber  int    `json:"adultNumber"  dc:"成人人数"`
	ChildNumber  int    `json:"childNumber"  dc:"儿童人数"`
	InfantNumber int    `json:"infantNumber"  dc:"婴儿人数"`
	IsLease      int    `json:"isLease"  dc:"是否短租 1、否 2、是"`
	RegionId     int    `json:"regionId"  dc:"地区ID"`
}

type PropertySearchRes struct {
	List []*PropertyOnlineSearchItem
}

type PropertyOnlineSearchItem struct {
	Id              int         `json:"id"                  dc:"主键"`
	Uid             string      `json:"uid"                 dc:"物业ID"`
	Name            string      `json:"name"                dc:"物业名称"`
	Description     string      `json:"description"         dc:"物业名称"`
	TagList         []string    `json:"tagListArr"          dc:"物业标签"`
	LinePrice       float64     `json:"linePrice"           dc:"划线价"`
	Price           float64     `json:"price"               dc:"预估价格"`
	Icon            string      `json:"icon"                dc:"图标"`
	Lat             string      `json:"lat"                 dc:"纬度"`
	Lng             string      `json:"lng"                 dc:"经度"`
	AddressDetail   string      `json:"addressDetail"       dc:"定位百度详细地址"`
	GgLat           string      `json:"ggLat"               dc:"谷歌纬度"`
	GgLng           string      `json:"ggLng"               dc:"谷歌经度"`
	GgAddressDetail string      `json:"ggAddressDetail"     dc:"定位google详细地址"`
	Cover           string      `json:"cover"               dc:"封面"`
	Currency        string      `json:"currency"            dc:"货币"`
	Language        string      `json:"language"            dc:"语言"`
	TimeZone        string      `json:"timeZone"            dc:"时区"`
	Address         string      `json:"address"             dc:"地址"`
	CreatedAt       *gtime.Time `json:"createdAt"           dc:"创建时间"`
	Inventory       int         `json:"inventory"           dc:"当前房间库存-如果入住房间数大于该值应该设置为售罄，不可预约"`
	ContactName     string      `json:"contactName"         dc:"联系人"`
	Phone           string      `json:"phone"               dc:"联系方式"`
	ContactEmail    string      `json:"contactEmail"        dc:"邮箱"`
	IsDisable       bool        `json:"isDisable"            dc:"是否禁用"`
	Occupancy       int         `json:"occupancy"            dc:"入住人数"`
	MinDaysNotice   int         `json:"minDaysNotice"        dc:"短租模式最小预定区间"`
	RegionId        int         `json:"regionId"             dc:"所属地区"`
	RegionInfo      *struct {
		gmeta.Meta `orm:"table:hg_pms_property_region"`
		Id         int    `json:"id"        orm:"id"         description:"主键"`
		Name       string `json:"name"      orm:"name"       description:"区域名称（多语言）"`
	} `json:"regionInfo" orm:"with:id=region_id"  dc:"所属地区"`
}

type PropertyViewReq struct {
	g.Meta `path:"/pmsProperty/view" method:"post" tags:"APP_HOTEL" summary:"物业_详情"`
	input_hotel.PmsPropertyViewInp
}

type PropertyViewRes struct {
	*input_hotel.PmsPropertyViewModel
	CollectId int     `json:"collectId"  dc:"收藏数据ID"`
	JaRate    float64 `json:"jaRate"           dc:"日元汇率"`
}
