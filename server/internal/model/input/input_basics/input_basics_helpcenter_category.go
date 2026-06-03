package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenterCategoryUpdateFields 修改帮助中心分类字段过滤
type PmsHelpcenterCategoryUpdateFields struct {
	Name     string `json:"name"     dc:"分类名称"`
	Language string `json:"language" dc:"语言"`
	Sort     int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// PmsHelpcenterCategoryInsertFields 新增帮助中心分类字段过滤
type PmsHelpcenterCategoryInsertFields struct {
	Name     string `json:"name"     dc:"分类名称"`
	Language string `json:"language" dc:"语言"`
	Sort     int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// PmsHelpcenterCategoryEditInp 修改/新增帮助中心分类
type PmsHelpcenterCategoryEditInp struct {
	entity.PmsHelpcenterCategory
}

func (in *PmsHelpcenterCategoryEditInp) Filter(ctx context.Context) (err error) {
	if in.Name == "" {
		err = gerror.New("名称不能为空")
		return
	}
	if in.Language == "" {
		err = gerror.New("请选择语言")
		return
	}
	return
}

type PmsHelpcenterCategoryEditModel struct{}

// PmsHelpcenterCategoryDeleteInp 删除帮助中心分类
type PmsHelpcenterCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsHelpcenterCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterCategoryDeleteModel struct{}

// PmsHelpcenterCategoryViewInp 获取指定帮助中心分类信息
type PmsHelpcenterCategoryViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsHelpcenterCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterCategoryViewModel struct {
	entity.PmsHelpcenterCategory
}

// PmsHelpcenterCategoryListInp 获取帮助中心分类列表
type PmsHelpcenterCategoryListInp struct {
	input_form.PageReq
	Language string `json:"language"                     dc:"语言"`
}

// PmsHelpcenterCategoryAllInp 获取帮助中心分类全部列表
type PmsHelpcenterCategoryAllInp struct {
	Language string `json:"language"                     dc:"语言"`
}

func (in *PmsHelpcenterCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterCategoryListModel struct {
	Id        int         `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"分类名称"`
	Language  string      `json:"language"  dc:"语言"`
	Sort      int         `json:"sort"      dc:"排序(越大越靠前)"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}

type PmsHelpcenterCategoryAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"分类名称"`
}
