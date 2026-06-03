package notify

import (
	"APT/internal/dao"
	"APT/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"

	"APT/api/notify/hook"
)

func (c *ControllerHook) Aladdin(ctx context.Context, req *hook.AladdinReq) (res *hook.AladdinRes, err error) {
	var (
		r           = ghttp.RequestFromCtx(ctx)
		CarOrder    *entity.CarOrder
		gHttpClient = g.Client()
		response    *gclient.Response
		Logger      = g.Log().Path("logs/HOOK/ALADDIN")
	)
	Logger.Info(ctx, gvar.New(r.GetBody()).String())
	Logger.Info(ctx, req)
	if err = dao.CarOrder.Ctx(ctx).Where(g.Map{
		dao.CarOrder.Columns().InnnOrderId: req.OrderId,
		dao.CarOrder.Columns().InnnOrderNo: req.OrderNo,
	}).Scan(&CarOrder); err != nil {
		return
	}
	if g.IsEmpty(CarOrder) {
		err = gerror.New("不存在订单")
		return
	}
	gHttpClient.SetHeader("Content-Type", "application/json")
	if response, err = gHttpClient.Post(ctx, "http://127.0.0.1:9000/car/order/endServing", g.MapStrAny{
		"orderSn": CarOrder.OrderSn,
	}); err != nil {
		return
	}
	Logger.Info(ctx, response.Raw())
	if gjson.New(response.ReadAll()).Get("code").Int() == 0 {
		return
	}
	res = new(hook.AladdinRes)
	return
}
