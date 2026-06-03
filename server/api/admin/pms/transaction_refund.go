package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_refund"
	"github.com/gogf/gf/v2/frame/g"
)

type TransactionRefundListReq struct {
	g.Meta `path:"/pmsTransactionRefund/list" method:"get" tags:"ADMIN_PMS" summary:"退款流水_列表"`
	input_refund.PmsTransactionRefundListInp
}

type TransactionRefundListRes struct {
	input_form.PageRes
	List []*input_refund.PmsTransactionRefundListModel `json:"list"   dc:"数据列表"`
}

type TransactionRefundExportReq struct {
	g.Meta `path:"/pmsTransactionRefund/export" method:"get" tags:"ADMIN_PMS" summary:"退款流水_导出"`
	input_refund.PmsTransactionRefundListExportInp
}

type TransactionRefundExportRes struct{}

type TransactionRefundViewReq struct {
	g.Meta `path:"/pmsTransactionRefund/view" method:"get" tags:"ADMIN_PMS" summary:"退款流水_详情"`
	input_refund.PmsTransactionRefundViewInp
}

type TransactionRefundViewRes struct {
	*input_refund.PmsTransactionRefundViewModel
}
