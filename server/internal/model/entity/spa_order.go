// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SpaOrder is the golang structure for table spa_order.
type SpaOrder struct {
	Id                      int64       `json:"id"                      orm:"id"                         description:""`
	IspId                   int         `json:"ispId"                   orm:"isp_id"                     description:"服务商ID"`
	OrderSn                 string      `json:"orderSn"                 orm:"order_sn"                   description:"订单编号"`
	OutOrderSn              string      `json:"outOrderSn"              orm:"out_order_sn"               description:"三方订单号"`
	ServiceType             int         `json:"serviceType"             orm:"service_type"               description:"1到店  2上门"`
	ConfirmType             int         `json:"confirmType"             orm:"confirm_type"               description:"1手动确认  2自动确认"`
	MemberId                uint        `json:"memberId"                orm:"member_id"                  description:"用户ID"`
	BookDate                string      `json:"bookDate"                orm:"book_date"                  description:"预定日期"`
	BookTime                string      `json:"bookTime"                orm:"book_time"                  description:"预定时间"`
	BookStartTime           *gtime.Time `json:"bookStartTime"           orm:"book_start_time"            description:"预定开始日期时间"`
	BookEndTime             *gtime.Time `json:"bookEndTime"             orm:"book_end_time"              description:"预定结束日期时间"`
	TechnicianGoTime        *gtime.Time `json:"technicianGoTime"        orm:"technician_go_time"         description:"技师上门出发时间"`
	MemberArriveTime        *gtime.Time `json:"memberArriveTime"        orm:"member_arrive_time"         description:"到店时间"`
	ActualStartTime         *gtime.Time `json:"actualStartTime"         orm:"actual_start_time"          description:"实际开始日期时间"`
	ActualEndTime           *gtime.Time `json:"actualEndTime"           orm:"actual_end_time"            description:"实际结束日期时间"`
	BookingName             string      `json:"bookingName"             orm:"booking_name"               description:"预定人姓名"`
	PhoneArea               string      `json:"phoneArea"               orm:"phone_area"                 description:"手机区号"`
	BookingMobile           string      `json:"bookingMobile"           orm:"booking_mobile"             description:"预定人手机"`
	BookingEmail            string      `json:"bookingEmail"            orm:"booking_email"              description:"预定人邮箱"`
	ServiceId               int         `json:"serviceId"               orm:"service_id"                 description:"服务ID"`
	OrderAmount             float64     `json:"orderAmount"             orm:"order_amount"               description:"订单金额"`
	CouponAmount            float64     `json:"couponAmount"            orm:"coupon_amount"              description:"优惠券抵扣金额"`
	BalAmount               float64     `json:"balAmount"               orm:"bal_amount"                 description:"积分抵扣金额"`
	GoodsId                 uint        `json:"goodsId"                 orm:"goods_id"                   description:"项目ID"`
	GoodsNum                int         `json:"goodsNum"                orm:"goods_num"                  description:"项目数量"`
	TechnicianIds           string      `json:"technicianIds"           orm:"technician_ids"             description:"技师ID ,分隔"`
	IsReturn                int         `json:"isReturn"                orm:"is_return"                  description:"是否退单中 1是  2否"`
	PayModel                int         `json:"payModel"                orm:"pay_model"                  description:"1、余额支付 2、组合支付 3、纯外部支付"`
	PayTime                 *gtime.Time `json:"payTime"                 orm:"pay_time"                   description:"支付时间"`
	PayStatus               string      `json:"payStatus"               orm:"pay_status"                 description:"订单付款状态"`
	StoreId                 uint        `json:"storeId"                 orm:"store_id"                   description:"到店门店ID"`
	PropertyId              uint        `json:"propertyId"              orm:"property_id"                description:"上门物业ID"`
	RoomNo                  string      `json:"roomNo"                  orm:"room_no"                    description:"上门房间号"`
	OrderStatus             string      `json:"orderStatus"             orm:"order_status"               description:"订单状态"`
	ConfirmTime             *gtime.Time `json:"confirmTime"             orm:"confirm_time"               description:"订单确认时间"`
	ConfirmRefuseReason     string      `json:"confirmRefuseReason"     orm:"confirm_refuse_reason"      description:"审核拒绝原因"`
	DispatchStatus          string      `json:"dispatchStatus"          orm:"dispatch_status"            description:"订单调度状态"`
	DispatchTime            *gtime.Time `json:"dispatchTime"            orm:"dispatch_time"              description:"订单调度时间"`
	DispatchDesc            string      `json:"dispatchDesc"            orm:"dispatch_desc"              description:"订单调度备注"`
	DispatchOperatorId      uint        `json:"dispatchOperatorId"      orm:"dispatch_operator_id"       description:"调度操作人ID"`
	ExpirationTime          int         `json:"expirationTime"          orm:"expiration_time"            description:"订单过期时间"`
	CancelTime              *gtime.Time `json:"cancelTime"              orm:"cancel_time"                description:"取消时间"`
	StartServeImages        string      `json:"startServeImages"        orm:"start_serve_images"         description:"开始服务图集"`
	EndServeImages          string      `json:"endServeImages"          orm:"end_serve_images"           description:"服务结束图集"`
	SettlementRate          float64     `json:"settlementRate"          orm:"settlement_rate"            description:"结算比例"`
	SettlementStatus        string      `json:"settlementStatus"        orm:"settlement_status"          description:"WAIT 等待结算   SUCCESS  结算处理成功    FAIL  结算处理失败"`
	SettlementAmount        float64     `json:"settlementAmount"        orm:"settlement_amount"          description:"结算金额"`
	SettlementTime          *gtime.Time `json:"settlementTime"          orm:"settlement_time"            description:"结算时间"`
	SettlementOrderId       uint        `json:"settlementOrderId"       orm:"settlement_order_id"        description:"结算单ID"`
	MemberMessage           string      `json:"memberMessage"           orm:"member_message"             description:"购买人留言信息"`
	MemberMessageJa         string      `json:"memberMessageJa"         orm:"member_message_ja"          description:"购买人留言信息日语"`
	RefundStatus            string      `json:"refundStatus"            orm:"refund_status"              description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	RefundFee               float64     `json:"refundFee"               orm:"refund_fee"                 description:"退款手续费"`
	RefundTime              *gtime.Time `json:"refundTime"              orm:"refund_time"                description:"退款时间"`
	RefundAmount            float64     `json:"refundAmount"            orm:"refund_amount"              description:"已退款总金额"`
	RefundBalAmount         float64     `json:"refundBalAmount"         orm:"refund_bal_amount"          description:"已退款积分"`
	RefundCouponAmount      float64     `json:"refundCouponAmount"      orm:"refund_coupon_amount"       description:"已退款优惠券"`
	AdminRefundAmount       float64     `json:"adminRefundAmount"       orm:"admin_refund_amount"        description:"后台已退款总金额"`
	AdminRefundBalAmount    float64     `json:"adminRefundBalAmount"    orm:"admin_refund_bal_amount"    description:"后台已退款积分"`
	AdminRefundCouponAmount float64     `json:"adminRefundCouponAmount" orm:"admin_refund_coupon_amount" description:"后台已退款优惠券"`
	RefundReason            string      `json:"refundReason"            orm:"refund_reason"              description:"退款原因"`
	Referrer                int         `json:"referrer"                orm:"referrer"                   description:"推荐人"`
	RebateRate              float64     `json:"rebateRate"              orm:"rebate_rate"                description:"分佣比例"`
	RebateStatus            string      `json:"rebateStatus"            orm:"rebate_status"              description:"WAIT 等待处理佣金   SUCCESS   佣金处理成功    FAIL  佣金处理失败"`
	RebateAmount            float64     `json:"rebateAmount"            orm:"rebate_amount"              description:"分佣结算金额"`
	RebateTime              *gtime.Time `json:"rebateTime"              orm:"rebate_time"                description:"分佣结算时间"`
	IsGetOpen               string      `json:"isGetOpen"               orm:"is_get_open"                description:"是否开启积分获取"`
	IsPayOpen               string      `json:"isPayOpen"               orm:"is_pay_open"                description:"是否开启积分抵扣"`
	SpaGetRateVip           float64     `json:"spaGetRateVip"           orm:"spa_get_rate_vip"           description:"按摩结算积分比例"`
	SpaGetRateScene         float64     `json:"spaGetRateScene"         orm:"spa_get_rate_scene"         description:"场景结算积分比例"`
	SpaGetScoreStatus       string      `json:"spaGetScoreStatus"       orm:"spa_get_score_status"       description:"'WAIT','SUCCESS','FAIL'"`
	SpaGetAmount            float64     `json:"spaGetAmount"            orm:"spa_get_amount"             description:"结算积分金额"`
	ExpValue                float64     `json:"expValue"                orm:"exp_value"                  description:"结算的经验值"`
	ExpTime                 *gtime.Time `json:"expTime"                 orm:"exp_time"                   description:"经验结算时间"`
	AdminCancelReason       string      `json:"adminCancelReason"       orm:"admin_cancel_reason"        description:"后台取消原因"`
	AdminCancelNum          int         `json:"adminCancelNum"          orm:"admin_cancel_num"           description:"后台取消次数"`
	AbnormalStatus          int         `json:"abnormalStatus"          orm:"abnormal_status"            description:"异常单状态（1：不是异常单  2异常待处理  3异常已处理）"`
	AbnormalReason          string      `json:"abnormalReason"          orm:"abnormal_reason"            description:"异常处理原因"`
	AbnormalOperatorId      uint        `json:"abnormalOperatorId"      orm:"abnormal_operator_id"       description:"异常处理操作人ID"`
	AbnormalTime            *gtime.Time `json:"abnormalTime"            orm:"abnormal_time"              description:"异常处理时间"`
	CreatedAt               *gtime.Time `json:"createdAt"               orm:"created_at"                 description:"创建时间"`
	UpdatedAt               *gtime.Time `json:"updatedAt"               orm:"updated_at"                 description:"更新时间"`
	IsFx                    string      `json:"isFx"                    orm:"is_fx"                      description:"是否是分销订单"`
}
