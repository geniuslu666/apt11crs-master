package notify

import (
	"APT/internal/library/paycloud"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
	"os"

	"APT/api/notify/test"
)

func (c *ControllerTest) PaycloudCardPay(ctx context.Context, req *test.PaycloudCardPayReq) (res *test.PaycloudCardPayRes, err error) {
	var (
		paycloudClient *paycloud.Client
		CreateWebOrder paycloud.CreateWebPayOrderResponse
		r              = ghttp.RequestFromCtx(ctx)
	)
	if paycloudClient, err = paycloud.NewClient(ctx, ""); err != nil {
		return
	}
	paycloudClient.Endpoint = "https://open.n-age.co.jp/api/entry"
	paycloudClient.AppID = "wz715fc0d10ee9d156"
	paycloudClient.MerchantNo = "312100007235"
	paycloudClient.StoreNo = "4122000030"
	paycloudClient.SubAppid = ""
	paycloudClient.PrivateKey = os.Getenv("PAYCLOUD_PRIVATE_KEY")
	if CreateWebOrder, err = paycloudClient.WebOrder(ctx, &paycloud.CreateWebPayOrderParams{
		MerchantOrderNo: guid.S(),
		OrderAmount:     req.Amount,
	}); err != nil {
		return
	}
	r.Response.RedirectTo(CreateWebOrder.Data.PayUrl)
	return
}
