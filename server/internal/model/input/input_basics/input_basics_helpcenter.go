package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsHelpcenterUpdateFields 修改帮助中心字段过滤
type PmsHelpcenterUpdateFields struct {
	Language   string `json:"language"   dc:"语言"`
	CategoryId int    `json:"categoryId" dc:"分类ID"`
	Title      string `json:"title"      dc:"标题"`
	Content    string `json:"content"    dc:"内容"`
	Sort       int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// PmsHelpcenterInsertFields 新增帮助中心字段过滤
type PmsHelpcenterInsertFields struct {
	Language   string `json:"language"   dc:"语言"`
	CategoryId int    `json:"categoryId" dc:"分类ID"`
	Title      string `json:"title"      dc:"标题"`
	Content    string `json:"content"    dc:"内容"`
	Sort       int    `json:"sort"     dc:"排序(越大越靠前)"`
}

// PmsHelpcenterEditInp 修改/新增帮助中心
type PmsHelpcenterEditInp struct {
	entity.PmsHelpcenter
}

func (in *PmsHelpcenterEditInp) Filter(ctx context.Context) (err error) {
	if in.Language == "" {
		err = gerror.New("请选择语言")
		return
	}
	if in.CategoryId == 0 {
		err = gerror.New("请选择分类")
		return
	}
	if in.Title == "" {
		err = gerror.New("标题不能为空")
		return
	}
	if in.Content == "" {
		err = gerror.New("内容不能为空")
		return
	}
	return
}

type PmsHelpcenterEditModel struct{}

// PmsHelpcenterDeleteInp 删除帮助中心
type PmsHelpcenterDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsHelpcenterDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterDeleteModel struct{}

// PmsHelpcenterViewInp 获取指定帮助中心信息
type PmsHelpcenterViewInp struct {
	Id int `json:"id" v:"required#id_cannot_be_empty" dc:"id"`
}

func (in *PmsHelpcenterViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterViewModel struct {
	entity.PmsHelpcenter
	PmsHelpcenterCategoryName string `json:"pmsHelpcenterCategoryName" dc:"分类名称"`
}

// PmsHelpcenterListInp 获取帮助中心列表
type PmsHelpcenterListInp struct {
	input_form.PageReq
	Title                     string        `json:"title"                     dc:"标题"`
	CreatedAt                 []*gtime.Time `json:"createdAt"                 dc:"created_at"`
	CategoryId                uint          `json:"categoryId"                dc:"分类ID"`
	Language                  string        `json:"language"                  dc:"语言"`
	Keyword                   string        `json:"keyword"                dc:"关键字"`
	PmsHelpcenterCategoryName string        `json:"pmsHelpcenterCategoryName" dc:"分类名称"`
}

func (in *PmsHelpcenterListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsHelpcenterListModel struct {
	Id                        int         `json:"id"                        dc:"id"`
	Language                  string      `json:"language"                  dc:"语言"`
	CategoryId                int         `json:"categoryId"                dc:"分类ID"`
	Title                     string      `json:"title"                     dc:"标题"`
	Sort                      int         `json:"sort"      dc:"排序(越大越靠前)"`
	CreatedAt                 *gtime.Time `json:"createdAt"                 dc:"created_at"`
	UpdatedAt                 *gtime.Time `json:"updatedAt"                 dc:"updated_at"`
	PmsHelpcenterCategoryName string      `json:"pmsHelpcenterCategoryName" dc:"分类名称"`
}
