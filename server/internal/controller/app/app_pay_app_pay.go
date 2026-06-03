package app

import (
	"APT/internal/library/cache"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"time"

	"APT/api/app/pay"
)

func (c *ControllerPay) ThirdPay(ctx context.Context, req *pay.ThirdPayReq) (res *pay.ThirdPayRes, err error) {
	res = new(pay.ThirdPayRes)
	if res.SelectThirdPayModel, err = service.PayService().SelectThirdPay(ctx, req.SelectThirdPayInp); err != nil {
		return
	}
	return
}
func (c *ControllerPay) InnerPay(ctx context.Context, req *pay.InnerPayReq) (res *pay.InnerPayRes, err error) {
	res = new(pay.InnerPayRes)
	if res.BalancePayModel, err = service.PayService().BalancePay(ctx, req.BalancePayInp); err != nil {
		return
	}
	return
}
func (c *ControllerPay) QueryStatus(ctx context.Context, req *pay.QueryStatusReq) (res *pay.QueryStatusRes, err error) {
	var (
		cacheData *gvar.Var
	)
	res = new(pay.QueryStatusRes)

	for i := 0; i < 15; i++ {
		cacheData, err = cache.Instance().Get(ctx, "PayOrderSn_"+req.OrderSn)

		if err == nil && !g.IsEmpty(cacheData.String()) {
			res.Status = cacheData.String()
			break
		}

		time.Sleep(time.Second * 2)
	}
	if g.IsEmpty(res.Status) {
		res.Status = "FAIL"
	}
	return
}
