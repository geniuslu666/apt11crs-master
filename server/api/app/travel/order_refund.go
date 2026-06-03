package travel

import (
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

type ApplyOrderRefundDetailReq struct {
	g.Meta `path:"/travel/applyRefundDetail" method:"post" tags:"APP_TRAVEL"  summary:"[一日游]_预退款详情"`
	input_travel.TravelOrderApplyRefundDetailInp
}

type ApplyOrderRefundDetailRes struct {
	*input_travel.TravelOrderApplyRefundDetailModel
}

type RefundOrderReq struct {
	g.Meta `path:"/travel/refundOrder" method:"post" tags:"APP_TRAVEL"  summary:"[一日游]_订单退款"`
	input_travel.TravelOrderApplyRefundDetailInp
}

type RefundOrderRes struct{}

type OrderRefundDetailReq struct {
	g.Meta `path:"/travel/refundOrderDetail" method:"post" tags:"APP_TRAVEL"  summary:"[一日游]_订单退款详情"`
	input_travel.TravelOrderApplyRefundDetailInp
}
type OrderRefundDetailRes struct {
	*input_travel.RefundDetailModel
}

type RefundDetailModel struct {
}
