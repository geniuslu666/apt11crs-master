// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBalanceChange is the golang structure of table hg_pms_balance_change for DAO operations like Where/Data.
type PmsBalanceChange struct {
	g.Meta      `orm:"table:hg_pms_balance_change, do:true"`
	Id          interface{} // 主键
	MemberId    interface{} // 会员ID
	Scene       interface{} // 场景值   HOTEL    酒店   SYSTEM 系统
	Type        interface{} // 金额变动方式   CONSUME 消费   REFUND   退款    AWARD  奖励    BROKERAGE   佣金    SYS 系统调整
	ChangePrice interface{} // 变更金额
	OrderSn     interface{} // 订单号
	Des         interface{} // 消费描述（后端展示）
	Reason      interface{} // 原因（前端展示）
	OperatorId  interface{} // 操作员ID
	MdCode      interface{} // 注册设备码
	MpModel     interface{} // 注册设备型号
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
