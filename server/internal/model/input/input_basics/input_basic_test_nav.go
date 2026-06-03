// Package sysin

package input_basics

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTestNavUpdateFields 修改导航字段过滤
type PmsTestNavUpdateFields struct {
	Name    string `json:"name"      dc:"导航名称"`
	Image   string `json:"image"     dc:"图标"`
	AppLink string `json:"appLink"   dc:"app跳转链接"`
	WxLink  string `json:"wxLink"    dc:"微信跳转链接"`
	Sort    int    `json:"sort"      dc:"排序(越大越靠前)"`
	Status  uint   `json:"status"    dc:"状态1、启用 2、禁用"`
}

// PmsTestNavInsertFields 新增导航字段过滤
type PmsTestNavInsertFields struct {
	Name    string `json:"name"      dc:"导航名称"`
	Image   string `json:"image"     dc:"图标"`
	AppLink string `json:"appLink"   dc:"app跳转链接"`
	WxLink  string `json:"wxLink"    dc:"微信跳转链接"`
	Sort    int    `json:"sort"      dc:"排序(越大越靠前)"`
	Status  uint   `json:"status"    dc:"状态1、启用 2、禁用"`
}

// PmsTestNavEditInp 修改/新增导航
type PmsTestNavEditInp struct {
	entity.PmsTestNav
}

func (in *PmsTestNavEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsTestNavEditModel struct{}

// PmsTestNavDeleteInp 删除导航
type PmsTestNavDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsTestNavDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTestNavDeleteModel struct{}

// PmsTestNavViewInp 获取指定导航信息
type PmsTestNavViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsTestNavViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTestNavViewModel struct {
	entity.PmsTestNav
}

// PmsTestNavListInp 获取导航列表
type PmsTestNavListInp struct {
	input_form.PageReq
	Status int `json:"status" dc:"状态: 1、启用 2、禁用"`
}

func (in *PmsTestNavListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTestNavListModel struct {
	Id       int         `json:"id"       dc:"id"`
	Name     string      `json:"name"        dc:"名称"`
	Image    string      `json:"image"     dc:"图标"`
	AppLink  string      `json:"appLink"   dc:"app跳转链接"`
	WxLink   string      `json:"wxLink"    dc:"微信跳转链接"`
	Sort     int         `json:"sort"      dc:"排序(越大越靠前)"`
	Status   uint        `json:"status"    dc:"状态1、启用 2、禁用"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}

type PmsTestNavApiListModel struct {
	Id      int    `json:"id"       dc:"id"`
	Name    string `json:"name"        dc:"名称"`
	Image   string `json:"image"     dc:"图标"`
	AppLink string `json:"appLink"   dc:"app跳转链接"`
	WxLink  string `json:"wxLink"    dc:"微信跳转链接"`
}

// PmsTestNavAllInp 获取导航全部列表
type PmsTestNavAllInp struct {
	Status int `json:"status" dc:"状态1、启用 2、禁用"`
}

type PmsTestNavAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"导航名称"`
}

// PmsTestNavStatusInp 更新导航状态
type PmsTestNavStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsTestNavStatusInp) Filter(ctx context.Context) (err error) {
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

type PmsTestNavStatusModel struct{}

// PmsTestNavSwitchInp 更新状态
type PmsTestNavSwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *PmsTestNavSwitchInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.Status) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type PmsTestNavSwitchModel struct{}

// PmsTestNavSortInp 编辑排序
type PmsTestNavSortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *PmsTestNavSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsTestNavSortModel struct{}
