package input_refund

type PmsOrderRefundInp struct {
	OrderSn string `json:"ordersn" v:"required-without:changeOrdersn#订单号未知" dc:"订单号"`
	//CancelOrderSn string  `json:"cancelOrderSn" dc:"订单预取消退款订单号"`
	RefundAmount  float64 `json:"refundAmount" v:"required|regex:^[0-9]([0-9]+)?(.[0-9]{1,2})?$#退款金额未知|退款金额格式错误" dc:"退款金额"`
	IsCancelOrder string  `json:"isCancelOrder" d:"N"  dc:"是否取消订单并退款 Y 是  N 否（默认）"`
	Reason        string  `json:"reason"   dc:"退款并取消原因"`
}

type RefundAmountInp struct {
	OrderSn      string  `json:"ordersn" v:"required#订单号未知" dc:"订单号"`
	RefundAmount float64 `json:"refundAmount" v:"required#退款金额未知" dc:"退款金额"`
	OperateType  string  `json:"operateType" description:"操作员类型"`
	OperateId    int     `json:"operateId"   description:"操作员ID"`
	Remark       string  `json:"remark" dc:"备注"`
}

type RefundAmountList struct {
	MemberId         int     `json:"memberId" v:"required#用户ID未知" dc:"用户ID"`
	OrderSn          string  `json:"orderSn"          dc:"订单号"`
	OrderType        string  `json:"orderType"        dc:"订单类型"`
	TransactionSn    string  `json:"transactionSn"    dc:"支付流水号"`
	PaymentRequestId string  `json:"paymentRequestId" dc:"第三方支付流水号"`
	CaptureId        string  `json:"captureId"        dc:"paypal capture_id"`
	Scene            string  `json:"scene"            dc:"场景值"`
	PayChannel       string  `json:"payChannel"       dc:"支付渠道"`
	PayType          string  `json:"payType"          dc:"支付方式"`
	Amount           float64 `json:"amount"           dc:"总金额"`
	PayParams        string  `json:"payParams"        dc:"支付参数"`
	PriceCurrency    string  `json:"priceCurrency"    dc:"币种"`
	PayAmount        float64 `json:"payAmount"        dc:"支付金额"`
	PayCharge        float64 `json:"payCharge"        dc:"支付税率"`
	RefundAmount     float64 `json:"refundAmount"     dc:"已退款金额"`
	BalanceAmount    float64 `json:"balanceAmount"    dc:"可退金额"`
	WantRefundAmount float64 `json:"wantRefundAmount" dc:"要退金额"`
}

type RefundInp struct {
	MemberId         int     `json:"memberId" v:"required#用户ID未知" dc:"用户ID"`
	OrderSn          string  `json:"orderSn" v:"required#订单号未知" dc:"订单号"`
	CancelOrderSn    string  `json:"cancelOrderSn" v:"required#订单预取消退款订单号未知" dc:"订单预取消退款订单号"`
	TransactionSn    string  `json:"transactionSn" v:"required#支付流水号未知" dc:"支付流水号"`
	PayChannel       string  `json:"payChannel" v:"required|in:BAL,Alipay+,WeChatPay,Paypal,PaypalCard,StripeCard#支付渠道未知|支付渠道不支持" dc:"支付渠道"`
	RefundType       string  `json:"refundType" v:"required|in:BAL,Alipay+,WeChatPay,Paypal,PaypalCard,StripeCard#退款方式未知|退款方式不支持" dc:"退款方式"`
	RefundAmount     float64 `json:"refundAmount" dc:"退款金额"`
	Scene            string  `json:"scene" v:"required#场景未知" dc:"场景"`
	PriceCurrency    string  `json:"priceCurrency" v:"required#币种未知" dc:"币种"`
	IsAllRefund      bool    `json:"isAllRefund" dc:"是否全部退款"`
	CaptureId        string  `json:"captureId" v:"required#paypal capture_id未知" dc:"paypal capture_id"`
	PaymentRequestId string  `json:"paymentRequestId" v:"required#第三方支付流水号未知" dc:"第三方支付流水号"`
	OperateType      string  `json:"operateType" description:"操作员类型"`
	OperateId        int     `json:"operateId"   description:"操作员ID"`
	Remark           string  `json:"remark" dc:"备注"`
}
