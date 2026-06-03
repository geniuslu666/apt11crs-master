package hotel

import (
	"APT/internal/model/input/input_hotel"
	"github.com/gogf/gf/v2/frame/g"
)

type PreOrderCreateReq struct {
	g.Meta `path:"/v2/home/createPreOrder" method:"post" tags:"APP_HOTEL_V2" summary:"酒店订单_预创建订单"`
	input_hotel.PreCreateOrderInp
}

type PreOrderCreateRes struct {
	*input_hotel.PreCreateOrderModel
}

type PreMoreOrderCreateReq struct {
	g.Meta `path:"/v2/home/more/createPreOrder" method:"post" tags:"APP_HOTEL_V2" summary:"酒店订单_预创建订单_多间房"`
	input_hotel.PreMoreCreateOrderInp
}

type PreMoreOrderCreateRes struct {
	*input_hotel.PreCreateOrderModel
}

type PreOrderCreateDetailReq struct {
	g.Meta `path:"/v2/home/preOrderDetail" method:"post" tags:"APP_HOTEL_V2" summary:"酒店订单_预订单详情"`
	input_hotel.PreOrderDetailInp
}

type PreOrderCreateDetailRes struct {
	*input_hotel.PreOrderDetailModel
}
