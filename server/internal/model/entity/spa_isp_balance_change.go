// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaIspBalanceChange is the golang structure for table spa_isp_balance_change.
type SpaIspBalanceChange struct {
	Id                int         `json:"id"                orm:"id"                  description:"主键"`
	IspId             int         `json:"ispId"             orm:"isp_id"              description:"服务商ID"`
	Type              string      `json:"type"              orm:"type"                description:"金额变动方式   SETTLEMENT 结算   VERIFY核账    RETURN_ORDER退单    SYS 系统调整"`
	ChangePrice       float64     `json:"changePrice"       orm:"change_price"        description:"变更金额"`
	SettlementOrderId int         `json:"settlementOrderId" orm:"settlement_order_id" description:"结算单ID"`
	WithdrawOrderId   int         `json:"withdrawOrderId"   orm:"withdraw_order_id"   description:"提现单ID"`
	Des               string      `json:"des"               orm:"des"                 description:"描述"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
}
