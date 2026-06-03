// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaIsp is the golang structure of table hg_spa_isp for DAO operations like Where/Data.
type SpaIsp struct {
	g.Meta                `orm:"table:hg_spa_isp, do:true"`
	Id                    interface{} //
	Name                  interface{} // 服务商名称
	Status                interface{} // 状态1、启用 2、禁用
	WorkStatus            interface{} // 工作状态
	MemberId              interface{} // 会员ID
	SettlementObject      interface{} // 结算对象
	SettlementType        interface{} // 服务分成类型 1跟随系统  2自定义
	SettlementId          interface{} // 结算模式ID
	SettlementRate        interface{} // 服务分成%
	SettlementOrderNum    interface{} // 预约单已结算数量
	SettlementOrderAmount interface{} // 已结算预约单订单金额
	TotalSettlementAmount interface{} // 已结算金额
	Balance               interface{} // 余额
	VerifyMoney           interface{} // 已核账金额
	ApplyWithdrawBalance  interface{} // 提现中余额
	WithdrawBalance       interface{} // 已提现余额
	SendSmsPhone          interface{} // 发送短信电话
	CreateAt              *gtime.Time // 创建时间
	UpdateAt              *gtime.Time // 更新时间
	DeletedAt             *gtime.Time //
}
