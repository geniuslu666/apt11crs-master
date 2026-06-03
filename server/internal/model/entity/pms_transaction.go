// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PmsTransaction is the golang structure for table pms_transaction.
type PmsTransaction struct {
	Id               int         `json:"id"               orm:"id"                 description:"主键"`
	OrderSn          string      `json:"orderSn"          orm:"order_sn"           description:"订单号"`
	ChangeOrderSn    string      `json:"changeOrderSn"    orm:"change_order_sn"    description:"变更订单号"`
	OrderType        string      `json:"orderType"        orm:"order_type"         description:""`
	TransactionSn    string      `json:"transactionSn"    orm:"transaction_sn"     description:"支付流水号"`
	PaymentRequestId string      `json:"paymentRequestId" orm:"payment_request_id" description:"第三方支付流水号"`
	CaptureId        string      `json:"captureId"        orm:"capture_id"         description:"paypal capture_id"`
	Scene            string      `json:"scene"            orm:"scene"              description:"场景值"`
	PayChannel       string      `json:"payChannel"       orm:"pay_channel"        description:"SYSTEM 系统积分  PAYCLOUD   paycloud第三方支付平台"`
	PayType          string      `json:"payType"          orm:"pay_type"           description:"支付方式   BAL 余额"`
	OpenId           string      `json:"openId"           orm:"open_id"            description:"用户ID"`
	Amount           float64     `json:"amount"           orm:"amount"             description:"总金额"`
	PayParams        string      `json:"payParams"        orm:"pay_params"         description:"支付参数"`
	PriceCurrency    string      `json:"priceCurrency"    orm:"price_currency"     description:"币种"`
	PayAmount        float64     `json:"payAmount"        orm:"pay_amount"         description:"支付金额"`
	PayCharge        float64     `json:"payCharge"        orm:"pay_charge"         description:"支付税率"`
	PayStatus        string      `json:"payStatus"        orm:"pay_status"         description:"支付状态  WAIT 等待支付、DONE 完成支付、CANCEL 取消支付"`
	PayTime          *gtime.Time `json:"payTime"          orm:"pay_time"           description:"支付时间"`
	ExpiredTime      *gtime.Time `json:"expiredTime"      orm:"expired_time"       description:"过期时间"`
	RefundAmount     float64     `json:"refundAmount"     orm:"refund_amount"      description:"退款金额"`
	RefundStatus     string      `json:"refundStatus"     orm:"refund_status"      description:"退款状态   WAIT 未退款   PART  部分退款 DONE 全部退款"`
	ScenePayRate     float64     `json:"scenePayRate"     orm:"scene_pay_rate"     description:"场景积分抵扣比例"`
	Level            int         `json:"level"            orm:"level"              description:"等级"`
	ExchangeRate     float64     `json:"exchangeRate"     orm:"exchange_rate"      description:"积分汇率"`
	CouponId         int         `json:"couponId"         orm:"coupon_id"          description:"优惠券ID"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:"删除时间"`
	Remark           string      `json:"remark"           orm:"remark"             description:"备注"`
	IsFx             string      `json:"isFx"             orm:"is_fx"              description:"是否是分销订单"`
}
