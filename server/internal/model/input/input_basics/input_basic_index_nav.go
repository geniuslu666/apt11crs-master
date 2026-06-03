// Package sysin

package input_basics

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_language"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsIndexNavUpdateFields 修改导航字段过滤
type PmsIndexNavUpdateFields struct {
	Name         string `json:"name"      dc:"导航名称(多语言)"`
	Tag          string `json:"tag"      dc:"标签"`
	Image        string `json:"image"     dc:"图标"`
	AppLink      string `json:"appLink"   dc:"app跳转链接"`
	WxLink       string `json:"wxLink"    dc:"微信跳转链接"`
	Sort         int    `json:"sort"      dc:"排序(越大越靠前)"`
	Status       uint   `json:"status"    dc:"状态1、启用 2、禁用"`
	MinappStatus int    `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Chain        string `json:"chain"        dc:"内外联"`
	LinkOpenType int    `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// PmsIndexNavInsertFields 新增导航字段过滤
type PmsIndexNavInsertFields struct {
	Name         string `json:"name"      dc:"导航名称(多语言)"`
	Tag          string `json:"tag"      dc:"标签"`
	Image        string `json:"image"     dc:"图标"`
	AppLink      string `json:"appLink"   dc:"app跳转链接"`
	WxLink       string `json:"wxLink"    dc:"微信跳转链接"`
	Sort         int    `json:"sort"      dc:"排序(越大越靠前)"`
	Status       uint   `json:"status"    dc:"状态1、启用 2、禁用"`
	MinappStatus int    `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Chain        string `json:"chain"        dc:"内外联"`
	LinkOpenType int    `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// PmsIndexNavEditInp 修改/新增导航
type PmsIndexNavEditInp struct {
	entity.PmsIndexNav
	NameLanguage input_language.LanguageModel `json:"nameLanguage"          dc:"多语言名称"`
	TagLanguage  input_language.LanguageModel `json:"tagLanguage"          dc:"多语言标签"`
}

func (in *PmsIndexNavEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsIndexNavEditModel struct{}

// PmsIndexNavDeleteInp 删除导航
type PmsIndexNavDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsIndexNavDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsIndexNavDeleteModel struct{}

// PmsIndexNavViewInp 获取指定导航信息
type PmsIndexNavViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsIndexNavViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsIndexNavViewModel struct {
	entity.PmsIndexNav
	NameLanguage []*input_hotel.LanguageType `json:"nameLanguage"         dc:"名称"   orm:"with:uuid=name"`
	TagLanguage  []*input_hotel.LanguageType `json:"tagLanguage"         dc:"标签"   orm:"with:uuid=tag"`
}

// PmsIndexNavListInp 获取导航列表
type PmsIndexNavListInp struct {
	input_form.PageReq
	Status int `json:"status" dc:"状态: 1、启用 2、禁用"`
}

func (in *PmsIndexNavListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsIndexNavListModel struct {
	Id           int         `json:"id"       dc:"id"`
	Name         string      `json:"name"        dc:"名称"`
	Tag          string      `json:"tag"        dc:"标签"`
	Image        string      `json:"image"     dc:"图标"`
	AppLink      string      `json:"appLink"   dc:"app跳转链接"`
	WxLink       string      `json:"wxLink"    dc:"微信跳转链接"`
	Sort         int         `json:"sort"      dc:"排序(越大越靠前)"`
	Status       uint        `json:"status"    dc:"状态1、启用 2、禁用"`
	MinappStatus int         `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Chain        string      `json:"chain"        dc:"内外联"`
	CreateAt     *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt" dc:"更新时间"`
}

type PmsIndexNavApiListModel struct {
	Id           int    `json:"id"       dc:"id"`
	Name         string `json:"name"        dc:"名称"`
	Tag          string `json:"tag"        dc:"标签"`
	Image        string `json:"image"     dc:"图标"`
	Chain        string `json:"chain"        dc:"内外联"`
	AppLink      string `json:"appLink"   dc:"app跳转链接"`
	WxLink       string `json:"wxLink"    dc:"微信跳转链接"`
	LinkOpenType int    `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// PmsIndexNavAllInp 获取导航全部列表
type PmsIndexNavAllInp struct {
	Status       int `json:"status" dc:"状态1、启用 2、禁用"`
	MinappStatus int `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
}

type PmsIndexNavAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"导航名称"`
}

// PmsIndexNavStatusInp 更新导航状态
type PmsIndexNavStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsIndexNavStatusInp) Filter(ctx context.Context) (err error) {
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

type PmsIndexNavStatusModel struct{}

// PmsIndexNavSwitchInp 更新状态
type PmsIndexNavSwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *PmsIndexNavSwitchInp) Filter(ctx context.Context) (err error) {
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

type PmsIndexNavSwitchModel struct{}

// PmsIndexNavSortInp 编辑排序
type PmsIndexNavSortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *PmsIndexNavSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsIndexNavSortModel struct{}
