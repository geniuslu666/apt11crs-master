// Package sysin

package input_hotel

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_language"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCancelRateUpdateFields 修改取消政策字段过滤
type PmsCancelRateUpdateFields struct {
	Mode      string `json:"mode"      dc:"规则模式"`
	StartDays int    `json:"startDays" dc:"开始天数"`
	EndDays   int    `json:"endDays"   dc:"结束天数"`
	Rate      int    `json:"rate"      dc:"费率"`
	Name      string `json:"name"      dc:"规则名"`
	Sort      int    `json:"sort"      dc:"排序规则  从小到大"`
}

// PmsCancelRateInsertFields 新增取消政策字段过滤
type PmsCancelRateInsertFields struct {
	Mode      string `json:"mode"      dc:"规则模式"`
	StartDays int    `json:"startDays" dc:"开始天数"`
	EndDays   int    `json:"endDays"   dc:"结束天数"`
	Rate      int    `json:"rate"      dc:"费率"`
	Name      string `json:"name"      dc:"规则名"`
	Sort      int    `json:"sort"      dc:"排序规则  从小到大"`
}

// PmsCancelRateEditInp 修改/新增取消政策
type PmsCancelRateEditInp struct {
	Id           int                          `json:"id"        orm:"id"         description:""`
	Mode         string                       `json:"mode"      orm:"mode"       description:"规则模式"`
	StartDays    int                          `json:"startDays" orm:"start_days" description:"开始天数"`
	EndDays      int                          `json:"endDays"   orm:"end_days"   description:"结束天数"`
	Rate         *int                         `json:"rate"      orm:"rate"       description:"费率"`
	Name         string                       `json:"name"      orm:"name"       description:"规则名"`
	Sort         int                          `json:"sort"      orm:"sort"       description:"排序规则  从小到大"`
	NameLanguage input_language.LanguageModel `json:"nameLanguage" dc:"多语言物业名称"`
}

func (in *PmsCancelRateEditInp) Filter(ctx context.Context) (err error) {
	// 验证规则模式
	if err := g.Validator().Rules("required").Data(in.Mode).Messages("规则模式不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证开始天数
	if err := g.Validator().Rules("required").Data(in.StartDays).Messages("开始天数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证结束天数
	if err := g.Validator().Rules("required").Data(in.EndDays).Messages("结束天数不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证费率
	if err := g.Validator().Rules("required").Data(in.Rate).Messages("费率不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证规则名
	if err := g.Validator().Rules("required").Data(in.Name).Messages("规则名不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序规则  从小到大
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序规则  从小到大不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type PmsCancelRateEditModel struct{}

// PmsCancelRateDeleteInp 删除取消政策
type PmsCancelRateDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsCancelRateDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCancelRateDeleteModel struct{}

// PmsCancelRateViewInp 获取指定取消政策信息
type PmsCancelRateViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsCancelRateViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCancelRateViewModel struct {
	entity.PmsCancelRate
}

// PmsCancelRateListInp 获取取消政策列表
type PmsCancelRateListInp struct {
	input_form.PageReq
}

func (in *PmsCancelRateListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCancelRateListModel struct {
	Id           int             `json:"id"            dc:"id"`
	Mode         string          `json:"mode"          dc:"规则模式"`
	StartDays    int             `json:"startDays"     dc:"开始天数"`
	EndDays      int             `json:"endDays"       dc:"结束天数"`
	Rate         int             `json:"rate"          dc:"费率"`
	Name         string          `json:"name"          dc:"规则名"`
	NameLanguage []*LanguageType `json:"nameLanguage"  dc:"多语言物业名称"   orm:"with:uuid=name"`
	Sort         int             `json:"sort"          dc:"排序规则  从小到大"`
	CreatedAt    *gtime.Time     `json:"createdAt"     dc:"created_at"`
	UpdatedAt    *gtime.Time     `json:"updatedAt"     dc:"updated_at"`
}

// PmsCancelRateMaxSortInp 获取取消政策最大排序
type PmsCancelRateMaxSortInp struct{}

func (in *PmsCancelRateMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsCancelRateMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}
