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

// IndexBannerUpdateFields 修改首页轮播字段过滤
type IndexBannerUpdateFields struct {
	Language     string `json:"language"     dc:"语言"`
	BannerImage  string `json:"bannerImage"  dc:"轮播图"`
	Chain        string `json:"chain"        dc:"内外联"`
	Path         string `json:"path"         dc:"链接内容"`
	BannerStatus int    `json:"bannerStatus" dc:"状态"`
	MinappStatus int    `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Sort         int    `json:"sort"       dc:"排序(越大越靠前)"`
	LinkOpenType string `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// IndexBannerInsertFields 新增首页轮播字段过滤
type IndexBannerInsertFields struct {
	Language     string `json:"language"     dc:"语言"`
	BannerImage  string `json:"bannerImage"  dc:"轮播图"`
	Chain        string `json:"chain"        dc:"内外联"`
	Path         string `json:"path"         dc:"链接内容"`
	BannerStatus int    `json:"bannerStatus" dc:"状态"`
	MinappStatus int    `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Sort         int    `json:"sort"       dc:"排序(越大越靠前)"`
	LinkOpenType int    `json:"linkOpenType"          dc:"外链打开方式 1-内部webview 2-外部浏览器"`
}

// IndexBannerEditInp 修改/新增首页轮播
type IndexBannerEditInp struct {
	entity.IndexBanner
}

func (in *IndexBannerEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsIndexBannerEditModel struct{}

// IndexBannerDeleteInp 删除首页轮播
type IndexBannerDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *IndexBannerDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsIndexBannerDeleteModel struct{}

// IndexBannerViewInp 获取指定首页轮播信息
type IndexBannerViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *IndexBannerViewInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexBannerViewModel struct {
	entity.IndexBanner
}

// IndexBannerListInp 获取首页轮播列表
type IndexBannerListInp struct {
	input_form.PageReq
	Id           int    `json:"id"           dc:"id"`
	Language     string `json:"language"     dc:"语言"`
	Model        string `json:"model"        dc:"模块"`
	Chain        string `json:"chain"        dc:"内外联"`
	BannerStatus int    `json:"bannerStatus" dc:"状态: 1、启用 2、禁用"`
	MinappStatus int    `json:"minappStatus" dc:"状态: 1、小程序显示 2、小程序不显示"`
}

func (in *IndexBannerListInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexBannerListModel struct {
	Id           int         `json:"id"           dc:"id"`
	Language     string      `json:"language"     dc:"语言"`
	BannerImage  string      `json:"bannerImage"  dc:"轮播图"`
	Model        string      `json:"model"        dc:"模块"`
	Chain        string      `json:"chain"        dc:"内外联"`
	Path         string      `json:"path"         dc:"链接内容"`
	BannerStatus int         `json:"bannerStatus" dc:"状态"`
	MinappStatus int         `json:"minappStatus" dc:"是否排除小程序显示 1-不排除 2-排除"`
	Sort         int         `json:"sort"       dc:"排序(越大越靠前)"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}

type IndexBannerApiListModel struct {
	Id           int         `json:"id"           dc:"id"`
	Language     string      `json:"language"     dc:"语言"`
	BannerImage  string      `json:"bannerImage"  dc:"轮播图"`
	Model        string      `json:"model"        dc:"模块"`
	Chain        string      `json:"chain"        dc:"内外联"`
	Path         string      `json:"path"         dc:"链接内容"`
	WxPath       string      `json:"wxPath"       dc:"微信链接内容"`
	AppPath      string      `json:"appPath"      dc:"app链接内容"`
	BannerStatus int         `json:"bannerStatus" dc:"状态"`
	CreateAt     *gtime.Time `json:"createAt"     dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     dc:"更新时间"`
}

// IndexBannerStatusInp 更新Banner状态
type IndexBannerStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *IndexBannerStatusInp) Filter(ctx context.Context) (err error) {
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

type IndexBannerStatusModel struct{}

// IndexBannerSwitchInp 更新状态
type IndexBannerSwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *IndexBannerSwitchInp) Filter(ctx context.Context) (err error) {
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

type IndexBannerSwitchModel struct{}

// IndexBannerSortInp 编辑排序
type IndexBannerSortInp struct {
	Id   interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
	Sort int         `json:"sort"        dc:"排序"`
}

func (in *IndexBannerSortInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexBannerSortModel struct{}

// IndexBannerMaxSortInp 获取Banner最大排序
type IndexBannerMaxSortInp struct{}

func (in *IndexBannerMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type IndexBannerMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}
