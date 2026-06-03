package hotel

import (
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type PreOrderRefundDetailReq struct {
	g.Meta `path:"/home/cancelOrderDetail" method:"post" tags:"APP_HOTEL"  summary:"酒店订单_预退款详情"`
	input_hotel.PreRefundIn
}

type PreOrderRefundDetailRes struct {
	*input_hotel.PreRefundOut
}

type RefundOrderReq struct {
	g.Meta           `path:"/home/cancelOrder" method:"post" tags:"APP_HOTEL"  summary:"酒店订单_订单退款"`
	PreCancelOrderSn string `json:"preCancelOrderSn" v:"required#order_pre_refund_order_number_miss" dc:"订单预退款订单号"`
	OrderSn          string `json:"orderSn" v:"required#order_number_miss" dc:"订单号"`
}

type RefundOrderRes struct {
	CancelOrderStatus string `json:"cancelOrderStatus" dc:"订单退款状态"`
	CancelOrderInfo   []*struct {
	} `json:"cancelOrderInfo" dc:"订单退款信息"`
}

type OrderRefundDetailReq struct {
	g.Meta `path:"/home/refundOrderDetail" method:"post" tags:"APP_HOTEL"  summary:"酒店订单_订单退款详情"`
	*input_hotel.RefundDetailInp
}
type OrderRefundDetailRes struct {
	*input_hotel.RefundDetailModel
}
