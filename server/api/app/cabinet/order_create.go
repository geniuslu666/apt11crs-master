package cabinet

import (
	"APT/internal/model/input/input_cabinet"

	"github.com/gogf/gf/v2/frame/g"
)

type PreOrderCreateReq struct {
	g.Meta `path:"/cabinet/createPreOrder" method:"post" tags:"APP_CABINET" summary:"储物柜订单_预创建订单(扫码拿到ID后就请求)"`
	input_cabinet.PreCreateOrderInp
}

type PreOrderCreateRes struct {
	*input_cabinet.PreCreateOrderModel
}

type PreOrderCreateDetailReq struct {
	g.Meta `path:"/cabinet/preOrderDetail" method:"post" tags:"APP_CABINET" summary:"储物柜订单_预订单详情(进入选格口页面的时候先请求)"`
	input_cabinet.PreOrderDetailInp
}

type PreOrderCreateDetailRes struct {
	*input_cabinet.PreOrderDetailModel
}

type PrePayInfoReq struct {
	g.Meta `path:"/cabinet/payInfo" method:"post" tags:"APP_CABINET" summary:"储物柜订单_支付信息(选择格口、选择优惠券、选择积分的时候请求)"`
	*input_cabinet.PrePayInfoInp
}

type PrePayInfoRes struct {
	*input_cabinet.OrderPayInfoModel
}

type CreateOrderReq struct {
	g.Meta `path:"/cabinet/createOrder" method:"post" tags:"APP_CABINET" summary:"储物柜订单_创建订单"`
	*input_cabinet.CreateOrderInp
}

type CreateOrderRes struct {
	*input_cabinet.CreateOrderModel
}
