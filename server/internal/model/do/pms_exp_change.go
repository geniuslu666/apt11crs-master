// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsExpChange is the golang structure of table hg_pms_exp_change for DAO operations like Where/Data.
type PmsExpChange struct {
	g.Meta     `orm:"table:hg_pms_exp_change, do:true"`
	Id         interface{} //
	Scene      interface{} // 场景值   HOTEL 酒店  SYSTEM 系统
	Exp        interface{} // 变动金额
	OrderSn    interface{} // 订单号
	Des        interface{} // 描述
	MemberId   interface{} // 会员ID
	OperatorId interface{} // 操作员ID
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
