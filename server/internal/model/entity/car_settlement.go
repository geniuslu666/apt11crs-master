// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarSettlement is the golang structure for table car_settlement.
type CarSettlement struct {
	Id              int64       `json:"id"              orm:"id"                description:""`
	Name            string      `json:"name"            orm:"name"              description:"名称"`
	Type            int         `json:"type"            orm:"type"              description:"结算方式  1无需结算 2按周期自动结算  3手动申请结算"`
	Cycle           int         `json:"cycle"           orm:"cycle"             description:"结算周期  1每日结算  2每周结算  3每月结算"`
	Rate            float64     `json:"rate"            orm:"rate"              description:"结算比例"`
	Cost            string      `json:"cost"            orm:"cost"              description:"结算成本控制  1结算优惠券  2结算积分抵扣  3结算扣除佣金"`
	IsAuditWithdraw int         `json:"isAuditWithdraw" orm:"is_audit_withdraw" description:"是否提现审核  1开启  2关闭"`
	Status          uint        `json:"status"          orm:"status"            description:"1、启用 2、禁用"`
	CreateAt        *gtime.Time `json:"createAt"        orm:"create_at"         description:"创建时间"`
	UpdateAt        *gtime.Time `json:"updateAt"        orm:"update_at"         description:"更新时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:""`
}
