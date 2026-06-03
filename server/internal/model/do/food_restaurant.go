// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodRestaurant is the golang structure of table hg_food_restaurant for DAO operations like Where/Data.
type FoodRestaurant struct {
	g.Meta                  `orm:"table:hg_food_restaurant, do:true"`
	Id                      interface{} //
	OrderMode               interface{} // 预定模式
	ToretaId                interface{} // Toreta的餐厅ID
	Name                    interface{} // 名称
	CuisineIds              interface{} // 菜系（多选）
	LabelIds                interface{} // 标签（多选）
	CooperateTypeId         interface{} // 合作类型ID
	Logo                    interface{} // LOGO
	Images                  interface{} // 图集
	Content                 interface{} // 简介
	Phone                   interface{} // 联系电话
	OpenTime                interface{} // 营业时间
	AreaPid                 interface{} // 地区省级ID
	AreaId                  interface{} // 地区市级ID
	DetailAddress           interface{} // 详细地址
	GgLat                   interface{} // 谷歌纬度
	GgLng                   interface{} // 谷歌经度
	Lat                     interface{} // 纬度
	Lng                     interface{} // 经度
	OpenStatus              interface{} // 营业状态
	OrderTimeType           interface{} // 预约日期  1每天 2自定义
	OrderTimeWeek           interface{} // 预约日期自定义周数
	RestTimeWeek            interface{} // 定休日周数
	RestTimeDate            interface{} // 定休日固定日期
	OrderConfirmType        interface{} // 预约确定模式 1手动确认 2自动确认(废除)
	OrderConfirmDayType     interface{} // 1 手动确认在前   2自动确认在前
	OrderConfirmBeforeDays  interface{} // 多少天内手动确认
	OrderConfirmTime        interface{} // 预约确认自定义分钟数
	OpenTimeType            interface{} // 1 按时段  2 按时间点
	AmOpenTime              interface{} // 早市开始时间
	AmCloseTime             interface{} // 早市结束时间
	PmOpenTime              interface{} // 晚市开始时间
	PmCloseTime             interface{} // 晚市结束时间
	TimePoints              interface{} // 时间点
	TimeDuration            interface{} // 时间间隔
	OrderMaxType            interface{} // 最大容纳数  1按时段  2按时间点
	TimeDurationMax         interface{} // 时段最大容纳预定数量  0不限制
	DayTimeLimit            interface{} // 当天预约处理截止时间
	AdvanceOrderDay         interface{} // 至少提前预约天数  0当日可约
	MaxOrderDay             interface{} // 最长预约天数
	CancelPolicyOpen        interface{} // 是否允许取消  1允许  2不允许
	BeforeConfirmCancelRate interface{} // 订单确认前取消费用为订单的%，0则免费
	AfterConfirmCancelDay   interface{} // 订单确认后距离到店多少天
	AfterConfirmCancelRate1 interface{} // 订单确认后距离到店天数前取消费用为订单的%，0则免费
	AfterConfirmCancelRate2 interface{} // 订单确认后距离到店天数后取消费用为订单的%，0则免费
	MaxSeat                 interface{} // 最大席位数
	CanSmoking              interface{} // 是否允许抽烟  1允许  2不允许
	SettlementType          interface{} // 门店抽成类型 1跟随系统  2自定义
	SettlementId            interface{} // 结算模式ID
	SettlementRate          interface{} // 门店结算比例
	Desc                    interface{} // 开业/停业原因
	TotalOrderNum           interface{} // 预约单总数量（包含退款）
	TotalOrderAmount        interface{} // 预约单总金额（包含退款）
	PayOrderNum             interface{} // 预约单支付数量（不包含退款）
	PayOrderAmount          interface{} // 预约单支付金额（不包含退款）
	WaitConfirmOrderNum     interface{} // 预约单待确认数量（已支付待确认）
	SettlementOrderNum      interface{} // 预约单已结算数量
	SettlementOrderAmount   interface{} // 已结算预约单订单金额
	TotalSettlementAmount   interface{} // 已结算金额
	VerifyCode              interface{} // 核销码
	CanOrder                interface{} // 是否开放预定  1开放  2关闭
	Sort                    interface{} // 排序(越大越靠前)
	DepositRate             interface{} // 定金比例
	Account                 interface{} // 账号
	PasswordHash            interface{} // 密码
	Salt                    interface{} // 密码盐
	PasswordResetToken      interface{} // 密码重置令牌
	ToretaCancelEnable      interface{} // Toreta是否允许取消 1 允许  2 不允许
	ToretaCancelLimitDay    interface{} // Toreta允许取消几天前
	ToretaCancelLimitTime   interface{} // Toreta允许取消时间前
	MaxDatetimeOrderOpen    interface{} // 日期限制 1 开启  2关闭
	MaxDatetimeOrderDate    interface{} // 限制的日期 逗号分隔
	CreateAt                *gtime.Time // 创建时间
	UpdateAt                *gtime.Time // 更新时间
	DeletedAt               *gtime.Time //
}
