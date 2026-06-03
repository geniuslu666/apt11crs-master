package input_basics

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsWithdrawDisagreeFields 拒绝提现字段过滤
type PmsWithdrawDisagreeFields struct {
	WithdrawStatus string `json:"withdrawStatus"   dc:"提现状态"`
	ApplyRemark    string `json:"applyRemark"      dc:"审核备注"`
}

// PmsWithdrawAgreeFields 同意提现字段过滤
type PmsWithdrawAgreeFields struct {
	WithdrawStatus string      `json:"withdrawStatus"   dc:"提现状态"`
	Transfer       int         `json:"transfer"   dc:"转账状态状态"`
	ApplyAt        *gtime.Time `json:"applyAt"        dc:"审核时间"`
}

// PmsWithdrawTransferFields 提现转账字段过滤
type PmsWithdrawTransferFields struct {
	Transfer int `json:"transfer"   dc:"转账状态状态"`
}

// PmsWithdrawViewInp 获取指定提现申请表信息
type PmsWithdrawViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsWithdrawViewInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawViewModel struct {
	entity.PmsWithdraw
	PmsStaffName   string `json:"pmsStaffName"   dc:"员工姓名"`
	PmsChannelName string `json:"pmsChannelName" dc:"渠道姓名"`
}

// PmsWithdrawListInp 获取提现申请表列表
type PmsWithdrawListInp struct {
	input_form.PageReq
	Type           string        `json:"type"           dc:"类型"`
	CreatedAt      []*gtime.Time `json:"createdAt"      dc:"创建时间"`
	PmsStaffName   string        `json:"pmsStaffName"   dc:"员工姓名"`
	PmsChannelName string        `json:"pmsChannelName" dc:"渠道姓名"`
}

func (in *PmsWithdrawListInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawListModel struct {
	Id             int         `json:"id"             dc:"id"`
	Type           string      `json:"type"           dc:"类型"`
	StaffId        int         `json:"staffId"        dc:"员工ID"`
	ChannelId      int         `json:"channelId"      dc:"渠道ID"`
	WithdrawSn     string      `json:"withdrawSn"     dc:"提现单号"`
	WithdrawStatus string      `json:"withdrawStatus" dc:"提现状态"`
	WithdrawAmount float64     `json:"withdrawAmount" dc:"提现金额"`
	ArrivalAmount  float64     `json:"arrivalAmount"  dc:"到账金额"`
	ServiceCharge  float64     `json:"serviceCharge"  dc:"提现手续费比例"`
	Transfer       int         `json:"transfer"       dc:"1、未转账 2、已转账"`
	ApplyRemark    string      `json:"applyRemark"    dc:"审核备注"`
	ApplyAt        *gtime.Time `json:"applyAt"        dc:"审核时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
	PmsStaffName   string      `json:"pmsStaffName"   dc:"员工姓名"`
	PmsChannelName string      `json:"pmsChannelName" dc:"渠道姓名"`
}

// PmsWithdrawAgreeInp 同意提现
type PmsWithdrawAgreeInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsWithdrawAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawAgreeModel struct{}

// PmsWithdrawDisagreeInp 拒绝提现
type PmsWithdrawDisagreeInp struct {
	Id          interface{} `json:"id" v:"required#id不能为空" dc:"id"`
	ApplyRemark string      `json:"type"           dc:"类型"`
}

func (in *PmsWithdrawDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawDisagreeModel struct{}

// PmsWithdrawTransferInp 提现转账
type PmsWithdrawTransferInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *PmsWithdrawTransferInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawTransferModel struct{}
