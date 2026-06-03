// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsChannel is the golang structure for table pms_channel.
type PmsChannel struct {
	Id                   int         `json:"id"                   orm:"id"                     description:""`
	Name                 string      `json:"name"                 orm:"name"                   description:"渠道姓名"`
	Phone                string      `json:"phone"                orm:"phone"                  description:"手机号"`
	Email                string      `json:"email"                orm:"email"                  description:"邮箱"`
	Rate                 float64     `json:"rate"                 orm:"rate"                   description:"返佣比例"`
	Status               int         `json:"status"               orm:"status"                 description:"状态   1、 开启   2、禁用"`
	Remark               string      `json:"remark"               orm:"remark"                 description:"备注"`
	Balance              float64     `json:"balance"              orm:"balance"                description:"可提现账户"`
	AllBalance           float64     `json:"allBalance"           orm:"all_balance"            description:"总账户"`
	ApplyWithdrawBalance float64     `json:"applyWithdrawBalance" orm:"apply_withdraw_balance" description:"提现中余额"`
	WithdrawBalance      float64     `json:"withdrawBalance"      orm:"withdraw_balance"       description:"已提现余额"`
	MinWithdrawalAmount  float64     `json:"minWithdrawalAmount"  orm:"min_withdrawal_amount"  description:"最低可提现额"`
	ServiceCharge        float64     `json:"serviceCharge"        orm:"service_charge"         description:"手续费"`
	AfterDay             int         `json:"afterDay"             orm:"after_day"              description:"预计提现时间周期"`
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""`
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""`
	DeletedAt            *gtime.Time `json:"deletedAt"            orm:"deleted_at"             description:""`
}
