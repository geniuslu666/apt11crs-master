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

// SpaBannerUpdateFields 修改Banner字段过滤
type SpaBannerUpdateFields struct {
	Language    string `json:"language"    dc:"语言"`
	BannerImage string `json:"bannerImage" dc:"轮播图"`
	Status      int    `json:"status"          dc:"状态"`
	Sort        int    `json:"sort"       dc:"排序(越大越靠前)"`
}

// SpaBannerInsertFields 新增Banner字段过滤
type SpaBannerInsertFields struct {
	Language    string `json:"language"    dc:"语言"`
	BannerImage string `json:"bannerImage" dc:"轮播图"`
	Status      int    `json:"status"          dc:"状态"`
	Sort        int    `json:"sort"       dc:"排序(越大越靠前)"`
}

// SpaBannerEditInp 修改/新增Banner
type SpaBannerEditInp struct {
	entity.SpaBanner
}

func (in *SpaBannerEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaBannerEditModel struct{}

// SpaBannerDeleteInp 删除Banner
type SpaBannerDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaBannerDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaBannerDeleteModel struct{}

// SpaBannerViewInp 获取指定Banner信息
type SpaBannerViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaBannerViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaBannerViewModel struct {
	entity.SpaBanner
}

// SpaBannerListInp 获取Banner列表
type SpaBannerListInp struct {
	input_form.PageReq
}

func (in *SpaBannerListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaBannerListModel struct {
	Id          int64       `json:"id"              dc:"id"`
	Language    string      `json:"language"    dc:"语言"`
	BannerImage string      `json:"bannerImage" dc:"轮播图"`
	Status      int         `json:"status"          dc:"状态"`
	Sort        int         `json:"sort"       dc:"排序(越大越靠前)"`
	CreateAt    *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// SpaBannerMaxSortInp 获取Banner最大排序
type SpaBannerMaxSortInp struct{}

func (in *SpaBannerMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaBannerMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// SpaBannerStatusInp 更新Banner状态
type SpaBannerStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SpaBannerStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaBannerStatusModel struct{}
