// Package sysin

package input_food

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSettlementUpdateFields 修改结算模式字段过滤
type FoodSettlementUpdateFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// FoodSettlementInsertFields 新增结算模式字段过滤
type FoodSettlementInsertFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// FoodSettlementEditInp 修改/新增结算模式
type FoodSettlementEditInp struct {
	entity.FoodSettlement
}

func (in *FoodSettlementEditInp) Filter(ctx context.Context) (err error) {

	return
}

type FoodSettlementEditModel struct{}

// FoodSettlementDeleteInp 删除结算模式
type FoodSettlementDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSettlementDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementDeleteModel struct{}

// FoodSettlementViewInp 获取指定结算模式信息
type FoodSettlementViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSettlementViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementViewModel struct {
	entity.FoodSettlement
}

// FoodSettlementListInp 获取结算模式列表
type FoodSettlementListInp struct {
	input_form.PageReq
}

func (in *FoodSettlementListInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementListModel struct {
	Id              int64       `json:"id"              dc:"id"`
	Name            string      `json:"name"            dc:"名称"`
	Type            int         `json:"type"            dc:"结算方式"`
	Cycle           int         `json:"cycle"           dc:"结算周期"`
	Rate            float64     `json:"rate"            dc:"结算比例"`
	Cost            string      `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int         `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int         `json:"status"          dc:"状态"`
	CreateAt        *gtime.Time `json:"createAt"        dc:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        dc:"更新时间"`
}

// FoodSettlementStatusInp 更新结算模式状态
type FoodSettlementStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *FoodSettlementStatusInp) Filter(ctx context.Context) (err error) {
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

type FoodSettlementStatusModel struct{}

// FoodSettlementSwitchInp 更新结算模式开关状态
type FoodSettlementSwitchInp struct {
	input_form.SwitchReq
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *FoodSettlementSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type FoodSettlementSwitchModel struct{}
