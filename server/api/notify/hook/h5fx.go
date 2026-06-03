package hook

import "github.com/gogf/gf/v2/frame/g"

type PaySuccessNotifyReq struct {
	g.Meta      `path:"/h5fx/paySuccessNotify" method:"post" tags:"H5FX" summary:"支付成功回调"`
	OrderNo     string `json:"order_no" dc:"分销中心业务订单号，跳转至收银台时会携带"`
	TradeNo     string `json:"trade_no" dc:"分销中心业务交易号，跳转至收银台时会携带"`
	TradeAmount string `json:"trade_amount" dc:"订单支付金额"`
	PayTime     string `json:"pay_time" dc:"支付完成时间，格式yyyy-MM-dd HH:mm:ss"`
}
type PaySuccessNotifyRes struct {
}

type RefundSuccessNotifyReq struct {
	g.Meta             `path:"/h5fx/refundSuccessNotify" method:"post" tags:"H5FX" summary:"退款成功回调"`
	OrderNo            string `json:"orderNo" dc:"分销中心业务订单号，跳转至收银台时会携带"`
	TradeNo            string `json:"tradeNo" dc:"分销中心业务交易号，跳转至收银台时会携带"`
	RefundOrderNo      string `json:"refundOrderNo" dc:"分销中心退款订单号，标识一次退款请求"`
	ThirdRefundOrderNo string `json:"thirdRefundOrderNo" dc:"接入方受理成功后，向平台返回接入方的退款单号"`
	RefundAmount       string `json:"refundAmount" dc:"订单退款金额，订单退款通知时携带"`
	RefundSuccessTime  string `json:"refundSuccessTime" dc:"退款成功时间  格式yyyy-MM-dd HH:mm:ss"`
}
type RefundSuccessNotifyRes struct {
}
