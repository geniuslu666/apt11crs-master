// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CarDriverBalanceChange is the golang structure for table car_driver_balance_change.
type CarDriverBalanceChange struct {
	Id                int         `json:"id"                orm:"id"                  description:"主键"`
	DriverId          int         `json:"driverId"          orm:"driver_id"           description:"司机ID"`
	Type              string      `json:"type"              orm:"type"                description:"金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整"`
	ChangePrice       float64     `json:"changePrice"       orm:"change_price"        description:"变更金额"`
	SettlementOrderId int         `json:"settlementOrderId" orm:"settlement_order_id" description:"接送机结算单ID"`
	WithdrawOrderId   int         `json:"withdrawOrderId"   orm:"withdraw_order_id"   description:"提现单ID"`
	ReturnOrderId     int         `json:"returnOrderId"     orm:"return_order_id"     description:"退单ID"`
	Des               string      `json:"des"               orm:"des"                 description:"描述"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
}
