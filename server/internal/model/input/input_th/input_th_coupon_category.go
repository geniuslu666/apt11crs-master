package input_th

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
)

// ThCouponCategoryUpdateFields 修改礼品券分类字段过滤
type ThCouponCategoryUpdateFields struct {
	Name string `json:"name"     dc:"分类名称"`
	Sort int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// ThCouponCategoryInsertFields 新增礼品券分类字段过滤
type ThCouponCategoryInsertFields struct {
	Name string `json:"name"     dc:"分类名称"`
	Sort int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// ThCouponCategoryEditInp 修改/新增礼品券分类
type ThCouponCategoryEditInp struct {
	entity.ThCouponCategory
}

func (in *ThCouponCategoryEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}
	return
}

type ThCouponCategoryEditModel struct{}

// ThCouponCategoryDeleteInp 删除礼品券分类
type ThCouponCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThCouponCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponCategoryDeleteModel struct{}

// ThCouponCategoryViewInp 获取指定礼品券分类信息
type ThCouponCategoryViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThCouponCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponCategoryViewModel struct {
	entity.ThCouponCategory
}

// ThCouponCategoryListInp 获取礼品券分类列表
type ThCouponCategoryListInp struct {
	input_form.PageReq
}

// ThCouponCategoryAllInp 获取礼品券分类全部列表
type ThCouponCategoryAllInp struct {
}

func (in *ThCouponCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThCouponCategoryListModel struct {
	Id        int         `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"分类名称"`
	Sort      int         `json:"sort"      dc:"排序(越大越靠前)"`
	Status    uint        `json:"status"    dc:"1、启用 2、禁用"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}

type ThCouponCategoryAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"分类名称"`
}

// ThCouponCategorySwitchInp 更新状态
type ThCouponCategorySwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *ThCouponCategorySwitchInp) Filter(ctx context.Context) (err error) {
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

type ThCouponCategorySwitchModel struct{}
