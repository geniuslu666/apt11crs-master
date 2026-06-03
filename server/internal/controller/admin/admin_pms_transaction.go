package admin

import (
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) TransactionList(ctx context.Context, req *pms.TransactionListReq) (res *pms.TransactionListRes, err error) {
	list, totalCount, err := service.PayService().List(ctx, &req.PmsTransactionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*input_pay.PmsTransactionListModel{}
	}

	res = new(pms.TransactionListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}
func (c *ControllerPms) TransactionExport(ctx context.Context, req *pms.TransactionExportReq) (res *pms.TransactionExportRes, err error) {
	err = service.PayService().Export(ctx, &req.PmsTransactionListExportInp)
	return
}
func (c *ControllerPms) TransactionView(ctx context.Context, req *pms.TransactionViewReq) (res *pms.TransactionViewRes, err error) {
	data, err := service.PayService().View(ctx, &req.PmsTransactionViewInp)
	if err != nil {
		return
	}

	res = new(pms.TransactionViewRes)
	res.PmsTransactionViewModel = data
	return
}
