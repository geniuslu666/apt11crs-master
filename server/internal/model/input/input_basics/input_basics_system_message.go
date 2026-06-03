package input_basics

import (
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// SendMessageInp 发送消息
type SendMessageInp struct {
	SystemMessageTitle   map[string]string `json:"systemMessageTitle" v:"required#系统消息标题不能为空" dc:"系统消息标题"`
	SystemMessageContent map[string]string `json:"systemMessageContent" v:"required#系统消息内容不能为空" dc:"系统消息内容"`
	MemberId             int               `json:"memberId" v:"required#会员ID不能为空" dc:"会员ID"`
	Language             string            `json:"language" v:"required#语言不能为空" dc:"语言"`
	Scene                string            `json:"scene" v:"required#scene不能为空" dc:"scene"`
	Type                 string            `json:"type" v:"required#type不能为空" dc:"type"`
	AppPushData          map[string]string `json:"appPushData" dc:"appPushData"`
	AppLink              string            `json:"appLink"  dc:"appLink"`
	WxLink               string            `json:"wxLink"  dc:"wxLink"`
	EnablePush           bool              `json:"enablePush"  dc:"enablePush"`
	EnableSms            bool              `json:"enableSms" dc:"enableSms"`
	PushTitle            string            `json:"pushTitle" dc:"pushTitle"`
	PushContent          string            `json:"pushContent" dc:"pushContent"`
	OperatorId           int               `json:"operatorId" dc:"operatorId"`
	OperatorRole         string            `json:"operatorRole" dc:"operatorRole"`
	OrderSn              string            `json:"orderSn" dc:"orderSn"`
	ShowIndex            bool              `json:"showIndex" dc:"是否显示在首页"`
}

func (in *SendMessageInp) Filter(ctx context.Context) (err error) {
	return
}

// MessageAppListInp 获取系统消息列表
type MessageAppListInp struct {
	input_form.PageReq
	MemberId  int    `json:"memberId"      dc:"用户"`
	Scene     string `json:"scene"      dc:"场景【system-系统|hotel-酒店|food-餐饮|car-接送机|spa-按摩】"`
	Type      string `json:"type"      dc:"类型【order-订单流转|im-聊天】"`
	UnRead    bool   `json:"unRead" dc:"是否查询未读消息"`
	ShowIndex bool   `json:"showIndex" dc:"是否显示在首页"`
}

func (in *MessageAppListInp) Filter(ctx context.Context) (err error) {
	return
}

type MessageAppListModel struct {
	Id          int    `json:"id" dc:"消息ID"`
	Scene       string `json:"scene"             dc:"场景：system-系统，hotel-酒店，food-餐饮，car-接送机，spa-按摩"`
	Type        string `json:"type"              dc:"类型：order-订单消息，im-聊天消息 (scene为system时，此字段不用管)"`
	Title       string `json:"title" dc:"标题"`
	Content     string `json:"content" dc:"内容"`
	WxLink      string `json:"wxLink" dc:"微信跳转链接"`
	AppLinkType int    `json:"appLinkType" dc:"app跳转类型 【0-不带参数跳转|1-带string参数跳转|2-带param参数跳转|3-跳转网页|4-无需跳转】"`
	AppLink     string `json:"appLink" dc:"app跳转链接"`
	UrlParam    string `json:"urlParam" dc:"app跳转参数"`
	IsRead      bool   `json:"isRead" dc:"是否已读"`
	OrderSn     string `json:"orderSn" dc:"订单号"`
	// ImInfo      *struct {
	// 	Avatar   string `json:"avatar" dc:"头像"`
	// 	Nickname string `json:"nickname" dc:"昵称"`
	// } `json:"imInfo" dc:"聊天消息发送者信息"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// type MessageAppListHotelItem struct {
// 	OrderSn      string      `json:"orderSn"             dc:"订单号"`
// 	Puid         string      `json:"puid"                dc:"物业ID"`
// 	CheckinDate  *gtime.Time `json:"checkinDate"     dc:"入住日期"`
// 	CheckoutDate *gtime.Time `json:"checkoutDate"     dc:"退房日期"`
// 	CheckinTime  string      `json:"checkinTime"      dc:"入住时间，24小时格式"`
// 	CheckoutTime string      `json:"checkoutTime"     dc:"退房时间，24小时格式"`
// 	CreatedAt    *gtime.Time `json:"createdAt"              dc:"订单创建时间"`
// 	PropertyInfo *struct {
// 		gmeta.Meta `orm:"table:hg_pms_property"`
// 		Uid        string `json:"uid"                  dc:"在API合作伙伴系统中的物业ID"`
// 		Name       string `json:"name"                dc:"物业名称"`
// 	} `json:"propertyInfo" dc:"物业信息" orm:"with:uid=puid"`
// }

// type MessageAppListFoodItem struct {
// 	OrderSn        string      `json:"orderSn"             dc:"订单号"`
// 	RestaurantId   int         `json:"restaurantId"            dc:"餐厅ID"`
// 	GoodsId        uint        `json:"goodsId"                  dc:"套餐ID"`
// 	BookDate       string      `json:"bookDate"                 dc:"预定日期"`
// 	BookTime       string      `json:"bookTime"                 dc:"预定时间"`
// 	CreatedAt      *gtime.Time `json:"createdAt"              dc:"订单创建时间"`
// 	RestaurantInfo *struct {
// 		gmeta.Meta `orm:"table:hg_food_restaurant"`
// 		Id         int    `json:"id"            dc:""`
// 		Name       string `json:"name"       dc:"餐厅名称"`
// 	} `json:"restaurantInfo" orm:"with:id=restaurant_id"  dc:"餐厅信息"`
// 	GoodsInfo *struct {
// 		gmeta.Meta `orm:"table:hg_food_goods"`
// 		Id         int    `json:"id"            dc:""`
// 		GoodsName  string `json:"goodsName"       dc:"套餐名称"`
// 	} `json:"goodsInfo" orm:"with:id=goods_id"  dc:"套餐信息"`
// }

// type MessageAppListSpaItem struct {
// 	OrderSn     string      `json:"orderSn"             dc:"订单号"`
// 	ServiceType int         `json:"serviceType"             dc:"1到店  2上门"`
// 	ServiceId   int         `json:"serviceId"               dc:"服务ID"`
// 	GoodsId     uint        `json:"goodsId"                 dc:"项目ID"`
// 	BookDate    string      `json:"bookDate"                 dc:"预定日期"`
// 	BookTime    string      `json:"bookTime"                 dc:"预定时间"`
// 	CreatedAt   *gtime.Time `json:"createdAt"              dc:"订单创建时间"`
// 	ServiceInfo *struct {
// 		gmeta.Meta `orm:"table:hg_spa_service"`
// 		Id         int    `json:"id"            dc:""`
// 		Name       string `json:"name"       dc:"服务名称"`
// 	} `json:"serviceInfo" orm:"with:id=service_id"  dc:"服务信息"`
// 	GoodsInfo *struct {
// 		gmeta.Meta `orm:"table:hg_spa_service_goods"`
// 		Id         int    `json:"id"            dc:""`
// 		GoodsName  string `json:"goodsName"       dc:"项目名称"`
// 	} `json:"goodsInfo" orm:"with:id=goods_id"  dc:"项目信息"`
// }

// type MessageAppListCarItem struct {
// 	OrderSn             string      `json:"orderSn"             dc:"订单号"`
// 	ServiceType         string      `json:"serviceType"             dc:"服务类型"`
// 	StartAddressId      int         `json:"startAddressId"          dc:"出发地ID"`
// 	EndAddressId        int         `json:"endAddressId"            dc:"目的地ID"`
// 	BookDate            string      `json:"bookDate"                 dc:"预定日期"`
// 	BookTime            string      `json:"bookTime"                 dc:"预定时间"`
// 	CreatedAt           *gtime.Time `json:"createdAt"              dc:"订单创建时间"`
// 	StartServiceAddress *struct {
// 		gmeta.Meta `orm:"table:hg_car_address"`
// 		Id         int    `json:"id"            dc:""`
// 		SubName    string `json:"subName"       dc:"地点名称"`
// 	} `json:"startServiceAddress" orm:"with:id=start_address_id"  dc:"出发地"`
// 	EndServiceAddress *struct {
// 		gmeta.Meta `orm:"table:hg_car_address"`
// 		Id         int    `json:"id"            dc:""`
// 		SubName    string `json:"subName"       dc:"地点名称"`
// 	} `json:"endServiceAddress" orm:"with:id=end_address_id"  dc:"目的地"`
// }

type LanguageJson struct {
	Zh   string `json:"zh" dc:"简体中文"`
	En   string `json:"en" dc:"英文"`
	Ja   string `json:"ja" dc:"日语"`
	Ko   string `json:"ko" dc:"韩语"`
	ZhCn string `json:"zh_CN" dc:"繁体中文"`
}

type AppUrlParamModel struct {
	Type   string `json:"type"`
	String string `json:"string"`
	Param  string `json:"param"`
}

// MessageAppReadInp 消息已读
type MessageAppReadInp struct {
	Id       int `json:"id" v:"required#id不能为空" dc:"消息id"`
	MemberId int `json:"member_id" dc:"会员ID(不用传)"`
}

func (in *MessageAppReadInp) Filter(ctx context.Context) (err error) {
	return
}

type MessageAppReadModel struct{}

type MessageAppLatestModel struct {
	SystemMessage *MessageAppLatestItem `json:"systemMessage" dc:"系统消息"`
	HotelMessage  *MessageAppLatestItem `json:"hotelMessage" dc:"民宿消息"`
	FoodMessage   *MessageAppLatestItem `json:"foodMessage" dc:"餐厅消息"`
	SpaMessage    *MessageAppLatestItem `json:"spaMessage" dc:"按摩消息"`
	CarMessage    *MessageAppLatestItem `json:"carMessage" dc:"接送机消息"`
}

type MessageAppLatestItem struct {
	Id          int         `json:"id" dc:"消息ID"`
	Title       string      `json:"title" dc:"标题"`
	Content     string      `json:"content" dc:"内容"`
	WxLink      string      `json:"wxLink" dc:"微信跳转链接"`
	AppLinkType int         `json:"appLinkType" dc:"app跳转类型 【0-不带参数跳转|1-带string参数跳转|2-带param参数跳转|3-跳转网页|4-无需跳转】"`
	AppLink     string      `json:"appLink" dc:"app跳转链接"`
	UrlParam    string      `json:"urlParam" dc:"app跳转参数"`
	IsRead      bool        `json:"isRead" dc:"是否已读"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
	UnReadNum   int         `json:"unReadNum" dc:"同场景下未读消息数"`
}
