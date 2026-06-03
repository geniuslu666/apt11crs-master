package mlilifeWxPay

// NotifyRes 支付回调响应结果
type NotifyRes struct {
	// 交易类型（must）
	TradeType string `json:"trade_type"` // JSAPI|NATIVE|APP|MICROPAY

	// 银行类型（must）
	BankType string `json:"bank_type"` // 如OTHERS表示非标准银行类型

	// 公众账号ID（must）
	AppId string `json:"appid"`

	// 用户唯一标识（must）
	OpenID string `json:"openid"`

	// 订单总金额（must）
	TotalFee int `json:"total_fee"` // 单位：该币种最小计算单位

	// 币种类型（must）
	FeeType string `json:"fee_type"` // 如JPY表示日元

	// 微信订单号（must）
	TransactionID string `json:"transaction_id"`

	// 支付完成时间（must）
	TimeEnd string `json:"time_end"` // 格式: YYYYMMDDHHMMSS

	// 交易状态（must）
	TradeState string `json:"trade_state"` // 如SUCCESS、NOTPAY等

	// 子商户公众账号ID（must）
	SubAppId string `json:"sub_appid"`

	// 用户子标识（must）
	SubOpenID string `json:"sub_openid"`

	// 是否关注子公众号（must）
	SubIsSubscribe string `json:"sub_is_subscribe"`

	// 实际支付金额（must）
	CashFee string `json:"cash_fee"` // 可能包含小数，用字符串存储

	// 现金货币类型（must）
	CashFeeType string `json:"cash_fee_type"`

	// 支付状态描述（must）
	TradeStateDesc string `json:"trade_state_desc"`

	// 支付费率（must）
	Rate string `json:"rate"` // 建议按字符串存储避免精度损失

	// 系统PID（must）
	Pid string `json:"pid"`

	// 商户订单号（must）
	OutTradeNo string `json:"out_trade_no"`

	// 随机加密串（must）
	NonceStr string `json:"nonce_str"`

	// 签名（must）
	Sign string `json:"sign"`

	// 附加数据（must）
	Attach string `json:"attach"` // 商户自定义数据
}
