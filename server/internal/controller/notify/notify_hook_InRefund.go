package notify

import (
	"APT/internal/model/input/input_refund"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/notify/hook"
)

func (c *ControllerHook) InRefund(ctx context.Context, req *hook.InRefundReq) (res *hook.InRefundRes, err error) {
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.Refund().RefundOrder(ctx, &input_refund.RefundAmountInp{
			OrderSn:      req.OrderSn,
			RefundAmount: req.Amount,
		}, tx)
		return
	}); err != nil {
		return
	}
	return
}
