package admin

import (
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) TransactionRefundList(ctx context.Context, req *pms.TransactionRefundListReq) (res *pms.TransactionRefundListRes, err error) {
	list, totalCount, err := service.RefundTransaction().List(ctx, &req.PmsTransactionRefundListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_refund.PmsTransactionRefundListModel{}
	}

	res = new(pms.TransactionRefundListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) TransactionRefundExport(ctx context.Context, req *pms.TransactionRefundExportReq) (res *pms.TransactionRefundExportRes, err error) {
	err = service.RefundTransaction().Export(ctx, &req.PmsTransactionRefundListExportInp)
	return
}
func (c *ControllerPms) TransactionRefundView(ctx context.Context, req *pms.TransactionRefundViewReq) (res *pms.TransactionRefundViewRes, err error) {
	data, err := service.RefundTransaction().View(ctx, &req.PmsTransactionRefundViewInp)
	if err != nil {
		return
	}

	res = new(pms.TransactionRefundViewRes)
	res.PmsTransactionRefundViewModel = data
	return
}
