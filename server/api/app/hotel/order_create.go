package hotel

import (
	"APT/internal/model/input/input_hotel"
	"APT/internal/model/input/input_pay"
	"github.com/gogf/gf/v2/frame/g"
)

type PreOrderCreateReq struct {
	g.Meta `path:"/home/createPreOrder" method:"post" tags:"APP_HOTEL" summary:"酒店订单_预创建订单"`
	input_hotel.PreCreateOrderInp
}

type PreOrderCreateRes struct {
	*input_hotel.PreCreateOrderModel
}

type PreOrderCreateDetailReq struct {
	g.Meta `path:"/home/preOrderDetail" method:"post" tags:"APP_HOTEL" summary:"酒店订单_预订单详情"`
	input_hotel.PreOrderDetailInp
}

type PreOrderCreateDetailRes struct {
	*input_hotel.PreOrderDetailModel
}

type PrePayInfoReq struct {
	g.Meta `path:"/home/payInfo" method:"post" tags:"APP_HOTEL" summary:"酒店订单_支付信息"`
	*input_pay.PrePayInfoInp
}

type PrePayInfoRes struct {
	*input_pay.PrePayInfoModel
}

type CreateOrderReq struct {
	g.Meta `path:"/home/createOrder" method:"post" tags:"APP_HOTEL" summary:"酒店订单_创建订单"`
	*input_hotel.CreateOrderInp
}

type CreateOrderRes struct {
	*input_hotel.CreateOrderModel
}
