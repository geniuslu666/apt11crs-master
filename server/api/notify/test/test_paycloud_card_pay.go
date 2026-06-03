package test

import "github.com/gogf/gf/v2/frame/g"

type PaycloudCardPayReq struct {
	g.Meta `path:"/test/paycloud/cardPay" method:"get" tags:"NOTIFY_KEFU" summary:"信用卡测试支付"`
	Amount float64 `json:"amount" v:"required#please_enter_the_amount" dc:"金额"`
}

type PaycloudCardPayRes struct {
}
