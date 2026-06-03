// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsStaff is the golang structure of table hg_pms_staff for DAO operations like Where/Data.
type PmsStaff struct {
	g.Meta               `orm:"table:hg_pms_staff, do:true"`
	Id                   interface{} //
	Name                 interface{} // 员工姓名
	Department           interface{} // 员工部门
	Phone                interface{} // 手机号
	Email                interface{} // 邮箱
	Rate                 interface{} // 返佣比例
	Status               interface{} // 状态   1、 开启   2、禁用
	Remark               interface{} // 备注
	Balance              interface{} // 可提现账户
	AllBalance           interface{} // 总账户
	ApplyWithdrawBalance interface{} // 提现中余额
	WithdrawBalance      interface{} // 已提现余额
	MinWithdrawalAmount  interface{} // 最低可提现额
	ServiceCharge        interface{} // 手续费
	AfterDay             interface{} // 预计提现时间周期
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
	DeletedAt            *gtime.Time //
}
