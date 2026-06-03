// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsPricePlan is the golang structure of table hg_pms_price_plan for DAO operations like Where/Data.
type PmsPricePlan struct {
	g.Meta          `orm:"table:hg_pms_price_plan, do:true"`
	Id              interface{} //
	PlanName        interface{} // 价格plan名称
	PlanShowName    interface{} // 展示标签 多语言
	PropertyId      interface{} // 物业ID
	RoomTypeId      interface{} // 房型ID
	BookingDays     interface{} // 预订天数
	MemberGroupId   interface{} // 用户组
	MemberLevelId   interface{} // 会员等级
	IsCancel        interface{} // 是否可取消   Y  是  N  否
	IsOpenPriceMode interface{} // 是否开启价格模式
	PriceMode       interface{} // 模式    +  贵   - 便宜
	PriceStandard   interface{} // 基准  PERCENT 倍率  AMOUNT  金额
	PlanValue       interface{} // 价格基准值
	PlanTips        interface{} // 价格计划提示
	PricePlanStatus interface{} // Y 开启  N 关闭
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
