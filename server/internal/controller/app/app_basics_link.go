package app

import (
	"APT/api/app/basics"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerBasics) Link(ctx context.Context, req *basics.LinkReq) (res *basics.LinkRes, err error) {
	tplContent := `referrer_id:{{.referrer_id}}, name:{{.name}}`
	if err = ghttp.RequestFromCtx(ctx).Response.WriteTplContent(tplContent, g.Map{
		"referrer_id": req.ReferrerId,
	}); err != nil {
		return
	}
	return
}
func (c *ControllerBasics) WxReturnUrl(ctx context.Context, req *basics.WxReturnUrlReq) (res *basics.WxReturnUrlRes, err error) {
	var (
		r = ghttp.RequestFromCtx(ctx)
	)
	r.Response.RedirectTo("apt11://createOrderReturnWx")
	return
}
