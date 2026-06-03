// Package sysin

package input_food

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

// FoodLabelUpdateFields 修改餐厅标签字段过滤
type FoodLabelUpdateFields struct {
	Type   int    `json:"type"      dc:"1 餐厅  2 套餐"`
	Name   string `json:"name" dc:"标签名称"`
	Status int    `json:"status"    dc:"状态"`
	Sort   int    `json:"sort"      dc:"排序"`
}

// FoodLabelInsertFields 新增餐厅标签字段过滤
type FoodLabelInsertFields struct {
	Type   int    `json:"type"      dc:"1 餐厅  2 套餐"`
	Name   string `json:"name" dc:"标签名称"`
	Status int    `json:"status"    dc:"状态"`
	Sort   int    `json:"sort"      dc:"排序"`
}

// FoodLabelEditInp 修改/新增餐厅标签
type FoodLabelEditInp struct {
	entity.FoodLabel
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言标签名称"`
}

func (in *FoodLabelEditInp) Filter(ctx context.Context) (err error) {
	//// 验证标签名称
	//if err := g.Validator().Rules("required").Data(in.LabelName).Messages("标签名称不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}
	//
	//// 验证排序
	//if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
	//	return err.Current()
	//}

	return
}

type FoodLabelEditModel struct{}

// FoodLabelDeleteInp 删除餐厅标签
type FoodLabelDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodLabelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodLabelDeleteModel struct{}

// FoodLabelViewInp 获取指定餐厅标签信息
type FoodLabelViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodLabelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodLabelViewModel struct {
	entity.FoodLabel
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"房型名称"   orm:"with:uuid=name"`
}

// FoodLabelListInp 获取餐厅标签列表
type FoodLabelListInp struct {
	input_form.PageReq
	Name   string `json:"name" dc:"标签名称"`
	Status int    `json:"status"    dc:"状态"`
	Type   int    `json:"type"      dc:"1 餐厅  2 套餐"`
}

func (in *FoodLabelListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodLabelListModel struct {
	Id       int         `json:"id"        dc:"id"`
	Type     int         `json:"type"      dc:"1 餐厅  2 套餐"`
	Name     string      `json:"name" dc:"标签名称"`
	Status   int         `json:"status"    dc:"状态"`
	Sort     int         `json:"sort"      dc:"排序"`
	CreateAt *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt"  dc:"更新时间"`
}

// FoodLabelMaxSortInp 获取餐厅标签最大排序
type FoodLabelMaxSortInp struct{}

func (in *FoodLabelMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodLabelMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// FoodLabelStatusInp 更新餐厅标签状态
type FoodLabelStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodLabelStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodLabelStatusModel struct{}
