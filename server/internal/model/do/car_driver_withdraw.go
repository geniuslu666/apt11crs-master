// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriverWithdraw is the golang structure of table hg_car_driver_withdraw for DAO operations like Where/Data.
type CarDriverWithdraw struct {
	g.Meta         `orm:"table:hg_car_driver_withdraw, do:true"`
	Id             interface{} //
	Type           interface{} // 类型
	DriverId       interface{} // 司机ID
	WithdrawSn     interface{} // 提现单号
	WithdrawStatus interface{} // 提现状态
	WithdrawAmount interface{} // 提现金额
	ArrivalAmount  interface{} // 到账金额
	ServiceCharge  interface{} // 提现手续费比例
	Transfer       interface{} // 1、未转账 2、已转账
	ApplyRemark    interface{} // 审核备注
	ApplyAt        *gtime.Time // 审核时间
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
