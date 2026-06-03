package input_spa

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// SpaIspUpdateFields 修改服务商管理字段过滤
type SpaIspUpdateFields struct {
	Name             string  `json:"name"                    dc:"真实姓名"`
	Status           int     `json:"status"                  dc:"状态"`
	SettlementId     int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType   int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate   float64 `json:"settlementRate"          dc:"分成结算比例"`
	SettlementObject string  `json:"settlementObject"      dc:"结算对象"`
}

// SpaIspInsertFields 新增服务商管理字段过滤
type SpaIspInsertFields struct {
	Name             string  `json:"name"                    dc:"真实姓名"`
	Status           int     `json:"status"                  dc:"状态"`
	SettlementId     int     `json:"settlementId"            dc:"结算模式ID"`
	SettlementType   int     `json:"settlementType"          dc:"分成抽成类型 1跟随系统  2自定义"`
	SettlementRate   float64 `json:"settlementRate"          dc:"分成结算比例"`
	SettlementObject string  `json:"settlementObject"      dc:"结算对象"`
}

// SpaIspEditInp 修改/新增服务商管理
type SpaIspEditInp struct {
	entity.SpaIsp
}

func (in *SpaIspEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaIspEditModel struct{}

// SpaIspDeleteInp 删除服务商管理
type SpaIspDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaIspDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaIspDeleteModel struct{}

// SpaIspViewInp 获取指定服务商管理信息
type SpaIspViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaIspViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaIspViewModel struct {
	entity.SpaIsp
}

// SpaIspListInp 获取服务商管理列表
type SpaIspListInp struct {
	input_form.PageReq
	Keywords   string        `json:"keywords" dc:"服务商名称"`
	Status     int           `json:"status"             dc:"状态"`
	WorkStatus string        `json:"workStatus"             dc:"工作状态"`
	CreateAt   []*gtime.Time `json:"createAt" dc:"加入时间"`
}

func (in *SpaIspListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaIspListModel struct {
	Id           int64       `json:"id"                      dc:"id"`
	Name         string      `json:"name"                dc:"真实姓名"`
	Status       int         `json:"status"    dc:"状态"`
	WorkStatus   string      `json:"workStatus"       dc:"工作状态"`
	MemberId     int         `json:"memberId"         dc:"会员ID"`
	SendSmsPhone string      `json:"sendSmsPhone"                dc:"发送短信电话"`
	CreateAt     *gtime.Time `json:"createAt"                dc:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"                dc:"更新时间"`
	MemberDetail *struct {
		gmeta.Meta `orm:"table:hg_pms_member"`
		Id         int    `json:"id"    orm:"id"      dc:"id"`
		FullName   string `json:"fullName"     orm:"full_name"     dc:"全名"`
		Avatar     string `json:"avatar"       orm:"avatar"        dc:"头像"`
		Phone      string `json:"phone"        orm:"phone"         dc:"手机号"`
		PhoneArea  string `json:"phoneArea"    orm:"phone_area"    dc:"手机区号"`
		MemberNo   string `json:"memberNo"    orm:"member_no"    dc:"会员号"`
	} `json:"memberDetail" orm:"with:id=member_id" dc:"会员信息"`
}

type SpaIspAllListModel struct {
	Id   int64  `json:"id"                      dc:"id"`
	Name string `json:"name"                dc:"真实姓名"`
}

// SpaIspStatusInp 更新服务商状态
type SpaIspStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *SpaIspStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaIspStatusModel struct{}

// SpaIspWorkStatusInp 更新工作状态
type SpaIspWorkStatusInp struct {
	Id         int    `json:"id" v:"required#id不能为空" dc:"id"`
	WorkStatus string `json:"workStatus" dc:"状态"`
}

func (in *SpaIspWorkStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("id不能为空")
		return
	}

	if g.IsEmpty(in.WorkStatus) {
		err = gerror.New("状态不能为空")
		return
	}

	return
}

type SpaIspWorkStatusModel struct{}

// SpaIspBindInp 绑定用户
type SpaIspBindInp struct {
	Id       int `json:"id" dc:"ID"`
	MemberId int `json:"memberId" dc:"会员ID"`
}

func (in *SpaIspBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.MemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type SpaIspBindModel struct{}

// SpaIspUnbindInp 解绑渠道管理
type SpaIspUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaIspUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaIspUnbindModel struct{}
