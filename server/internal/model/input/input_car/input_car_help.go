package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// CarHelpUpdateFields 修改Help字段过滤
type CarHelpUpdateFields struct {
	Language string `json:"language"    dc:"语言"`
	Image    string `json:"image" dc:"主图"`
	Content  string `json:"content"    dc:"内容"`
}

// CarHelpInsertFields 新增Help字段过滤
type CarHelpInsertFields struct {
	Language string `json:"language"    dc:"语言"`
	Image    string `json:"image" dc:"主图"`
	Content  string `json:"content"    dc:"内容"`
}

// CarHelpEditInp 修改/新增Help
type CarHelpEditInp struct {
	entity.CarHelp
}

func (in *CarHelpEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarHelpEditModel struct{}

// CarHelpDeleteInp 删除Help
type CarHelpDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarHelpDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarHelpDeleteModel struct{}

// CarHelpViewInp 获取指定Help信息
type CarHelpViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarHelpViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarHelpViewModel struct {
	entity.CarHelp
}

// CarHelpListInp 获取Help列表
type CarHelpListInp struct {
	input_form.PageReq
}

func (in *CarHelpListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarHelpListModel struct {
	Id       int64       `json:"id"              dc:"id"`
	Language string      `json:"language"    dc:"语言"`
	Image    string      `json:"image" dc:"主图"`
	Content  string      `json:"content"    dc:"内容"`
	CreateAt *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// CarHelpLanguageListInp 获取Banner最大排序
type CarHelpLanguageListInp struct{}

func (in *CarHelpLanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarHelpLanguageListModel struct {
	Language string `json:"language"    dc:"语言"`
}
