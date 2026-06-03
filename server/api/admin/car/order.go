package car

import (
	"APT/internal/model/input/input_car"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/carOrder/list" method:"get" tags:"ADMIN_CAR" summary:"获取接送机预订单列表"`
	input_car.CarOrderListInp
}

type OrderListRes struct {
	input_form.PageRes
	List []*input_car.CarOrderListModel `json:"list"   dc:"数据列表"`
}

type OrderViewReq struct {
	g.Meta `path:"/carOrder/view" method:"get" tags:"ADMIN_CAR" summary:"获取接送机预订单指定信息"`
	input_car.CarOrderViewInp
}

type OrderViewRes struct {
	*input_car.CarOrderViewModel
}

type OrderConfirmAgreeReq struct {
	g.Meta `path:"/carOrder/confirmAgree" method:"post" tags:"ADMIN_CAR" summary:"订单-确认同意"`
	input_car.CarOrderConfirmAgreeInp
}

type OrderConfirmAgreeRes struct {
}

type OrderConfirmDisagreeReq struct {
	g.Meta `path:"/carOrder/confirmDisagree" method:"post" tags:"ADMIN_CAR" summary:"订单-确认拒绝"`
	input_car.CarOrderConfirmDisagreeInp
}

type OrderConfirmDisagreeRes struct {
}

type OrderDriverReq struct {
	g.Meta `path:"/carOrder/driverList" method:"get" tags:"ADMIN_CAR" summary:"获取满足订单条件的司机列表"`
	input_car.CarOrderDriverInp
}

type OrderDriverRes struct {
	input_form.PageRes
	List []*input_car.CarOrderDriverModel `json:"list"   dc:"数据列表"`
}

type OrderDispatchReq struct {
	g.Meta `path:"/carOrder/dispatch" method:"post" tags:"ADMIN_CAR" summary:"订单-调度"`
	input_car.CarOrderDispatchInp
}

type OrderDispatchRes struct {
}

type SettleOrderListReq struct {
	g.Meta `path:"/carOrder/settleOrderList" method:"get" tags:"ADMIN_CAR" summary:"获取结算订单列表"`
	input_car.SettleCarOrderListInp
}

type SettleOrderListRes struct {
	input_form.PageRes
	List []*input_car.SettleCarOrderListModel `json:"list"   dc:"数据列表"`
}

type OrderRefundReq struct {
	g.Meta `path:"/carOrder/refund" method:"post" tags:"ADMIN_CAR" summary:"订单退款"`
	input_car.CarOrderRefundInp
}

type OrderRefundRes struct {
}

type OrderGoOutReq struct {
	g.Meta `path:"/carOrder/goOut" method:"post" tags:"ADMIN_CAR" summary:"订单_司机出发"`
	input_car.CarOrderGoOutInp
}

type OrderGoOutRes struct {
}

type OrderStartServiceReq struct {
	g.Meta `path:"/carOrder/startService" method:"post" tags:"ADMIN_CAR" summary:"订单_开始服务"`
	input_car.CarOrderStartServiceInp
}

type OrderStartServiceRes struct {
}

type OrderEndServiceReq struct {
	g.Meta `path:"/carOrder/endService" method:"post" tags:"ADMIN_CAR" summary:"订单_结束服务"`
	input_car.CarOrderEndServiceInp
}

type OrderEndServiceRes struct {
}

type OrderAbnormalReq struct {
	g.Meta `path:"/carOrder/abnormal" method:"post" tags:"ADMIN_CAR" summary:"订单-异常处理"`
	input_car.CarOrderAbnormalInp
}

type OrderAbnormalRes struct {
}

type OrderExportReq struct {
	g.Meta `path:"/carOrder/export" method:"post" tags:"ADMIN_CAR" summary:"订单-导出"`
	input_car.CarOrderExportInp
}

type OrderExportRes struct{}

type OrderExportListReq struct {
	g.Meta `path:"/carOrder/exportList" method:"get" tags:"ADMIN_CAR" summary:"获取导出记录列表"`
	input_car.CarOrderExportListInp
}

type OrderExportListRes struct {
	input_form.PageRes
	List []*input_car.CarOrderExportListModel `json:"list"   dc:"数据列表"`
}

type OrderTransferReq struct {
	g.Meta `path:"/carOrder/transfer" method:"post" tags:"ADMIN_CAR" summary:"订单-转单"`
	input_car.CarOrderTransferInp
}

type OrderTransferRes struct {
}
