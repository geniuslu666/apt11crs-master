// Package sysin

package input_app_member

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// PmsMemberLogViewInp 获取指定会员登录日志信息
type PmsMemberLogViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsMemberLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberLogViewModel struct {
	entity.PmsMemberLog
}

// PmsMemberLogListInp 获取会员登录日志列表
type PmsMemberLogListInp struct {
	input_form.PageReq
	Id                 int           `json:"id"                 dc:"id"`
	LoginTime          []*gtime.Time `json:"loginTime"          dc:"登录时间"`
	LoginType          string        `json:"loginType"          dc:"登录方式"`
	CreatedAt          []*gtime.Time `json:"createdAt"          dc:"created_at"`
	PmsMemberMemberNo  string        `json:"pmsMemberMemberNo"  dc:"会员号"`
	PmsMemberFullName  string        `json:"pmsMemberFullName"  dc:"全名"`
	PmsMemberPhone     string        `json:"pmsMemberPhone"     dc:"手机号"`
	PmsMemberPhoneArea string        `json:"pmsMemberPhoneArea" dc:"手机区号"`
	PmsMemberMail      string        `json:"pmsMemberMail"      dc:"邮箱"`
	PmsMemberSource    string        `json:"pmsMemberSource"    dc:"注册来源"`
}

func (in *PmsMemberLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsMemberLogListModel struct {
	Id                 int         `json:"id"                 dc:"id"`
	MemberId           int         `json:"memberId"           dc:"会员ID"`
	LoginTime          *gtime.Time `json:"loginTime"          dc:"登录时间"`
	LoginType          string      `json:"loginType"          dc:"登录方式"`
	LoginIp            string      `json:"loginIp"            dc:"登录IP"`
	ExpirTime          *gtime.Time `json:"expirTime"          dc:"过期时间"`
	PmsMemberMemberNo  string      `json:"pmsMemberMemberNo"  dc:"会员号"`
	PmsMemberAvatar    string      `json:"pmsMemberAvatar"    dc:"头像"`
	PmsMemberFullName  string      `json:"pmsMemberFullName"  dc:"全名"`
	PmsMemberPhone     string      `json:"pmsMemberPhone"     dc:"手机号"`
	PmsMemberPhoneArea string      `json:"pmsMemberPhoneArea" dc:"手机区号"`
	PmsMemberMail      string      `json:"pmsMemberMail"      dc:"邮箱"`
	PmsMemberSource    string      `json:"pmsMemberSource"    dc:"注册来源"`
	MdCode             string      `json:"mdCode"    dc:"注册设备号"`
	MpModel            string      `json:"mpModel"    dc:"注册手机型号"`
}

// PmsMemberLogExportModel 导出会员登录日志
type PmsMemberLogExportModel struct {
	Id        int         `json:"id"        dc:"id"`
	MemberId  int         `json:"memberId"  dc:"会员ID"`
	LoginTime *gtime.Time `json:"loginTime" dc:"登录时间"`
	LoginType string      `json:"loginType" dc:"登录方式"`
	LoginIp   string      `json:"loginIp"   dc:"登录IP"`
	ExpirTime *gtime.Time `json:"expirTime" dc:"过期时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"created_at"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"updated_at"`
}
