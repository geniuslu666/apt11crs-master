package input_car

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_form"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// DriverWithdrawDisagreeFields 拒绝提现字段过滤
type DriverWithdrawDisagreeFields struct {
	WithdrawStatus string `json:"withdrawStatus"   dc:"提现状态"`
	ApplyRemark    string `json:"applyRemark"      dc:"审核备注"`
}

// DriverWithdrawAgreeFields 同意提现字段过滤
type DriverWithdrawAgreeFields struct {
	WithdrawStatus string      `json:"withdrawStatus"   dc:"提现状态"`
	Transfer       int         `json:"transfer"   dc:"转账状态状态"`
	ApplyAt        *gtime.Time `json:"applyAt"        dc:"审核时间"`
}

// DriverWithdrawTransferFields 提现转账字段过滤
type DriverWithdrawTransferFields struct {
	Transfer int `json:"transfer"   dc:"转账状态状态"`
}

// DriverWithdrawListInp 获取提现申请表列表
type DriverWithdrawListInp struct {
	input_form.PageReq
	Type           string        `json:"type"           dc:"类型"`
	WithdrawSn     string        `json:"withdrawSn"     dc:"提现单号"`
	DriverId       int           `json:"driverId"       dc:"司机ID"`
	DriverName     string        `json:"driverName"     dc:"司机名称"`
	WithdrawStatus string        `json:"withdrawStatus" dc:"提现状态"`
	CreatedAt      []*gtime.Time `json:"createdAt"      dc:"创建时间"`
}

func (in *DriverWithdrawListInp) Filter(ctx context.Context) (err error) {
	return
}

type DriverWithdrawListModel struct {
	Id             int         `json:"id"              dc:""`
	Type           string      `json:"type"            dc:"类型"`
	DriverId       int         `json:"driverId"        dc:"司机ID"`
	WithdrawSn     string      `json:"withdrawSn"      dc:"提现单号"`
	WithdrawStatus string      `json:"withdrawStatus"  dc:"提现状态"`
	WithdrawAmount float64     `json:"withdrawAmount"  dc:"提现金额"`
	ArrivalAmount  float64     `json:"arrivalAmount"   dc:"到账金额"`
	ServiceCharge  float64     `json:"serviceCharge"   dc:"提现手续费比例"`
	Transfer       int         `json:"transfer"        dc:"1、未转账 2、已转账"`
	ApplyRemark    string      `json:"applyRemark"     dc:"审核备注"`
	ApplyAt        *gtime.Time `json:"applyAt"         dc:"审核时间"`
	CreatedAt      *gtime.Time `json:"createdAt"       dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"       dc:"更新时间"`
	DriverDetail   *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
}

// DriverWithdrawViewInp 获取指定提现申请表信息
type DriverWithdrawViewInp struct {
	Id int `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *DriverWithdrawViewInp) Filter(ctx context.Context) (err error) {
	return
}

type DriverWithdrawViewModel struct {
	entity.CarDriverWithdraw
	DriverDetail *struct {
		gmeta.Meta `orm:"table:hg_car_driver"`
		*entity.CarDriver
	} `json:"driverDetail" orm:"with:id=driver_id"`
}

// DriverWithdrawAgreeInp 同意提现
type DriverWithdrawAgreeInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *DriverWithdrawAgreeInp) Filter(ctx context.Context) (err error) {
	return
}

type PmsWithdrawAgreeModel struct{}

// DriverWithdrawDisagreeInp 拒绝提现
type DriverWithdrawDisagreeInp struct {
	Id          interface{} `json:"id" v:"required#id不能为空" dc:"id"`
	ApplyRemark string      `json:"type"           dc:"类型"`
}

func (in *DriverWithdrawDisagreeInp) Filter(ctx context.Context) (err error) {
	return
}

type DriverWithdrawDisagreeModel struct{}

// DriverWithdrawTransferInp 提现转账
type DriverWithdrawTransferInp struct {
	Id interface{} `json:"id" v:"required#id不能为空" dc:"id"`
}

func (in *DriverWithdrawTransferInp) Filter(ctx context.Context) (err error) {
	return
}

type DriverWithdrawTransferModel struct{}
