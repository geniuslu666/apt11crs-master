package app

import (
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/travel"
)

func (c *ControllerTravel) CancelOrderNoPay(ctx context.Context, req *travel.CancelOrderNoPayReq) (res *travel.CancelOrderNoPayRes, err error) {
	if err = service.TravelOrder().OrderExpiration(ctx, req.OrderSn); err != nil {
		// 订单过期处理失败
		err = gerror.New(gi18n.T(ctx, "order_expire_handle_failed"))
		return
	}
	return
}
