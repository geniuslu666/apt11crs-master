package spa

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_spa"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/spaOrder/list" method:"get" tags:"ADMIN_SPA" summary:"获取按摩预订单列表"`
	input_spa.SpaOrderListInp
}

type OrderListRes struct {
	input_form.PageRes
	List []*input_spa.SpaOrderListModel `json:"list"   dc:"数据列表"`
}

type OrderViewReq struct {
	g.Meta `path:"/spaOrder/view" method:"get" tags:"ADMIN_SPA" summary:"获取按摩预订单指定信息"`
	input_spa.SpaOrderViewInp
}

type OrderViewRes struct {
	*input_spa.SpaOrderViewModel
}

type OrderConfirmAgreeReq struct {
	g.Meta `path:"/spaOrder/confirmAgree" method:"post" tags:"ADMIN_SPA" summary:"订单_确认同意"`
	input_spa.SpaOrderConfirmAgreeInp
}

type OrderConfirmAgreeRes struct {
}

type OrderConfirmDisagreeReq struct {
	g.Meta `path:"/spaOrder/confirmDisagree" method:"post" tags:"ADMIN_SPA" summary:"订单_确认拒绝"`
	input_spa.SpaOrderConfirmDisagreeInp
}

type OrderConfirmDisagreeRes struct {
}

type OrderTechnicianReq struct {
	g.Meta `path:"/spaOrder/technicianList" method:"get" tags:"ADMIN_SPA" summary:"获取满足订单条件的技师列表"`
	input_spa.SpaOrderTechnicianInp
}

type OrderTechnicianRes struct {
	input_form.PageRes
	List []*input_spa.SpaOrderTechnicianModel `json:"list"   dc:"数据列表"`
}

type OrderDispatchReq struct {
	g.Meta `path:"/spaOrder/dispatch" method:"post" tags:"ADMIN_SPA" summary:"订单_调度"`
	input_spa.SpaOrderDispatchInp
}

type OrderDispatchRes struct {
}

// TechnicianOrderListReq 技师结算订单列表
type TechnicianOrderListReq struct {
	g.Meta `path:"/spaOrder/technicianOrderList" method:"get" tags:"ADMIN_SPA" summary:"技师结算订单列表"`
	input_spa.SettleSpaOrderTechnicianListInp
}

type TechnicianOrderListRes struct {
	input_form.PageRes
	List []*input_spa.SettleSpaOrderTechnicianListModel `json:"list"   dc:"数据列表"`
}

type OrderGoOutReq struct {
	g.Meta `path:"/spaOrder/goOut" method:"post" tags:"ADMIN_SPA" summary:"订单_出发"`
	input_spa.SpaOrderGoOutInp
}

type OrderGoOutRes struct {
}

type OrderArriveReq struct {
	g.Meta `path:"/spaOrder/arrive" method:"post" tags:"ADMIN_SPA" summary:"订单_客人到店"`
	input_spa.SpaOrderGoOutInp
}

type OrderArriveRes struct {
}

type OrderStartServiceReq struct {
	g.Meta `path:"/spaOrder/startService" method:"post" tags:"ADMIN_SPA" summary:"订单_开始服务"`
	input_spa.SpaOrderStartServiceInp
}

type OrderStartServiceRes struct {
}

type OrderEndServiceReq struct {
	g.Meta `path:"/spaOrder/endService" method:"post" tags:"ADMIN_SPA" summary:"订单_结束服务"`
	input_spa.SpaOrderEndServiceInp
}

type OrderEndServiceRes struct {
}

type OrderCancelPayReq struct {
	g.Meta `path:"/spaOrder/cancelPay" method:"post" tags:"ADMIN_SPA" summary:"订单_取消"`
	input_spa.SpaOrderCancelPayInp
}

type OrderCancelPayRes struct {
}

type OrderExportReq struct {
	g.Meta `path:"/spaOrder/export" method:"post" tags:"ADMIN_SPA" summary:"订单-导出"`
	input_spa.SpaOrderExportInp
}

type OrderExportRes struct{}

type OrderExportListReq struct {
	g.Meta `path:"/spaOrder/exportList" method:"get" tags:"ADMIN_SPA" summary:"获取导出记录列表"`
	input_spa.SpaOrderExportListInp
}

type OrderExportListRes struct {
	input_form.PageRes
	List []*input_spa.SpaOrderExportListModel `json:"list"   dc:"数据列表"`
}

type OrderAbnormalReq struct {
	g.Meta `path:"/spaOrder/abnormal" method:"post" tags:"ADMIN_SPA" summary:"订单-异常处理"`
	input_spa.SpaOrderAbnormalInp
}

type OrderAbnormalRes struct {
}
