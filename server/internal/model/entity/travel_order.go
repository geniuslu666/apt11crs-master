// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TravelOrder is the golang structure for table travel_order.
type TravelOrder struct {
	Id                      uint64      `json:"id"                      orm:"id"                         description:""`
	OrderSn                 string      `json:"orderSn"                 orm:"order_sn"                   description:"预约单号"`
	ProductId               uint64      `json:"productId"               orm:"product_id"                 description:"产品ID"`
	SkuId                   uint64      `json:"skuId"                   orm:"sku_id"                     description:"SKUID"`
	MemberId                uint64      `json:"memberId"                orm:"member_id"                  description:"会员ID"`
	BookingName             string      `json:"bookingName"             orm:"booking_name"               description:"预订人姓名"`
	FirstName               string      `json:"firstName"               orm:"first_name"                 description:"订单预定人名"`
	LastName                string      `json:"lastName"                orm:"last_name"                  description:"订单预定人姓"`
	PhoneArea               string      `json:"phoneArea"               orm:"phone_area"                 description:"手机区号"`
	BookingMobile           string      `json:"bookingMobile"           orm:"booking_mobile"             description:"预订人电话"`
	BookingEmail            string      `json:"bookingEmail"            orm:"booking_email"              description:"预定人邮箱"`
	BookingNum              uint        `json:"bookingNum"              orm:"booking_num"                description:"预约人数"`
	BookDate                *gtime.Time `json:"bookDate"                orm:"book_date"                  description:"预约日期"`
	OrderAmount             float64     `json:"orderAmount"             orm:"order_amount"               description:"订单金额（元）"`
	CouponAmount            float64     `json:"couponAmount"            orm:"coupon_amount"              description:"优惠券抵扣金额"`
	BalAmount               float64     `json:"balAmount"               orm:"bal_amount"                 description:"积分抵扣金额"`
	OrderStatus             string      `json:"orderStatus"             orm:"order_status"               description:"订单状态"`
	PayModel                int         `json:"payModel"                orm:"pay_model"                  description:"1、余额支付 2、组合支付 3、纯外部支付"`
	PayStatus               string      `json:"payStatus"               orm:"pay_status"                 description:"订单付款状态"`
	PayTime                 *gtime.Time `json:"payTime"                 orm:"pay_time"                   description:"支付时间"`
	VerifyStaffId           uint64      `json:"verifyStaffId"           orm:"verify_staff_id"            description:"核销人员ID"`
	VerifyTime              *gtime.Time `json:"verifyTime"              orm:"verify_time"                description:"核销时间"`
	ExpirationTime          int         `json:"expirationTime"          orm:"expiration_time"            description:"订单过期时间"`
	CancelTime              *gtime.Time `json:"cancelTime"              orm:"cancel_time"                description:"取消时间"`
	CancelFee               float64     `json:"cancelFee"               orm:"cancel_fee"                 description:"取消手续费（元）"`
	RefundStatus            string      `json:"refundStatus"            orm:"refund_status"              description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundAmount            float64     `json:"refundAmount"            orm:"refund_amount"              description:"退款金额（元）"`
	RefundTime              *gtime.Time `json:"refundTime"              orm:"refund_time"                description:"退款时间"`
	RefundBalAmount         float64     `json:"refundBalAmount"         orm:"refund_bal_amount"          description:"已退款积分"`
	RefundCouponAmount      float64     `json:"refundCouponAmount"      orm:"refund_coupon_amount"       description:"已退款优惠券"`
	RefundReason            string      `json:"refundReason"            orm:"refund_reason"              description:"退款原因"`
	CreatedAt               *gtime.Time `json:"createdAt"               orm:"created_at"                 description:"创建时间"`
	UpdatedAt               *gtime.Time `json:"updatedAt"               orm:"updated_at"                 description:"更新时间"`
	Referrer                int         `json:"referrer"                orm:"referrer"                   description:"推荐人"`
	RebateRate              float64     `json:"rebateRate"              orm:"rebate_rate"                description:"分佣比例"`
	RebateStatus            string      `json:"rebateStatus"            orm:"rebate_status"              description:"WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败"`
	RebateAmount            float64     `json:"rebateAmount"            orm:"rebate_amount"              description:"分佣结算金额"`
	RebateTime              *gtime.Time `json:"rebateTime"              orm:"rebate_time"                description:"分佣结算时间"`
	IsGetOpen               string      `json:"isGetOpen"               orm:"is_get_open"                description:"是否开启积分获取"`
	IsPayOpen               string      `json:"isPayOpen"               orm:"is_pay_open"                description:"是否开启积分抵扣"`
	TravelGetRateVip        float64     `json:"travelGetRateVip"        orm:"travel_get_rate_vip"        description:"结算积分比例"`
	TravelGetRateScene      float64     `json:"travelGetRateScene"      orm:"travel_get_rate_scene"      description:"场景结算积分比例"`
	TravelGetScoreStatus    string      `json:"travelGetScoreStatus"    orm:"travel_get_score_status"    description:"'WAIT','SUCCESS','FAIL'"`
	TravelGetAmount         float64     `json:"travelGetAmount"         orm:"travel_get_amount"          description:"结算积分金额"`
	ExpValue                float64     `json:"expValue"                orm:"exp_value"                  description:"结算的经验值"`
	ExpTime                 *gtime.Time `json:"expTime"                 orm:"exp_time"                   description:"经验结算时间"`
	AdminRefundAmount       float64     `json:"adminRefundAmount"       orm:"admin_refund_amount"        description:"后台已退款总金额"`
	AdminRefundBalAmount    float64     `json:"adminRefundBalAmount"    orm:"admin_refund_bal_amount"    description:"后台已退款积分"`
	AdminRefundCouponAmount float64     `json:"adminRefundCouponAmount" orm:"admin_refund_coupon_amount" description:"后台已退款优惠券"`
	AdminCancelReason       string      `json:"adminCancelReason"       orm:"admin_cancel_reason"        description:"后台取消原因"`
	AdminCancelNum          int         `json:"adminCancelNum"          orm:"admin_cancel_num"           description:"后台取消次数"`
	IsFx                    string      `json:"isFx"                    orm:"is_fx"                      description:"是否是分销订单"`
}
