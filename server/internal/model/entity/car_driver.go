// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriver is the golang structure for table car_driver.
type CarDriver struct {
	Id                    int         `json:"id"                    orm:"id"                      description:""`
	CarId                 int         `json:"carId"                 orm:"car_id"                  description:"车辆ID"`
	IsLeader              int         `json:"isLeader"              orm:"is_leader"               description:"是否是车队长 1是  2否"`
	Name                  string      `json:"name"                  orm:"name"                    description:"司机真实姓名"`
	Nickname              string      `json:"nickname"              orm:"nickname"                description:"司机昵称"`
	CooperateTypeId       int         `json:"cooperateTypeId"       orm:"cooperate_type_id"       description:"合作类型ID"`
	Sex                   int         `json:"sex"                   orm:"sex"                     description:"1、男 2、女"`
	Phone                 string      `json:"phone"                 orm:"phone"                   description:"手机号"`
	PhoneArea             string      `json:"phoneArea"             orm:"phone_area"              description:"手机区号"`
	Photo                 string      `json:"photo"                 orm:"photo"                   description:"照片"`
	Age                   int         `json:"age"                   orm:"age"                     description:"年龄"`
	WorkYears             int         `json:"workYears"             orm:"work_years"              description:"从业年数"`
	Language              string      `json:"language"              orm:"language"                description:"语言能力"`
	Status                uint        `json:"status"                orm:"status"                  description:"状态1、启用 2、禁用"`
	WorkStatus            string      `json:"workStatus"            orm:"work_status"             description:"工作状态"`
	MemberId              int         `json:"memberId"              orm:"member_id"               description:"会员ID"`
	QualityMaterials      string      `json:"qualityMaterials"      orm:"quality_materials"       description:"资质信息(多图)"`
	SettlementType        int         `json:"settlementType"        orm:"settlement_type"         description:"服务分成类型 1跟随系统  2自定义"`
	SettlementId          int         `json:"settlementId"          orm:"settlement_id"           description:"结算模式ID"`
	SettlementRate        float64     `json:"settlementRate"        orm:"settlement_rate"         description:"服务分成%"`
	TotalOrderNum         int         `json:"totalOrderNum"         orm:"total_order_num"         description:"预约单总数量（包含退款）"`
	TotalOrderAmount      float64     `json:"totalOrderAmount"      orm:"total_order_amount"      description:"预约单总金额（包含退款）"`
	PayOrderNum           int         `json:"payOrderNum"           orm:"pay_order_num"           description:"预约单支付数量（不包含退款）"`
	PayOrderAmount        float64     `json:"payOrderAmount"        orm:"pay_order_amount"        description:"预约单支付金额（不包含退款）"`
	SettlementOrderNum    int         `json:"settlementOrderNum"    orm:"settlement_order_num"    description:"预约单已结算数量"`
	SettlementOrderAmount float64     `json:"settlementOrderAmount" orm:"settlement_order_amount" description:"已结算预约单订单金额"`
	TotalSettlementAmount float64     `json:"totalSettlementAmount" orm:"total_settlement_amount" description:"已结算金额"`
	Balance               float64     `json:"balance"               orm:"balance"                 description:"余额"`
	VerifyMoney           float64     `json:"verifyMoney"           orm:"verify_money"            description:"已核账金额"`
	ApplyWithdrawBalance  float64     `json:"applyWithdrawBalance"  orm:"apply_withdraw_balance"  description:"提现中余额"`
	WithdrawBalance       float64     `json:"withdrawBalance"       orm:"withdraw_balance"        description:"已提现余额"`
	CreateAt              *gtime.Time `json:"createAt"              orm:"create_at"               description:"创建时间"`
	UpdateAt              *gtime.Time `json:"updateAt"              orm:"update_at"               description:"更新时间"`
	DeletedAt             *gtime.Time `json:"deletedAt"             orm:"deleted_at"              description:""`
}
