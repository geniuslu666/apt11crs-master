package input_spa

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SpaServiceUpdateFields 修改服务管理字段过滤
type SpaServiceUpdateFields struct {
	IspId         int    `json:"ispId"       dc:"服务商ID"`
	Name          string `json:"name"                    dc:"服务名称(多语言)"`
	SubName       string `json:"subName"              dc:"副标题(多语言)"`
	LabelIds      string `json:"labelIds"                dc:"标签（多选）"`
	Images        string `json:"images"                  dc:"图集"`
	Channel       int    `json:"channel"         dc:"服务渠道"`
	PropertyIds   string `json:"propertyIds"                dc:"物业（多选）"`
	ServiceStatus int    `json:"serviceStatus"                  dc:"状态"`
	Sort          int    `json:"sort"      dc:"排序"`
	Content       string `json:"content"                 dc:"服务详情"`
}

// SpaServiceInsertFields 新增服务管理字段过滤
type SpaServiceInsertFields struct {
	IspId         int    `json:"ispId"       dc:"服务商ID"`
	Name          string `json:"name"                    dc:"服务名称(多语言)"`
	SubName       string `json:"subName"              dc:"副标题(多语言)"`
	LabelIds      string `json:"labelIds"                dc:"标签（多选）"`
	Images        string `json:"images"                  dc:"图集"`
	Channel       int    `json:"channel"         dc:"服务渠道"`
	PropertyIds   string `json:"propertyIds"                dc:"物业（多选）"`
	ServiceStatus int    `json:"serviceStatus"                  dc:"状态"`
	Sort          int    `json:"sort"      dc:"排序"`
	Content       string `json:"content"                 dc:"服务详情"`
}

// SpaServiceEditInp 修改/新增服务管理
type SpaServiceEditInp struct {
	entity.SpaService
	NameLanguage    input_language.LanguageModel `json:"nameLanguage"          dc:"多语言服务名称"`
	SubNameLanguage input_language.LanguageModel `json:"subNameLanguage"          dc:"多语言服务简介"`
	ContentLanguage input_language.LanguageModel `json:"contentLanguage"          dc:"多语言服务详情"`
	PriceList       []*struct {
		Id           int                          `json:"id"  dc:"ID"`
		GoodsName    string                       `json:"goodsName" dc:"服务套餐名称(多语言)"`
		Image        string                       `json:"image"        orm:"image"           description:"图片"`
		DetailImage  string                       `json:"detailImage"          description:"详情图片"`
		NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言服务名称"`
		Price        float64                      `json:"price"           dc:"价格"`
		Duration     int                          `json:"duration"  dc:"时长(分钟)"`
		Status       uint                         `json:"status"  dc:"状态"`
	} `json:"priceList"          dc:"价格列表"`
}

func (in *SpaServiceEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaServiceEditModel struct{}

// SpaServiceDeleteInp 删除服务管理
type SpaServiceDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaServiceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaServiceDeleteModel struct{}

// SpaServiceViewInp 获取指定服务管理信息
type SpaServiceViewInp struct {
	Id         int64 `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool  `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *SpaServiceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaServiceViewModel struct {
	entity.SpaService
	NameLanguage    []*input_hotel.LanguageType `json:"nameLanguage"         dc:"服务名称"   orm:"with:uuid=name"`
	SubNameLanguage []*input_hotel.LanguageType `json:"subNameLanguage"         dc:"服务简介"   orm:"with:uuid=sub_name"`
	ContentLanguage []*input_hotel.LanguageType `json:"contentLanguage"         dc:"服务详情"   orm:"with:uuid=content"`
	PriceList       []*struct {
		gmeta.Meta   `orm:"table:hg_spa_service_goods"`
		Id           int64                       `json:"id"        dc:"Id"`
		ServiceId    int64                       `json:"serviceId"  dc:"服务ID"`
		GoodsName    string                      `json:"goodsName"        dc:"项目名"`
		Image        string                      `json:"image"        orm:"image"           description:"图片"`
		DetailImage  string                      `json:"detailImage"     description:"详情图片"`
		NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"服务名称"   orm:"with:uuid=goods_name"`
		Price        float64                     `json:"price"     dc:"套餐售价"`
		Duration     int                         `json:"duration"  dc:"时长(分钟)"`
		Status       int                         `json:"status"  dc:"状态 1-启用 2-禁用"`
	} `json:"priceList" orm:"with:service_id=id" dc:"价格列表"`
}

// SpaServiceListInp 获取服务管理列表
type SpaServiceListInp struct {
	input_form.PageReq
	Name          string `json:"name" dc:"服务名称"`
	ServiceStatus int    `json:"serviceStatus"             dc:"状态"`
	ServiceIds    string `json:"serviceIds"   dc:"餐厅Ids"`
}

func (in *SpaServiceListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaServiceListModel struct {
	Id        int64  `json:"id"                      dc:"id"`
	IspId     int    `json:"ispId"                 dc:"服务商ID"`
	Name      string `json:"name"         dc:"服务名称(多语言)"`
	Images    string `json:"images"       dc:"图集"`
	LabelIds  string `json:"labelIds"                dc:"标签（多选）"`
	LabelList []*struct {
		gmeta.Meta  `orm:"table:hg_spa_service_label"`
		ServiceId   int64 `json:"serviceId"  orm:"service_id"     description:"服务ID"`
		LabelId     int   `json:"labelId"        orm:"label_id"           description:"标签ID"`
		LabelDetail *struct {
			gmeta.Meta `orm:"table:hg_spa_label"`
			*entity.SpaLabel
		} `json:"labelDetail" orm:"with:id=label_id"  dc:"标签信息"`
	} `json:"labelList" orm:"with:service_id=id" dc:"标签列表"`
	PriceList []*struct {
		gmeta.Meta `orm:"table:hg_spa_service_goods"`
		ServiceId  int64   `json:"serviceId"  orm:"service_id"     description:"服务ID"`
		GoodsName  string  `json:"goodsName"        orm:"goods_name"           description:"项目名"`
		Image      string  `json:"image"        orm:"image"           description:"图片"`
		Price      float64 `json:"price"     orm:"price"      description:"套餐售价"`
		Duration   int     `json:"duration"  orm:"duration"   description:"时长(分钟)"`
		Status     uint    `json:"status"  orm:"status"  dc:"状态"`
	} `json:"priceList" orm:"with:service_id=id" dc:"项目列表"`
	SpaIspDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_isp"`
		Id         int    `json:"id"      orm:"id"    dc:""`
		Name       string `json:"name"    orm:"name"  dc:"服务商名称"`
	} `json:"spaIspDetail" orm:"with:id=isp_id" dc:"服务商信息"`
	Sort         int         `json:"sort"         orm:"sort"          description:"排序(越大越靠前)"`
	ServiceState int         `json:"serviceState" orm:"service_state" description:"状态（1-立即上架 2-放入仓库）"`
	CreateAt     *gtime.Time `json:"createAt"                dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"                dc:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"               dc:"删除时间"`
}

type SpaServiceAllListModel struct {
	Id           int64  `json:"id"                      dc:"id"`
	IspId        int    `json:"ispId"                 dc:"服务商ID"`
	Name         string `json:"name"         dc:"服务名称(多语言)"`
	SpaIspDetail *struct {
		gmeta.Meta `orm:"table:hg_spa_isp"`
		Id         int    `json:"id"      orm:"id"    dc:""`
		Name       string `json:"name"    orm:"name"  dc:"服务商名称"`
	} `json:"spaIspDetail" orm:"with:id=isp_id" dc:"服务商信息"`
	Sort         int `json:"sort"          dc:"排序(越大越靠前)"`
	ServiceState int `json:"serviceState"  dc:"状态（1-立即上架 2-放入仓库）"`
}

// SpaServiceMaxSortInp 获取服务最大排序
type SpaServiceMaxSortInp struct{}

func (in *SpaServiceMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaServiceMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// SpaServiceStatusInp 更新服务状态
type SpaServiceStatusInp struct {
	Id            int `json:"id" v:"required#id不能为空" dc:"id"`
	ServiceStatus int `json:"serviceStatus" dc:"状态"`
}

func (in *SpaServiceStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.ServiceStatus <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.ServiceStatus) {
		err = gerror.New("状态不正确")
		return
	}

	return
}

type SpaServiceStatusModel struct{}
