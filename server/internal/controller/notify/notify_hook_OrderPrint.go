package notify

import (
	"APT/internal/model/input/input_basics"
	"APT/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"APT/api/notify/hook"
)

func (c *ControllerHook) OrderPrint(ctx context.Context, req *hook.OrderPrintReq) (res *hook.OrderPrintRes, err error) {

	g.Log().Path("logs/HOOK/PRINTER_HOOK").Info(ctx, req)
	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if req.Scene == "SPA" {
			g.Log().Path("logs/HOOK/PRINTER_HOOK").Info(ctx, "按摩订单：")
			err = service.BasicsPrinter().PrinterSpaOrder(ctx, &input_basics.PrinterCarOrderInp{
				OrderSn: req.OrderSn,
			})

			return
		}
		if req.Scene == "CAR" {
			err = service.BasicsPrinter().PrinterCarOrder(ctx, &input_basics.PrinterCarOrderInp{
				OrderSn: req.OrderSn,
			})
			return
		}
		if req.Scene == "FOOD" {
			err = service.BasicsPrinter().PrinterFoodOrder(ctx, &input_basics.PrinterCarOrderInp{
				OrderSn: req.OrderSn,
			})
		}
		return
	}); err != nil {
		g.Log().Path("logs/HOOK/PRINTER_HOOK").Errorf(ctx, err.Error())
		return
	}
	return
}
