package cabinet

import (
	"APT/internal/model/input/input_cabinet"
	"APT/internal/model/input/input_form"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderListReq struct {
	g.Meta `path:"/cabinetOrder/list" method:"get" tags:"ADMIN_CABINET" summary:"储物柜订单列表"`
	input_cabinet.OrderListInp
}

type OrderListRes struct {
	input_form.PageRes
	List []*input_cabinet.OrderListModel `json:"list"   dc:"数据列表"`
}

type OrderViewReq struct {
	g.Meta `path:"/cabinetOrder/view" method:"get" tags:"ADMIN_CABINET" summary:"储物柜订单详情"`
	input_cabinet.OrderViewInp
}

type OrderViewRes struct {
	*input_cabinet.OrderViewModel
}

type OrderExportReq struct {
	g.Meta `path:"/cabinetOrder/export" method:"post" tags:"ADMIN_CABINET" summary:"订单-导出"`
	input_cabinet.OrderExportInp
}

type OrderExportRes struct{}

type OrderExportListReq struct {
	g.Meta `path:"/cabinetOrder/exportList" method:"get" tags:"ADMIN_CABINET" summary:"获取导出记录列表"`
	input_cabinet.OrderExportListInp
}

type OrderExportListRes struct {
	input_form.PageRes
	List []*input_cabinet.OrderExportListModel `json:"list"   dc:"数据列表"`
}

type OrderCompleteReq struct {
	g.Meta `path:"/cabinetOrder/complete" method:"post" tags:"ADMIN_CABINET" summary:"订单-完成"`
	input_cabinet.OrderCompleteInp
}

type OrderCompleteRes struct{}

type OrderRefundReq struct {
	g.Meta `path:"/cabinetOrder/refund" method:"post" tags:"ADMIN_CABINET" summary:"订单退款"`
	input_cabinet.CabinetOrderRefundInp
}

type OrderRefundRes struct {
}
