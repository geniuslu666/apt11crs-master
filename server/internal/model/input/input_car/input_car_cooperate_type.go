// Package sysin

package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarCooperateTypeUpdateFields 修改订餐合作类型字段过滤
type CarCooperateTypeUpdateFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
	IsThird  int    `json:"isThird"  dc:"是否第三方: 1、是 2、否"`
}

// CarCooperateTypeInsertFields 新增订餐合作类型字段过滤
type CarCooperateTypeInsertFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
	IsThird  int    `json:"isThird"  dc:"是否第三方: 1、是 2、否"`
}

// CarCooperateTypeEditInp 修改/新增订餐合作类型
type CarCooperateTypeEditInp struct {
	entity.CarCooperateType
}

func (in *CarCooperateTypeEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarCooperateTypeEditModel struct{}

// CarCooperateTypeDeleteInp 删除订餐合作类型
type CarCooperateTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarCooperateTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCooperateTypeDeleteModel struct{}

// CarCooperateTypeViewInp 获取指定订餐合作类型信息
type CarCooperateTypeViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarCooperateTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCooperateTypeViewModel struct {
	entity.CarCooperateType
}

// CarCooperateTypeListInp 获取订餐合作类型列表
type CarCooperateTypeListInp struct {
	input_form.PageReq
	TypeName string `json:"typeName"     dc:"合作类型"`
	Status   int    `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *CarCooperateTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarCooperateTypeListModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	IsThird  int         `json:"isThird"  dc:"是否第三方: 1、是 2、否"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// CarCooperateTypeExportModel 导出订餐合作类型
type CarCooperateTypeExportModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	IsThird  int         `json:"isThird"  dc:"是否第三方: 1、是 2、否"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// CarCooperateTypeStatusInp 更新订餐合作类型状态
type CarCooperateTypeStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *CarCooperateTypeStatusInp) Filter(ctx context.Context) (err error) {
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

type CarCooperateTypeStatusModel struct{}
