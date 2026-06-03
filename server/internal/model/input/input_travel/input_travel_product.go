package input_travel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

// TravelProductUpdateFields 修改产品字段过滤
type TravelProductUpdateFields struct {
	Title           string  `json:"title"           dc:"标题（多语言UUID）"`
	SubTitle        string  `json:"subTitle"           dc:"副标题（多语言UUID）"`
	ContentZh       string  `json:"contentZh"       dc:"内容（中文 编辑器）"`
	ContentEn       string  `json:"contentEn"       dc:"内容（英文 编辑器）"`
	ContentKo       string  `json:"contentKo"       dc:"内容（韩文 编辑器）"`
	ContentJa       string  `json:"contentJa"       dc:"内容（日文 编辑器）"`
	ContentTw       string  `json:"contentTw"       dc:"内容（繁文 编辑器）"`
	ListImage       string  `json:"listImage"       dc:"列表图"`
	CarouselImages  string  `json:"carouselImages"  dc:"轮播图（JSON 数组）"`
	DailyCapacity   int     `json:"dailyCapacity"   dc:"每日最大接待人数"`
	Price           float64 `json:"price"           dc:"售价（JPY）"`
	MeetingPlace    string  `json:"meetingPlace"    dc:"集合地点"`
	MeetingTime     string  `json:"meetingTime"     dc:"集合时间（HH:MM）"`
	GgLat           string  `json:"ggLat"          dc:"谷歌纬度"`
	GgLng           string  `json:"ggLng"          dc:"谷歌经度"`
	MaxBookDays     int     `json:"maxBookDays"     dc:"最大可预约天数"`
	AdvanceBookDays int     `json:"advanceBookDays" dc:"至少提前预订天数"`
	TripPlanningZh  string  `json:"tripPlanningZh"       dc:"行程规划（中文 编辑器）"`
	TripPlanningEn  string  `json:"tripPlanningEn"       dc:"行程规划（英文 编辑器）"`
	TripPlanningJa  string  `json:"tripPlanningJa"       dc:"行程规划（日文 编辑器）"`
	TripPlanningKo  string  `json:"tripPlanningKo"       dc:"行程规划（韩文 编辑器）"`
	TripPlanningTw  string  `json:"tripPlanningTw"       dc:"行程规划（繁文 编辑器）"`
	BookingNotesZh  string  `json:"bookingNotesZh"    dc:"预约须知（中文 编辑器）"`
	BookingNotesEn  string  `json:"bookingNotesEn"    dc:"预约须知（英文 编辑器）"`
	BookingNotesJa  string  `json:"bookingNotesJa"    dc:"预约须知（日文 编辑器）"`
	BookingNotesKo  string  `json:"bookingNotesKo"    dc:"预约须知（韩文 编辑器）"`
	BookingNotesTw  string  `json:"bookingNotesTw"    dc:"预约须知（繁文 编辑器）"`
	Status          int     `json:"status"          dc:"状态（1启用 2禁用）"`
	Sort            int     `json:"sort"            dc:"排序（越大越靠前）"`
	Stock           uint    `json:"stock"           dc:"总库存"`
	ContactMobile   string  `json:"contactMobile" dc:"联系电话"`
}

// TravelProductInsertFields 新增产品字段过滤
type TravelProductInsertFields struct {
	Title           string  `json:"title"           dc:"标题（多语言UUID）"`
	SubTitle        string  `json:"subTitle"           dc:"副标题（多语言UUID）"`
	ContentZh       string  `json:"contentZh"       dc:"内容（中文 编辑器）"`
	ContentEn       string  `json:"contentEn"       dc:"内容（英文 编辑器）"`
	ContentKo       string  `json:"contentKo"       dc:"内容（韩文 编辑器）"`
	ContentJa       string  `json:"contentJa"       dc:"内容（日文 编辑器）"`
	ContentTw       string  `json:"contentTw"       dc:"内容（繁文 编辑器）"`
	ListImage       string  `json:"listImage"       dc:"列表图"`
	CarouselImages  string  `json:"carouselImages"  dc:"轮播图（JSON 数组）"`
	DailyCapacity   int     `json:"dailyCapacity"   dc:"每日最大接待人数"`
	Price           float64 `json:"price"           dc:"售价（JPY）"`
	MeetingPlace    string  `json:"meetingPlace"    dc:"集合地点"`
	MeetingTime     string  `json:"meetingTime"     dc:"集合时间（HH:MM）"`
	GgLat           string  `json:"ggLat"          dc:"谷歌纬度"`
	GgLng           string  `json:"ggLng"          dc:"谷歌经度"`
	MaxBookDays     int     `json:"maxBookDays"     dc:"最大可预约天数"`
	AdvanceBookDays int     `json:"advanceBookDays" dc:"至少提前预订天数"`
	TripPlanningZh  string  `json:"tripPlanningZh"       dc:"行程规划（中文 编辑器）"`
	TripPlanningEn  string  `json:"tripPlanningEn"       dc:"行程规划（英文 编辑器）"`
	TripPlanningJa  string  `json:"tripPlanningJa"       dc:"行程规划（日文 编辑器）"`
	TripPlanningKo  string  `json:"tripPlanningKo"       dc:"行程规划（韩文 编辑器）"`
	TripPlanningTw  string  `json:"tripPlanningTw"       dc:"行程规划（繁文 编辑器）"`
	BookingNotesZh  string  `json:"bookingNotesZh"    dc:"预约须知（中文 编辑器）"`
	BookingNotesEn  string  `json:"bookingNotesEn"    dc:"预约须知（英文 编辑器）"`
	BookingNotesJa  string  `json:"bookingNotesJa"    dc:"预约须知（日文 编辑器）"`
	BookingNotesKo  string  `json:"bookingNotesKo"    dc:"预约须知（韩文 编辑器）"`
	BookingNotesTw  string  `json:"bookingNotesTw"    dc:"预约须知（繁文 编辑器）"`
	Status          int     `json:"status"          dc:"状态（1启用 2禁用）"`
	Sort            int     `json:"sort"            dc:"排序（越大越靠前）"`
	Stock           uint    `json:"stock"           dc:"总库存"`
	ContactMobile   string  `json:"contactMobile" dc:"联系电话"`
}

// TravelProductListInp 获取一日游产品列表
type TravelProductListInp struct {
	input_form.PageReq
	Title     string   `json:"title"     dc:"标题关键词"`
	Status    int      `json:"status"    dc:"状态（0全部 1启用 2禁用）"`
	CreatedAt []string `json:"createdAt" dc:"创建时间范围 [开始, 结束]"`
}

// TravelProductListModel 产品列表展示模型
type TravelProductListModel struct {
	Id            int64   `json:"id"            dc:"产品ID"`
	Title         string  `json:"title"         dc:"标题"`
	SubTitle      string  `json:"subTitle"         dc:"副标题"`
	Price         float64 `json:"price"         dc:"售价（JPY）"`
	DailyCapacity int     `json:"dailyCapacity" dc:"每日接待人数"`
	Sort          int     `json:"sort"          dc:"排序"`
	Status        int     `json:"status"        dc:"状态（1启用 2禁用）"`
	Stock         uint    `json:"stock"           dc:"总库存"`
	CreatedAt     string  `json:"createdAt"     dc:"创建时间"`
	SkuList       []*struct {
		gmeta.Meta `orm:"table:hg_travel_product_sku"`
		Id         uint64  `json:"id"            dc:""`
		ProductId  uint    `json:"productId"     dc:"产品ID"`
		Price      float64 `json:"price"         dc:"售价（元）"`
	} `json:"skuList" orm:"with:product_id=id" dc:"SKU列表"`
}

// TravelProductAppListInp 获取员工活动表列表
type TravelProductAppListInp struct {
	input_form.PageReq
}

func (in *TravelProductAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type TravelProductAppListModel struct {
	Id        int64       `json:"id"        dc:"id"`
	Title     string      `json:"title"     dc:"标题"`
	SubTitle  string      `json:"subTitle"  dc:"副标题"`
	ListImage string      `json:"listImage" dc:"列表图（单图 URL）"`
	Price     float64     `json:"price"     dc:"售价（JPY）"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// TravelProductViewInp 获取产品详情
type TravelProductViewInp struct {
	Id         int64 `json:"id"         v:"required#请选择要查看的产品" dc:"产品ID"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言原始数据"`
}

// TravelProductViewModel 产品详情模型
type TravelProductViewModel struct {
	entity.TravelProduct
	TitleLanguage    []*input_hotel.LanguageType `json:"titleLanguage"        dc:"标题多语言"   orm:"with:uuid=title"`
	SubTitleLanguage []*input_hotel.LanguageType `json:"subTitleLanguage" dc:"副标题多语言" orm:"with:uuid=sub_title"`
	SkuList          []*struct {
		gmeta.Meta    `orm:"table:hg_travel_product_sku"`
		Id            uint64                      `json:"id"            dc:""`
		ProductId     uint                        `json:"productId"     dc:"产品ID"`
		Name          string                      `json:"name"          dc:"车型名称（默认语言；多语言存 hg_pms_language）"`
		Price         float64                     `json:"price"         dc:"售价（元）"`
		DailyCapacity uint                        `json:"dailyCapacity" dc:"每日最大接待人数"`
		Status        int                         `json:"status"        dc:"状态（1启用 2禁用）"`
		Sort          int                         `json:"sort"          dc:"排序（越大越靠前）"`
		SalesNum      uint                        `json:"salesNum"      dc:"已售"`
		CreatedAt     *gtime.Time                 `json:"createdAt"     dc:"创建时间"`
		NameLanguage  []*input_hotel.LanguageType `json:"nameLanguage"  dc:"名称"   orm:"with:uuid=name"`
	} `json:"skuList" orm:"with:product_id=id" dc:"SKU列表"`
}

// TravelProductAppViewInp 获取指定产品表信息-APP
type TravelProductAppViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TravelProductAppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TravelProductAppViewModel struct {
	Id               uint64   `json:"id"              dc:""`
	Title            string   `json:"title"           dc:"标题（默认语言；多语言存 hg_pms_language）"`
	SubTitle         string   `json:"subTitle"        dc:"副标题（默认语言；多语言存 hg_pms_language）"`
	Content          string   `json:"content"       dc:"内容"`
	CarouselImages   []string `json:"carouselImages"  dc:"轮播图（多图，JSON）"`
	DailyCapacity    uint     `json:"dailyCapacity"   dc:"每日最大接待人数"`
	Price            float64  `json:"price"           dc:"售价（JPY）"`
	MeetingPlace     string   `json:"meetingPlace"    dc:"集合地点"`
	MeetingTime      string   `json:"meetingTime"     dc:"集合时间（格式：HH:MM）"`
	GgLat            string   `json:"ggLat"           dc:"谷歌纬度"`
	GgLng            string   `json:"ggLng"           dc:"谷歌经度"`
	MaxBookDays      uint     `json:"maxBookDays"     dc:"最大可预约天数"`
	AdvanceBookDays  uint     `json:"advanceBookDays" dc:"至少提前预订天数"`
	TripPlanning     string   `json:"tripPlanning"    dc:"行程规划"`
	BookingNotes     string   `json:"bookingNotes"    dc:"预约须知"`
	OrderRule        string   `json:"orderRule"       dc:"活动规则"`
	Stock            uint     `json:"stock"           dc:"总库存"`
	SalesNum         uint     `json:"salesNum"        dc:"已售"`
	IsStockExhausted bool     `json:"isStockExhausted" dc:"库存是否已耗尽"`
	DateList         []string `json:"dateList"        dc:"可预约日期列表"`
}

// TravelProductEditInp 新增/编辑产品
type TravelProductEditInp struct {
	entity.TravelProduct
	TitleLanguage    input_language.LanguageModel `json:"titleLanguage"        dc:"标题多语言"`
	SubTitleLanguage input_language.LanguageModel `json:"subTitleLanguage" dc:"副标题多语言"`
	SkuList          []*TravelProductSkuInp       `json:"skuList" dc:"SKU列表"`
}

// TravelProductSkuInp SKU输入结构
type TravelProductSkuInp struct {
	Id            uint                         `json:"id"            dc:"SKU ID"`
	NameLanguage  input_language.LanguageModel `json:"nameLanguage"   dc:"车型名称多语言"`
	Price         float64                      `json:"price"         dc:"售价（JPY）"`
	DailyCapacity int                          `json:"dailyCapacity" dc:"每日最大接待人数"`
	Status        int                          `json:"status"        dc:"状态（1启用 2禁用）"`
	Sort          int                          `json:"sort"          dc:"排序（越大越靠前）"`
}

// TravelProductDeleteInp 删除/恢复产品
type TravelProductDeleteInp struct {
	Id []int64 `json:"id" v:"required#请选择要操作的产品" dc:"产品ID列表"`
}

// TravelProductStatusInp 更新产品状态
type TravelProductStatusInp struct {
	Id     int64 `json:"id"     v:"required#请选择要操作的产品" dc:"产品ID"`
	Status int   `json:"status" v:"in:1,2#状态值错误" dc:"状态（1启用 2禁用）"`
}

// TravelProductSkuStockInp 查询SKU库存
type TravelProductSkuStockInp struct {
	ProductId uint64 `json:"productId" v:"required#产品ID不能为空" dc:"产品ID"`
	BookDate  string `json:"bookDate"  v:"required#预约日期不能为空|date-format:Y-m-d#预约日期格式错误" dc:"预约日期（格式：Y-m-d）"`
}

func (in *TravelProductSkuStockInp) Filter(ctx context.Context) (err error) {
	return
}

// TravelProductSkuStockModel SKU库存信息
type TravelProductSkuStockModel struct {
	Id             uint64  `json:"id"              dc:"SKU ID"`
	Name           string  `json:"name"            dc:"车型名称"`
	Price          float64 `json:"price"           dc:"售价（元）"`
	DailyCapacity  uint    `json:"dailyCapacity"   dc:"每日最大接待人数"`
	BookedNum      uint    `json:"bookedNum"       dc:"已预订人数"`
	RemainingSeats uint    `json:"remainingSeats"  dc:"剩余座位数"`
}
