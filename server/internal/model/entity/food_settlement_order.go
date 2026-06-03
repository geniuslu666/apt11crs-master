// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodSettlementOrder is the golang structure for table food_settlement_order.
type FoodSettlementOrder struct {
	Id               int64       `json:"id"               orm:"id"                description:""`
	OrderSn          string      `json:"orderSn"          orm:"order_sn"          description:"结算单编号"`
	RestaurantId     int         `json:"restaurantId"     orm:"restaurant_id"     description:"餐厅ID"`
	OrderAmount      float64     `json:"orderAmount"      orm:"order_amount"      description:"订单总额"`
	SettlementAmount float64     `json:"settlementAmount" orm:"settlement_amount" description:"结算总额"`
	Status           string      `json:"status"           orm:"status"            description:"结算状态"`
	StartTime        *gtime.Time `json:"startTime"        orm:"start_time"        description:"账期开始时间"`
	EndTime          *gtime.Time `json:"endTime"          orm:"end_time"          description:"账期结束时间"`
	VerifyStatus     string      `json:"verifyStatus"     orm:"verify_status"     description:"核账状态"`
	VerifyTime       *gtime.Time `json:"verifyTime"       orm:"verify_time"       description:"核帐时间"`
	VerifyImg        string      `json:"verifyImg"        orm:"verify_img"        description:"核账凭证"`
	VerifyDesc       string      `json:"verifyDesc"       orm:"verify_desc"       description:"核账说明"`
	VerifyOperateId  string      `json:"verifyOperateId"  orm:"verify_operate_id" description:"核账操作人ID"`
	CreateAt         *gtime.Time `json:"createAt"         orm:"create_at"         description:"创建时间"`
	UpdateAt         *gtime.Time `json:"updateAt"         orm:"update_at"         description:"更新时间"`
}
