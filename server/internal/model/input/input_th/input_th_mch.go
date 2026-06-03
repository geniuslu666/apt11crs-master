// Package sysin

package input_th

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// ThMchUpdateFields 修改商户字段过滤
type ThMchUpdateFields struct {
	CategoryId  uint   `json:"categoryId"  dc:"分类ID"`
	Name        string `json:"name"       dc:"名称"`
	Logo        string `json:"logo"        dc:"LOGO"`
	ContactInfo string `json:"contactInfo" dc:"联系信息"`
}

// ThMchInsertFields 新增商户字段过滤
type ThMchInsertFields struct {
	CategoryId  uint   `json:"categoryId"  dc:"分类ID"`
	Name        string `json:"name"       dc:"名称"`
	Logo        string `json:"logo"        dc:"LOGO"`
	ContactInfo string `json:"contactInfo" dc:"联系信息"`
}

// ThMchEditInp 修改/新增商户
type ThMchEditInp struct {
	entity.ThMch
}

func (in *ThMchEditInp) Filter(ctx context.Context) (err error) {

	return
}

type ThMchEditModel struct{}

// ThMchDeleteInp 删除商户
type ThMchDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThMchDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchDeleteModel struct{}

// ThMchViewInp 获取指定商户信息
type ThMchViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *ThMchViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchViewModel struct {
	entity.ThMch
	ThMchCategoryName string `json:"thMchCategoryName" dc:"分类名称"`
}

// ThMchListInp 获取商户列表
type ThMchListInp struct {
	input_form.PageReq
	Name        string `json:"name"       dc:"名称"`
	CategoryId  uint   `json:"categoryId"  dc:"分类ID"`
	ContactInfo string `json:"contactInfo" dc:"联系信息"`
	Status      int    `json:"status" dc:"状态1、启用 2、禁用"`
	MchIds      []int  `json:"mchIds" dc:"商户id数组"`
}

func (in *ThMchListInp) Filter(ctx context.Context) (err error) {
	return
}

type ThMchListModel struct {
	Id                int         `json:"id"       dc:"id"`
	Name              string      `json:"name"        dc:"名称"`
	StoreNum          int         `json:"storeNum"  dc:"门店数量"`
	ContactInfo       string      `json:"contactInfo" dc:"联系信息"`
	Status            int         `json:"status"   dc:"状态1、启用 2、禁用"`
	StoreOnNum        int         `json:"storeOnNum"  dc:"启用中门店数量"`
	StoreOffNum       int         `json:"storeOffNum" dc:"禁用中门店数量"`
	ThMchCategoryName string      `json:"thMchCategoryName" dc:"分类名称"`
	CreatedAt         *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// ThMchAllInp 获取商户全部列表
type ThMchAllInp struct {
}

type ThMchAllModel struct {
	Id   int    `json:"id"        dc:"id"`
	Name string `json:"name"      dc:"商户名称"`
}

// ThMchStatusInp 更新商户状态
type ThMchStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *ThMchStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type ThMchStatusModel struct{}

// ThMchSwitchInp 更新状态
type ThMchSwitchInp struct {
	Id     int  `json:"id" v:"required#id不能为空" dc:"id"`
	Status uint `json:"status"    dc:"1、启用 2、禁用"`
}

func (in *ThMchSwitchInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.Status) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type ThMchSwitchModel struct{}
