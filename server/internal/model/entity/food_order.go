// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FoodOrder is the golang structure for table food_order.
type FoodOrder struct {
	Id                       int64       `json:"id"                       orm:"id"                         description:""`
	OrderType                string      `json:"orderType"                orm:"order_type"                 description:"订单类型"`
	OrderSn                  string      `json:"orderSn"                  orm:"order_sn"                   description:"订单编号"`
	OutOrderSn               string      `json:"outOrderSn"               orm:"out_order_sn"               description:"三方订单号"`
	MemberId                 uint        `json:"memberId"                 orm:"member_id"                  description:"用户ID"`
	RestaurantId             int         `json:"restaurantId"             orm:"restaurant_id"              description:"餐厅ID"`
	OrderAmount              float64     `json:"orderAmount"              orm:"order_amount"               description:"订单金额"`
	CouponAmount             float64     `json:"couponAmount"             orm:"coupon_amount"              description:"优惠券抵扣金额"`
	BalAmount                float64     `json:"balAmount"                orm:"bal_amount"                 description:"积分抵扣金额"`
	GoodsId                  uint        `json:"goodsId"                  orm:"goods_id"                   description:"套餐ID"`
	BookingName              string      `json:"bookingName"              orm:"booking_name"               description:"预定人姓名"`
	FirstName                string      `json:"firstName"                orm:"first_name"                 description:"订单预定人姓"`
	LastName                 string      `json:"lastName"                 orm:"last_name"                  description:"订单预定人名"`
	GoodsNum                 int         `json:"goodsNum"                 orm:"goods_num"                  description:"套餐数量"`
	PhoneArea                string      `json:"phoneArea"                orm:"phone_area"                 description:"手机区号"`
	BookingMobile            string      `json:"bookingMobile"            orm:"booking_mobile"             description:"预定人手机"`
	BookingEmail             string      `json:"bookingEmail"             orm:"booking_email"              description:"预定人邮箱"`
	BookDate                 string      `json:"bookDate"                 orm:"book_date"                  description:"预定日期"`
	BookTime                 string      `json:"bookTime"                 orm:"book_time"                  description:"预定时间"`
	BookDatetime             *gtime.Time `json:"bookDatetime"             orm:"book_datetime"              description:"预定日期时间"`
	SeatId                   int         `json:"seatId"                   orm:"seat_id"                    description:"座位ID"`
	PayModel                 int         `json:"payModel"                 orm:"pay_model"                  description:"1、余额支付 2、组合支付 3、纯外部支付"`
	PayTime                  *gtime.Time `json:"payTime"                  orm:"pay_time"                   description:"支付时间"`
	OrderStatus              string      `json:"orderStatus"              orm:"order_status"               description:"订单付款状态"`
	BookingStatus            string      `json:"bookingStatus"            orm:"booking_status"             description:"订单预定状态"`
	BookingTime              *gtime.Time `json:"bookingTime"              orm:"booking_time"               description:"订单确认时间"`
	ConfirmRefuseReason      string      `json:"confirmRefuseReason"      orm:"confirm_refuse_reason"      description:"审核拒绝原因"`
	VerifyStatus             string      `json:"verifyStatus"             orm:"verify_status"              description:"订单核销状态"`
	VerifyCode               string      `json:"verifyCode"               orm:"verify_code"                description:"核销码"`
	VerifyTime               *gtime.Time `json:"verifyTime"               orm:"verify_time"                description:"核销时间"`
	ExpirationTime           int         `json:"expirationTime"           orm:"expiration_time"            description:"订单过期时间"`
	SettlementRate           float64     `json:"settlementRate"           orm:"settlement_rate"            description:"结算比例"`
	SettlementStatus         string      `json:"settlementStatus"         orm:"settlement_status"          description:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount         float64     `json:"settlementAmount"         orm:"settlement_amount"          description:"结算金额"`
	SettlementTime           *gtime.Time `json:"settlementTime"           orm:"settlement_time"            description:"结算时间"`
	SettlementOrderId        uint        `json:"settlementOrderId"        orm:"settlement_order_id"        description:"结算单ID"`
	SettlementType           int         `json:"settlementType"           orm:"settlement_type"            description:"结算方式  1无需结算 2按周期自动结算  3手动申请结算"`
	SettlementCycle          int         `json:"settlementCycle"          orm:"settlement_cycle"           description:"结算周期  1每日结算  2每周结算  3每月结算"`
	MemberMessage            string      `json:"memberMessage"            orm:"member_message"             description:"购买人留言信息"`
	MemberMessageJa          string      `json:"memberMessageJa"          orm:"member_message_ja"          description:"购买人留言信息日语版"`
	RestaurantMessage        string      `json:"restaurantMessage"        orm:"restaurant_message"         description:"餐厅留言信息"`
	BookingCount             int         `json:"bookingCount"             orm:"booking_count"              description:"预定人数"`
	RefundStatus             string      `json:"refundStatus"             orm:"refund_status"              description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundRate               float64     `json:"refundRate"               orm:"refund_rate"                description:"退款比例（废弃）"`
	RefundFee                float64     `json:"refundFee"                orm:"refund_fee"                 description:"退款手续费"`
	RefundTime               *gtime.Time `json:"refundTime"               orm:"refund_time"                description:"退款时间"`
	RefundAmount             float64     `json:"refundAmount"             orm:"refund_amount"              description:"已退款总金额"`
	RefundBalAmount          float64     `json:"refundBalAmount"          orm:"refund_bal_amount"          description:"已退款积分"`
	RefundCouponAmount       float64     `json:"refundCouponAmount"       orm:"refund_coupon_amount"       description:"已退款优惠券"`
	AdminRefundAmount        float64     `json:"adminRefundAmount"        orm:"admin_refund_amount"        description:"后台已退款总金额"`
	AdminRefundBalAmount     float64     `json:"adminRefundBalAmount"     orm:"admin_refund_bal_amount"    description:"后台已退款积分"`
	AdminRefundCouponAmount  float64     `json:"adminRefundCouponAmount"  orm:"admin_refund_coupon_amount" description:"后台已退款优惠券"`
	AdminCancelReason        string      `json:"adminCancelReason"        orm:"admin_cancel_reason"        description:"后台取消原因"`
	ActivityId               int         `json:"activityId"               orm:"activity_id"                description:"活动ID"`
	AdminCancelNum           int         `json:"adminCancelNum"           orm:"admin_cancel_num"           description:"后台取消次数"`
	Referrer                 int         `json:"referrer"                 orm:"referrer"                   description:"推荐人"`
	RebateRate               float64     `json:"rebateRate"               orm:"rebate_rate"                description:"分佣比例"`
	RebateStatus             string      `json:"rebateStatus"             orm:"rebate_status"              description:"WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败"`
	RebateAmount             float64     `json:"rebateAmount"             orm:"rebate_amount"              description:"分佣结算金额"`
	RebateTime               *gtime.Time `json:"rebateTime"               orm:"rebate_time"                description:"分佣结算时间"`
	IsGetOpen                string      `json:"isGetOpen"                orm:"is_get_open"                description:"是否开启积分获取"`
	IsPayOpen                string      `json:"isPayOpen"                orm:"is_pay_open"                description:"是否开启积分抵扣"`
	FoodGetRateVip           float64     `json:"foodGetRateVip"           orm:"food_get_rate_vip"          description:"餐饮结算积分比例"`
	FoodGetRateScene         float64     `json:"foodGetRateScene"         orm:"food_get_rate_scene"        description:"场景结算积分比例"`
	FoodGetScoreStatus       string      `json:"foodGetScoreStatus"       orm:"food_get_score_status"      description:"'WAIT','SUCCESS','FAIL'"`
	FoodGetAmount            float64     `json:"foodGetAmount"            orm:"food_get_amount"            description:"结算积分金额"`
	ExpValue                 float64     `json:"expValue"                 orm:"exp_value"                  description:"结算的经验值"`
	ExpTime                  *gtime.Time `json:"expTime"                  orm:"exp_time"                   description:"经验结算时间"`
	DepositRate              float64     `json:"depositRate"              orm:"deposit_rate"               description:"定金比例"`
	DepositAmount            float64     `json:"depositAmount"            orm:"deposit_amount"             description:"定金"`
	DepositPayStatus         string      `json:"depositPayStatus"         orm:"deposit_pay_status"         description:"定金付款状态"`
	DepositPayTime           *gtime.Time `json:"depositPayTime"           orm:"deposit_pay_time"           description:"定金支付时间"`
	DepositCancelTime        *gtime.Time `json:"depositCancelTime"        orm:"deposit_cancel_time"        description:"定金取消时间"`
	DepositRefundTime        *gtime.Time `json:"depositRefundTime"        orm:"deposit_refund_time"        description:"定金退款时间"`
	DepositExpirationTime    int         `json:"depositExpirationTime"    orm:"deposit_expiration_time"    description:"定金支付过期时间"`
	RemainPayStatus          string      `json:"remainPayStatus"          orm:"remain_pay_status"          description:"尾款付款状态"`
	RemainCancelTime         *gtime.Time `json:"remainCancelTime"         orm:"remain_cancel_time"         description:"尾款取消时间"`
	RemainRefundTime         *gtime.Time `json:"remainRefundTime"         orm:"remain_refund_time"         description:"尾款退款时间"`
	DepositCancelSource      string      `json:"depositCancelSource"      orm:"deposit_cancel_source"      description:"定金取消来源"`
	DepositCancelReason      string      `json:"depositCancelReason"      orm:"deposit_cancel_reason"      description:"定金取消原因"`
	RemainCancelSource       string      `json:"remainCancelSource"       orm:"remain_cancel_source"       description:"尾款取消来源"`
	RemainCancelReason       string      `json:"remainCancelReason"       orm:"remain_cancel_reason"       description:"尾款取消原因"`
	PayStep                  string      `json:"payStep"                  orm:"pay_step"                   description:""`
	OldBookDate              string      `json:"oldBookDate"              orm:"old_book_date"              description:"原预定日期"`
	OldBookTime              string      `json:"oldBookTime"              orm:"old_book_time"              description:"原预定时间"`
	OldBookDatetime          *gtime.Time `json:"oldBookDatetime"          orm:"old_book_datetime"          description:"原预定日期时间"`
	OldBookingCount          int         `json:"oldBookingCount"          orm:"old_booking_count"          description:"原预定人数"`
	ToretaReservationNo      int         `json:"toretaReservationNo"      orm:"toreta_reservation_no"      description:"Toreta预约号"`
	ToretaReservationId      string      `json:"toretaReservationId"      orm:"toreta_reservation_id"      description:"Toreta预约ID"`
	ToretaReservationStatus  int         `json:"toretaReservationStatus"  orm:"toreta_reservation_status"  description:"Toreta预约状态（0：未来店，1：到店，2：预约取消，3：未到，4：部分到店，5：网页取消，6：已用餐，7：已完成结账，8：已重置）"`
	ToretaReservationEndtime *gtime.Time `json:"toretaReservationEndtime" orm:"toreta_reservation_endtime" description:"Toreta预约结束时间"`
	GoodsIsNoPay             int         `json:"goodsIsNoPay"             orm:"goods_is_no_pay"            description:"套餐是否无需支付：0-否，1-是"`
	ToretaHasTimeLimit       int         `json:"toretaHasTimeLimit"       orm:"toreta_has_time_limit"      description:"Toreta是否有时间限制  1-true  2-false"`
	ToretaEndTime            int         `json:"toretaEndTime"            orm:"toreta_end_time"            description:"Toreta用餐结束时间"`
	CreatedAt                *gtime.Time `json:"createdAt"                orm:"created_at"                 description:"创建时间"`
	UpdatedAt                *gtime.Time `json:"updatedAt"                orm:"updated_at"                 description:"更新时间"`
	IsFx                     string      `json:"isFx"                     orm:"is_fx"                      description:"是否是分销订单"`
}
