// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppReservation is the golang structure for table pms_app_reservation.
type PmsAppReservation struct {
	Id               int         `json:"id"               orm:"id"                  description:"主键"`
	Uid              string      `json:"uid"              orm:"uid"                 description:"我方系统 ID"`
	Uuid             string      `json:"uuid"             orm:"uuid"                description:"airhost   ID"`
	Source           string      `json:"source"           orm:"source"              description:"来源"`
	SourceCode       string      `json:"sourceCode"       orm:"source_code"         description:"来源渠道"`
	SourceName       string      `json:"sourceName"       orm:"source_name"         description:"渠道名称"`
	MemberId         int         `json:"memberId"         orm:"member_id"           description:"用户ID"`
	Puid             string      `json:"puid"             orm:"puid"                description:"物业ID"`
	OrderSn          string      `json:"orderSn"          orm:"order_sn"            description:"系统订单号"`
	OrderIndex       int         `json:"orderIndex"       orm:"order_index"         description:"订单索引"`
	OutOrderSn       string      `json:"outOrderSn"       orm:"out_order_sn"        description:"三方订单号"`
	RoomType         string      `json:"roomType"         orm:"room_type"           description:"房型信息，参考房型uid"`
	RoomUnit         string      `json:"roomUnit"         orm:"room_unit"           description:"房间单元的id, uid或组合"`
	RatePlanId       string      `json:"ratePlanId"       orm:"rate_plan_id"        description:"费率ID"`
	CheckinDate      *gtime.Time `json:"checkinDate"      orm:"checkin_date"        description:"入住日期"`
	CheckoutDate     *gtime.Time `json:"checkoutDate"     orm:"checkout_date"       description:"退房日期"`
	CheckinTime      string      `json:"checkinTime"      orm:"checkin_time"        description:"入住时间，24小时格式"`
	CheckoutTime     string      `json:"checkoutTime"     orm:"checkout_time"       description:"退房时间，24小时格式"`
	Status           string      `json:"status"           orm:"status"              description:"预订状态（确认/confirmed、取消/cancelled）"`
	CheckinStatus    string      `json:"checkinStatus"    orm:"checkin_status"      description:"入住状态  before_checkin  在入住之前  checked_in  已入住  checked_out  已退房"`
	OrderStatus      string      `json:"orderStatus"      orm:"order_status"        description:"WAIT_PAY、待支付 CANCEL、支付过期 HAVE_PAID、支付成功"`
	MainGuest        string      `json:"mainGuest"        orm:"main_guest"          description:"住宿人编号"`
	AdultCount       int         `json:"adultCount"       orm:"adult_count"         description:"成人数量"`
	ChildCount       int         `json:"childCount"       orm:"child_count"         description:"儿童数量"`
	InfantCount      int         `json:"infantCount"      orm:"infant_count"        description:"婴儿数量"`
	PricePlanId      int         `json:"pricePlanId"      orm:"price_plan_id"       description:"价格plan ID"`
	ChangeAmount     float64     `json:"changeAmount"     orm:"change_amount"       description:"变动金额"`
	PricePlanInfo    *gjson.Json `json:"pricePlanInfo"    orm:"price_plan_info"     description:"价格plan内容"`
	BookingFee       float64     `json:"bookingFee"       orm:"booking_fee"         description:"预订费"`
	ChannelFee       float64     `json:"channelFee"       orm:"channel_fee"         description:"渠道费"`
	CleaningFee      float64     `json:"cleaningFee"      orm:"cleaning_fee"        description:"清洁费"`
	CancellationFee  float64     `json:"cancellationFee"  orm:"cancellation_fee"    description:"取消费，仅在取消时适用"`
	Charges          string      `json:"charges"          orm:"charges"             description:"费用详情，参考Charge对象"`
	GuestRemarks     string      `json:"guestRemarks"     orm:"guest_remarks"       description:"备注"`
	CancelRemake     string      `json:"cancelRemake"     orm:"cancel_remake"       description:"取消原因"`
	ExpValue         float64     `json:"expValue"         orm:"exp_value"           description:"结算的经验值"`
	ExpTime          *gtime.Time `json:"expTime"          orm:"exp_time"            description:"经验结算时间"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"          description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"          description:""`
	IsChangeGuest    string      `json:"isChangeGuest"    orm:"is_change_guest"     description:"是否变更入住人信息 Y 是 N 否"`
	IsChangeGuestId  int         `json:"isChangeGuestId"  orm:"is_change_guest_id"  description:"入住人信息变更ID"`
	IsChangePeople   string      `json:"isChangePeople"   orm:"is_change_people"    description:"是否变更入住人数 Y 是 N 否"`
	IsChangePeopleId int         `json:"isChangePeopleId" orm:"is_change_people_id" description:"入住人数信息变更ID"`
	IsChangeDate     string      `json:"isChangeDate"     orm:"is_change_date"      description:"是否变更入住日期 Y 是 N 否"`
	IsChangeDateId   int         `json:"isChangeDateId"   orm:"is_change_date_id"   description:"入住日期变更ID"`
	IsFx             string      `json:"isFx"             orm:"is_fx"               description:"是否是分销订单"`
}
