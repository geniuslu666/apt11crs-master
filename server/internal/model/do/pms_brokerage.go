// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsBrokerage is the golang structure of table hg_pms_brokerage for DAO operations like Where/Data.
type PmsBrokerage struct {
	g.Meta    `orm:"table:hg_pms_brokerage, do:true"`
	Id        interface{} //
	Identity  interface{} // 类型    CHANNEL   渠道   STAFF  员工
	Scene     interface{} // 场景值   HOTEL 酒店  SYSTEM 系统
	Type      interface{} // REBATE 返佣  WITHDRAW  提现  SYS 系统调整
	Balance   interface{} // 变动金额
	StaffId   interface{} // 员工ID
	ChannelId interface{} // 渠道ID
	MemberId  interface{} // 会员ID
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
