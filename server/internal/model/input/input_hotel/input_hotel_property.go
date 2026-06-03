// Package sysin

package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// PmsPropertyUpdateFields 修改物业字段过滤
type PmsPropertyUpdateFields struct {
	Name             string `json:"name"     dc:"物业名称"`
	Icon             string `json:"icon"     dc:"图标"`
	Cover            string `json:"cover"    dc:"封面"`
	Currency         string `json:"currency" dc:"货币"`
	Language         string `json:"language" dc:"语言"`
	TimeZone         string `json:"timeZone" dc:"时区"`
	Address          string `json:"address"  dc:"地址"`
	LeaseClose       int    `json:"leaseClose"           dc:"1、开启物业短租  2、关闭物业短租"`
	BookingClose     int    `json:"bookingClose"         dc:"1、开启预订模式  2、关闭预订模式"`
	BatchReservation string `json:"batchReservation"     dc:"是否开启多房型预定 （Y 开启 N  关闭）"`
}

// PmsPropertyInsertFields 新增物业字段过滤
type PmsPropertyInsertFields struct {
	Name             string `json:"name"     dc:"物业名称"`
	Icon             string `json:"icon"     dc:"图标"`
	Cover            string `json:"cover"    dc:"封面"`
	Currency         string `json:"currency" dc:"货币"`
	Language         string `json:"language" dc:"语言"`
	TimeZone         string `json:"timeZone" dc:"时区"`
	Address          string `json:"address"  dc:"地址"`
	LeaseClose       int    `json:"leaseClose"           dc:"1、开启物业短租  2、关闭物业短租"`
	BookingClose     int    `json:"bookingClose"         dc:"1、开启预订模式  2、关闭预订模式"`
	BatchReservation string `json:"batchReservation"     dc:"是否开启多房型预定 （Y 开启 N  关闭）"`
}

// PmsPropertyEditInp 修改/新增物业
type PmsPropertyEditInp struct {
	entity.PmsProperty
	ShowType int `json:"showType"           dc:"1、全部可见  2、部分可见"`
}

func (in *PmsPropertyEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsPropertyEditModel struct{}

// PmsPropertyEditGroupInp 物业绑定会员分组
type PmsPropertyEditGroupInp struct {
	Id       int   `json:"id" v:"required#主键不能为空" dc:"主键"`
	GroupIds []int `json:"groupIds" dc:"开放显示的会员分组IDS"`
}

func (in *PmsPropertyEditGroupInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsPropertyEditGroupModel struct{}

// PmsPropertyDeleteInp 删除物业
type PmsPropertyDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsPropertyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsPropertyDeleteModel struct{}

// PmsPropertyRecycleInp 物业回收
type PmsPropertyRecycleInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsPropertyRecycleInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsPropertyRecycleModel struct{}

// PmsPropertyViewInp 获取指定物业信息
type PmsPropertyViewInp struct {
	Id         int    `json:"id" v:"required#the_primary_key_cannot_be_empty" dc:"主键"`
	Uid        string `json:"uid" dc:"唯一标识"`
	IsLanguage bool   `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *PmsPropertyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsPropertyViewModel struct {
	Id                   int         `json:"id"                   orm:"id"                      description:"主键"`
	Uid                  string      `json:"uid"                  orm:"uid"                     description:"在API合作伙伴系统中的物业ID"`
	Icon                 string      `json:"icon"                 orm:"icon"                    description:"图标"`
	Name                 string      `json:"name"                 orm:"name"                    description:"物业名称   多语言"`
	Style                string      `json:"style"                orm:"style"                   description:"物业类型"`
	Cover                string      `json:"cover"                orm:"cover"                   description:"封面"`
	Currency             string      `json:"currency"             orm:"currency"                description:"物业的默认货币"`
	Language             string      `json:"language"             orm:"language"                description:"物业的默认自动消息语言"`
	TimeZone             string      `json:"timeZone"             orm:"time_zone"               description:"标准时区名称"`
	Address              string      `json:"address"              orm:"address"                 description:"地址描述"`
	Lat                  string      `json:"lat"                  orm:"lat"                     description:"纬度"`
	Lng                  string      `json:"lng"                  orm:"lng"                     description:"经度"`
	AddressDetail        string      `json:"addressDetail"        orm:"address_detail"          description:"百度详细地址"`
	GgAddressDetail      string      `json:"ggAddressDetail"      orm:"gg_address_detail"       description:"google详细地址"`
	GgLat                string      `json:"ggLat"                orm:"gg_lat"                  description:"谷歌纬度"`
	GgLng                string      `json:"ggLng"                orm:"gg_lng"                  description:"谷歌经度"`
	TagList              *gjson.Json `json:"tagList"              orm:"tag_list"                description:"标签多语言"`
	RoomDes              string      `json:"roomDes"              orm:"room_des"                description:"房间描述"`
	Surroundings         string      `json:"surroundings"         orm:"surroundings"            description:"周边环境"`
	Description          string      `json:"description"          orm:"description"             description:"物业描述  多语言"`
	BusStation           string      `json:"busStation"           orm:"bus_station"             description:"公交站"`
	Subway               string      `json:"subway"               orm:"subway"                  description:"地铁站"`
	ContactName          string      `json:"contactName"          orm:"contact_name"            description:"联系人"`
	Phone                string      `json:"phone"                orm:"phone"                   description:"联系方式"`
	ContactEmail         string      `json:"contactEmail"         orm:"contact_email"           description:"邮箱"`
	MaxDaysNotice        int         `json:"maxDaysNotice"        orm:"max_days_notice"         description:"最大预定区间"`
	MinDaysNotice        int         `json:"minDaysNotice"        orm:"min_days_notice"         description:"短租模式最小预定区间"`
	MinutesAfterCheckout int         `json:"minutesAfterCheckout" orm:"minutes_after_checkout"  description:"退房后 分钟"`
	MinutesBeforeCheckin int         `json:"minutesBeforeCheckin" orm:"minutes_before_checkin"  description:"入住前 分钟"`
	BookingLeadTimeLabel string      `json:"bookingLeadTimeLabel" orm:"booking_lead_time_label" description:"预约期限"`
	TurnoverDays         uint        `json:"turnoverDays"         orm:"turnover_days"           description:"周转天数"`
	LinePrice            float64     `json:"linePrice"            orm:"line_price"              description:"单价"`
	Price                float64     `json:"price"                orm:"price"                   description:"单价"`
	CheckinAt            string      `json:"checkinAt"            orm:"checkin_at"              description:"入住时间"`
	CheckoutAt           string      `json:"checkoutAt"           orm:"checkout_at"             description:"退房时间"`
	CancelPolicy         string      `json:"cancelPolicy"         orm:"cancel_policy"           description:"取消政策 多语言"`
	GalleryImages        *gjson.Json `json:"galleryImages"        orm:"gallery_images"          description:"画廊图片"`
	GalleryCover         string      `json:"galleryCover"         orm:"gallery_cover"           description:"画廊封面"`
	RequiredBook         string      `json:"requiredBook"         orm:"required_book"           description:""`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`
	LeaseClose           int         `json:"leaseClose"           orm:"lease_close"             description:"1、开启物业短租  2、关闭物业短租"`
	BookingClose         int         `json:"bookingClose"         orm:"booking_close"           description:"1、开启预订模式  2、关闭预订模式"`
	CheckInGuide         string      `json:"checkInGuide"         orm:"check_in_guide"          description:"入住指南"`
	BatchReservation     string      `json:"batchReservation"     orm:"batch_reservation"       description:"是否开启多房型预定 （Y 开启 N  关闭）"`
	RegionId             uint        `json:"regionId"             orm:"region_id"               description:"地区"`
	RegionInfo           *struct {
		gmeta.Meta `orm:"table:hg_pms_property_region"`
		Id         int    `json:"id"        orm:"id"         description:"主键"`
		Name       string `json:"name"      orm:"name"       description:"区域名称（多语言）"`
	} `json:"regionInfo" orm:"with:id=region_id"  dc:"所属地区"`
}
type PmsPropertyViewLannguageModel struct {
	entity.PmsProperty
	NameLanguage         []*LanguageType `json:"nameLanguage"         dc:"物业名称"   orm:"with:uuid=name"`
	DescriptionLanguage  []*LanguageType `json:"descriptionLanguage"  dc:"物业描述"   orm:"with:uuid=description"`
	StyleLanguage        []*LanguageType `json:"styleLanguage"        dc:"物业描述"   orm:"with:uuid=style"`
	TagListLanguage      []*LanguageType `json:"tagListLanguage"     dc:"物业描述"   orm:"with:uuid=tag_list"`
	RoomDesLanguage      []*LanguageType `json:"roomDesLanguage"      dc:"物业描述"   orm:"with:uuid=room_des"`
	SurroundingsLanguage []*LanguageType `json:"surroundingsLanguage" dc:"物业描述"   orm:"with:uuid=surroundings"`
	BusStationLanguage   []*LanguageType `json:"busStationLanguage"   dc:"物业描述"   orm:"with:uuid=bus_station"`
	SubwayLanguage       []*LanguageType `json:"subwayLanguage"       dc:"物业描述"   orm:"with:uuid=subway"`
	CancelPolicyLanguage []*LanguageType `json:"cancelPolicyLanguage" dc:"物业描述"   orm:"with:uuid=cancel_policy"`
	AddressLanguage      []*LanguageType `json:"addressLanguage"      dc:"物业描述"   orm:"with:uuid=address"`
	RequiredBookLanguage []*LanguageType `json:"requiredBookLanguage" dc:"物业描述"   orm:"with:uuid=required_book"`
	CheckInGuideLanguage []*LanguageType `json:"checkInGuideLanguage" dc:"入住指南"   orm:"with:uuid=check_in_guide"`
	TagListArr           []string        `json:"tagListArr"         dc:"物业标签"`
	RegionInfo           *struct {
		gmeta.Meta         `orm:"table:hg_pms_property_region"`
		Id                 int             `json:"id"        orm:"id"         description:"主键"`
		Name               string          `json:"name"      orm:"name"       description:"区域名称（多语言）"`
		RegionNameLanguage []*LanguageType `json:"regionNameLanguage" dc:"地区"   orm:"with:uuid=name"`
	} `json:"regionInfo" orm:"with:id=region_id"  dc:"所属地区"`
}

// PmsPropertyListInp 获取物业列表
type PmsPropertyListInp struct {
	input_form.PageReq
	Uid            string        `json:"uid"       dc:"物业ID"`
	Name           string        `json:"name"      dc:"物业名称"`
	CreatedAt      []*gtime.Time `json:"createdAt" dc:"创建时间"`
	IsFindClose    bool          `json:"isFindClose" dc:"是否查询关闭"`
	Sort           string        `json:"sort"      dc:"排序"`
	ContainDelete  bool          `json:"containDelete" dc:"是否包含删除 true 包含"`
	PropertyStatus int           `json:"propertyStatus" dc:"物业状态 1-已启用 2-已关闭 3-已删除 0-全部"`
}

func (in *PmsPropertyListInp) Filter(ctx context.Context) (err error) {
	return
}

type LanguageType struct {
	gmeta.Meta `orm:"table:hg_pms_language"`
	Uuid       string `json:"uuid"   description:""`
	Content    string `json:"content"   description:"语言内容"`
	Language   string `json:"language"  description:"语言"`
}

type PmsPropertyListModel struct {
	Id                  int             `json:"id"                  dc:"主键"`
	Uid                 string          `json:"uid"                 dc:"物业ID"`
	GroupIds            string          `json:"groupIds"            dc:"绑定的会员分组ID"`
	Name                string          `json:"name"                dc:"物业名称"`
	Description         string          `json:"description"         dc:"物业名称"`
	NameLanguage        []*LanguageType `json:"nameLanguage"        dc:"多语言物业名称"   orm:"with:uuid=name"`
	DescriptionLanguage []*LanguageType `json:"descriptionLanguage" dc:"多语言物业描述"   orm:"with:uuid=description"`
	AddressLanguage     []*LanguageType `json:"addressLanguage"     dc:"多语言物业地址"   orm:"with:uuid=address"`
	TagList             string          `json:"tag_list"            dc:"物业标签"`
	TagListArr          []string        `json:"tag_list_arr"        dc:"物业标签"`
	Lat                 string          `json:"lat"                 dc:"纬度"`
	Lng                 string          `json:"lng"                 dc:"经度"`
	AddressDetail       string          `json:"addressDetail"       dc:"百度详细地址"`
	GgLat               string          `json:"ggLat"               dc:"谷歌纬度"`
	GgLng               string          `json:"ggLng"               dc:"谷歌经度"`
	GgAddressDetail     string          `json:"ggAddressDetail"     dc:"google详细地址"`
	TagListLanguage     []*LanguageType `json:"taglistLanguage"     dc:"物业标签多语言"   orm:"with:uuid=tag_list"`
	LinePrice           float64         `json:"line_price"          dc:"划线价"`
	Price               float64         `json:"price"               dc:"预估价格"`
	Icon                string          `json:"icon"                dc:"图标"`
	Cover               string          `json:"cover"               dc:"封面"`
	Currency            string          `json:"currency"            dc:"货币"`
	Language            string          `json:"language"            dc:"语言"`
	TimeZone            string          `json:"timeZone"            dc:"时区"`
	Address             string          `json:"address"             dc:"地址"`
	CreatedAt           *gtime.Time     `json:"createdAt"           dc:"创建时间"`
	RoomTypeNum         int             `json:"roomTypeNum"         dc:"房型数量"`
	RoomUnitNum         int             `json:"roomUnitNum"         dc:"房间数量"`
	MinDaysNotice       int             `json:"minDaysNotice"       dc:"短租模式最小预定区间"`
	Close               int             `json:"close"               dc:"1、关闭该物业  2、开启该物业"`
	Sort                int             `json:"sort"                dc:"排序 越大越靠前"`
	BookingClose        int             `json:"bookingClose"        dc:"1、开启预订模式  2、关闭预订模式"`
	LeaseClose          int             `json:"leaseClose"          dc:"1、开启物业短租  2、关闭物业短租"`
	SpaCanOrder         int             `json:"spaCanOrder"         dc:"是否开放预定按摩服务：1-开放 2-关闭"`
	AccessPass          string          `json:"accessPass"           dc:"门禁密码"`
	DeletedAt           *gtime.Time     `json:"deletedAt"           dc:"删除时间"`
}

type PmsPropertyAppIndexModel struct {
	Id            int             `json:"id"                  dc:"主键"`
	Uid           string          `json:"uid"                 dc:"物业ID"`
	Name          string          `json:"name"                dc:"物业名称"`
	NameLanguage  []*LanguageType `json:"nameLanguage"        dc:"多语言物业名称"   orm:"with:uuid=name"`
	AddressDetail string          `json:"addressDetail"       dc:"百度详细地址"`
	Cover         string          `json:"cover"               dc:"封面"`
	Sort          int             `json:"sort"                dc:"排序 越大越靠前"`
}

// PmsPropertyAllInp 获取所有物业列表
type PmsPropertyAllInp struct {
	Uid         string `json:"uid"       dc:"物业ID"`
	Name        string `json:"name"      dc:"物业名称"`
	Ids         string `json:"ids"       dc:"物业ids"`
	IsFindClose bool   `json:"isFindClose" dc:"是否查询关闭"`
}

func (in *PmsPropertyAllInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsPropertyAllModel struct {
	Id        int         `json:"id"                  dc:"主键"`
	Uid       string      `json:"uid"                 dc:"物业ID"`
	Name      string      `json:"name"                dc:"物业名称"`
	CreatedAt *gtime.Time `json:"createdAt"           dc:"创建时间"`
	Close     int         `json:"close"               dc:"1、关闭该物业  2、开启该物业"`
}

// PmsPropertyExportModel 导出物业
type PmsPropertyExportModel struct {
	Id        int         `json:"id"        dc:"主键"`
	Uid       string      `json:"uid"       dc:"物业ID"`
	Name      string      `json:"name"      dc:"物业名称"`
	Icon      string      `json:"icon"      dc:"图标"`
	Cover     string      `json:"cover"     dc:"封面"`
	Currency  string      `json:"currency"  dc:"货币"`
	Language  string      `json:"language"  dc:"语言"`
	TimeZone  string      `json:"timeZone"  dc:"时区"`
	Address   string      `json:"address"   dc:"地址"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// PmsPropertySortInp 编辑物业排序
type PmsPropertySortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *PmsPropertySortInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsPropertySortModel struct{}

// PropertySwitchSpaCanOrderInp 更新预定状态
type PropertySwitchSpaCanOrderInp struct {
	Id          int `json:"id" v:"required#id不能为空" dc:"id"`
	SpaCanOrder int `json:"spaCanOrder" dc:"是否开放预定按摩服务：1-开放 2-关闭"`
}

// PropertyUpdateAccessPassInp 更新门禁密码
type PropertyUpdateAccessPassInp struct {
	Id         int    `json:"id" v:"required#id不能为空" dc:"id"`
	AccessPass string `json:"accessPass" dc:"门禁密码"`
}
