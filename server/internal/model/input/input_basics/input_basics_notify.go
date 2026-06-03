package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsNotifyUpdateFields 修改站内信字段过滤
type PmsNotifyUpdateFields struct {
	MemberId      int    `json:"memberId"      dc:"会员ID"`
	NotifyTitle   string `json:"notifyTitle"   dc:"通知标题"`
	NotifyType    string `json:"notifyType"    dc:"通知类型"`
	NotifyContent string `json:"notifyContent" dc:"通知内容"`
	IsRead        string `json:"isRead"        dc:"Y 已读  N 未读"`
}

// PmsNotifyInsertFields 新增站内信字段过滤
type PmsNotifyInsertFields struct {
	MemberId      int    `json:"memberId"      dc:"会员ID"`
	NotifyTitle   string `json:"notifyTitle"   dc:"通知标题"`
	NotifyType    string `json:"notifyType"    dc:"通知类型"`
	NotifyContent string `json:"notifyContent" dc:"通知内容"`
	IsRead        string `json:"isRead"        dc:"Y 已读  N 未读"`
}

// PmsNotifyEditInp 修改/新增站内信
type PmsNotifyEditInp struct {
	entity.PmsNotify
}

func (in *PmsNotifyEditInp) Filter(ctx context.Context) (err error) {

	return
}

type PmsNotifyEditModel struct{}

// PmsNotifyDeleteInp 删除站内信
type PmsNotifyDeleteInp struct {
	Id interface{} `json:"id" v:"required#id_cannot_be_empty" dc:"id"`
}

func (in *PmsNotifyDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsNotifyDeleteModel struct{}

// PmsNotifyViewInp 获取指定站内信信息
type PmsNotifyViewInp struct {
	Id int `json:"id" v:"required#id_cannot_be_empty" dc:"id"`
}

func (in *PmsNotifyViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsNotifyViewModel struct {
	entity.PmsNotify
}

// PmsNotifyListInp 获取站内信列表
type PmsNotifyListInp struct {
	input_form.PageReq
	Id          int    `json:"id"          dc:"id"`
	MemberId    int    `json:"memberId"    dc:"会员ID"`
	NotifyTitle string `json:"notifyTitle" dc:"通知标题"`
	NotifyType  string `json:"notifyType"  dc:"通知类型"`
	IsRead      string `json:"isRead"      dc:"Y 已读  N 未读"`
}

func (in *PmsNotifyListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsNotifyListModel struct {
	Id            int         `json:"id"            dc:"id"`
	MemberId      int         `json:"memberId"      dc:"会员ID"`
	NotifyTitle   string      `json:"notifyTitle"   dc:"通知标题"`
	NotifyType    string      `json:"notifyType"    dc:"通知类型"`
	NotifyContent string      `json:"notifyContent" dc:"通知内容"`
	IsRead        string      `json:"isRead"        dc:"Y 已读  N 未读"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
