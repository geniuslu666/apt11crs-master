// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/util/gmeta"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberIntentionViewInp 获取指定会员意向表信息
type PmsMemberIntentionViewInp struct {
	Id int `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *PmsMemberIntentionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberIntentionViewModel struct {
	entity.PmsMemberIntention
	Member *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
}

// PmsMemberIntentionListInp 获取会员意向表列表
type PmsMemberIntentionListInp struct {
	input_form.PageReq
	Name     string        `json:"name"     dc:"称呼"`
	Sex      int           `json:"sex"      dc:"1、男 2、女"`
	Phone    string        `json:"phone"    dc:"手机号"`
	Mail     string        `json:"mail"     dc:"邮箱"`
	CreateAt []*gtime.Time `json:"createAt" dc:"创建时间"`
}

func (in *PmsMemberIntentionListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberIntentionListModel struct {
	Id        int         `json:"id"        dc:"主键ID"`
	MemberId  int         `json:"memberId"  dc:"意向用户ID"`
	Name      string      `json:"name"      dc:"称呼"`
	Sex       int         `json:"sex"       dc:"1、男 2、女"`
	Country   string      `json:"country"   dc:"国家/地区"`
	Language  string      `json:"language"      dc:"语言"`
	Phone     string      `json:"phone"     dc:"手机号"`
	PhoneArea string      `json:"phoneArea" dc:"手机区号"`
	Mail      string      `json:"mail"      dc:"邮箱"`
	Remark    string      `json:"remark"      dc:"备注"`
	WechatNo  string      `json:"wechatNo"      dc:"微信号"`
	LineId    string      `json:"lineId"      dc:"lineId"`
	CreateAt  *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  dc:"修改时间"`
	Member    *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"member" orm:"with:id=memberId" dc:"会员信息"`
}

// PmsMemberIntentionExportModel 导出会员意向表
type PmsMemberIntentionExportModel struct {
	Id        int         `json:"id"        dc:"主键ID"`
	MemberId  int         `json:"memberId"  dc:"意向用户ID"`
	Name      string      `json:"name"      dc:"称呼"`
	Sex       int         `json:"sex"       dc:"1、男 2、女"`
	Country   string      `json:"country"   dc:"国家/地区"`
	Phone     string      `json:"phone"     dc:"手机号"`
	PhoneArea string      `json:"phoneArea" dc:"手机区号"`
	Mail      string      `json:"mail"      dc:"邮箱"`
	Remark    string      `json:"remark"      dc:"备注"`
	WechatNo  string      `json:"wechatNo"      dc:"微信号"`
	LineId    string      `json:"lineId"      dc:"lineId"`
	CreateAt  *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  dc:"修改时间"`
}

// PmsMemberIntentionUpdateFields 修改会员意向表字段过滤
type PmsMemberIntentionUpdateFields struct {
	MemberId  int         `json:"memberId"  dc:"意向用户ID"`
	Name      string      `json:"name"      dc:"称呼"`
	Sex       int         `json:"sex"       dc:"1、男 2、女"`
	Country   string      `json:"country"   dc:"国家/地区"`
	Phone     string      `json:"phone"     dc:"手机号"`
	PhoneArea string      `json:"phoneArea" dc:"手机区号"`
	Mail      string      `json:"mail"      dc:"邮箱"`
	Remark    string      `json:"remark"    dc:"备注"`
	WechatNo  string      `json:"wechatNo"      dc:"微信号"`
	LineId    string      `json:"lineId"      dc:"lineId"`
	CreateAt  *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  dc:"修改时间"`
}

// PmsMemberIntentionInsertFields 新增会员意向表字段过滤
type PmsMemberIntentionInsertFields struct {
	MemberId  int         `json:"memberId"  dc:"意向用户ID"`
	Name      string      `json:"name"      dc:"称呼"`
	Sex       int         `json:"sex"       dc:"1、男 2、女"`
	Country   string      `json:"country"   dc:"国家/地区"`
	Phone     string      `json:"phone"     dc:"手机号"`
	PhoneArea string      `json:"phoneArea" dc:"手机区号"`
	Mail      string      `json:"mail"      dc:"邮箱"`
	Remark    string      `json:"remark"    dc:"备注"`
	WechatNo  string      `json:"wechatNo"      dc:"微信号"`
	LineId    string      `json:"lineId"      dc:"lineId"`
	CreateAt  *gtime.Time `json:"createAt"  dc:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  dc:"修改时间"`
}

// PmsMemberIntentionEditInp 修改/新增会员意向表
type PmsMemberIntentionEditInp struct {
	entity.PmsMemberIntention
}

func (in *PmsMemberIntentionEditInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberIntentionEditModel struct{}
