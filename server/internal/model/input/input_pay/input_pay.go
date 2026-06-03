package input_pay

import (
	"APT/internal/model/entity"
	"APT/internal/model/input/input_hotel"
)

type PayInfoInp struct {
	OrderAmount float64 `json:"orderAmount" dc:"订单总金额"`
	MemberId    int     `json:"memberId"    dc:"会员ID"`
	IsBalance   bool    `json:"isBalance"   dc:"是否使用余额进行支付"`
	CouponId    int     `json:"couponId"    dc:"参与支付的优惠券ID"`
	Scene       string  `json:"scene"       dc:"支付场景   HOTEL  FOOD "`
}

type PayInfoModel struct {
	PayModel      int     `json:"payModel"      dc:"支付方式-1纯余额支付-2余额加外部支付-3纯外部支付"`
	MemberBalance float64 `json:"memberBalance" dc:"用户当前余额"`
	AllAmount     float64 `json:"allAmount"     dc:"订单总金额"`
	Score         float64 `json:"score"         dc:"订单可用积分上限"`
	BalanceAmount float64 `json:"balanceAmount" dc:"余额支付金额"`
	ThirdAmount   float64 `json:"thirdAmount"   dc:"第三方支付金额"`
	BalanceConfig *input_hotel.BalanceConfig
	CouponAmount  float64 `json:"couponAmount"  dc:"优惠券抵用金额"`
	CouponName    string  `json:"couponName"    dc:"优惠券名称"`
	CouponId      int     `json:"couponId"      dc:"优惠券ID"`
}

type BookingPriceInp struct {
	RoomInfos   *input_hotel.RoomItems
	CheckinAt   string `json:"checkin_at" dc:"入住日期"`
	CheckoutAt  string `json:"checkout_at" dc:"退房日期"`
	Puid        string `json:"puid" dc:"物业ID"`
	IsPricePlan bool
	PricePlanId int `json:"pricePlanId"    dc:"价格planID"`
}

type BookingPriceModel struct {
	MinTotalPrice        float64             `json:"minTotalPrice" dc:"房间最低总价格"`
	TotalPrice           float64             `json:"totalPrice" dc:"房间默认基础价格的总价"`
	PlanPrice            []*BookingPriceItem `json:"planPrice" dc:"价格plan明细"`
	PlanPriceTotalAmount float64
	OrderRooms           *input_hotel.OrderRooms
	PricePercent         int64 `json:"pricePercent" dc:"价格百分比"`
}

type BookingPriceItem struct {
	*entity.PmsPricePlan
	PlanShowNameLanguage []*input_hotel.LanguageType `json:"planShowNameLanguage"     dc:"多语言物业地址"   orm:"with:uuid=plan_show_name"`
	TotalAmount          float64                     `json:"totalAmount" dc:"房间总价"`
}

type PrePayInfoInp struct {
	PreOrderSn string `json:"preOrderSn"    dc:"预订单号"`
	IsBalance  int    `json:"isBalance"     dc:"是否使用余额支付【1启用，2禁用】"`
	CouponId   int    `json:"couponId"      dc:"使用的优惠券ID"`
}

type PrePayInfoModel struct {
	*input_hotel.OrderPayInfoModel
}

type SelectThirdPayInp struct {
	OrderSn               string `json:"OrderSn" v:"required-without:changeOrderSn#payment_order_number_unknown" dc:"支付订单号"`
	PayChannel            string `json:"payChannel" v:"required|in:WeChatMiniPay,WeChatPay,Alipay+,Paypal,PaypalCard,StripeCard,Paypay_app,Paypay_h5,H5_FX#payment_method_unknown|payment_method_error" dc:"支付方式 WeChatPay,Alipay+,Paypal,PaypalCard,StripeCard,Paypay_h5,Paypay_app,H5_FX"`
	PaySuccessCallbackUrl string `json:"paySuccessNotifyUrl" dc:"支付成功回调地址"`
	PayFailCallbackUrl    string `json:"PayFailNotifyUrl" dc:"支付失败回调地址"`
	PaypalCardPayParams
}

type SelectThirdPayModel struct {
	PaypayUrl          string `json:"payPayUrl" dc:"paypay支付地址"`
	StripePayUrl       string `json:"stripePayUrl" dc:"stripe支付地址"`
	PaypalPayUrl       string `json:"paypalPayUrl" dc:"paypal支付地址"`
	WebPayUrl          string `json:"webPayUrl"    dc:"web支付地址"`
	AppPayParams       string `json:"appPayParams" dc:"微信APP支付参数包"`
	WxMiniPayParams    string `json:"wxMiniPayParams" dc:"微信小程序支付参数包"`
	AppWechatPayParams struct {
		Appid     string `json:"appid"`
		Partnerid string `json:"partnerid"`
		Prepayid  string `json:"prepayid"`
		Package   string `json:"package"`
		Noncestr  string `json:"noncestr"`
		Timestamp string `json:"timestamp"`
		Sign      string `json:"sign"`
	} `json:"appWechatPayParams" dc:"微信支付APP支付参数包"`
}

type BalancePayInp struct {
	OrderSn           string `json:"orderSn" v:"required" dc:"订单号"`
	BalancePayOrderSn string `json:"payOrderSn" dc:"余额支付订单号"`
}
type BalancePayModel struct {
}

type HotelCreateOrderInp struct {
	OrderSn                 string  `json:"orderSn" v:"required" dc:"订单号"`
	PmsTransactionBalAmount float64 `json:"pmsTransactionBalAmount" dc:"支付金额"`
	Booker                  string  `json:"booker" v:"required" dc:"预订人"`
}

type HotelCreateOrderModel struct {
}

type CabinetCreateOrderInp struct {
	OrderSn                 string  `json:"orderSn" v:"required" dc:"订单号"`
	PmsTransactionBalAmount float64 `json:"pmsTransactionBalAmount" dc:"支付金额"`
}

type CabinetCreateOrderModel struct {
}

type TravelCreateOrderInp struct {
	OrderSn                 string  `json:"orderSn" v:"required" dc:"订单号"`
	PmsTransactionBalAmount float64 `json:"pmsTransactionBalAmount" dc:"支付金额"`
}

type TravelCreateOrderModel struct {
}

type ThirdPayInp struct {
	TransEndTime      string `json:"trans_end_time"`
	Charset           string `json:"charset"`
	StoreNo           string `json:"store_no"`
	PayScenario       string `json:"pay_scenario"`
	Sign              string `json:"sign"`
	TransFeeC         string `json:"trans_fee_c"`
	DiscountBpc       string `json:"discount_bpc"`
	VatAmount         string `json:"vat_amount"`
	OrderAmount       string `json:"order_amount"`
	AppId             string `json:"app_id"`
	SignType          string `json:"sign_type"`
	TransStatus       int    `json:"trans_status"`
	PriceCurrency     string `json:"price_currency"`
	TerminalSn        string `json:"terminal_sn"`
	TransType         int    `json:"trans_type"`
	Timestamp         string `json:"timestamp"`
	TransNo           string `json:"trans_no"`
	MerchantNo        string `json:"merchant_no"`
	Method            string `json:"method"`
	PayUserAccountId  string `json:"pay_user_account_id"`
	Format            string `json:"format"`
	PayMethodId       string `json:"pay_method_id"`
	TransAmount       string `json:"trans_amount"`
	HttpRequestId     string `json:"http_request_id"`
	Version           string `json:"version"`
	PayChannelTransNo string `json:"pay_channel_trans_no"`
	PaidAmount        string `json:"paid_amount"`
	DiscountBmopc     string `json:"discount_bmopc"`
	CaptureId         string `json:"capture_id"`
}

type ThirdPayModel struct {
}

type PaypalCardPayParams struct {
	CardName    string  `json:"card_name" dc:"持卡人姓名[paypal信用卡支付必传]"`
	CardNumber  string  `json:"card_number" dc:"卡号[paypal信用卡支付必传]"`
	CardExp     string  `json:"card_expiry_month" dc:"卡过期时间格式YYYY-MM[paypal信用卡支付必传]"`
	CardCvv     string  `json:"card_cvv" dc:"卡CVV校验码[paypal信用卡支付必传]"`
	Address     string  `json:"address" dc:"详细地址[paypal信用卡支付必传]"`
	PostalCode  string  `json:"postal_code" dc:"邮政编码[paypal信用卡支付必传]"`
	CountryCode string  `json:"country_code" dc:"国家二字码[paypal信用卡支付必传]"`
	Amount      float64 `json:"amount" dc:"无需传递此参数"`
}
