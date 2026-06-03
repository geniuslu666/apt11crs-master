package input_basics

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsChannelUpdateFields 修改渠道管理字段过滤
type PmsChannelUpdateFields struct {
	Name                string  `json:"name"   dc:"姓名"`
	Phone               string  `json:"phone"  dc:"手机号"`
	Email               string  `json:"email"  dc:"邮箱"`
	Rate                float64 `json:"rate"   dc:"返佣比例"`
	Status              int     `json:"status" dc:"状态"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64 `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64 `json:"afterDay"   dc:"预计可提现周期"`
	Remark              string  `json:"remark" dc:"备注"`
}

// PmsChannelInsertFields 新增渠道管理字段过滤
type PmsChannelInsertFields struct {
	Name                string  `json:"name"   dc:"姓名"`
	Phone               string  `json:"phone"  dc:"手机号"`
	Email               string  `json:"email"  dc:"邮箱"`
	Rate                float64 `json:"rate"   dc:"返佣比例"`
	Status              int     `json:"status" dc:"状态"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64 `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64 `json:"afterDay"   dc:"预计可提现周期"`
	Remark              string  `json:"remark" dc:"备注"`
}

// PmsChannelEditInp 修改/新增渠道管理
type PmsChannelEditInp struct {
	entity.PmsChannel
}

func (in *PmsChannelEditInp) Filter(ctx context.Context) (err error) {
	// 验证邮箱
	if in.Email != "" {
		if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
			return err.Current()
		}
	}

	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsChannelEditModel struct{}

// PmsChannelBindInp 渠道绑定会员
type PmsChannelBindInp struct {
	Id          int `json:"id" dc:"ID"`
	PmsMemberId int `json:"pmsMemberId" dc:"会员ID"`
}

func (in *PmsChannelBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.PmsMemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type PmsChannelBindModel struct{}

// PmsChannelUnbindInp 解绑渠道管理
type PmsChannelUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsChannelUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsChannelUnbindModel struct{}

// PmsChannelDeleteInp 删除渠道管理
type PmsChannelDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsChannelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsChannelDeleteModel struct{}

// PmsChannelViewInp 获取指定渠道管理信息
type PmsChannelViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsChannelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsChannelViewModel struct {
	entity.PmsChannel
	PmsMemberId       int    `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo string `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName string `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone    string `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail     string `json:"pmsMemberMail"           dc:"邮箱"`
}

// PmsChannelListInp 获取渠道管理列表
type PmsChannelListInp struct {
	input_form.PageReq
	Name      string        `json:"name"      dc:"姓名"`
	Phone     string        `json:"phone"     dc:"手机号"`
	Email     string        `json:"email"     dc:"邮箱"`
	Rate      []float64     `json:"rate"      dc:"返佣比例"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *PmsChannelListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsChannelListModel struct {
	Id                  int         `json:"id"        dc:"id"`
	Name                string      `json:"name"      dc:"姓名"`
	Phone               string      `json:"phone"     dc:"手机号"`
	Email               string      `json:"email"     dc:"邮箱"`
	Rate                float64     `json:"rate"      dc:"返佣比例"`
	Status              int         `json:"status"    dc:"状态"`
	Balance             float64     `json:"balance"    dc:"可提现账户"`
	AllBalance          float64     `json:"allBalance" dc:"总账户"`
	MinWithdrawalAmount float64     `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64     `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64     `json:"afterDay"   dc:"预计可提现周期"`
	CreatedAt           *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updatedAt" dc:"updated_at"`
	PmsMemberId         int         `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo   string      `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName   string      `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone      string      `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail       string      `json:"pmsMemberMail"           dc:"邮箱"`
}

// PmsChannelStatusInp 更新渠道管理状态
type PmsChannelStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsChannelStatusInp) Filter(ctx context.Context) (err error) {
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

type PmsChannelStatusModel struct{}
