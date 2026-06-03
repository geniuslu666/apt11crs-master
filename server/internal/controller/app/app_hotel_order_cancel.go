package app

import (
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/errors/gerror"

	"APT/api/app/hotel"
)

func (c *ControllerHotel) CancelOrderNoPay(ctx context.Context, req *hotel.CancelOrderNoPayReq) (res *hotel.CancelOrderNoPayRes, err error) {
	if err = service.HotelService().OrderExpiration(ctx, req.OrderSn); err != nil {
		// 订单过期处理失败
		err = gerror.New(gi18n.T(ctx, "order_expire_handle_failed"))
		return
	}
	return
}
