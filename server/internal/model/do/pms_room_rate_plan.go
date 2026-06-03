// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsRoomRatePlan is the golang structure of table hg_pms_room_rate_plan for DAO operations like Where/Data.
type PmsRoomRatePlan struct {
	g.Meta     `orm:"table:hg_pms_room_rate_plan, do:true"`
	Id         interface{} //
	Tuid       interface{} // 房型ID
	TName      interface{} // 房型名
	RateName   interface{} // 费率名称
	RatePlanId interface{} // 价格计划
	IsUse      interface{} // 是否使用
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
