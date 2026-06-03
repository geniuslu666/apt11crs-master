// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberGroupUpdateFields 修改会员分组字段过滤
type PmsMemberGroupUpdateFields struct {
	MemberGroup string `json:"memberGroup" dc:"会员分组名称"`
}

// PmsMemberGroupInsertFields 新增会员分组字段过滤
type PmsMemberGroupInsertFields struct {
	MemberGroup string `json:"memberGroup" dc:"会员分组名称"`
}

// PmsMemberGroupEditInp 修改/新增会员分组
type PmsMemberGroupEditInp struct {
	entity.PmsMemberGroup
}

func (in *PmsMemberGroupEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsMemberGroupEditModel struct{}

// PmsMemberGroupDeleteInp 删除会员分组
type PmsMemberGroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#会员分组ID不能为空" dc:"会员分组ID"`
}

func (in *PmsMemberGroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberGroupDeleteModel struct{}

// PmsMemberGroupViewInp 获取指定会员分组信息
type PmsMemberGroupViewInp struct {
	Id int `json:"id" v:"required#会员分组ID不能为空" dc:"会员分组ID"`
}

func (in *PmsMemberGroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberGroupViewModel struct {
	entity.PmsMemberGroup
}

// PmsMemberGroupListInp 获取会员分组列表
type PmsMemberGroupListInp struct {
	input_form.PageReq
	Id        int           `json:"id"        dc:"会员分组ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *PmsMemberGroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberGroupListModel struct {
	Id          int         `json:"id"          dc:"会员分组ID"`
	MemberGroup string      `json:"memberGroup" dc:"会员分组名称"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// PmsMemberGroupAllInp 获取会员等级列表
type PmsMemberGroupAllInp struct {
	MemberGroup string `json:"memberGroup" dc:"会员分组名称"`
}

func (in *PmsMemberGroupAllInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberGroupAllModel struct {
	Id          int         `json:"id"          dc:"会员分组ID"`
	MemberGroup string      `json:"memberGroup" dc:"会员分组名称"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// PmsMemberGroupExportModel 导出会员分组
type PmsMemberGroupExportModel struct {
	Id          int         `json:"id"          dc:"会员分组ID"`
	MemberGroup string      `json:"memberGroup" dc:"会员分组名称"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}
