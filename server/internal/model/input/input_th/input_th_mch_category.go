package input_th

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchCategoryUpdateFields 修改商户分类字段过滤
type ThMchCategoryUpdateFields struct {
	Name string `json:"name"     dc:"分类名称"`
	Sort int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// ThMchCategoryInsertFields 新增商户分类字段过滤
type ThMchCategoryInsertFields struct {
	Name string `json:"name"     dc:"分类名称"`
	Sort int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// ThMchCategoryEditInp 修改/新增商户分类
type ThMchCategoryEditInp struct {
	entity.ThMchCategory
}

func (in *ThMchCategoryEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}
	return
}

type ThMchCategoryEditModel struct{}

// ThMchCategoryDeleteInp 删除商户分类
type ThMchCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThMchCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchCategoryDeleteModel struct{}

// ThMchCategoryViewInp 获取指定商户分类信息
type ThMchCategoryViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThMchCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchCategoryViewModel struct {
	entity.ThMchCategory
}

// ThMchCategoryListInp 获取商户分类列表
type ThMchCategoryListInp struct {
	input_form.PageReq
}

// ThMchCategoryAllInp 获取商户分类全部列表
type ThMchCategoryAllInp struct {
}

func (in *ThMchCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchCategoryListModel struct {
	Id        int         `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"分类名称"`
	Sort      int         `json:"sort"      dc:"排序(越大越靠前)"`
	Status    uint        `json:"status"    dc:"1、启用 2、禁用"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}

type ThMchCategoryAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"分类名称"`
}

// ThMchCategorySwitchInp 更新状态
type ThMchCategorySwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *ThMchCategorySwitchInp) Filter(ctx context.Context) (err error) {
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

type ThMchCategorySwitchModel struct{}
