// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriverBalanceChange is the golang structure of table hg_car_driver_balance_change for DAO operations like Where/Data.
type CarDriverBalanceChange struct {
	g.Meta            `orm:"table:hg_car_driver_balance_change, do:true"`
	Id                interface{} // 主键
	DriverId          interface{} // 司机ID
	Type              interface{} // 金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整
	ChangePrice       interface{} // 变更金额
	SettlementOrderId interface{} // 接送机结算单ID
	WithdrawOrderId   interface{} // 提现单ID
	ReturnOrderId     interface{} // 退单ID
	Des               interface{} // 描述
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
}
