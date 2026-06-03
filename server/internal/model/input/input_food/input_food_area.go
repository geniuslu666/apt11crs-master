// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodAreaUpdateFields 修改区域管理字段过滤
type FoodAreaUpdateFields struct {
	Pid        int64       `json:"pid"        description:"上级角色ID"`
	Level      int         `json:"level"      description:"关系树等级"`
	Tree       string      `json:"tree"       description:"关系树"`
	AreaName   string      `json:"areaName"   dc:"区域名称"`
	AreaStatus int         `json:"areaStatus" dc:"1、启用 2、禁用"`
	CreateAt   *gtime.Time `json:"createAt"   dc:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   dc:"更新时间"`
}

// FoodAreaInsertFields 新增区域管理字段过滤
type FoodAreaInsertFields struct {
	Pid        int64       `json:"pid"        description:"上级角色ID"`
	Level      int         `json:"level"      description:"关系树等级"`
	Tree       string      `json:"tree"       description:"关系树"`
	AreaName   string      `json:"areaName"   dc:"区域名称"`
	AreaStatus int         `json:"areaStatus" dc:"1、启用 2、禁用"`
	CreateAt   *gtime.Time `json:"createAt"   dc:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   dc:"更新时间"`
}

// FoodAreaEditInp 修改/新增区域管理
type FoodAreaEditInp struct {
	entity.FoodArea
}

func (in *FoodAreaEditInp) Filter(ctx context.Context) (err error) {
	// 验证区域名称
	if err := g.Validator().Rules("required").Data(in.AreaName).Messages("区域名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FoodAreaEditModel struct{}

// FoodAreaDeleteInp 删除区域管理
type FoodAreaDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodAreaDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodAreaDeleteModel struct{}

// FoodAreaViewInp 获取指定区域管理信息
type FoodAreaViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodAreaViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodAreaViewModel struct {
	entity.FoodArea
}

// FoodAreaListInp 获取区域管理列表
type FoodAreaListInp struct {
	input_form.PageReq
	Level  int `json:"level"                     dc:"等级"`
	Status int `json:"status" dc:"状态1、启用 2、禁用"`
}

func (in *FoodAreaListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodAreaListModel struct {
	List []*AreaTree `json:"list"`
}

type AreaTree struct {
	entity.FoodArea
	Children []*AreaTree `json:"children"  dc:"子级"`
}

// FoodAreaStatusInp 更新区域状态
type FoodAreaStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *FoodAreaStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodAreaStatusModel struct{}
