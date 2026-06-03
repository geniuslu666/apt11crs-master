package input_car

import (
	"APT/internal/consts"
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"APT/utility/validate"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarSettlementUpdateFields 修改结算模式字段过滤
type CarSettlementUpdateFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// CarSettlementInsertFields 新增结算模式字段过滤
type CarSettlementInsertFields struct {
	Name            string  `json:"name"            dc:"名称"`
	Type            int     `json:"type"            dc:"结算方式"`
	Cycle           int     `json:"cycle"           dc:"结算周期"`
	Rate            float64 `json:"rate"            dc:"结算比例"`
	Cost            string  `json:"cost"            dc:"结算成本控制"`
	IsAuditWithdraw int     `json:"isAuditWithdraw" dc:"是否提现审核"`
	Status          int     `json:"status"          dc:"状态"`
}

// CarSettlementEditInp 修改/新增结算模式
type CarSettlementEditInp struct {
	entity.CarSettlement
}

func (in *CarSettlementEditInp) Filter(ctx context.Context) (err error) {

	return
}

type CarSettlementEditModel struct{}

// CarSettlementDeleteInp 删除结算模式
type CarSettlementDeleteInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarSettlementDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementDeleteModel struct{}

// CarSettlementViewInp 获取指定结算模式信息
type CarSettlementViewInp struct {
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarSettlementViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementViewModel struct {
	entity.CarSettlement
}

// CarSettlementListInp 获取结算模式列表
type CarSettlementListInp struct {
	input_form.PageReq
}

func (in *CarSettlementListInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementListModel struct {
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

// CarSettlementStatusInp 更新结算模式状态
type CarSettlementStatusInp struct {
	Id     int64 `json:"id" v:"required#id不能为空" dc:"id"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CarSettlementStatusInp) Filter(ctx context.Context) (err error) {
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

type CarSettlementStatusModel struct{}

// CarSettlementSwitchInp 更新结算模式开关状态
type CarSettlementSwitchInp struct {
	input_form.SwitchReq
	Id int64 `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *CarSettlementSwitchInp) Filter(ctx context.Context) (err error) {
	return
}

type CarSettlementSwitchModel struct{}
