// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsCharge is the golang structure of table hg_pms_charge for DAO operations like Where/Data.
type PmsCharge struct {
	g.Meta         `orm:"table:hg_pms_charge, do:true"`
	Id             interface{} // 主键
	Uid            interface{} // 三方系统 ID
	AirUid         interface{} // airhost uid
	Date           *gtime.Time // 费用日期，可以是发生日期或记账日期
	Name           interface{} // 名称
	FeeType        interface{} // 费用类型（预订费、餐饮费、清洁费、取消费、其他费用）
	Amount         interface{} // 金额
	Currency       interface{} // 货币
	OriginalAmount interface{} // 原价
	Description    interface{} // 费用详情
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
