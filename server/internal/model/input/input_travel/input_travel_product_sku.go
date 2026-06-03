package input_travel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

// TravelProductSkuUpdateFields 修改车型字段过滤
type TravelProductSkuUpdateFields struct {
	ProductId        uint    `json:"productId"     dc:"产品ID"`
	Price            float64 `json:"price"         dc:"售价（JPY）"`
	DailyCapacity    int     `json:"dailyCapacity" dc:"每日最大接待人数"`
	Status           int     `json:"status"        dc:"状态（1启用 2禁用）"`
	Sort             int     `json:"sort"          dc:"排序（越大越靠前）"`
	MeetingPlace     string  `json:"meetingPlace"  dc:"集合地点"`
	MeetingTime      string  `json:"meetingTime"   dc:"集合时间（HH:MM）"`
	GgLat            string  `json:"ggLat"         dc:"谷歌纬度"`
	GgLng            string  `json:"ggLng"         dc:"谷歌经度"`
	ContactMobile    string  `json:"contactMobile" dc:"联系电话"`
	AllowCancel      int     `json:"allowCancel"      dc:"是否允许取消"`
	FreeCancelHours  int     `json:"freeCancelHours"  dc:"几小时前免费"`
	CancelFeePercent int     `json:"cancelFeePercent" dc:"取消费率%"`
}

// TravelProductSkuInsertFields 新增车型字段过滤
type TravelProductSkuInsertFields struct {
	ProductId        uint    `json:"productId"     dc:"产品ID"`
	Name             string  `json:"name"          dc:"车型名称（多语言UUID）"`
	Price            float64 `json:"price"         dc:"售价（JPY）"`
	DailyCapacity    int     `json:"dailyCapacity" dc:"每日最大接待人数"`
	Status           int     `json:"status"        dc:"状态（1启用 2禁用）"`
	Sort             int     `json:"sort"          dc:"排序（越大越靠前）"`
	MeetingPlace     string  `json:"meetingPlace"  dc:"集合地点"`
	MeetingTime      string  `json:"meetingTime"   dc:"集合时间（HH:MM）"`
	GgLat            string  `json:"ggLat"         dc:"谷歌纬度"`
	GgLng            string  `json:"ggLng"         dc:"谷歌经度"`
	ContactMobile    string  `json:"contactMobile" dc:"联系电话"`
	AllowCancel      int     `json:"allowCancel"      dc:"是否允许取消"`
	FreeCancelHours  int     `json:"freeCancelHours"  dc:"几小时前免费"`
	CancelFeePercent int     `json:"cancelFeePercent" dc:"取消费率%"`
}

// TravelProductSkuEditInp 新增/编辑车型
type TravelProductSkuEditInp struct {
	entity.TravelProductSku
	NameLanguage input_language.LanguageModel `json:"nameLanguage" dc:"车型名称多语言"`
}

func (in *TravelProductSkuEditInp) Filter(ctx context.Context) (err error) {
	if in.ProductId <= 0 {
		err = gerror.New("产品ID不能为空")
		return
	}
	return
}

type TravelProductSkuEditModel struct{}

// TravelProductSkuDeleteInp 删除车型
type TravelProductSkuDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *TravelProductSkuDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TravelProductSkuDeleteModel struct{}

// TravelProductSkuViewInp 获取指定车型信息
type TravelProductSkuViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言原始数据"`
}

func (in *TravelProductSkuViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TravelProductSkuViewModel struct {
	entity.TravelProductSku
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage" dc:"车型名称多语言" orm:"with:uuid=name"`
}

// TravelProductSkuListInp 获取车型列表
type TravelProductSkuListInp struct {
	input_form.PageReq
	ProductId uint   `json:"productId" dc:"产品ID"`
	Name      string `json:"name"      dc:"车型名称"`
}

func (in *TravelProductSkuListInp) Filter(ctx context.Context) (err error) {
	return
}

type TravelProductSkuListModel struct {
	Id            uint64  `json:"id"            dc:"id"`
	ProductId     uint    `json:"productId"     dc:"产品ID"`
	Name          string  `json:"name"          dc:"车型名称"`
	Price         float64 `json:"price"         dc:"售价（JPY）"`
	DailyCapacity uint    `json:"dailyCapacity" dc:"每日最大接待人数"`
	Status        int     `json:"status"        dc:"状态（1启用 2禁用）"`
	Sort          int     `json:"sort"          dc:"排序"`
	CreatedAt     string  `json:"createdAt"     dc:"创建时间"`
}

// TravelProductSkuStatusInp 更新车型状态
type TravelProductSkuStatusInp struct {
	Id     int `json:"id"     v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *TravelProductSkuStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}
	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}
	return
}

type TravelProductSkuStatusModel struct{}
