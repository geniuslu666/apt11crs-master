// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppReservationChange is the golang structure of table hg_pms_app_reservation_change for DAO operations like Where/Data.
type PmsAppReservationChange struct {
	g.Meta          `orm:"table:hg_pms_app_reservation_change, do:true"`
	Id              interface{} // 主键
	ChangeOrderSn   interface{} // 变更订单号
	OrderId         interface{} // 订单ID
	OrderSn         interface{} // 预订订单号
	OutOrderSn      interface{} // 预订外部订单号
	ChangeType      interface{} // 变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更
	OldCheckinDate  *gtime.Time // 入住日期
	OldCheckoutDate *gtime.Time // 退房日期
	NewCheckinDate  *gtime.Time // 入住日期
	NewCheckoutDate *gtime.Time // 退房日期
	OldMainGuest    *gjson.Json // 住宿人编号
	NewMainGuest    *gjson.Json // 住宿人编号
	OldAdultCount   interface{} // 成人数量
	NewAdultCount   interface{} // 成人数量
	OldChildCount   interface{} // 儿童数量
	NewChildCount   interface{} // 儿童数量
	OldInfantCount  interface{} // 婴儿数量
	NewInfantCount  interface{} // 婴儿数量
	ChangeStatus    interface{} // 变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败
	ChangeAmount    interface{} // 变动金额
	SubmitDate      *gtime.Time // 提交变更时间
	DoneDate        *gtime.Time // 变更成功时间
	OldOrderPrice   interface{} // 原订单价格
	NewOrderPrice   interface{} // 变更后订单价格
	ExpirationTime  interface{} // 订单过期时间
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
	PricePercent    interface{} // 全局溢价比例
	IsFx            interface{} // 是否是分销订单
}
