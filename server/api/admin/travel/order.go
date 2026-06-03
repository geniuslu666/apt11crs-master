package travel

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_travel"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderListReq 获取一日游订单列表
type OrderListReq struct {
	g.Meta `path:"/travel/order/list" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游订单列表"`
	input_travel.TravelOrderListInp
}

type OrderListRes struct {
	input_form.PageRes
	List []*input_travel.TravelOrderListModel `json:"list" dc:"数据列表"`
}

// OrderViewReq 获取一日游订单详情
type OrderViewReq struct {
	g.Meta `path:"/travel/order/view" method:"get" tags:"ADMIN_TRAVEL" summary:"获取一日游订单详情"`
	input_travel.TravelOrderViewInp
}

type OrderViewRes struct {
	*input_travel.TravelOrderViewModel
}

// OrderRefundReq 一日游订单退款
type OrderRefundReq struct {
	g.Meta `path:"/travel/order/refund" method:"post" tags:"ADMIN_TRAVEL" summary:"一日游订单退款"`
	input_travel.TravelOrderRefundInp
}

type OrderRefundRes struct{}
