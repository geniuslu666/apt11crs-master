package input_language

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LanguageModel struct {
	EN   string `json:"en"`
	ZH   string `json:"zh"`
	JA   string `json:"ja"`
	KO   string `json:"ko"`
	ZHCN string `json:"zh_CN"`
}

type LoadLanguage struct {
	Uuid string `json:"uuid"`
	Tag  string `json:"tag"`
	Type string `json:"type"`
	Key  string `json:"key"`
}

// PmsLanguageUpdateFields 修改语言字典字段过滤
type PmsLanguageUpdateFields struct {
	Uuid     string `json:"uuid"     dc:"标签ID"`
	Tag      string `json:"tag"      dc:"语言标签"`
	Type     string `json:"type"     dc:"类型"`
	Key      string `json:"key"      dc:"标识"`
	Language string `json:"language" dc:"语言"`
	Content  string `json:"content"  dc:"内容"`
}

// PmsLanguageInsertFields 新增语言字典字段过滤
type PmsLanguageInsertFields struct {
	Uuid     string `json:"uuid"     dc:"标签ID"`
	Tag      string `json:"tag"      dc:"语言标签"`
	Type     string `json:"type"     dc:"类型"`
	Key      string `json:"key"      dc:"标识"`
	Language string `json:"language" dc:"语言"`
	Content  string `json:"content"  dc:"内容"`
}

// PmsLanguageEditInp 修改/新增语言字典
type PmsLanguageEditInp struct {
	entity.PmsLanguage
}

func (in *PmsLanguageEditInp) Filter(ctx context.Context) (err error) {
	// 验证语言标签
	if err := g.Validator().Rules("required").Data(in.Tag).Messages("语言标签不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证类型
	if err := g.Validator().Rules("required").Data(in.Type).Messages("类型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:table,manage,mobile").Data(in.Type).Messages("类型值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证标识
	if err := g.Validator().Rules("required").Data(in.Key).Messages("标识不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证语言
	if err := g.Validator().Rules("required").Data(in.Language).Messages("语言不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:en,zh,ja,ko").Data(in.Language).Messages("语言值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsLanguageEditModel struct{}

// PmsLanguageDeleteInp 删除语言字典
type PmsLanguageDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsLanguageDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsLanguageDeleteModel struct{}

// PmsLanguageViewInp 获取指定语言字典信息
type PmsLanguageViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsLanguageViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsLanguageViewModel struct {
	entity.PmsLanguage
}

// PmsLanguageListInp 获取语言字典列表
type PmsLanguageListInp struct {
	input_form.PageReq
	Id       int    `json:"id"       dc:"id"`
	Uuid     string `json:"uuid"     dc:"标签ID"`
	Tag      string `json:"tag"      dc:"语言标签"`
	Type     string `json:"type"     dc:"类型"`
	Language string `json:"language" dc:"语言"`
	Content  string `json:"content"  dc:"内容"`
}

func (in *PmsLanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsLanguageListModel struct {
	Id       int    `json:"id"       dc:"id"`
	Tag      string `json:"tag"      dc:"语言标签"`
	Type     string `json:"type"     dc:"类型"`
	Key      string `json:"key"      dc:"标识"`
	Language string `json:"language" dc:"语言"`
	Content  string `json:"content"  dc:"内容"`
}

// PmsLanguageExportModel 导出语言字典
type PmsLanguageExportModel struct {
	Id       int         `json:"id"       dc:"id"`
	Uuid     string      `json:"uuid"     dc:"标签ID"`
	Tag      string      `json:"tag"      dc:"语言标签"`
	Type     string      `json:"type"     dc:"类型"`
	Key      string      `json:"key"      dc:"标识"`
	Language string      `json:"language" dc:"语言"`
	Content  string      `json:"content"  dc:"内容"`
	CreateAt *gtime.Time `json:"createAt" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" dc:"更新时间"`
}
