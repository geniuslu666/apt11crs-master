package pms

import (
	"APT/internal/model/input/input_refund"
	"github.com/gogf/gf/v2/frame/g"
)

type RefundOrderReq struct {
	g.Meta `path:"/payRefund/paycloud" method:"post" tags:"ADMIN_PMS" summary:"交易退款_支付云退款接口"`
	*input_refund.PmsOrderRefundInp
}

type RefundOrderRes struct{}
