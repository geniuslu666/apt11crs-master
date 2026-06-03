package websocket

import (
	"APT/api/websocket/basics"
	"APT/internal/websocket"
	"APT/utility/simple"
	"context"
)

func (c *ControllerBasics) SendToTag(ctx context.Context, req *basics.SendToTagReq) (res *basics.SendToTagRes, err error) {
	simple.SafeGo(ctx, func(ctx context.Context) {
		websocket.SendToTag(req.Tag, &websocket.WResponse{
			Event: req.Response.Event,
			Data:  req.Response,
		})
	})
	return
}
