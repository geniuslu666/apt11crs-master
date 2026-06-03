package input_food

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// FoodMaintenanceUpdateFields 修改Maintenance字段过滤
type FoodMaintenanceUpdateFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
}

// FoodMaintenanceInsertFields 新增Maintenance字段过滤
type FoodMaintenanceInsertFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
}

// FoodMaintenanceEditInp 修改/新增Maintenance
type FoodMaintenanceEditInp struct {
	entity.FoodMaintenance
}

func (in *FoodMaintenanceEditInp) Filter(ctx context.Context) (err error) {

	return
}

type FoodMaintenanceEditModel struct{}

// FoodMaintenanceDeleteInp 删除Maintenance
type FoodMaintenanceDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodMaintenanceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodMaintenanceDeleteModel struct{}

// FoodMaintenanceViewInp 获取指定Maintenance信息
type FoodMaintenanceViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodMaintenanceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodMaintenanceViewModel struct {
	entity.FoodMaintenance
}

// FoodMaintenanceListInp 获取Maintenance列表
type FoodMaintenanceListInp struct {
	input_form.PageReq
}

func (in *FoodMaintenanceListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodMaintenanceListModel struct {
	Id        int64       `json:"id"              dc:"id"`
	Language  string      `json:"language"    dc:"语言"`
	IsDefault uint        `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string      `json:"content"    dc:"内容"`
	CreateAt  *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// FoodMaintenanceLanguageListInp 获取Banner最大排序
type FoodMaintenanceLanguageListInp struct{}

func (in *FoodMaintenanceLanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodMaintenanceLanguageListModel struct {
	Language string `json:"language"    dc:"语言"`
}
