package app

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"APT/api/app/basics"
)

func (c *ControllerBasics) Chat(ctx context.Context, req *basics.ChatReq) (res *basics.ChatRes, err error) {
	var (
		Host = g.Cfg().MustGet(ctx, "kefu.kefuUrlHost").String()
	)
	switch req.Lang {
	case "zh":
		req.Lang = "cn"
		break
	}
	UrlPath := fmt.Sprintf("/chatIndex?kefu_id=%s&ent_id=%d&lang=%s", req.KefuId, req.EntId, req.Lang)
	ghttp.RequestFromCtx(ctx).Response.RedirectTo(Host + UrlPath)
	return
}
