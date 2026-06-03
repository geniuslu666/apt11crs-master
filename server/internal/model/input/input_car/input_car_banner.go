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

// CarBannerUpdateFields 修改Banner字段过滤
type CarBannerUpdateFields struct {
	Language    string `json:"language"    dc:"语言"`
	BannerImage string `json:"bannerImage" dc:"轮播图"`
	Status      int    `json:"status"          dc:"状态"`
	Sort        int    `json:"sort"       dc:"排序(越大越靠前)"`
}

// CarBannerInsertFields 新增Banner字段过滤
type CarBannerInsertFields struct {
	Language    string `json:"language"    dc:"语言"`
	BannerImage string `json:"bannerImage" dc:"轮播图"`
	Status      int    `json:"status"          dc:"状态"`
	Sort        int    `json:"sort"       dc:"排序(越大越靠前)"`
}

// CarBannerEditInp 修改/新增Banner
type CarBannerEditInp struct {
	entity.CarBanner
}

func (in *CarBannerEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarBannerEditModel struct{}

// CarBannerDeleteInp 删除Banner
type CarBannerDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarBannerDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarBannerDeleteModel struct{}

// CarBannerViewInp 获取指定Banner信息
type CarBannerViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarBannerViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarBannerViewModel struct {
	entity.CarBanner
}

// CarBannerListInp 获取Banner列表
type CarBannerListInp struct {
	input_form.PageReq
}

func (in *CarBannerListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarBannerListModel struct {
	Id          int64       `json:"id"              dc:"id"`
	Language    string      `json:"language"    dc:"语言"`
	BannerImage string      `json:"bannerImage" dc:"轮播图"`
	Status      int         `json:"status"          dc:"状态"`
	Sort        int         `json:"sort"       dc:"排序(越大越靠前)"`
	CreateAt    *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt    *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// CarBannerMaxSortInp 获取Banner最大排序
type CarBannerMaxSortInp struct{}

func (in *CarBannerMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CarBannerMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CarBannerStatusInp 更新Banner状态
type CarBannerStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CarBannerStatusInp) Filter(ctx context.Context) (err error) {
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

type CarBannerStatusModel struct{}
