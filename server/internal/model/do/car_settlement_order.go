// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CarSettlementOrder is the golang structure of table hg_car_settlement_order for DAO operations like Where/Data.
type CarSettlementOrder struct {
	g.Meta           `orm:"table:hg_car_settlement_order, do:true"`
	Id               interface{} //
	OrderSn          interface{} // 结算单编号
	DriverId         interface{} // 司机ID
	OrderAmount      interface{} // 订单总额
	SettlementAmount interface{} // 结算总额
	Status           interface{} // 结算状态
	StartTime        *gtime.Time // 账期开始时间
	EndTime          *gtime.Time // 账期结束时间
	VerifyStatus     interface{} // 核账状态
	VerifyTime       *gtime.Time // 核帐时间
	VerifyImg        interface{} // 核账凭证
	VerifyDesc       interface{} // 核账说明
	VerifyOperateId  interface{} // 核账操作人ID
	CreateAt         *gtime.Time // 创建时间
	UpdateAt         *gtime.Time // 更新时间
}
