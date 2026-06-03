// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsAppStay is the golang structure for table pms_app_stay.
type PmsAppStay struct {
	Id                  int         `json:"id"                  orm:"id"                     description:"APP订单主键"`
	Uid                 string      `json:"uid"                 orm:"uid"                    description:"我方系统 ID"`
	Uuid                string      `json:"uuid"                orm:"uuid"                   description:"三方系统 ID"`
	Puid                string      `json:"puid"                orm:"puid"                   description:"物业ID"`
	Source              string      `json:"source"              orm:"source"                 description:"订单来源"`
	MemberId            int         `json:"memberId"            orm:"member_id"              description:"用户ID"`
	OrderSn             string      `json:"orderSn"             orm:"order_sn"               description:"订单号"`
	OutOrderSn          string      `json:"outOrderSn"          orm:"out_order_sn"           description:"三方订单号"`
	Booker              string      `json:"booker"              orm:"booker"                 description:"预定人"`
	OrderAmount         float64     `json:"orderAmount"         orm:"order_amount"           description:"订单金额"`
	PayModel            int         `json:"payModel"            orm:"pay_model"              description:"1、余额支付 2、组合支付 3、纯外部支付"`
	OrderStatus         string      `json:"orderStatus"         orm:"order_status"           description:"订单付款状态"`
	ExpirationTime      int         `json:"expirationTime"      orm:"expiration_time"        description:"订单过期时间"`
	CancelTime          *gtime.Time `json:"cancelTime"          orm:"cancel_time"            description:"未支付取消时间"`
	RefundStatus        string      `json:"refundStatus"        orm:"refund_status"          description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundTime          *gtime.Time `json:"refundTime"          orm:"refund_time"            description:"退款时间"`
	RefundAmount        float64     `json:"refundAmount"        orm:"refund_amount"          description:"已退款金额"`
	CleanFee            float64     `json:"cleanFee"            orm:"clean_fee"              description:"取消费用"`
	CheckInDate         string      `json:"checkInDate"         orm:"check_in_date"          description:"入住时间"`
	CheckOutDate        string      `json:"checkOutDate"        orm:"check_out_date"         description:"退房时间"`
	CancelRate          string      `json:"cancelRate"          orm:"cancel_rate"            description:"退款政策"`
	Referrer            int         `json:"referrer"            orm:"referrer"               description:"推荐人"`
	RebateRate          float64     `json:"rebateRate"          orm:"rebate_rate"            description:"分佣比例"`
	RebateStatus        string      `json:"rebateStatus"        orm:"rebate_status"          description:"WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败"`
	RebateAmount        float64     `json:"rebateAmount"        orm:"rebate_amount"          description:"分佣结算金额"`
	RebateTime          *gtime.Time `json:"rebateTime"          orm:"rebate_time"            description:"分佣结算时间"`
	IsGetOpen           string      `json:"isGetOpen"           orm:"is_get_open"            description:"是否开启积分获取"`
	IsPayOpen           string      `json:"isPayOpen"           orm:"is_pay_open"            description:"是否开启积分抵扣"`
	HotelGetRateVip     float64     `json:"hotelGetRateVip"     orm:"hotel_get_rate_vip"     description:"酒店结算积分比例"`
	HotelGetRateScene   float64     `json:"hotelGetRateScene"   orm:"hotel_get_rate_scene"   description:"场景结算积分比例"`
	HotelGetScoreStatus string      `json:"hotelGetScoreStatus" orm:"hotel_get_score_status" description:"'WAIT','SUCCESS','FAIL'"`
	HotelGetAmount      float64     `json:"hotelGetAmount"      orm:"hotel_get_amount"       description:"结算积分金额"`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"             description:"创建时间"`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"             description:"更新时间"`
	PricePercent        int64       `json:"pricePercent"        orm:"price_percent"          description:"全局溢价比例"`
	TotalAmount         float64     `json:"totalAmount"         orm:"total_amount"           description:"订单总价"`
	ChangeAmount        float64     `json:"changeAmount"        orm:"change_amount"          description:"优惠金额"`
	IsFx                string      `json:"isFx"                orm:"is_fx"                  description:"是否是分销订单"`
}
