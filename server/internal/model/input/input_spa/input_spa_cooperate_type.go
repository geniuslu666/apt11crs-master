package input_spa

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaCooperateTypeUpdateFields 修改按摩营业类型字段过滤
type SpaCooperateTypeUpdateFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
}

// SpaCooperateTypeInsertFields 新增按摩营业类型字段过滤
type SpaCooperateTypeInsertFields struct {
	TypeName string `json:"typeName" dc:"合作类型"`
	Status   int    `json:"status"   dc:"状态1、启用 2、禁用"`
}

// SpaCooperateTypeEditInp 修改/新增按摩营业类型
type SpaCooperateTypeEditInp struct {
	entity.SpaCooperateType
}

func (in *SpaCooperateTypeEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaCooperateTypeEditModel struct{}

// SpaCooperateTypeDeleteInp 删除按摩营业类型
type SpaCooperateTypeDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaCooperateTypeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaCooperateTypeDeleteModel struct{}

// SpaCooperateTypeViewInp 获取指定按摩营业类型信息
type SpaCooperateTypeViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaCooperateTypeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaCooperateTypeViewModel struct {
	entity.SpaCooperateType
}

// SpaCooperateTypeListInp 获取按摩营业类型列表
type SpaCooperateTypeListInp struct {
	input_form.PageReq
	TypeName string `json:"typeName"     dc:"合作类型"`
	Status   int    `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *SpaCooperateTypeListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaCooperateTypeListModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// SpaCooperateTypeExportModel 导出按摩营业类型
type SpaCooperateTypeExportModel struct {
	Id       int         `json:"id"       dc:"id"`
	TypeName string      `json:"typeName" dc:"合作类型"`
	Status   int         `json:"status"   dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

// SpaCooperateTypeStatusInp 更新按摩营业类型状态
type SpaCooperateTypeStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *SpaCooperateTypeStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaCooperateTypeStatusModel struct{}
