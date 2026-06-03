package input_spa

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaSettlementUpdateFields 修改结算模式字段过滤
type SpaSettlementUpdateFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// SpaSettlementInsertFields 新增结算模式字段过滤
type SpaSettlementInsertFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// SpaSettlementEditInp 修改/新增结算模式
type SpaSettlementEditInp struct {
	entity.SpaSettlement
}

func (in *SpaSettlementEditInp) Filter(ctx context.Context) (err error) {

	return
}

type SpaSettlementEditModel struct{}

// SpaSettlementDeleteInp 删除结算模式
type SpaSettlementDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaSettlementDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaSettlementDeleteModel struct{}

// SpaSettlementViewInp 获取指定结算模式信息
type SpaSettlementViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaSettlementViewInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaSettlementViewModel struct {
	entity.SpaSettlement
}

// SpaSettlementListInp 获取结算模式列表
type SpaSettlementListInp struct {
	input_form.PageReq
}

func (in *SpaSettlementListInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaSettlementListModel struct {
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

// SpaSettlementStatusInp 更新结算模式状态
type SpaSettlementStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *SpaSettlementStatusInp) Filter(ctx context.Context) (err error) {
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

type SpaSettlementStatusModel struct{}

// SpaSettlementSwitchInp 更新结算模式开关状态
type SpaSettlementSwitchInp struct {
	input_form.SwitchReq
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *SpaSettlementSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type SpaSettlementSwitchModel struct{}
