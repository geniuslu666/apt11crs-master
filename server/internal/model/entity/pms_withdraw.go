// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsWithdraw is the golang structure for table pms_withdraw.
type PmsWithdraw struct {
	Id             int         `json:"id"             orm:"id"              description:""`
	Type           string      `json:"type"           orm:"type"            description:"类型"`
	StaffId        int         `json:"staffId"        orm:"staff_id"        description:"员工ID"`
	ChannelId      int         `json:"channelId"      orm:"channel_id"      description:"渠道ID"`
	WithdrawSn     string      `json:"withdrawSn"     orm:"withdraw_sn"     description:"提现单号"`
	WithdrawStatus string      `json:"withdrawStatus" orm:"withdraw_status" description:"提现状态"`
	WithdrawAmount float64     `json:"withdrawAmount" orm:"withdraw_amount" description:"提现金额"`
	ArrivalAmount  float64     `json:"arrivalAmount"  orm:"arrival_amount"  description:"到账金额"`
	ServiceCharge  float64     `json:"serviceCharge"  orm:"service_charge"  description:"提现手续费比例"`
	Transfer       int         `json:"transfer"       orm:"transfer"        description:"1、未转账 2、已转账"`
	ApplyRemark    string      `json:"applyRemark"    orm:"apply_remark"    description:"审核备注"`
	ApplyAt        *gtime.Time `json:"applyAt"        orm:"apply_at"        description:"审核时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`
}
