// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaIsp is the golang structure for table spa_isp.
type SpaIsp struct {
	Id                    int         `json:"id"                    orm:"id"                      description:""`
	Name                  string      `json:"name"                  orm:"name"                    description:"服务商名称"`
	Status                uint        `json:"status"                orm:"status"                  description:"状态1、启用 2、禁用"`
	WorkStatus            string      `json:"workStatus"            orm:"work_status"             description:"工作状态"`
	MemberId              int         `json:"memberId"              orm:"member_id"               description:"会员ID"`
	SettlementObject      string      `json:"settlementObject"      orm:"settlement_object"       description:"结算对象"`
	SettlementType        int         `json:"settlementType"        orm:"settlement_type"         description:"服务分成类型 1跟随系统  2自定义"`
	SettlementId          int         `json:"settlementId"          orm:"settlement_id"           description:"结算模式ID"`
	SettlementRate        float64     `json:"settlementRate"        orm:"settlement_rate"         description:"服务分成%"`
	SettlementOrderNum    int         `json:"settlementOrderNum"    orm:"settlement_order_num"    description:"预约单已结算数量"`
	SettlementOrderAmount float64     `json:"settlementOrderAmount" orm:"settlement_order_amount" description:"已结算预约单订单金额"`
	TotalSettlementAmount float64     `json:"totalSettlementAmount" orm:"total_settlement_amount" description:"已结算金额"`
	Balance               float64     `json:"balance"               orm:"balance"                 description:"余额"`
	VerifyMoney           float64     `json:"verifyMoney"           orm:"verify_money"            description:"已核账金额"`
	ApplyWithdrawBalance  float64     `json:"applyWithdrawBalance"  orm:"apply_withdraw_balance"  description:"提现中余额"`
	WithdrawBalance       float64     `json:"withdrawBalance"       orm:"withdraw_balance"        description:"已提现余额"`
	SendSmsPhone          string      `json:"sendSmsPhone"          orm:"send_sms_phone"          description:"发送短信电话"`
	CreateAt              *gtime.Time `json:"createAt"              orm:"create_at"               description:"创建时间"`
	UpdateAt              *gtime.Time `json:"updateAt"              orm:"update_at"               description:"更新时间"`
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"              description:""`
}
