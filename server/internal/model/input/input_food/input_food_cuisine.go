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

// FoodCuisineUpdateFields 修改餐厅菜系字段过滤
type FoodCuisineUpdateFields struct {
	CuisineName string `json:"cuisineName" dc:"菜系名称"`
	Pic         string `json:"pic"                    dc:"图片"`
	Status      int    `json:"status"      dc:"状态"`
	Sort        int    `json:"sort"      dc:"排序"`
}

// FoodCuisineInsertFields 新增餐厅菜系字段过滤
type FoodCuisineInsertFields struct {
	CuisineName string `json:"cuisineName" dc:"菜系名称"`
	Pic         string `json:"pic"                    dc:"图片"`
	Status      int    `json:"status"      dc:"状态"`
	Sort        int    `json:"sort"      dc:"排序"`
}

// FoodCuisineEditInp 修改/新增餐厅菜系
type FoodCuisineEditInp struct {
	entity.FoodCuisine
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言套餐名称"`
}

func (in *FoodCuisineEditInp) Filter(ctx context.Context) (err error) {

	return
}

type FoodCuisineEditModel struct{}

// FoodCuisineDeleteInp 删除餐厅菜系
type FoodCuisineDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodCuisineDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCuisineDeleteModel struct{}

// FoodCuisineViewInp 获取指定餐厅菜系信息
type FoodCuisineViewInp struct {
	Id         int  `json:"id" v:"required#id不能为空" dc:"id"`
	IsLanguage bool `json:"isLanguage" dc:"是否获取多语言数据"`
}

func (in *FoodCuisineViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCuisineViewModel struct {
	entity.FoodCuisine
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"房型名称"   orm:"with:uuid=cuisine_name"`
}

// FoodCuisineListInp 获取餐厅菜系列表
type FoodCuisineListInp struct {
	input_form.PageReq
	CuisineName string `json:"cuisineName" dc:"菜系名称"`
	Status      int    `json:"status"      dc:"状态"`
}

func (in *FoodCuisineListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCuisineListModel struct {
	Id          int         `json:"id"          dc:"id"`
	CuisineName string      `json:"cuisineName" dc:"菜系名称"`
	Status      int         `json:"status"      dc:"状态"`
	Sort        int         `json:"sort"      dc:"排序"`
	CreateAt    *gtime.Time `json:"createAt"    dc:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"    dc:"更新时间"`
}

// FoodCuisineMaxSortInp 获取餐厅菜系最大排序
type FoodCuisineMaxSortInp struct{}

func (in *FoodCuisineMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCuisineMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// FoodCuisineStatusInp 更新餐厅菜系状态
type FoodCuisineStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodCuisineStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodCuisineStatusModel struct{}
