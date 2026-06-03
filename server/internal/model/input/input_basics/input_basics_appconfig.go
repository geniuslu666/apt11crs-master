package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppconfigUpdateFields 修改APP配置字段过滤
type PmsAppconfigUpdateFields struct {
	Name     string `json:"name"     dc:"名称"`
	Key      string `json:"key"      dc:"配置项"`
	Value    string `json:"value"    dc:"配置信息"`
	Language string `json:"language" dc:"语言"`
}

// PmsAppconfigInsertFields 新增APP配置字段过滤
type PmsAppconfigInsertFields struct {
	Name     string `json:"name"     dc:"名称"`
	Key      string `json:"key"      dc:"配置项"`
	Value    string `json:"value"    dc:"配置信息"`
	Language string `json:"language" dc:"语言"`
}

// PmsAppconfigEditInp 修改/新增APP配置
type PmsAppconfigEditInp struct {
	entity.PmsAppconfig
}

func (in *PmsAppconfigEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsAppconfigEditModel struct{}

// PmsAppconfigDeleteInp 删除APP配置
type PmsAppconfigDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsAppconfigDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppconfigDeleteModel struct{}

// PmsAppconfigViewInp 获取指定APP配置信息
type PmsAppconfigViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsAppconfigViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppconfigViewModel struct {
	entity.PmsAppconfig
}

// PmsAppconfigListInp 获取APP配置列表
type PmsAppconfigListInp struct {
	input_form.PageReq
	Id        int           `json:"id"        dc:"id"`
	Name      string        `json:"name"      dc:"名称"`
	Key       string        `json:"key"       dc:"配置项"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"created_at"`
}

func (in *PmsAppconfigListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsAppconfigListModel struct {
	Id        int         `json:"id"        dc:"id"`
	Name      string      `json:"name"      dc:"名称"`
	Key       string      `json:"key"       dc:"配置项"`
	Value     string      `json:"value"     dc:"配置信息"`
	Language  string      `json:"language"  dc:"语言"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}
