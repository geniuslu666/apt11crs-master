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

// PmsStaffUpdateFields 修改员工管理字段过滤
type PmsStaffUpdateFields struct {
	Name                string  `json:"name"       dc:"员工姓名"`
	Department          string  `json:"department" dc:"员工部门"`
	Phone               string  `json:"phone"      dc:"手机号"`
	Email               string  `json:"email"      dc:"邮箱"`
	Rate                float64 `json:"rate"       dc:"返佣比例"`
	Status              int     `json:"status"     dc:"状态   1、 开启   2、禁用"`
	Remark              string  `json:"remark"     dc:"备注"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64 `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64 `json:"afterDay"   dc:"预计可提现周期"`
}

// PmsStaffInsertFields 新增员工管理字段过滤
type PmsStaffInsertFields struct {
	Name                string  `json:"name"       dc:"员工姓名"`
	Department          string  `json:"department" dc:"员工部门"`
	Phone               string  `json:"phone"      dc:"手机号"`
	Email               string  `json:"email"      dc:"邮箱"`
	Rate                float64 `json:"rate"       dc:"返佣比例"`
	Status              int     `json:"status"     dc:"状态   1、 开启   2、禁用"`
	Remark              string  `json:"remark"     dc:"备注"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64 `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64 `json:"afterDay"   dc:"预计可提现周期"`
}

// PmsStaffEditInp 修改/新增员工管理
type PmsStaffEditInp struct {
	entity.PmsStaff
}

func (in *PmsStaffEditInp) Filter(ctx context.Context) (err error) {
	// 验证邮箱
	if in.Email != "" {
		if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
			return err.Current()
		}
	}

	// 验证状态   1、 开启   2、禁用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态   1、 开启   2、禁用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态   1、 开启   2、禁用值不正确").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsStaffEditModel struct{}

// PmsStaffBindInp 员工绑定会员
type PmsStaffBindInp struct {
	Id          int `json:"id" dc:"ID"`
	PmsMemberId int `json:"pmsMemberId" dc:"会员ID"`
}

func (in *PmsStaffBindInp) Filter(ctx context.Context) (err error) {
	// 验证会员
	if in.PmsMemberId == 0 {
		err = gerror.New("请选择会员")
		return
	}

	return
}

type PmsStaffBindModel struct{}

// PmsStaffUnbindInp 解绑员工管理
type PmsStaffUnbindInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsStaffUnbindInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsStaffUnbindModel struct{}

// PmsStaffDeleteInp 删除员工管理
type PmsStaffDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsStaffDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsStaffDeleteModel struct{}

// PmsStaffViewInp 获取指定员工管理信息
type PmsStaffViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsStaffViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsStaffViewModel struct {
	entity.PmsStaff
	PmsMemberId       int    `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo string `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName string `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone    string `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail     string `json:"pmsMemberMail"           dc:"邮箱"`
}

// PmsStaffListInp 获取员工管理列表
type PmsStaffListInp struct {
	input_form.PageReq
	Name       string        `json:"name"       dc:"员工姓名"`
	Department string        `json:"department" dc:"员工部门"`
	Phone      string        `json:"phone"      dc:"手机号"`
	Email      string        `json:"email"      dc:"邮箱"`
	Rate       []float64     `json:"rate"       dc:"返佣比例"`
	Status     int           `json:"status"     dc:"状态   1、 开启   2、禁用"`
	CreatedAt  []*gtime.Time `json:"createdAt"  dc:"created_at"`
}

func (in *PmsStaffListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsStaffListModel struct {
	Id                  int         `json:"id"         dc:"id"`
	Name                string      `json:"name"       dc:"员工姓名"`
	Department          string      `json:"department" dc:"员工部门"`
	Phone               string      `json:"phone"      dc:"手机号"`
	Email               string      `json:"email"      dc:"邮箱"`
	Rate                float64     `json:"rate"       dc:"返佣比例"`
	Status              int         `json:"status"     dc:"状态   1、 开启   2、禁用"`
	Balance             float64     `json:"balance"    dc:"可提现账户"`
	AllBalance          float64     `json:"allBalance" dc:"总账户"`
	MinWithdrawalAmount float64     `json:"minWithdrawalAmount" dc:"最低可提现金额"`
	ServiceCharge       float64     `json:"serviceCharge" dc:"手续费"`
	AfterDay            float64     `json:"afterDay"   dc:"预计可提现周期"`
	CreatedAt           *gtime.Time `json:"createdAt"  dc:"created_at"`
	UpdatedAt           *gtime.Time `json:"updatedAt"  dc:"updated_at"`
	PmsMemberId         int         `json:"pmsMemberId"        dc:"会员ID"`
	PmsMemberMemberNo   string      `json:"pmsMemberMemberNo"        dc:"会员号"`
	PmsMemberFullName   string      `json:"pmsMemberFullName"    dc:"全名"`
	PmsMemberPhone      string      `json:"pmsMemberPhone"           dc:"手机号"`
	PmsMemberMail       string      `json:"pmsMemberMail"           dc:"邮箱"`
}

// PmsStaffStatusInp 更新员工管理状态
type PmsStaffStatusInp struct {
	Id     int `json:"id" v:"required#id不能为空" dc:"id"`
	Status int `json:"status" dc:"状态"`
}

func (in *PmsStaffStatusInp) Filter(ctx context.Context) (err error) {
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

type PmsStaffStatusModel struct{}
