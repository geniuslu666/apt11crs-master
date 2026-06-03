package pay

import (
	"APT/internal/model/input/input_pay"
	"github.com/gogf/gf/v2/frame/g"
)

type ThirdPayReq struct {
	g.Meta `path:"/home/selectPayMode" method:"post" tags:"APP_PAY" summary:"支付_选择支付方式"`
	*input_pay.SelectThirdPayInp
}

type ThirdPayRes struct {
	*input_pay.SelectThirdPayModel
}

type InnerPayReq struct {
	g.Meta `path:"/home/payOrder" method:"post" tags:"APP_PAY" summary:"支付_支付订单"`
	*input_pay.BalancePayInp
}

type InnerPayRes struct {
	*input_pay.BalancePayModel
}

type QueryStatusReq struct {
	g.Meta  `path:"/home/checkOrderSuccess" method:"post" tags:"APP_PAY" summary:"支付_询问支付状态"`
	OrderSn string `json:"orderSn" v:"required#order_number_unknown" dc:"订单号"`
}

type QueryStatusRes struct {
	Status string `json:"status" dc:"订单状态【SUCCESS、已支付 FAIL、已取消 HC_DATE_FAIL、续住失败】"`
}
