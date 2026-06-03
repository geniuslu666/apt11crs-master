package notify

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"APT/api/notify/paypay"
)

func (c *ControllerPaypay) DoPayPayCallback(ctx context.Context, req *paypay.DoPayPayCallbackReq) (res *paypay.DoPayPayCallbackRes, err error) {
	var (
		r    = ghttp.RequestFromCtx(ctx)
		body = r.GetBody()
	)
	g.Log().Debug(ctx, "paypay_回调参数", body)
	return
}
