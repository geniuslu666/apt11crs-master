// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaSettlement is the golang structure of table hg_spa_settlement for DAO operations like Where/Data.
type SpaSettlement struct {
	g.Meta    `orm:"table:hg_spa_settlement, do:true"`
	Id        interface{} //
	Name      interface{} // 名称
	Type      interface{} // 结算方式  1无需结算 2按周期自动结算
	Cycle     interface{} // 结算周期  1每日结算  2每周结算  3每月结算
	Rate      interface{} // 结算比例
	Cost      interface{} // 结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金
	Status    interface{} // 1、启用 2、禁用
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 更新时间
	DeletedAt *gtime.Time //
}
