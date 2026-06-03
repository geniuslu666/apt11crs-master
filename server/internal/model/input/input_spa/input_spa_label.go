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
)

// SpaLabelUpdateFields 修改服务标签字段过滤
type SpaLabelUpdateFields struct {
	LabelName    string `json:"labelName" dc:"标签名称"`
	LabelContent string `json:"labelContent" dc:"标签内容"`
	Status       int    `json:"status"    dc:"状态"`
	Sort         int    `json:"sort"      dc:"排序"`
}

// SpaLabelInsertFields 新增服务标签字段过滤
type SpaLabelInsertFields struct {
	LabelName    string `json:"labelName" dc:"标签名称"`
	LabelContent string `json:"labelContent" dc:"标签内容"`
	Status       int    `json:"status"    dc:"状态"`
	Sort         int    `json:"sort"      dc:"排序"`
}

// SpaLabelEditInp 修改/新增服务标签
type SpaLabelEditInp struct {
	entity.SpaLabel
	ContentLanguage input_language.LanguageModel `json:"contentLanguage"          dc:"多语言标签内容"`
}

func (in *SpaLabelEditInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaLabelEditModel struct{}

// SpaLabelDeleteInp 删除服务标签
type SpaLabelDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaLabelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaLabelDeleteModel struct{}

// SpaLabelViewInp 获取指定服务标签信息
type SpaLabelViewInp struct {
	Id         int  `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *SpaLabelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaLabelViewModel struct {
	entity.SpaLabel
	ContentLanguage []*input_hotel.LanguageType `json:"contentLanguage"         dc:"标签内容"   orm:"with:uuid=label_content"`
}

// SpaLabelListInp 获取服务标签列表
type SpaLabelListInp struct {
	input_form.PageReq
	LabelName string `json:"labelName" dc:"标签名称"`
	Status    int    `json:"status"    dc:"状态"`
}

func (in *SpaLabelListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaLabelListModel struct {
	Id        int         `json:"id"        dc:"id"`
	LabelName string      `json:"labelName" dc:"标签名称"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int         `json:"sort"      dc:"排序"`
	CreateAt  *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  dc:"更新时间"`
}

// SpaLabelMaxSortInp 获取服务标签最大排序
type SpaLabelMaxSortInp struct{}

func (in *SpaLabelMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaLabelMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// SpaLabelStatusInp 更新服务标签状态
type SpaLabelStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *SpaLabelStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type SpaLabelStatusModel struct{}
