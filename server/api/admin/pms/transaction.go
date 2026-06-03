package pms

import (
	"APT/internal/model/input/input_form"
	"APT/internal/model/input/input_pay"
	"github.com/gogf/gf/v2/frame/g"
)

type TransactionListReq struct {
	g.Meta `path:"/pmsTransaction/list" method:"get" tags:"ADMIN_PMS" summary:"支付流水_列表"`
	input_pay.PmsTransactionListInp
}

type TransactionListRes struct {
	input_form.PageRes
	List []*input_pay.PmsTransactionListModel `json:"list"   dc:"数据列表"`
}

type TransactionExportReq struct {
	g.Meta `path:"/pmsTransaction/export" method:"get" tags:"ADMIN_PMS" summary:"支付流水_导出"`
	input_pay.PmsTransactionListExportInp
}

type TransactionExportRes struct{}

type TransactionViewReq struct {
	g.Meta `path:"/pmsTransaction/view" method:"get" tags:"ADMIN_PMS" summary:"支付流水_详情"`
	input_pay.PmsTransactionViewInp
}

type TransactionViewRes struct {
	*input_pay.PmsTransactionViewModel
}
