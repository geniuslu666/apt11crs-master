package travel

import (
	"APT/internal/model/input/input_pay"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

type PreOrderCreateReq struct {
	g.Meta `path:"/travel/createPreOrder" method:"post" tags:"APP_TRAVEL" summary:"[一日游]预创建订单"`
	input_travel.PreCreateOrderInp
}

type PreOrderCreateRes struct {
	*input_travel.PreCreateOrderModel
}

type PreOrderCreateDetailReq struct {
	g.Meta `path:"/travel/preOrderDetail" method:"post" tags:"APP_TRAVEL" summary:"[一日游]_预订单详情"`
	input_travel.PreOrderDetailInp
}

type PreOrderCreateDetailRes struct {
	*input_travel.PreOrderDetailModel
}

type PrePayInfoReq struct {
	g.Meta `path:"/travel/payInfo" method:"post" tags:"APP_TRAVEL" summary:"[一日游]_支付信息(选择优惠券、选择积分的时候请求)"`
	*input_pay.PrePayInfoInp
}

type PrePayInfoRes struct {
	*input_pay.PrePayInfoModel
}

type CreateOrderReq struct {
	g.Meta `path:"/travel/createOrder" method:"post" tags:"APP_TRAVEL" summary:"[一日游]_创建订单"`
	*input_travel.CreateOrderInp
}

type CreateOrderRes struct {
	*input_travel.CreateOrderModel
}
