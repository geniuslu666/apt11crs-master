// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriver is the golang structure of table hg_car_driver for DAO operations like Where/Data.
type CarDriver struct {
	g.Meta                `orm:"table:hg_car_driver, do:true"`
	Id                    interface{} //
	CarId                 interface{} // 车辆ID
	IsLeader              interface{} // 是否是车队长 1是  2否
	Name                  interface{} // 司机真实姓名
	Nickname              interface{} // 司机昵称
	CooperateTypeId       interface{} // 合作类型ID
	Sex                   interface{} // 1、男 2、女
	Phone                 interface{} // 手机号
	PhoneArea             interface{} // 手机区号
	Photo                 interface{} // 照片
	Age                   interface{} // 年龄
	WorkYears             interface{} // 从业年数
	Language              interface{} // 语言能力
	Status                interface{} // 状态1、启用 2、禁用
	WorkStatus            interface{} // 工作状态
	MemberId              interface{} // 会员ID
	QualityMaterials      interface{} // 资质信息(多图)
	SettlementType        interface{} // 服务分成类型 1跟随系统  2自定义
	SettlementId          interface{} // 结算模式ID
	SettlementRate        interface{} // 服务分成%
	TotalOrderNum         interface{} // 预约单总数量（包含退款）
	TotalOrderAmount      interface{} // 预约单总金额（包含退款）
	PayOrderNum           interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount        interface{} // 预约单支付金额（不包含退款）
	SettlementOrderNum    interface{} // 预约单已结算数量
	SettlementOrderAmount interface{} // 已结算预约单订单金额
	TotalSettlementAmount interface{} // 已结算金额
	Balance               interface{} // 余额
	VerifyMoney           interface{} // 已核账金额
	ApplyWithdrawBalance  interface{} // 提现中余额
	WithdrawBalance       interface{} // 已提现余额
	CreateAt              *gtime.Time // 创建时间
	UpdateAt              *gtime.Time // 更新时间
	DeletedAt             *gtime.Time //
}
