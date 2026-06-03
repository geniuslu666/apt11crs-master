package admin

import (
	"APT/internal/service"
	"context"

	"APT/api/admin/pms"
)

func (c *ControllerPms) FinanceStat(ctx context.Context, req *pms.FinanceStatReq) (res *pms.FinanceStatRes, err error) {
	data, err := service.BasicsFinance().Stat(ctx, &req.FinanceStatInp)

	res = new(pms.FinanceStatRes)
	res.FinanceStatModel = data
	return
}
func (c *ControllerPms) FinanceList(ctx context.Context, req *pms.FinanceListReq) (res *pms.FinanceListRes, err error) {
	res = new(pms.FinanceListRes)
	if res.FinanceListModel, err = service.BasicsFinance().List(ctx, &req.FinanceListInp); err != nil {
		return
	}
	res.FinanceListModel.PayListItem.PageRes.Pack(req, res.FinanceListModel.PayCount)
	res.FinanceListModel.RefundListItem.PageRes.Pack(req, res.FinanceListModel.RefundCount)
	return
}
func (c *ControllerPms) FinanceExport(ctx context.Context, req *pms.FinanceExportReq) (res *pms.FinanceExportRes, err error) {
	err = service.BasicsFinance().Export(ctx, &req.FinanceListInp)
	return
}
