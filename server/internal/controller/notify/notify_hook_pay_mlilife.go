package notify

import (
	"APT/api/notify/hook"
	"APT/internal/dao"
	"APT/internal/model/input/input_pay"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerHook) MlilifePay(ctx context.Context, req *hook.MlilifePayReq) (res *hook.MlilifePayRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
	)
	g.Log().Info(ctx, req)
	if req.TradeState == "REFUND" {
		return
	} else if req.TradeState == "SUCCESS" {
		if _, err = dao.PmsTransaction.Ctx(ctx).
			Where(dao.PmsTransaction.Columns().TransactionSn, req.OutTradeNo).Data(g.Map{
			dao.PmsTransaction.Columns().PaymentRequestId: req.TransactionID,
			dao.PmsTransaction.Columns().PayAmount:        req.TotalFee,
			dao.PmsTransaction.Columns().PayStatus:        "DONE",
		}).Update(); err != nil {
			return
		}
	} else {
		r.Response.WriteExit("fail")
	}

	if err = service.PayService().ThirdPayCompleted(ctx, &input_pay.ThirdPayInp{
		//TransNo:   req.Resource.SupplementaryData.RelatedIds.OrderID,
		TransNo: req.TransactionID,
	}); err != nil {
		g.Log().Path("logs/HOOK/PAYPAL_HOOK").Error(ctx, err)
		err = nil
	}
	r.Response.WriteExit("SUCCESS")
	return
}
