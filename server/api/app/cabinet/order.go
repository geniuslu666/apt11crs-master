package cabinet

import (
	"APT/internal/model/input/input_cabinet"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/cabinet/orderList" method:"post" tags:"APP_CABINET" summary:"储物柜订单_订单列表"`
	input_cabinet.OrderAppListInp
}

type OrderListRes struct {
	List  []*input_cabinet.OrderAppListModel `json:"list"   dc:"数据列表"`
	Count int                                `json:"count"   dc:"数据总数"`
}

type OrderViewReq struct {
	g.Meta `path:"/cabinet/orderDetail" method:"post" tags:"APP_CABINET" summary:"储物柜订单_订单详情"`
	input_cabinet.OrderAppViewInp
}

type OrderViewRes struct {
	*input_cabinet.OrderAppViewModel
}

type PayOvertimeInfoReq struct {
	g.Meta `path:"/cabinet/payOvertimeInfo" method:"post" tags:"APP_CABINET" summary:"储物柜订单_支付超时费页面信息"`
	input_cabinet.PayOvertimeInfoInp
}

type PayOvertimeInfoRes struct {
	*input_cabinet.PayOvertimeInfoModel
}

type PayOvertimeReq struct {
	g.Meta `path:"/cabinet/payOvertime" method:"post" tags:"APP_CABINET" summary:"储物柜订单_支付超时费"`
	*input_cabinet.PayOvertimeInp
}

type PayOvertimeRes struct {
	*input_cabinet.PayOvertimeModel
}

type CancelOrderNoPayReq struct {
	g.Meta  `path:"/cabinet/cancelOrder" method:"post" tags:"APP_CABINET" summary:"储物柜订单_取消未支付订单"`
	OrderSn string `json:"orderSn" dc:"订单号"`
}

type CancelOrderNoPayRes struct{}
