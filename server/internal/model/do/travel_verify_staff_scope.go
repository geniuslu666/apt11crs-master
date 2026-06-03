// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelVerifyStaffScope is the golang structure of table hg_travel_verify_staff_scope for DAO operations like Where/Data.
type TravelVerifyStaffScope struct {
	g.Meta    `orm:"table:hg_travel_verify_staff_scope, do:true"`
	Id        interface{} //
	StaffId   interface{} // 核销人员ID
	ProductId interface{} // 产品ID
	SkuId     interface{} // SKU ID（0=产品全部SKU）
	IsAll     interface{} // 是否全部可核销（1是 0否）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 软删除时间（NULL=正常）
}
