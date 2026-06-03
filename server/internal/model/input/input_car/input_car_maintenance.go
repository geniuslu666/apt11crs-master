package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// CarMaintenanceUpdateFields 修改Maintenance字段过滤
type CarMaintenanceUpdateFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
	Type      uint   `json:"type"    dc:"类型 1-全部 2-接送机 3-包机"`
}

// CarMaintenanceInsertFields 新增Maintenance字段过滤
type CarMaintenanceInsertFields struct {
	Language  string `json:"language"    dc:"语言"`
	IsDefault uint   `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string `json:"content"    dc:"内容"`
	Type      uint   `json:"type"    dc:"类型 1-全部 2-接送机 3-包机"`
}

// CarMaintenanceEditInp 修改/新增Maintenance
type CarMaintenanceEditInp struct {
	entity.CarMaintenance
}

func (in *CarMaintenanceEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarMaintenanceEditModel struct{}

// CarMaintenanceDeleteInp 删除Maintenance
type CarMaintenanceDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarMaintenanceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarMaintenanceDeleteModel struct{}

// CarMaintenanceViewInp 获取指定Maintenance信息
type CarMaintenanceViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarMaintenanceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarMaintenanceViewModel struct {
	entity.CarMaintenance
}

// CarMaintenanceListInp 获取Maintenance列表
type CarMaintenanceListInp struct {
	input_form.PageReq
	Type uint `json:"type"    dc:"类型 1-全部 2-接送机 3-包机"`
}

func (in *CarMaintenanceListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarMaintenanceListModel struct {
	Id        int64       `json:"id"              dc:"id"`
	Type      uint        `json:"type"    dc:"类型 1-全部 2-接送机 3-包机"`
	Language  string      `json:"language"    dc:"语言"`
	IsDefault uint        `json:"isDefault" dc:"是否默认  1是 2否"`
	Content   string      `json:"content"    dc:"内容"`
	CreateAt  *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// CarMaintenanceLanguageListInp 获取Banner最大排序
type CarMaintenanceLanguageListInp struct {
	Type uint `json:"type"    dc:"类型 1-全部 2-接送机 3-包机"`
}

func (in *CarMaintenanceLanguageListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarMaintenanceLanguageListModel struct {
	Language string `json:"language"    dc:"语言"`
}
