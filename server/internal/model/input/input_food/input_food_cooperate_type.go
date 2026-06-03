// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodCooperateTypeUpdateFields 修改订餐合作类型字段过滤
type FoodCooperateTypeUpdateFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
}

// FoodCooperateTypeInsertFields 新增订餐合作类型字段过滤
type FoodCooperateTypeInsertFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
}

// FoodCooperateTypeEditInp 修改/新增订餐合作类型
type FoodCooperateTypeEditInp struct {
	entity.FoodCooperateType
}

func (in *FoodCooperateTypeEditInp) Filter(ctx context.Context) (err error) {

	return
}

type FoodCooperateTypeEditModel struct{}

// FoodCooperateTypeDeleteInp 删除订餐合作类型
type FoodCooperateTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodCooperateTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCooperateTypeDeleteModel struct{}

// FoodCooperateTypeViewInp 获取指定订餐合作类型信息
type FoodCooperateTypeViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodCooperateTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCooperateTypeViewModel struct {
	entity.FoodCooperateType
}

// FoodCooperateTypeListInp 获取订餐合作类型列表
type FoodCooperateTypeListInp struct {
	input_form.PageReq
	TypeName string `json:"typeName"     dc:"合作类型"`
	Status   int    `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *FoodCooperateTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodCooperateTypeListModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// FoodCooperateTypeExportModel 导出订餐合作类型
type FoodCooperateTypeExportModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// FoodCooperateTypeStatusInp 更新订餐合作类型状态
type FoodCooperateTypeStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodCooperateTypeStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodCooperateTypeStatusModel struct{}
