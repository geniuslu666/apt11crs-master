// Package sysin

package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsGuestProfileUpdateFields 修改客户档案字段过滤
type PmsGuestProfileUpdateFields struct {
	Uid           string `json:"uid"           dc:"客户档案ID"`
	FirstName     string `json:"firstName"     dc:"名"`
	LastName      string `json:"lastName"      dc:"姓"`
	FirstNameKana string `json:"firstNameKana" dc:"名的假名"`
	LastNameKana  string `json:"lastNameKana"  dc:"姓的假名"`
	FullName      string `json:"fullName"      dc:"全名"`
	Language      string `json:"language"      dc:"语言"`
	Email         string `json:"email"         dc:"电子邮件"`
	Phone         string `json:"phone"         dc:"电话"`
	Nationality   string `json:"nationality"   dc:"国籍"`
	Address       string `json:"address"       dc:"地址"`
}

// PmsGuestProfileInsertFields 新增客户档案字段过滤
type PmsGuestProfileInsertFields struct {
	MemberId      string `json:"member_id"     dc:"会员ID"`
	Uid           string `json:"uid"           dc:"客户档案ID"`
	FirstName     string `json:"firstName"     dc:"名"`
	LastName      string `json:"lastName"      dc:"姓"`
	FirstNameKana string `json:"firstNameKana" dc:"名的假名"`
	LastNameKana  string `json:"lastNameKana"  dc:"姓的假名"`
	FullName      string `json:"fullName"      dc:"全名"`
	Language      string `json:"language"      dc:"语言"`
	Email         string `json:"email"         dc:"电子邮件"`
	Phone         string `json:"phone"         dc:"电话"`
	Nationality   string `json:"nationality"   dc:"国籍"`
	Address       string `json:"address"       dc:"地址"`
}

// PmsGuestProfileEditInp 修改/新增客户档案
type PmsGuestProfileEditInp struct {
	entity.PmsGuestProfile
}

func (in *PmsGuestProfileEditInp) Filter(ctx context.Context) (err error) {
	// 验证电子邮件
	if err := g.Validator().Rules("email").Data(in.Email).Messages("电子邮件不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsGuestProfileEditModel struct{}

// PmsGuestProfileDeleteInp 删除客户档案
type PmsGuestProfileDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsGuestProfileDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsGuestProfileDeleteModel struct{}

// PmsGuestProfileViewInp 获取指定客户档案信息
type PmsGuestProfileViewInp struct {
	Id int `json:"id" v:"required#主键不能为空" dc:"主键"`
}

func (in *PmsGuestProfileViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsGuestProfileViewModel struct {
	entity.PmsGuestProfile
}

// PmsGuestProfileListInp 获取客户档案列表
type PmsGuestProfileListInp struct {
	input_form.PageReq
	FullName    string        `json:"fullName"    dc:"全名"`
	Email       string        `json:"email"       dc:"电子邮件"`
	Phone       string        `json:"phone"       dc:"电话"`
	Nationality string        `json:"nationality" dc:"国籍"`
	CreatedAt   []*gtime.Time `json:"createdAt"   dc:"创建时间"`
}

func (in *PmsGuestProfileListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsGuestProfileListModel struct {
	Id            int         `json:"id"            dc:"主键"`
	Uid           string      `json:"uid"           dc:"客户档案ID"`
	FirstName     string      `json:"firstName"     dc:"名"`
	LastName      string      `json:"lastName"      dc:"姓"`
	FirstNameKana string      `json:"firstNameKana" dc:"名的假名"`
	LastNameKana  string      `json:"lastNameKana"  dc:"姓的假名"`
	FullName      string      `json:"fullName"      dc:"全名"`
	Language      string      `json:"language"      dc:"语言"`
	Email         string      `json:"email"         dc:"电子邮件"`
	Phone         string      `json:"phone"         dc:"电话"`
	Nationality   string      `json:"nationality"   dc:"国籍"`
	Address       string      `json:"address"       dc:"地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
}

// PmsGuestProfileExportModel 导出客户档案
type PmsGuestProfileExportModel struct {
	Id            int         `json:"id"            dc:"主键"`
	Uid           string      `json:"uid"           dc:"客户档案ID"`
	FirstName     string      `json:"firstName"     dc:"名"`
	LastName      string      `json:"lastName"      dc:"姓"`
	FirstNameKana string      `json:"firstNameKana" dc:"名的假名"`
	LastNameKana  string      `json:"lastNameKana"  dc:"姓的假名"`
	FullName      string      `json:"fullName"      dc:"全名"`
	Language      string      `json:"language"      dc:"语言"`
	Email         string      `json:"email"         dc:"电子邮件"`
	Phone         string      `json:"phone"         dc:"电话"`
	Nationality   string      `json:"nationality"   dc:"国籍"`
	Address       string      `json:"address"       dc:"地址"`
	CreatedAt     *gtime.Time `json:"createdAt"     dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     dc:"更新时间"`
}
