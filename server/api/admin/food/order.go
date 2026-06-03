package food

import (
	"APT/internal/model/input/input_food"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/foodOrder/list" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅套餐预订单列表"`
	input_food.FoodOrderListInp
}

type OrderListRes struct {
	input_form.PageRes
	List []*input_food.FoodOrderListModel `json:"list"   dc:"数据列表"`
}

type OrderViewReq struct {
	g.Meta `path:"/foodOrder/view" method:"get" tags:"ADMIN_FOOD" summary:"获取餐厅套餐预订单指定信息"`
	input_food.FoodOrderViewInp
}

type OrderViewRes struct {
	*input_food.FoodOrderViewModel
}

type OrderConfirmAgreeReq struct {
	g.Meta `path:"/foodOrder/confirmAgree" method:"post" tags:"ADMIN_FOOD" summary:"订单_确认同意"`
	input_food.FoodOrderConfirmAgreeInp
}

type OrderConfirmAgreeRes struct {
}

type OrderConfirmDisagreeReq struct {
	g.Meta `path:"/foodOrder/confirmDisagree" method:"post" tags:"ADMIN_FOOD" summary:"订单_确认拒绝"`
	input_food.FoodOrderConfirmDisagreeInp
}

type OrderConfirmDisagreeRes struct {
}

type SettleOrderListReq struct {
	g.Meta `path:"/foodOrder/settleOrderList" method:"get" tags:"ADMIN_FOOD" summary:"获取结算订单列表"`
	input_food.SettleFoodOrderListInp
}

type SettleOrderListRes struct {
	input_form.PageRes
	List []*input_food.SettleFoodOrderListModel `json:"list"   dc:"数据列表"`
}

type OrderCancelPayReq struct {
	g.Meta `path:"/foodOrder/cancelPay" method:"post" tags:"ADMIN_FOOD" summary:"订单_取消"`
	input_food.FoodOrderCancelPayInp
}

type OrderCancelPayRes struct {
}

type OrderExportReq struct {
	g.Meta `path:"/foodOrder/export" method:"post" tags:"ADMIN_FOOD" summary:"订单-导出"`
	input_food.FoodOrderExportInp
}

type OrderExportRes struct{}

type OrderExportListReq struct {
	g.Meta `path:"/foodOrder/exportList" method:"get" tags:"ADMIN_FOOD" summary:"获取导出记录列表"`
	input_food.FoodOrderExportListInp
}

type OrderExportListRes struct {
	input_form.PageRes
	List []*input_food.FoodOrderExportListModel `json:"list"   dc:"数据列表"`
}
