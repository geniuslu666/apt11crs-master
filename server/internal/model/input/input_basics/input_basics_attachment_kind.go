package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
)

// SysAttachmentKindUpdateFields 修改附件分类字段过滤
type SysAttachmentKindUpdateFields struct {
	Label string `json:"label" dc:"分类名称"`
	Key   string `json:"key"   dc:"分类键"`
	Value string `json:"value" dc:"分类值"`
	Icon  string `json:"icon"  dc:"分类图标"`
}

// SysAttachmentKindInsertFields 新增附件分类字段过滤
type SysAttachmentKindInsertFields struct {
	Label string `json:"label" dc:"分类名称"`
	Key   string `json:"key"   dc:"分类键"`
	Value string `json:"value" dc:"分类值"`
	Icon  string `json:"icon"  dc:"分类图标"`
}

// SysAttachmentKindEditInp 修改/新增附件分类
type SysAttachmentKindEditInp struct {
	entity.SysAttachmentKind
}

func (in *SysAttachmentKindEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SysAttachmentKindEditModel struct{}

// SysAttachmentKindDeleteInp 删除附件分类
type SysAttachmentKindDeleteInp struct {
	Id interface{} `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *SysAttachmentKindDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SysAttachmentKindDeleteModel struct{}

// SysAttachmentKindViewInp 获取指定附件分类信息
type SysAttachmentKindViewInp struct {
	Id int `json:"id" v:"required#编号不能为空" dc:"编号"`
}

func (in *SysAttachmentKindViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SysAttachmentKindViewModel struct {
	entity.SysAttachmentKind
}

// SysAttachmentKindListInp 获取附件分类列表
type SysAttachmentKindListInp struct {
	input_form.PageReq
	Id    int    `json:"id"    dc:"编号"`
	Label string `json:"label" dc:"分类名称"`
	Key   string `json:"key"   dc:"分类键"`
}

func (in *SysAttachmentKindListInp) Filter(ctx context.Context) (err error) {
	return
}

type SysAttachmentKindListModel struct {
	Id    int    `json:"id"    dc:"编号"`
	Label string `json:"label" dc:"分类名称"`
	Key   string `json:"key"   dc:"分类键"`
	Value string `json:"value" dc:"分类值"`
	Icon  string `json:"icon"  dc:"分类图标"`
}
