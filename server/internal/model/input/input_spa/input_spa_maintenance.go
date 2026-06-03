package input_spa

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// SpaMaintenanceUpdateFields 修改Maintenance字段过滤
type SpaMaintenanceUpdateFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
}

// SpaMaintenanceInsertFields 新增Maintenance字段过滤
type SpaMaintenanceInsertFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
}

// SpaMaintenanceEditInp 修改/新增Maintenance
type SpaMaintenanceEditInp struct {
	entity.SpaMaintenance
}

func (in *SpaMaintenanceEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaMaintenanceEditModel struct{}

// SpaMaintenanceDeleteInp 删除Maintenance
type SpaMaintenanceDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaMaintenanceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaMaintenanceDeleteModel struct{}

// SpaMaintenanceViewInp 获取指定Maintenance信息
type SpaMaintenanceViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaMaintenanceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaMaintenanceViewModel struct {
	entity.SpaMaintenance
}

// SpaMaintenanceListInp 获取Maintenance列表
type SpaMaintenanceListInp struct {
	input_form.PageReq
}

func (in *SpaMaintenanceListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaMaintenanceListModel struct {
	Id        int64       `json:"id"              dc:"id"`
	Language  string      `json:"language"    dc:"语言"`
	IsDefault uint        `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string      `json:"content"    dc:"内容"`
	CreateAt  *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// SpaMaintenanceLanguageListInp 获取Banner最大排序
type SpaMaintenanceLanguageListInp struct{}

func (in *SpaMaintenanceLanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaMaintenanceLanguageListModel struct {
	Language string `json:"language"    dc:"语言"`
}
