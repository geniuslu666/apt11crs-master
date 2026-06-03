// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppReservationChange is the golang structure for table pms_app_reservation_change.
type PmsAppReservationChange struct {
	Id              int         `json:"id"              orm:"id"                description:"主键"`
	ChangeOrderSn   string      `json:"changeOrderSn"   orm:"change_order_sn"   description:"变更订单号"`
	OrderId         int         `json:"orderId"         orm:"order_id"          description:"订单ID"`
	OrderSn         string      `json:"orderSn"         orm:"order_sn"          description:"预订订单号"`
	OutOrderSn      string      `json:"outOrderSn"      orm:"out_order_sn"      description:"预订外部订单号"`
	ChangeType      string      `json:"changeType"      orm:"change_type"       description:"变更内容  GUEST 预定人信息变更  PEOPLE   入住人数变更  DATE   日期变更"`
	OldCheckinDate  *gtime.Time `json:"oldCheckinDate"  orm:"old_checkin_date"  description:"入住日期"`
	OldCheckoutDate *gtime.Time `json:"oldCheckoutDate" orm:"old_checkout_date" description:"退房日期"`
	NewCheckinDate  *gtime.Time `json:"newCheckinDate"  orm:"new_checkin_date"  description:"入住日期"`
	NewCheckoutDate *gtime.Time `json:"newCheckoutDate" orm:"new_checkout_date" description:"退房日期"`
	OldMainGuest    *gjson.Json `json:"oldMainGuest"    orm:"old_main_guest"    description:"住宿人编号"`
	NewMainGuest    *gjson.Json `json:"newMainGuest"    orm:"new_main_guest"    description:"住宿人编号"`
	OldAdultCount   int         `json:"oldAdultCount"   orm:"old_adult_count"   description:"成人数量"`
	NewAdultCount   int         `json:"newAdultCount"   orm:"new_adult_count"   description:"成人数量"`
	OldChildCount   int         `json:"oldChildCount"   orm:"old_child_count"   description:"儿童数量"`
	NewChildCount   int         `json:"newChildCount"   orm:"new_child_count"   description:"儿童数量"`
	OldInfantCount  int         `json:"oldInfantCount"  orm:"old_infant_count"  description:"婴儿数量"`
	NewInfantCount  int         `json:"newInfantCount"  orm:"new_infant_count"  description:"婴儿数量"`
	ChangeStatus    string      `json:"changeStatus"    orm:"change_status"     description:"变动状态 ING   处理中   DONE   变更完成   FAIL   变更失败"`
	ChangeAmount    float64     `json:"changeAmount"    orm:"change_amount"     description:"变动金额"`
	SubmitDate      *gtime.Time `json:"submitDate"      orm:"submit_date"       description:"提交变更时间"`
	DoneDate        *gtime.Time `json:"doneDate"        orm:"done_date"         description:"变更成功时间"`
	OldOrderPrice   float64     `json:"oldOrderPrice"   orm:"old_order_price"   description:"原订单价格"`
	NewOrderPrice   float64     `json:"newOrderPrice"   orm:"new_order_price"   description:"变更后订单价格"`
	ExpirationTime  int         `json:"expirationTime"  orm:"expiration_time"   description:"订单过期时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:""`
	PricePercent    int64       `json:"pricePercent"    orm:"price_percent"     description:"全局溢价比例"`
	IsFx            string      `json:"isFx"            orm:"is_fx"             description:"是否是分销订单"`
}
